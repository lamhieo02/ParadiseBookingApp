package reportconvert

import (
	"paradise-booking/constant"
	"paradise-booking/entities"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"strings"
)

func ReportEntityToModel(data *entities.Report) *reportiomodel.GetReportResp {

	res := &reportiomodel.GetReportResp{
		ID:          data.Id,
		ObjectID:    data.ObjectID,
		ObjectType:  data.ObjectType,
		ObjectName:  constant.MapReportObjectType[data.ObjectType],
		Type:        data.Type,
		Description: data.Description,
		StatusID:    data.StatusID,
		StatusName:  constant.MapReportStatus[data.StatusID],
		Images:      []string{},
		Videos:      []string{},
		UserID:      data.UserID,
	}

	if data.Images != "" {
		res.Images = strings.Split(data.Images, ",")
	}

	if data.Videos != "" {
		res.Videos = strings.Split(data.Videos, ",")
	}

	return res
}
