import { useEffect, useState } from "react";
import { getUserURLs as apiGetUserURLs } from "../api/url.api";
import type { UserLink } from "../models/url.model";

export function useURL() {
  const [userURL, setUserURL] = useState<UserLink[]>([]);
  const [limit, setLimit] = useState<number>(5);
  const [page, setPage] = useState<number>(0);
  const [hasNextPage, setHasNextPage] = useState<boolean>(false);
  const [cursors, setCursors] = useState<Record<number, number | undefined>>({
    0: undefined,
  });
  const [activeCursor, setActiveCursor] = useState<number | undefined>(
    undefined,
  );
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    let isMounted = true;

    apiGetUserURLs(limit + 1, activeCursor)
      .then((val) => {
        if (!isMounted) return;
        if (val.entries.length > limit) {
          setHasNextPage(true);
          setUserURL(val.entries.slice(0, limit));
        } else {
          setHasNextPage(false);
          setUserURL(val.entries);
        }
        if (val.page_info?.last_id) {
          setCursors((prev) => ({
            ...prev,
            [page + 1]: val.page_info.last_id,
          }));
        }
      })
      .catch((err) => {
        setError(err as Error);
      });

    return () => {
      isMounted = false;
    };
  }, [limit, activeCursor, page]);

  async function handleChangePage(
    _: React.MouseEvent<HTMLButtonElement> | null,
    newPage: number,
  ) {
    const nextCursor = cursors[newPage];

    setActiveCursor(nextCursor);
    setPage(newPage);
  }

  function handleChangeLimit(
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) {
    const newLimit = parseInt(event.target.value);

    setLimit(newLimit);
    setPage(0);
    setActiveCursor(undefined);
    setCursors({ 0: undefined });
  }

  return {
    userURL,
    limit,
    page,
    hasNextPage,
    error,
    handleChangePage,
    handleChangeLimit,
  };
}
