package transactions

import "amarfin/src/entity"

type (
	Transaction struct {
		ID           string
		Amount       uint64
		UserID       string
		Description  string
		SchedulerID  *string
		ScheduledFor *string
		entity.AuditTimestamp
	}
)
