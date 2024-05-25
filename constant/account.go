package constant

type Role int

const (
	UserRole Role = iota + 1
	VendorRole
	AdminRole
	GuiderRole
)

var MapRole = map[Role]string{
	UserRole:   "user",
	VendorRole: "vendor",
	AdminRole:  "admin",
	GuiderRole: "guider",
}

const (
	TypeVerifyEmail   = 1
	TypeResetPassword = 2
)

const (
	StatusActive   = 2
	StatusInactive = 1
)

const (
	LoginWithGoogle = 2
)
