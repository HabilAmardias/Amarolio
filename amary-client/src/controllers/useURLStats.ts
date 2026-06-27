import { useEffect, useState } from "react";
import { getURLMetadata, getURLStats } from "../api/url.api";
import type { VisitDashboardRes, UserLink } from "../models/url.model";

export function useURLStats(urlId: string | undefined) {
  const [stats, setStats] = useState<VisitDashboardRes | null>(null);
  const [urlDetails, setUrlDetails] = useState<UserLink | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<Error | null>(null);

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

  return { stats, urlDetails, loading, error, refetch: fetchStats };
}

