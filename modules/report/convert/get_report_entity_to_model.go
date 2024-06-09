package reportconvert

import (
	"paradise-booking/constant"
	"paradise-booking/entities"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"strings"
)

func ReportEntityToModel(data *entities.Report) *reportiomodel.GetReportResp {
	return &reportiomodel.GetReportResp{
		ID:          data.Id,
		ObjectID:    data.ObjectID,
		ObjectType:  data.ObjectType,
		ObjectName:  constant.MapReportObjectType[data.ObjectType],
		Type:        data.Type,
		Description: data.Description,
		StatusID:    data.StatusID,
		StatusName:  constant.MapReportStatus[data.StatusID],
		Videos:      strings.Split(data.Videos, ","),
		Images:      strings.Split(data.Images, ","),
	}
}
