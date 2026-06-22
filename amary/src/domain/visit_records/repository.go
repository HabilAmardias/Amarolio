package visitrecords

import (
	"amary/src/customerror"
	"amary/src/db"
	"amary/src/repository"
	"context"
)

type VisitRecordRepoImpl struct {
	handle db.DBTX
}

func NewVisitRecordRepo(handle db.DBTX) *VisitRecordRepoImpl {
	return &VisitRecordRepoImpl{handle}
}

func (vrr *VisitRecordRepoImpl) GetThisWeekCountGroupByDeviceAndDayOfWeek(ctx context.Context, urlID int64, count *[]DeviceDayOfWeekCount) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	query := `
	WITH week_days AS (
		SELECT generate_series(
			date_trunc('week', CURRENT_TIMESTAMP)::date,
			date_trunc('week', CURRENT_TIMESTAMP)::date + 6,
			'1 day'::interval
		) AS day
	),
	devices AS (
		SELECT DISTINCT device
		FROM visit_records
		WHERE url_id = :your_url_id AND deleted_at IS NULL
		-- Optional: add other global device lists if needed
	)
	SELECT
		d.device,
		TO_CHAR(wd.day, 'Day') AS day_of_week,
		COALESCE(COUNT(vr.id), 0) AS visit_count
	FROM devices d
	CROSS JOIN week_days wd
	LEFT JOIN visit_records vr
		ON vr.device = d.device
		AND vr.created_at::date = wd.day
		AND vr.url_id = :your_url_id
		AND vr.deleted_at IS NULL
	GROUP BY d.device, wd.day
	ORDER BY d.device, wd.day;
	`
	rows, err := driver.QueryContext(ctx, query, urlID)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer rows.Close()
	for rows.Next() {
		var item DeviceDayOfWeekCount
		if err := rows.Scan(
			&item.Device,
			&item.DayOfWeek,
			&item.Count,
		); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.DatabaseExecutionErr,
			)
		}
		*count = append(*count, item)
	}
	if err := rows.Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (vrr *VisitRecordRepoImpl) GetThisWeekCountGroupByDevice(ctx context.Context, urlID int64, count *[]DeviceCount) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	query := `
	SELECT
		device,
		COUNT(*) AS visit_count
	FROM visit_records
	WHERE
		url_id = $1
		AND deleted_at IS NULL
		AND created_at >= date_trunc('week', CURRENT_TIMESTAMP)
		AND created_at <  date_trunc('week', CURRENT_TIMESTAMP) + INTERVAL '7 days'
	GROUP BY device
	ORDER BY visit_count DESC;
	`

	rows, err := driver.QueryContext(ctx, query, urlID)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer rows.Close()
	for rows.Next() {
		var item DeviceCount
		if err := rows.Scan(
			&item.Device,
			&item.Count,
		); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.DatabaseExecutionErr,
			)
		}
		*count = append(*count, item)
	}
	if err := rows.Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (vrr *VisitRecordRepoImpl) GetThisWeekCount(ctx context.Context, urlID int64, count *int64) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}

	query := `
	SELECT COALESCE(COUNT(*), 0) AS total_visits_this_week
	FROM visit_records
	WHERE
		url_id = $1
		AND deleted_at IS NULL
		AND created_at >= date_trunc('week', CURRENT_TIMESTAMP)  -- Monday start
		AND created_at <  date_trunc('week', CURRENT_TIMESTAMP) + INTERVAL '7 days';
	`

	if err := driver.QueryRowContext(ctx, query, urlID).Scan(
		count,
	); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (vrr *VisitRecordRepoImpl) GetThisDayOfWeekCount(ctx context.Context, urlID int64, weekCounts *[]DayOfWeekCount) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	query := `
	WITH week_days AS (
		SELECT generate_series(
			date_trunc('week', CURRENT_TIMESTAMP)::date,
			date_trunc('week', CURRENT_TIMESTAMP)::date + 6,
			'1 day'::interval
		) AS day
	)
	SELECT
		TO_CHAR(wd.day, 'Day')   AS day_of_week,
		COUNT(vr.id)             AS visit_count
	FROM week_days wd
	LEFT JOIN visit_records vr
		ON vr.created_at::date = wd.day
		AND vr.url_id = $1
		AND vr.deleted_at IS NULL
	GROUP BY wd.day
	ORDER BY wd.day;
	`
	rows, err := driver.QueryContext(ctx, query, urlID)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer rows.Close()
	for rows.Next() {
		var item DayOfWeekCount
		if err := rows.Scan(
			&item.DayOfWeek,
			&item.Count,
		); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.DatabaseExecutionErr,
			)
		}
		*weekCounts = append(*weekCounts, item)
	}
	if err := rows.Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (vrr *VisitRecordRepoImpl) GetTodayCountGroupByDevice(ctx context.Context, urlID int64, count *[]DeviceCount) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}

	query := `
	SELECT
		device,
		COALESCE(COUNT(id), 0)
	FROM visit_records
	GROUP BY device
	HAVING url_id = $1
	AND YEAR(created_at) = YEAR(NOW())
	AND MONTH(created_at) = MONTH(NOW())
	AND DAY(created_at) = DAY(NOW())
	AND deleted_at IS NULL
	`

	rows, err := driver.QueryContext(ctx, query, urlID)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer rows.Close()

	for rows.Next() {
		var item DeviceCount
		if err := rows.Scan(
			&item.Device,
			&item.Count,
		); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.DatabaseExecutionErr,
			)
		}
		*count = append(*count, item)
	}

	if err := rows.Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}

	return nil
}

func (vrr *VisitRecordRepoImpl) GetTodayCount(ctx context.Context, urlID int64, count *int64) error {
	driver := vrr.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}

	query := `
	SELECT
		COALESCE(COUNT(id),0)
	FROM visit_records
	WHERE url_id = $1
	AND YEAR(created_at) = YEAR(NOW())
	AND MONTH(created_at) = MONTH(NOW())
	AND DAY(created_at) = DAY(NOW())
	AND deleted_at IS NULL
	`
	if err := driver.QueryRowContext(ctx, query, urlID).Scan(count); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (vrr *VisitRecordRepoImpl) InsertNewRecord(
	ctx context.Context,
	userID string,
	id int64,
	device string,
) error {
	query := `
	INSERT INTO visit_records (user_id, url_id, device)
	VALUES ($1, $2, $3)
	`

	_, err := vrr.handle.ExecContext(ctx, query, userID, id, device)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}
