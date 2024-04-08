package cache

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type CommentSto interface {
	CountCommentByPostReview(ctx context.Context, postReviewID int) (*int64, error)
}

type commentStoCache struct {
	store      CommentSto // mysql
	cacheStore Cache      // redis
}

func NewCommentStoCache(store CommentSto, cacheStore Cache) *commentStoCache {
	return &commentStoCache{store: store, cacheStore: cacheStore}
}

func (c *commentStoCache) CountCommentByPostReview(ctx context.Context, postReviewID int) (*int64, error) {
	comment := entities.Comment{}
	comment.PostReviewID = int64(postReviewID)
	key := comment.CacheKeyNumCommentByPostReview()
	var cmtCount *int64

	err := c.cacheStore.Get(ctx, key, &cmtCount) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if cmtCount != nil {
		return cmtCount, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.CountCommentByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24*5); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
