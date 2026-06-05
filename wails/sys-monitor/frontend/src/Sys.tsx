import { FC, PropsWithChildren, useEffect, useRef, useState } from "react";

import { sys } from "../wailsjs/go/models";

import { GetCPUUsage } from "../wailsjs/go/sys/Stats";
import { EventsEmit, EventsOn } from "../wailsjs/runtime/runtime";

export type TSysProps = Record<string, never>;

export const Sys: FC<PropsWithChildren<TSysProps>> = () => {
  const [stats, setStats] = useState<sys.CPUUsage | null>(null);
  const [isStreaming, setIsStreaming] = useState(false);
  // Use a ref to store the unsubscribe function for the CPU usage updates. This allows us to call the unsubscribe function from different places in the component without having to worry about stale closures or dependencies in useEffect.
  const unSubRef = useRef<null | (() => void)>(null);

  const updateStats = () => {
    GetCPUUsage().then((result) => {
      console.log("Received CPU usage from backend:", result);
      setStats(result);
    });
  };

  const subscribeToCpuUsageUpdates = () => {
    if (unSubRef.current) {
      return;
    }

    EventsEmit("cpuUsage:subscribe");
    unSubRef.current = EventsOn("cpuUsage", (nextStats: sys.CPUUsage) => {
      console.log("Received CPU usage update from backend:", nextStats);
      setStats(nextStats);
    });
    setIsStreaming(true);
  };

  const unsubscribeFromCpuUsageUpdates = () => {
    if (unSubRef.current) {
      unSubRef.current();
      unSubRef.current = null;
    }

    EventsEmit("cpuUsage:unsubscribe");
    setIsStreaming(false);
  };

  useEffect(() => () => unsubscribeFromCpuUsageUpdates(), []);

  return (
    <div className="">
      <h3>CPU Usage</h3>
      <div className="code-block">
        <code>{JSON.stringify(stats, null, 2)}</code>
      </div>

      <button
        className="btn"
        onClick={updateStats}
      >
        Update stats
      </button>

      <button
        className="btn"
        onClick={subscribeToCpuUsageUpdates}
        disabled={isStreaming}
      >
        Listen for updates
      </button>

      <button
        className="btn"
        onClick={unsubscribeFromCpuUsageUpdates}
      >
        Stop listening for updates
      </button>
    </div>
  );
};
