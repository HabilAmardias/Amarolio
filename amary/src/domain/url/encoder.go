package url

import (
	"amary/src/customerror"
	"strconv"

	"github.com/deatil/go-encoding/encoding"
)

type IDEncoder struct{}

func NewIDEncoder() *IDEncoder {
	return &IDEncoder{}
}

func (ide *IDEncoder) Encode(id int64) string {
	idStr := strconv.FormatInt(id, 10)
	return encoding.FromString(idStr).Base62Encode().ToString()
}

func (ide *IDEncoder) Decode(encodedID string) (int64, error) {
	idStr := encoding.FromString(encodedID).Base62Decode().ToString()
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return id, customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	return id, nil
}

func (ide *IDEncoder) DecodeMultipleIDs(ids []string) ([]int64, error) {
	res := []int64{}
	for _, eid := range ids {
		idStr := encoding.FromString(eid).Base62Decode().ToString()
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, customerror.NewError(
				"something went wrong",
				err,
				customerror.CommonErr,
			)
		}
		res = append(res, id)
	}
	return res, nil
}
