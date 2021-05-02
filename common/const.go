package common

const (
	DbTypeCategory = 1
	DbTypeUser     = 2
	DbTypePost     = 3
	DbTypeComment  = 4
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
