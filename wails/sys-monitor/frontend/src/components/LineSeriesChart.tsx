import { FC, useMemo } from "react";

export type LineSeriesChartProps = {
  data: number[];
  width?: number;
  height?: number;
  min?: number;
  max?: number;
  className?: string;
};

export const LineSeriesChart: FC<LineSeriesChartProps> = ({
  data,
  width = 320,
  height = 120,
  min = 0,
  max = 100,
  className,
}) => {
  const points = useMemo(() => {
    if (data.length === 0) {
      return "";
    }

    const safeRange = Math.max(1, max - min);
    const step = data.length > 1 ? width / (data.length - 1) : 0;

    return data
      .map((value, index) => {
        const clamped = Math.min(max, Math.max(min, value));
        const x = Math.round(index * step);
        const y = Math.round(height - ((clamped - min) / safeRange) * height);
        return `${x},${y}`;
      })
      .join(" ");
  }, [data, height, max, min, width]);

  return (
    <div className={className}>
      <svg
        className="series-chart"
        viewBox={`0 0 ${width} ${height}`}
        preserveAspectRatio="none"
        role="img"
        aria-label="CPU usage time series"
      >
        <polyline
          className="series-chart__line"
          fill="none"
          strokeWidth="2"
          points={points}
        />
      </svg>
    </div>
  );
};
