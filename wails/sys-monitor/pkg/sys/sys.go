// Package sys provides system monitoring helpers exposed to the frontend.
package sys

import (
	"context"
	"math"
	"sync"
	"sys-monitor/pkg/logger"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Stats exposes system metrics and streaming controls.
type Stats struct {
	ctx    context.Context
	log    *logger.RuntimeLogger
	mu     sync.Mutex
	cancel context.CancelFunc
	sub    func()
	unSub  func()
}

// NewStats constructs a Stats instance for Wails binding.
func NewStats() *Stats {
	return &Stats{}
}

// CPUUsage contains CPU usage data sent to the frontend.
type CPUUsage struct {
	Average int `json:"average"`
}

// Startup is called by Wails after the app initializes.
func (s *Stats) Startup(ctx context.Context) {
	s.ctx = ctx
	s.log = logger.NewRuntimeLogger(ctx, "sys:stats")
	s.log.Infof("System stats initialized")

	s.sub = runtime.EventsOn(ctx, "cpuUsage:subscribe", func(...interface{}) {
		s.startStream()
	})
	s.unSub = runtime.EventsOn(ctx, "cpuUsage:unsubscribe", func(...interface{}) {
		s.stopStream()
	})
}

// Shutdown is called by Wails before the app exits.
func (s *Stats) Shutdown(_ context.Context) {
	s.stopStream()
	if s.sub != nil {
		s.sub()
		s.sub = nil
	}
	if s.unSub != nil {
		s.unSub()
		s.unSub = nil
	}
}

// GetCPUUsage returns the current average CPU usage.
func (s *Stats) GetCPUUsage() *CPUUsage {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		if s.log != nil {
			s.log.Errorf("Error getting CPU usage: %s", err.Error())
		}
		return nil
	}
	return &CPUUsage{
		Average: int(math.Round(percentages[0])),
	}
}

func (s *Stats) startStream() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cancel != nil || s.ctx == nil {
		return
	}

	// Create a cancellable context for the CPU usage streaming goroutine. This context will be used later to signal the goroutine to stop when the stream is unsubscribed or when the application shuts down.
	streamCtx, cancel := context.WithCancel(s.ctx)
	s.cancel = cancel

	s.log.Infof("CPU usage stream started")

	// Start a goroutine to emit CPU usage every second until the context is cancelled
	// This will allow the frontend to receive real-time updates of CPU usage while the stream is active.
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			// If the streaming context is cancelled, exit the goroutine to stop emitting CPU usage updates to the frontend. This will effectively stop the real-time updates of CPU usage until the stream is started again.
			case <-streamCtx.Done():
				return
			case <-ticker.C:
				usage := s.GetCPUUsage()
				if usage == nil {
					continue
				}
				// Emit the CPU usage to the frontend using Wails runtime events. The event name is "cpuUsage" and the payload is the CPUUsage struct containing the average CPU usage percentage.
				// https://wails.io/docs/next/reference/runtime/events/#eventsemit
				runtime.EventsEmit(s.ctx, "cpuUsage", usage)
				if s.log != nil {
					s.log.Debugf("Current CPU usage: %d%%", usage.Average)
				}
			}
		}
	}()
}

func (s *Stats) stopStream() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cancel == nil {
		return
	}

	// Cancel the streaming context to stop the goroutine that emits CPU usage updates to the frontend. This will effectively stop the real-time updates of CPU usage until the stream is started again.
	s.cancel()
	s.cancel = nil
	s.log.Infof("CPU usage stream stopped")
}
