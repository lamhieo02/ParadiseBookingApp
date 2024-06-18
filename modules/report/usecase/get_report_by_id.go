package reportusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/account/convert"
	reportconvert "paradise-booking/modules/report/convert"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) GetReportByID(ctx context.Context, id int) (*reportiomodel.GetReportResp, error) {
	report, err := uc.reportSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	account, err := uc.accountCache.GetProfileByID(ctx, report.UserID)
	if err != nil {
		return nil, err
	}

	result := reportconvert.ReportEntityToModel(report)

	uc.getDataUser(result, account)

	uc.getObjectValue(ctx, result)
	return result, nil
}

func (uc *reportUseCase) getObjectValue(ctx context.Context, reportData *reportiomodel.GetReportResp) {
	var objectVL interface{}
	if reportData.ObjectType == constant.Report_Object_Type_Place {
		place, err := uc.placeCache.GetPlaceByID(ctx, reportData.ObjectID)
		if err != nil {
			return
		}

		objectValue := reportiomodel.ObjectValue{}
		objectValue.Title = place.Name
		objectValue.Address = place.Address
		objectValue.Cover = place.Cover
		objectValue.Description = place.Description

		objectVL = objectValue

	} else if reportData.ObjectType == constant.Report_Object_Type_Tour {
		postGuide, err := uc.postGuideCache.GetByID(ctx, reportData.ObjectID)
		if err != nil {
			return
		}

		objectValue := reportiomodel.ObjectValue{}
		objectValue.Title = postGuide.Title
		objectValue.Address = postGuide.Address
		objectValue.Cover = postGuide.Cover
		objectValue.Description = postGuide.Description

		objectVL = objectValue

	} else if reportData.ObjectType == constant.Report_Object_Type_Post_Review {
		postReview, err := uc.postReviewSto.GetByID(ctx, reportData.ObjectID)
		if err != nil {
			return
		}

		objectValue := reportiomodel.ObjectValue{}
		objectValue.Title = postReview.Title
		objectValue.Description = "TOPIC" + constant.MapCategoryIDToName[postReview.Topic]
		objectValue.Address = postReview.Country + " - " + postReview.State + " - " + postReview.District
		objectValue.Cover = postReview.Image

		objectVL = objectValue
	} else if reportData.ObjectType == constant.Report_Object_Type_Comment {
		comment, err := uc.commentSto.GetByID(ctx, reportData.ObjectID)
		if err != nil {
			return
		}

		objectVL = comment
	} else if reportData.ObjectType == constant.Report_Object_Type_Guider || reportData.ObjectType == constant.Report_Object_Type_User ||
		reportData.ObjectType == constant.Report_Object_Type_Vendor {
		account, err := uc.accountCache.GetProfileByID(ctx, reportData.ObjectID)
		if err != nil {
			return
		}

		objectVL = convert.ConvertAccountEntityToInfoResp(account)
	}

	reportData.ObjectValue = objectVL
}

func (uc *reportUseCase) getDataUser(reportData *reportiomodel.GetReportResp, account *entities.Account) {
	reportData.User.ID = account.Id
	reportData.User.Role = constant.MapRole[constant.Role(account.Role)]
	reportData.User.Username = account.Username
	reportData.User.FullName = account.FullName
	reportData.User.Email = account.Email
	reportData.User.Phone = account.Phone
	reportData.User.Address = account.Address
	reportData.User.DOB = account.Dob
	reportData.User.Avt = account.Avatar
	reportData.User.Bio = account.Bio
	reportData.User.Created = account.CreatedAt
	reportData.User.Updated = account.UpdatedAt
}
