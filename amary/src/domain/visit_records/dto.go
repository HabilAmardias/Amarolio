package visitrecords

type (
	DeviceCountRes struct {
		Device string `json:"device"`
		Count  int64  `json:"count"`
	}
	DayOfWeekCountRes struct {
		DayOfWeek string `json:"day_of_week"`
		Count     int64  `json:"count"`
	}
	DeviceDayOfWeekCountRes struct {
		Device    string `json:"device"`
		DayOfWeek string `json:"day_of_week"`
		Count     int64  `json:"count"`
	}
	VisitDashboardRes struct {
		TodayVisitCount        int64                     `json:"today_visit_count"`
		ThisWeekCount          int64                     `json:"this_week_count"`
		TodayDeviceCount       []DeviceCountRes          `json:"today_device_count"`
		ThisDayOfWeekCount     []DayOfWeekCountRes       `json:"this_day_of_week_count"`
		ThisWeekDeviceCount    []DeviceCountRes          `json:"this_week_device_count"`
		ThisWeekDOWDeviceCount []DeviceDayOfWeekCountRes `json:"this_week_dow_device_count"`
	}
)
