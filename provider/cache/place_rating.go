package cache

import (
	"context"
	"fmt"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type PlaceSto interface {
	GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error)
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type placeStoCache struct {
	store      PlaceSto // mysql
	cacheStore Cache    // redis
}

func NewPlaceStoCache(store PlaceSto, cacheStore Cache) *placeStoCache {
	return &placeStoCache{store: store, cacheStore: cacheStore}
}

func (c *placeStoCache) GetPlaceByID(ctx context.Context, placeId int) (*entities.Place, error) {
	place := entities.Place{}

	key := "place:" + fmt.Sprintf("%d", placeId)

	err := c.cacheStore.Get(ctx, key, &place) // get data from redis
	if err != nil {
		fmt.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if place.Id != 0 {
		return &place, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetPlaceByID(ctx, placeId)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24*5); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}

func (c *placeStoCache) GetRatingAverageByPlaceId(ctx context.Context, placeId int64) (*float64, error) {
	place := entities.Place{}
	place.Id = int(placeId)

	key := place.CacheKeyPlaceRating()
	var ratingAverage *float64

	err := c.cacheStore.Get(ctx, key, &ratingAverage) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if ratingAverage != nil {
		return ratingAverage, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetRatingAverageByPlaceId(ctx, placeId)
	if err != nil {
		return nil, err
	}

	if u == nil {
		defaulRating := 0.0
		u = &defaulRating
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
