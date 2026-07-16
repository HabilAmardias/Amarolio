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
	GetThisWeekCountGroupByDeviceAndDayOfWeek(ctx context.Context, urlID int64) ([]DeviceDayOfWeekCount, error)
	GetThisWeekCountGroupByDevice(ctx context.Context, urlID int64) ([]DeviceCount, error)
	GetThisWeekCount(ctx context.Context, urlID int64) (int64, error)
	GetThisDayOfWeekCount(ctx context.Context, urlID int64) ([]DayOfWeekCount, error)
	GetTodayCountGroupByDevice(ctx context.Context, urlID int64) ([]DeviceCount, error)
	GetTodayCount(ctx context.Context, urlID int64) (int64, error)
}

type VisitRecordCacheItf interface {
	Set(ctx context.Context, code string, ttl time.Duration, vd VisitDashboard) error
	Get(ctx context.Context, code string) (VisitDashboard, error)
}

type URLRepoItf interface {
	FindByCode(ctx context.Context, shortCode string) (url.URL, error)
}

type URLCacheItf interface {
	Get(ctx context.Context, encodedID string) (url.URL, error)
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

func (vrs *VisitRecordServiceImpl) GetVisitRecordSummary(ctx context.Context, userID string, code string) (VisitDashboard, error) {
	url, err := vrs.uc.Get(ctx, code)
	if err != nil {
		url, err = vrs.ur.FindByCode(ctx, code)
		if err != nil {
			return VisitDashboard{}, err
		}
	}

	if url.UserID == nil {
		return VisitDashboard{}, customerror.NewError(
			"Forbidden",
			errors.New("url does not belong to anyone"),
			customerror.ForbiddenAction,
		)
	}

	if *url.UserID != userID {
		return VisitDashboard{}, customerror.NewError(
			"Forbidden",
			errors.New("url does not belong to the user"),
			customerror.ForbiddenAction,
		)
	}

	// if exists on cache, return immediately, otherwise fetch from DB and refresh the cache
	vd, err := vrs.vc.Get(ctx, code)
	if err == nil {
		return vd, nil
	}

	g, newCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		count, err := vrs.vrr.GetTodayCount(newCtx, url.ID)
		if err == nil {
			vd.TodayVisitCount = count
		}
		return err
	})
	g.Go(func() error {
		count, err := vrs.vrr.GetThisWeekCount(newCtx, url.ID)
		if err == nil {
			vd.ThisWeekCount = count
		}
		return err
	})
	g.Go(func() error {
		count, err := vrs.vrr.GetTodayCountGroupByDevice(newCtx, url.ID)
		if err == nil {
			vd.TodayDeviceCount = count
		}
		return err
	})
	g.Go(func() error {
		count, err := vrs.vrr.GetThisDayOfWeekCount(newCtx, url.ID)
		if err == nil {
			vd.ThisDayOfWeekCount = count
		}
		return err
	})
	g.Go(func() error {
		count, err := vrs.vrr.GetThisWeekCountGroupByDevice(newCtx, url.ID)
		if err == nil {
			vd.ThisWeekDeviceCount = count
		}
		return err
	})
	g.Go(func() error {
		count, err := vrs.vrr.GetThisWeekCountGroupByDeviceAndDayOfWeek(newCtx, url.ID)
		if err == nil {
			vd.ThisWeekDOWDeviceCount = count
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return vd, err
	}

	go func(cd string, v VisitDashboard) {
		ttl := 30 * time.Second
		conCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		fun := func() error {
			return vrs.vc.Set(conCtx, cd, ttl, v)
		}

		services.WithErrorRetry(conCtx, fun, 100*time.Millisecond)
	}(code, vd)

	return vd, nil
}
