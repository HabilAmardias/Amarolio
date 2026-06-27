package visitrecords

import (
	"amary/src/customerror"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type VisitRecordCacheImpl struct {
	rc *redis.Client
}

func NewVisitRecordCache(rc *redis.Client) *VisitRecordCacheImpl {
	return &VisitRecordCacheImpl{rc}
}

func (vc *VisitRecordCacheImpl) Set(ctx context.Context, code string, ttl time.Duration, vd VisitDashboard) error {
	var (
		todayDeviceCount       []TaggedDeviceCount
		thisDayOfWeekCount     []TaggedDayOfWeekCount
		thisWeekDeviceCount    []TaggedDeviceCount
		thisWeekDOWDeviceCount []TaggedDeviceDayOfWeekCount
	)
	for _, el := range vd.TodayDeviceCount {
		todayDeviceCount = append(todayDeviceCount, TaggedDeviceCount(el))
	}
	for _, el := range vd.ThisDayOfWeekCount {
		thisDayOfWeekCount = append(thisDayOfWeekCount, TaggedDayOfWeekCount(el))
	}
	for _, el := range vd.ThisWeekDeviceCount {
		thisWeekDeviceCount = append(thisWeekDeviceCount, TaggedDeviceCount(el))
	}
	for _, el := range vd.ThisWeekDOWDeviceCount {
		thisWeekDOWDeviceCount = append(thisWeekDOWDeviceCount, TaggedDeviceDayOfWeekCount(el))
	}

	key := fmt.Sprintf("shorten_url:%s:dashboard", code)
	taggedData := TaggedVisitDashboard{
		TodayVisitCount:        vd.TodayVisitCount,
		ThisWeekCount:          vd.ThisWeekCount,
		TodayDeviceCount:       todayDeviceCount,
		ThisDayOfWeekCount:     thisDayOfWeekCount,
		ThisWeekDeviceCount:    thisWeekDeviceCount,
		ThisWeekDOWDeviceCount: thisWeekDOWDeviceCount,
	}
	b, err := json.Marshal(taggedData)
	if err != nil {
		return err
	}
	_, err = vc.rc.Set(ctx, key, string(b), ttl).Result()
	return err
}

func (vc *VisitRecordCacheImpl) Get(ctx context.Context, code string, vd *VisitDashboard) error {
	key := fmt.Sprintf("shorten_url:%s:dashboard", code)
	taggedData := new(TaggedVisitDashboard)

	val, err := vc.rc.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return customerror.NewError(
				"url not found",
				err,
				customerror.ItemNotFound,
			)
		}
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	if err := json.Unmarshal([]byte(val), taggedData); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}

	var (
		todayDeviceCount       []DeviceCount
		thisDayOfWeekCount     []DayOfWeekCount
		thisWeekDeviceCount    []DeviceCount
		thisWeekDOWDeviceCount []DeviceDayOfWeekCount
	)
	for _, el := range taggedData.TodayDeviceCount {
		todayDeviceCount = append(todayDeviceCount, DeviceCount(el))
	}
	for _, el := range taggedData.ThisDayOfWeekCount {
		thisDayOfWeekCount = append(thisDayOfWeekCount, DayOfWeekCount(el))
	}
	for _, el := range taggedData.ThisWeekDeviceCount {
		thisWeekDeviceCount = append(thisWeekDeviceCount, DeviceCount(el))
	}
	for _, el := range taggedData.ThisWeekDOWDeviceCount {
		thisWeekDOWDeviceCount = append(thisWeekDOWDeviceCount, DeviceDayOfWeekCount(el))
	}

	*vd = VisitDashboard{
		TodayVisitCount:        taggedData.TodayVisitCount,
		ThisWeekCount:          taggedData.ThisWeekCount,
		TodayDeviceCount:       todayDeviceCount,
		ThisDayOfWeekCount:     thisDayOfWeekCount,
		ThisWeekDeviceCount:    thisWeekDeviceCount,
		ThisWeekDOWDeviceCount: thisWeekDOWDeviceCount,
	}

	return nil
}
