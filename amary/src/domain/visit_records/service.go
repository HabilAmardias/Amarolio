package visitrecords

import (
	"amary/src/customerror"
	"amary/src/domain/url"
	"amary/src/services"
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

type VisitRecordRepoItf interface {
	GetThisWeekCountGroupByDeviceAndDayOfWeek(ctx context.Context, urlID int64, count *[]DeviceDayOfWeekCount) error
	GetThisWeekCountGroupByDevice(ctx context.Context, urlID int64, count *[]DeviceCount) error
	GetThisWeekCount(ctx context.Context, urlID int64, count *int64) error
	GetThisDayOfWeekCount(ctx context.Context, urlID int64, weekCounts *[]DayOfWeekCount) error
	GetTodayCountGroupByDevice(ctx context.Context, urlID int64, count *[]DeviceCount) error
	GetTodayCount(ctx context.Context, urlID int64, count *int64) error
}

type VisitRecordCacheItf interface {
	Set(ctx context.Context, code string, ttl time.Duration, vd VisitDashboard) error
	Get(ctx context.Context, code string, vd *VisitDashboard) error
}

type URLRepoItf interface {
	FindByCode(ctx context.Context, shortCode string, url *url.URL) error
}

type URLCacheItf interface {
	Get(ctx context.Context, encodedID string, url *url.URL) error
}

type VisitRecordServiceImpl struct {
	vrr VisitRecordRepoItf
	ur  URLRepoItf
	uc  URLCacheItf
	vc  VisitRecordCacheItf
}

func NewVisitRecordService(vrr VisitRecordRepoItf, ur URLRepoItf, uc URLCacheItf, vc VisitRecordCacheItf) *VisitRecordServiceImpl {
	return &VisitRecordServiceImpl{vrr, ur, uc, vc}
}

func (vrs *VisitRecordServiceImpl) GetVisitRecordSummary(ctx context.Context, userID string, code string) (*VisitDashboard, error) {
	var (
		todayVisitCount        *int64                 = new(int64)
		thisWeekCount          *int64                 = new(int64)
		todayDeviceCount       []DeviceCount          = make([]DeviceCount, 0)
		thisDayOfWeekCount     []DayOfWeekCount       = make([]DayOfWeekCount, 0)
		thisWeekDeviceCount    []DeviceCount          = make([]DeviceCount, 0)
		thisWeekDOWDeviceCount []DeviceDayOfWeekCount = make([]DeviceDayOfWeekCount, 0)
		url                    *url.URL               = new(url.URL)
		vd                     *VisitDashboard        = new(VisitDashboard)
	)

	if err := vrs.uc.Get(ctx, code, url); err != nil {
		if err := vrs.ur.FindByCode(ctx, code, url); err != nil {
			return nil, err
		}
	}

	if url.UserID == nil {
		return nil, customerror.NewError(
			"Forbidden",
			errors.New("url does not belong to anyone"),
			customerror.ForbiddenAction,
		)
	}

	if *url.UserID != userID {
		return nil, customerror.NewError(
			"Forbidden",
			errors.New("url does not belong to the user"),
			customerror.ForbiddenAction,
		)
	}

	// if exists on cache, return immediately, otherwise fetch from DB and refresh the cache
	if err := vrs.vc.Get(ctx, code, vd); err == nil {
		return vd, nil
	}

	g, newCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return vrs.vrr.GetTodayCount(newCtx, url.ID, todayVisitCount)
	})
	g.Go(func() error {
		return vrs.vrr.GetThisWeekCount(newCtx, url.ID, thisWeekCount)
	})
	g.Go(func() error {
		return vrs.vrr.GetTodayCountGroupByDevice(newCtx, url.ID, &todayDeviceCount)
	})
	g.Go(func() error {
		return vrs.vrr.GetThisDayOfWeekCount(newCtx, url.ID, &thisDayOfWeekCount)
	})
	g.Go(func() error {
		return vrs.vrr.GetThisWeekCountGroupByDevice(newCtx, url.ID, &thisWeekDeviceCount)
	})
	g.Go(func() error {
		return vrs.vrr.GetThisWeekCountGroupByDeviceAndDayOfWeek(newCtx, url.ID, &thisWeekDOWDeviceCount)
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	*vd = VisitDashboard{
		TodayVisitCount:        *todayVisitCount,
		ThisWeekCount:          *thisWeekCount,
		TodayDeviceCount:       todayDeviceCount,
		ThisDayOfWeekCount:     thisDayOfWeekCount,
		ThisWeekDeviceCount:    thisWeekDeviceCount,
		ThisWeekDOWDeviceCount: thisWeekDOWDeviceCount,
	}

	go func(cd string, v VisitDashboard) {
		ttl := 30 * time.Second
		conCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		fun := func() error {
			return vrs.vc.Set(conCtx, cd, ttl, v)
		}

		services.WithErrorRetry(conCtx, fun, 100*time.Millisecond)
	}(code, *vd)

	return vd, nil
}
