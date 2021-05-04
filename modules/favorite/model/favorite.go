package favoritemodel

import (
	"nhaancs/common"
	"time"
)

const EntityName = "Favorite"

type Favorite struct {
	CreatedAt  *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UserId     int         `json:"-" gorm:"column:user_id;"`
	FakeUserId *common.UID `json:"user_id" gorm:"-"`
	PostId     int         `json:"-" gorm:"column:user_id;"`
	FakePostId *common.UID `json:"post_id" gorm:"-"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func (data *Favorite) Mask(isAdmin bool) {
	data.FakeUserId = common.NewUID(uint32(data.UserId), common.DbTypeUser, 1)
	data.FakePostId = common.NewUID(uint32(data.PostId), common.DbTypePost, 1)
}

var (
	ErrFavoritePostIsMissing = common.NewCustomError(nil, "missing post", "ErrFavoritePostIsMissing")
	ErrFavoritePostIsInvalid = common.NewCustomError(nil, "invalid post", "ErrFavoritePostIsInvalid")
)
