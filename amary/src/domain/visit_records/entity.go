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
)
