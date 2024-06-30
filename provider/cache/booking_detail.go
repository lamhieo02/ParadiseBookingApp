package cache

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/entities"
	"time"
)

type bookingDetailSto interface {
	GetByBookingID(ctx context.Context, bookingId int) (res *entities.BookingDetail, err error)
}

type bookingDetailStoCache struct {
	store      bookingDetailSto // mysql
	cacheStore Cache            // redis
}

func NewBookingDetailStoCache(store bookingDetailSto, cacheStore Cache) *bookingDetailStoCache {
	return &bookingDetailStoCache{store: store, cacheStore: cacheStore}
}

func (c *bookingDetailStoCache) GetByBookingID(ctx context.Context, bookingId int) (*entities.BookingDetail, error) {
	data := entities.BookingDetail{}
	data.BookingId = bookingId

	key := data.CacheKey()
	var bookingDetail *entities.BookingDetail

	err := c.cacheStore.Get(ctx, key, &bookingDetail) // get data from redis
	if err != nil {
		log.Printf("Error when cache.Get() data: %v", err)
	}

	// if data is found in cache, then return the data
	if bookingDetail != nil {
		return bookingDetail, nil
	}

	// if data is not found in cache, then query in real database to find data
	u, err := c.store.GetByBookingID(ctx, bookingId)
	if err != nil {
		return nil, err
	}

	// save data to cache
	if err := c.cacheStore.Set(ctx, key, &u, time.Hour*24*5); err != nil {
		panic(common.NewCustomError(err, "Error when cache.Set() data"))
	}
	return u, err
}
