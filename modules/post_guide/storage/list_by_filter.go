package postguidestorage

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func (s *postGuideStorage) ListByFilter(ctx context.Context, paging *common.Paging, filter *postguideiomodel.Filter) ([]*entities.PostGuide, error) {
	db := s.db

	var data []*entities.PostGuide

	// db = db.Table(entities.PostGuide{}.TableName()).Order("id desc")
	db = db.Table(entities.PostGuide{}.TableName()).Order("id desc")

	if filter.PostOwnerId != 0 {
		db = db.Where("post_owner_id = ?", filter.PostOwnerId)
	}

	if filter.TopicID != 0 {
		db = db.Where("topic_id = ?", filter.TopicID)
	}

	if filter.Lat != 0 && filter.Lng != 0 {
		db = db.Where("lat = ? AND lng = ?", filter.Lat, filter.Lng)
	}

	if filter.State != "" {
		db = db.Where("state = ?", filter.State)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
