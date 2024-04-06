package postreviewstorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (s *postReviewStorage) ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error) {
	db := s.db

	var data []*entities.PostReview

	db = db.Table(entities.PostReview{}.TableName()).Where("post_owner_id = ?", accountID)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
