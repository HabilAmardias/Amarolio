package visitrecords

import (
	"amary/src/customerror"
	"amary/src/dto"
	"amary/src/handlers"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VisitRecordServiceItf interface {
	GetVisitRecordSummary(ctx context.Context, userID string, code string) (*VisitDashboard, error)
}

type VisitRecordHandlerImpl struct {
	vrs VisitRecordServiceItf
}

func NewVisitRecordHandler(vrs VisitRecordServiceItf) *VisitRecordHandlerImpl {
	return &VisitRecordHandlerImpl{vrs}
}

func (vrh *VisitRecordHandlerImpl) GetVisitRecordSummary(ctx *gin.Context) {
	code := ctx.Param("id")
	userID := handlers.GetAuthenticationPayload(ctx)
	if len(userID) == 0 {
		ctx.Error(customerror.NewError(
			"unauthorized",
			errors.New("user is not authorized"),
			customerror.Unauthenticate,
		))
		return
	}
	vd, err := vrh.vrs.GetVisitRecordSummary(ctx.Request.Context(), userID, code)
	if err != nil {
		ctx.Error(err)
		return
	}
	var (
		todayDeviceCount       []DeviceCountRes
		thisDayOfWeekCount     []DayOfWeekCountRes
		thisWeekDeviceCount    []DeviceCountRes
		thisWeekDOWDeviceCount []DeviceDayOfWeekCountRes
	)

	for _, el := range vd.TodayDeviceCount {
		todayDeviceCount = append(todayDeviceCount, DeviceCountRes(el))
	}
	for _, el := range vd.ThisDayOfWeekCount {
		thisDayOfWeekCount = append(thisDayOfWeekCount, DayOfWeekCountRes(el))
	}
	for _, el := range vd.ThisWeekDeviceCount {
		thisWeekDeviceCount = append(thisWeekDeviceCount, DeviceCountRes(el))
	}
	for _, el := range vd.ThisWeekDOWDeviceCount {
		thisWeekDOWDeviceCount = append(thisWeekDOWDeviceCount, DeviceDayOfWeekCountRes(el))
	}
	res := VisitDashboardRes{
		TodayVisitCount:        vd.TodayVisitCount,
		ThisWeekCount:          vd.ThisWeekCount,
		TodayDeviceCount:       todayDeviceCount,
		ThisDayOfWeekCount:     thisDayOfWeekCount,
		ThisWeekDeviceCount:    thisWeekDeviceCount,
		ThisWeekDOWDeviceCount: thisWeekDOWDeviceCount,
	}

	ctx.JSON(http.StatusOK, dto.ServerResponse[VisitDashboardRes]{
		Success: true,
		Data:    res,
	})
}
