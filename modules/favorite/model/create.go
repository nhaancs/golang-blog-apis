package favoritemodel

import (
	"nhaancs/common"
	"time"
)

type FavoriteCreate struct {
	CreatedAt  *time.Time  `json:"created_at" gorm:"column:created_at;"`
	UserId     int         `json:"-" gorm:"column:user_id;"`
	PostId     int         `json:"-" gorm:"column:user_id;"`
	FakePostId *common.UID `json:"post_id" gorm:"-"`
}

func (p *FavoriteCreate) Fulfill() {
	if p.FakePostId != nil {
		p.PostId = int(p.FakePostId.GetLocalID())
	}
}

func (FavoriteCreate) TableName() string {
	return Favorite{}.TableName()
}

func (res *FavoriteCreate) Validate() error {
	if res.FakePostId == nil {
		return ErrFavoritePostIsMissing
	}

	return nil
}

func (data *FavoriteCreate) Mask(isAdmin bool) {

}
