import { atom } from "jotai";

export interface FindCustomURLRequest {
  custom_code: string;
}

export interface ShortenRequest {
  url: string;
  duration: number | null; // null = no expiration (auth only)
  custom_code: string | null;
}

export interface ShortenResponse {
  url: string;
  original_url: string;
  expired_at: string | null;
}

export interface UserLink {
  id: number;
  user_id: string | null;
  short_url: string;
  code: string;
  url: string;
  created_at: string;
  expired_at: string | null;
}

const initialToken: string = "";
export const tokenAtom = atom<string>(initialToken);

const initialResult: ShortenResponse | null = null;
export const resultAtom = atom<ShortenResponse | null>(initialResult);

export interface DeviceCount {
  device: string;
  count: number;
}

export interface DayOfWeekCount {
  day_of_week: string;
  count: number;
}

export interface DayOfWeekDeviceCount {
  device: string;
  day_of_week: string;
  count: number;
}

export interface VisitDashboardRes {
  today_visit_count: number;
  this_week_count: number;
  today_device_count: DeviceCount[];
  this_day_of_week_count: DayOfWeekCount[];
  this_week_device_count: DeviceCount[];
  this_week_dow_device_count: DayOfWeekDeviceCount[];
}

