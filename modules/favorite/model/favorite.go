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
	PostId    int                `json:"-" gorm:"column:post_id;"`
	Post      *postmodel.Post    `json:"post" gorm:"preload:false;"`
}

func (Favorite) TableName() string {
	return "favorites"
}

func (data *Favorite) Mask(isAdmin bool) {
	if u := data.User; u != nil {
		u.Mask(isAdmin)
	}

	if p := data.Post; p != nil {
		p.Mask(isAdmin)
	}
}

var (
	// use error variables when you do not have a root error
	ErrFavoritePostIsMissing = common.NewCustomError(nil, "missing post", "ErrFavoritePostIsMissing")
	ErrFavoriteAPostTwice    = common.NewCustomError(nil, "you favorited this post", "ErrFavoriteAPostTwice")
)

// use error function when you want to capture the root error
func ErrFavoritePostIsInvalid(err error) *common.AppError {
	return common.NewCustomError(err, "invalid post", "ErrFavoritePostIsInvalid")
}
