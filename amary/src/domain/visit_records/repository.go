package visitrecords

import (
	"amary/src/customerror"
	"amary/src/db"
	"context"
)

type VisitRecordRepoImpl struct {
	handle db.DBTX
}

func NewVisitRecordRepo(handle db.DBTX) *VisitRecordRepoImpl {
	return &VisitRecordRepoImpl{handle}
}

func (vrr *VisitRecordRepoImpl) InsertNewRecord(
	ctx context.Context,
	userID string,
	id int64,
	device string,
	vr *VisitRecord,
) error {
	query := `
	INSERT INTO visit_records (user_id, url_id, device)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, url_id, device, created_at, updated_at, deleted_at
	`

	if err := vrr.handle.QueryRowContext(ctx, query, userID, id, device).Scan(
		&vr.ID,
		&vr.UserID,
		&vr.URLID,
		&vr.Device,
		&vr.CreatedAt,
		&vr.UpdatedAt,
		&vr.DeletedAt,
	); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}
