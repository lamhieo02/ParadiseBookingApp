package cache

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type LikePostReviewSto interface {
	CountLikeByPostReview(ctx context.Context, postReviewID int) (*int64, error)
}

type likePostReviewStoCache struct {
	store      LikePostReviewSto // mysql
	cacheStore Cache             // redis
}

func NewLikePostReviewStoCache(store LikePostReviewSto, cacheStore Cache) *likePostReviewStoCache {
	return &likePostReviewStoCache{store: store, cacheStore: cacheStore}
}

func (c *likePostReviewStoCache) CountLikeByPostReview(ctx context.Context, postReviewID int) (*int64, error) {
	likePostReview := entities.LikePostReview{}
	likePostReview.PostReviewId = postReviewID
	key := likePostReview.CacheKeyNumLikePostReview()
	var likeCount *int64

	err := c.cacheStore.Get(ctx, key, &likeCount) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if likeCount != nil {
		return likeCount, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.CountLikeByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24*5); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
