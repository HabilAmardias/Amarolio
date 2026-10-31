package entity

import (
	"amarfin/src/customerror"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Kalau daily expectnya local_time, start_at (date), end_at(date, nullable)
// Kalau monthly expectnya date_of_month (1-28), local_time, start_at (date), end_at (date, nullable)

type (
	DateTime struct {
		value time.Time
	}

	AuditTimestamp struct {
		CreatedAt string
		UpdatedAt string
		DeletedAt *string
	}
)

func NewDateTime(
	year int,
	month time.Month,
	day int,
	hour int,
	minute int,
	second int,
	tz string,
) (*DateTime, error) {
	if tz == "" || strings.TrimSpace(tz) != tz || tz == "Local" {
		return nil, customerror.NewError(
			"invalid timezone",
			fmt.Errorf("invalid timezone: %q", tz),
			customerror.InvalidAction,
		)
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, customerror.NewError(
			"invalid timezone",
			err,
			customerror.InvalidAction,
		)
	}

	return &DateTime{
		value: time.Date(
			year,
			month,
			day,
			hour,
			minute,
			second,
			0,
			loc,
		),
	}, nil
}

func (dt *DateTime) ToTimestampString() string {
	return dt.value.String()
}

func (dt *DateTime) Add(t time.Duration) {
	dt.value.Add(t)
}

func (dt *DateTime) GetDay() int {
	return dt.value.Day()
}

func (dt *DateTime) ToTimeString() string {
	hour := dt.value.Hour()
	minute := dt.value.Minute()
	second := dt.value.Second()

	hourString := strconv.Itoa(hour)
	if hour < 10 {
		hourString = fmt.Sprintf("0%s", hourString)
	}

	minuteString := strconv.Itoa(minute)
	if minute < 10 {
		minuteString = fmt.Sprintf("0%s", minuteString)
	}

	secondString := strconv.Itoa(second)
	if minute < 10 {
		secondString = fmt.Sprintf("0%s", secondString)
	}
	// return hh:mm:ss
	return fmt.Sprintf("%s:%s:%s", hourString, minuteString, secondString)
}

func (dt *DateTime) ToDateString() string {
	day := dt.value.Day()
	month := dt.value.Month()
	year := dt.value.Year()

	dayString := strconv.Itoa(day)
	if day < 10 {
		dayString = fmt.Sprintf("0%s", dayString)
	}

	monthString := strconv.Itoa(int(month))
	if month < 10 {
		monthString = fmt.Sprintf("0%s", monthString)
	}
	// return dd-mm-yyyy
	return fmt.Sprintf("%s-%s-%d", dayString, monthString, year)
}
