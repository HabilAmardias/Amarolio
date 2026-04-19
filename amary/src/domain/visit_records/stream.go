package visitrecords

import (
	"amary/src/customerror"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type VisitRecordStreamRepo struct {
	rc *redis.Client
}

func NewVisitRecordStream(rc *redis.Client) *VisitRecordStreamRepo {
	return &VisitRecordStreamRepo{rc}
}

func (vrs *VisitRecordStreamRepo) CreateGroup() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	streamKey := "streams:visit_records"
	groupKey := "visit-record-group"

	err := vrs.rc.XGroupCreate(ctx, streamKey, groupKey, "0").Err()

	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return nil
}

func (vrs *VisitRecordStreamRepo) Get() ([]redis.XStream, error) {
	streamKey := "streams:visit_records"
	groupKey := "visit-record-group"
	consumerName := "consumer-1"

	return vrs.rc.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Group:    groupKey,
		Consumer: consumerName,
		Streams:  []string{streamKey, ">"},
		Count:    10,
		Block:    5 * time.Second,
	}).Result()
}

func (vrs *VisitRecordStreamRepo) Send(ctx context.Context, userID string, urlID int64, device string) error {
	args := &redis.XAddArgs{
		Stream: "streams:visit_records",
		Approx: true,
		Values: map[string]any{
			"action":  "SEND_RECORD",
			"user_id": userID,
			"url_id":  urlID,
			"device":  device,
			"time":    time.Now().Format(time.RFC3339),
		},
	}
	return vrs.rc.XAdd(ctx, args).Err()
}

func (vrs *VisitRecordStreamRepo) parseVisitRecord(values map[string]any) VisitRecordStream {
	getString := func(key string) string {
		if v, ok := values[key].(string); ok {
			return v
		}
		return ""
	}

	getInt64 := func(key string) int64 {
		switch v := values[key].(type) {
		case int64:
			return v
		case string:
			i, _ := strconv.ParseInt(v, 10, 64)
			return i
		case float64: // Sometimes JSON numbers come as float
			return int64(v)
		default:
			return 0
		}
	}

	return VisitRecordStream{
		Action: getString("action"),
		UserID: getString("user_id"),
		URLID:  getInt64("url_id"),
		Device: getString("device"),
		Time:   getString("time"),
	}
}
