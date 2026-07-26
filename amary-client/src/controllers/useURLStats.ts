import { useEffect, useState, useRef } from "react";
import { getURLMetadata, getURLStats } from "../api/url.api";
import type { VisitDashboardRes, UserLink } from "../models/url/type";

import { Chart, registerables } from "chart.js";

Chart.register(...registerables);

export function useURLStats(urlId: string | undefined) {
  const [stats, setStats] = useState<VisitDashboardRes | null>(null);
  const [urlDetails, setUrlDetails] = useState<UserLink | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);


  const [deviceTab, setDeviceTab] = useState(0); // 0 = Today, 1 = This Week
  const [copied, setCopied] = useState(false);

  // Chart refs
  const deviceChartRef = useRef<HTMLCanvasElement | null>(null);
  const dailyChartRef = useRef<HTMLCanvasElement | null>(null);
  const stackedChartRef = useRef<HTMLCanvasElement | null>(null);

  // Chart instances to destroy on recreate
  const deviceChartInstance = useRef<Chart | null>(null);
  const dailyChartInstance = useRef<Chart | null>(null);
  const stackedChartInstance = useRef<Chart | null>(null);

  const handleChangeTab = (_: React.SyntheticEvent<Element, Event>, val: number) => {
    setDeviceTab(val);
  }

  const handleCopy = () => {
    if (urlDetails?.short_url) {
      navigator.clipboard.writeText(urlDetails.short_url);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  const fetchStats = async () => {
    if (!urlId) {
      setLoading(false);
      return;
    }
    setLoading(true);
    setError(null);
    try {
      const statsData = await getURLStats(urlId);
      setStats(statsData);

      const metaData = await getURLMetadata(urlId)
      setUrlDetails(metaData.url);
    } catch (err) {
      setError(err as Error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStats();
  }, [urlId]);

  // 2. Device Chart (Doughnut)
  useEffect(() => {
    if (loading || !stats || !deviceChartRef.current) return;

    const deviceData = deviceTab === 0 ? stats.today_device_count : stats.this_week_device_count;
    const labels = deviceData.map((d) => d.device);
    const counts = deviceData.map((d) => d.count);

    if (deviceChartInstance.current) {
      deviceChartInstance.current.destroy();
    }

    deviceChartInstance.current = new Chart(deviceChartRef.current, {
      type: "doughnut",
      data: {
        labels: labels,
        datasets: [
          {
            data: counts,
            backgroundColor: ["#c25e00", "#8b4513", "#f4a460"],
            borderWidth: 2,
            borderColor: "#ffffff",
          },
        ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: "bottom",
            labels: {
              font: {
                family: "'Merriweather', 'Georgia', serif",
              },
              color: "#3e2723",
            },
          },
        },
      },
    });

    return () => {
      if (deviceChartInstance.current) {
        deviceChartInstance.current.destroy();
      }
    };
  }, [stats, deviceTab, loading]);

  // 3. Daily Activity Chart (Bar)
  useEffect(() => {
    if (loading || !stats || !dailyChartRef.current) return;

    const labels = stats.this_day_of_week_count.map((d) => d.day_of_week);
    const counts = stats.this_day_of_week_count.map((d) => d.count);

    if (dailyChartInstance.current) {
      dailyChartInstance.current.destroy();
    }

    dailyChartInstance.current = new Chart(dailyChartRef.current, {
      type: "bar",
      data: {
        labels: labels,
        datasets: [
          {
            label: "Visits",
            data: counts,
            backgroundColor: "rgba(194, 94, 0, 0.85)",
            hoverBackgroundColor: "#c25e00",
            borderRadius: 6,
            borderWidth: 0,
          },
        ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false,
          },
        },
        scales: {
          x: {
            grid: {
              display: false,
            },
            ticks: {
              color: "#5d4037",
              font: {
                family: "'Merriweather', 'Georgia', serif",
                size: 11,
              },
            },
          },
          y: {
            beginAtZero: true,
            grid: {
              color: "rgba(139, 69, 19, 0.08)",
            },
            ticks: {
              color: "#5d4037",
              font: {
                size: 11,
              },
            },
          },
        },
      },
    });

    return () => {
      if (dailyChartInstance.current) {
        dailyChartInstance.current.destroy();
      }
    };
  }, [stats, loading]);

  // 4. Stacked Device Activity by Day
  useEffect(() => {
    if (loading || !stats || !stackedChartRef.current) return;

    const days = ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"];
    const devices = [... new Set(stats.today_device_count.map(d => d.device))]

    const datasets = devices.map((device, idx) => {
      const colors = ["#c25e00", "#8b4513", "#f4a460"];
      return {
        label: device,
        data: days.map((day) => {
          const match = stats.this_week_dow_device_count.find(
            (item) => item.day_of_week === day && item.device === device
          );
          return match ? match.count : 0;
        }),
        backgroundColor: colors[idx],
        borderRadius: 4,
      };
    });

    if (stackedChartInstance.current) {
      stackedChartInstance.current.destroy();
    }

    stackedChartInstance.current = new Chart(stackedChartRef.current, {
      type: "bar",
      data: {
        labels: days,
        datasets: datasets,
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: "bottom",
            labels: {
              font: {
                family: "'Merriweather', 'Georgia', serif",
              },
              color: "#3e2723",
            },
          },
        },
        scales: {
          x: {
            stacked: true,
            grid: {
              display: false,
            },
            ticks: {
              color: "#5d4037",
              font: {
                family: "'Merriweather', 'Georgia', serif",
                size: 11,
              },
            },
          },
          y: {
            stacked: true,
            beginAtZero: true,
            grid: {
              color: "rgba(139, 69, 19, 0.08)",
            },
            ticks: {
              color: "#5d4037",
              font: {
                size: 11,
              },
            },
          },
        },
      },
    });

    return () => {
      if (stackedChartInstance.current) {
        stackedChartInstance.current.destroy();
      }
    };
  }, [stats, loading]);

  return {
    stats,
    urlDetails,
    loading,
    error,
    copied,
    handleCopy,
    deviceTab,
    handleChangeTab,
    deviceChartRef,
    dailyChartRef,
    stackedChartRef,
  };
}

