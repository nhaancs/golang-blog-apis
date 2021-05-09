package favoritemodel

import (
	"time"
)

type FavoriteCreate struct {
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UserId    int        `json:"-" gorm:"column:user_id;"`
	PostId    int        `json:"-" gorm:"column:post_id;"`
}

func (FavoriteCreate) TableName() string {
	return Favorite{}.TableName()
}

func (res *FavoriteCreate) Validate() error {
	if res.PostId == 0 {
		return ErrFavoritePostIsMissing
	}

	return nil
}

func (data *FavoriteCreate) Mask(isAdmin bool) {

}
