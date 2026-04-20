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
