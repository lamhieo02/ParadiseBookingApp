package cache

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type PostGuideSto interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type postGuideStoCache struct {
	store      PostGuideSto // mysql
	cacheStore Cache        // redis
}

func NewPostGuideStoCache(store PostGuideSto, cacheStore Cache) *postGuideStoCache {
	return &postGuideStoCache{store: store, cacheStore: cacheStore}
}

func (c *postGuideStoCache) GetByID(ctx context.Context, id int) (*entities.PostGuide, error) {
	data := entities.PostGuide{}
	data.Id = id

	key := data.CacheKey()
	var postGuide *entities.PostGuide

	err := c.cacheStore.Get(ctx, key, &postGuide) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if postGuide != nil {
		return postGuide, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetByID(ctx, int(id))
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24*5); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
