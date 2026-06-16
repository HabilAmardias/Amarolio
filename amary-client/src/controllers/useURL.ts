import { useEffect, useState, useRef } from "react";
import { getUserURLs as apiGetUserURLs } from "../api/url.api";
import type { UserLink } from "../models/url.model";

type PaginationState = {
  page: number;
  cursor: number | undefined;
  limit: number;
};

export function useURL() {
  const [userURL, setUserURL] = useState<UserLink[]>([]);
  const [hasNextPage, setHasNextPage] = useState<boolean>(false);
  const [pagination, setPagination] = useState<PaginationState>({
    page: 0,
    cursor: undefined,
    limit: 5,
  });
  const [error, setError] = useState<Error | null>(null);

  const cursorsRef = useRef<Record<number, number | undefined>>({
    0: undefined,
  });

  useEffect(() => {
    let isMounted = true;

    apiGetUserURLs(pagination.limit + 1, pagination.cursor)
      .then((val) => {
        if (!isMounted) return;

        const hasNext = val.entries.length > pagination.limit;
        const pageItems = hasNext
          ? val.entries.slice(0, pagination.limit)
          : val.entries;

        setHasNextPage(hasNext);
        setUserURL(pageItems);

        const lastVisibleItem = pageItems[pageItems.length - 1];
        if (lastVisibleItem?.id) {
          cursorsRef.current = {
            ...cursorsRef.current,
            [pagination.page + 1]: lastVisibleItem.id,
          };
        }
      })
      .catch((err) => {
        setError(err as Error);
      });

    return () => {
      isMounted = false;
    };
  }, [pagination]); // ✅ only one state object drives everything

  function handleChangePage(
    _: React.MouseEvent<HTMLButtonElement> | null,
    newPage: number,
  ) {
    setPagination((prev) => ({
      ...prev,
      page: newPage,
      cursor: cursorsRef.current[newPage],
    }));
  }

  function handleChangeLimit(
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) {
    const newLimit = parseInt(event.target.value);
    cursorsRef.current = { 0: undefined };
    // ✅ single atomic update, one effect run
    setPagination({ page: 0, cursor: undefined, limit: newLimit });
  }

  return {
    userURL,
    limit: pagination.limit,
    page: pagination.page,
    hasNextPage,
    error,
    handleChangePage,
    handleChangeLimit,
  };
}
