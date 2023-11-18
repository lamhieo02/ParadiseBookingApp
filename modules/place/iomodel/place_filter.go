package iomodel

type Filter struct {
	VendorID *int    `json:"vendor_id" form:"vendor_id"`
	Country  *string `json:"country" form:"country"`
	State    *string `json:"state" form:"state"`
	City     *string `json:"city" form:"city"`
}
