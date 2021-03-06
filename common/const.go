package common

const (
	DbTypeCategory = 1
	DbTypeUser     = 2
	DbTypePost     = 3
	DbTypeComment  = 4
	DbTypeImage    = 5
)

const CurrentUser = "user"
const AdminRole = "admin"
const UserRole = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
