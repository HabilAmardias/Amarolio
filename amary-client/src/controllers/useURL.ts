import { useEffect, useState } from "react";
import { getUserURLs as apiGetUserURLs } from "../api/url.api";
import type { UserLink } from "../models/url.model";

export function useURL() {
  const [userURL, setUserURL] = useState<UserLink[]>([]);
  const [limit, setLimit] = useState<number>(25);
  const [totalRow, setTotalRow] = useState<number | undefined>(undefined);
  const [lastID, setLastID] = useState<number | undefined>(undefined);
  const [page, setPage] = useState<number>(0);
  const [cursors, setCursors] = useState<Record<number, number | undefined>>({
    0: undefined,
  });

  useEffect(() => {
    apiGetUserURLs().then((val) => {
      setUserURL(val.entries);
      setLimit(val.page_info.limit);
      setTotalRow(val.page_info.total_row);
      setLastID(val.page_info.last_id);
    });
  }, []);

  async function handleChangePage(
    _: React.MouseEvent<HTMLButtonElement> | null,
    newPage: number,
  ) {
    let targetCursorId: number | undefined;

    if (newPage > page) {
      targetCursorId = lastID;
      setCursors((prev) => ({ ...prev, [newPage]: lastID }));
    } else {
      targetCursorId = cursors[newPage];
    }

    const val = await apiGetUserURLs(limit, targetCursorId);

    setUserURL(val.entries);
    setTotalRow(val.page_info.total_row);
    setLastID(val.page_info.last_id);
    setPage(newPage);
  }

  async function handleChangeLimit(
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) {
    const newLimit = parseInt(event.target.value);

    setLimit(newLimit);
    setPage(0);
    setLastID(undefined);
    setCursors({ 0: undefined });

    const val = await apiGetUserURLs(newLimit, undefined);

    setUserURL(val.entries);
    setTotalRow(val.page_info.total_row);
    setLastID(val.page_info.last_id);
  }

  return {
    userURL,
    limit,
    page,
    totalRow,
    handleChangePage,
    handleChangeLimit,
  };
}
