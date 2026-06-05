package logger

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type RuntimeLogger struct {
	ctx context.Context
	id  string
}

func NewRuntimeLogger(ctx context.Context, id string) *RuntimeLogger {
	return &RuntimeLogger{ctx: ctx, id: id}
}

func (l *RuntimeLogger) Debugf(format string, args ...interface{}) {
	runtime.LogDebugf(l.ctx, l.id+": "+format, args...)
}

func (l *RuntimeLogger) Infof(format string, args ...interface{}) {
	runtime.LogInfof(l.ctx, l.id+": "+format, args...)
}

func (l *RuntimeLogger) Errorf(format string, args ...interface{}) {
	runtime.LogErrorf(l.ctx, l.id+": "+format, args...)
}
