package favoritemodel

import (
	"nhaancs/common"
	postmodel "nhaancs/modules/post/model"
	"time"
)

const EntityName = "Favorite"

type Favorite struct {
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	UserId    int                `json:"-" gorm:"column:user_id;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
	PostId    int                `json:"-" gorm:"column:user_id;"`
	Post      *postmodel.Post    `json:"post" gorm:"preload:false;"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func (data *Favorite) Mask(isAdmin bool) {
	if u := data.User; u != nil {
		u.Mask(isAdmin)
	}

	if p := data.User; p != nil {
		p.Mask(isAdmin)
	}
}

var (
	ErrFavoritePostIsMissing = common.NewCustomError(nil, "missing post", "ErrFavoritePostIsMissing")
	ErrFavoritePostIsInvalid = common.NewCustomError(nil, "invalid post", "ErrFavoritePostIsInvalid")
)
