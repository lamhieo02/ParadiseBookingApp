package entities

import "time"

type PlaceWishList struct {
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  *time.Time `json:"updated_at" gorm:"column:updated_at"`
	PlaceId    int        `json:"place_id" gorm:"column:place_id"`
	WishListId int        `json:"wishlist_id" gorm:"column:wishlist_id"`
}

func (PlaceWishList) TableName() string {
	return "place_wishlist"
}
