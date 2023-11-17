package constant

type Role int

const (
	UserRole Role = iota + 1
	VendorRole
	AdminRole
)
