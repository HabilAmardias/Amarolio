package visitrecords

import "time"

type (
	VisitRecord struct {
		ID        string
		UserID    string
		URLID     int64
		Device    string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}
	VisitRecordStream struct {
		Action string
		UserID string
		URLID  int64
		Device string
		Time   string
	}
	DeviceCount struct {
		Device string
		Count  int64
	}

	DayOfWeekCount struct {
		DayOfWeek string
		Count     int64
	}
	DeviceDayOfWeekCount struct {
		Device    string
		DayOfWeek string
		Count     int64
	}
	VisitDashboard struct {
		TodayVisitCount        int64
		ThisWeekCount          int64
		TodayDeviceCount       []DeviceCount
		ThisDayOfWeekCount     []DayOfWeekCount
		ThisWeekDeviceCount    []DeviceCount
		ThisWeekDOWDeviceCount []DeviceDayOfWeekCount
	}
	TaggedDeviceCount struct {
		Device string `json:"device"`
		Count  int64  `json:"count"`
	}
	TaggedDayOfWeekCount struct {
		DayOfWeek string `json:"day_of_week"`
		Count     int64  `json:"count"`
	}
	TaggedDeviceDayOfWeekCount struct {
		Device    string `json:"device"`
		DayOfWeek string `json:"day_of_week"`
		Count     int64  `json:"count"`
	}
	TaggedVisitDashboard struct {
		TodayVisitCount        int64                        `json:"today_visit_count"`
		ThisWeekCount          int64                        `json:"this_week_count"`
		TodayDeviceCount       []TaggedDeviceCount          `json:"today_device_count"`
		ThisDayOfWeekCount     []TaggedDayOfWeekCount       `json:"this_day_of_week_count"`
		ThisWeekDeviceCount    []TaggedDeviceCount          `json:"this_week_device_count"`
		ThisWeekDOWDeviceCount []TaggedDeviceDayOfWeekCount `json:"this_week_dow_device_count"`
	}
)
