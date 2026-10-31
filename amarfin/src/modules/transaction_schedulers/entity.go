package transactionschedulers

import "amarfin/src/entity"

type (
	TransactionScheduler struct {
		ID          string
		Frequency   string
		DateOfMonth *uint8
		TimeZone    string
		LocalTime   string
		StartAt     string
		EndAt       *string
		Amount      uint64
		UserID      string
		Description string
		NextRunAt   string
		entity.AuditTimestamp
	}
)

func (ts TransactionScheduler) ValidateFrequency() bool {
	return ts.Frequency == "monthly" || ts.Frequency == "daily"
}
