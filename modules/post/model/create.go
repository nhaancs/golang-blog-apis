package postmodel

import (
	"nhaancs/common"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type PostCreate struct {
	common.SQLCreateModel `json:",inline"`
	Title                 string        `json:"title" gorm:"column:title;"`
	Slug                  string        `json:"slug" gorm:"column:slug;"`
	ShortDesc             string        `json:"short_desc" gorm:"column:short_desc;"`
	Body                  string        `json:"body" gorm:"column:body;"`
	Image                 *common.Image `json:"image" gorm:"column:image;"`
	PublishedAt           *time.Time    `json:"published_at" gorm:"column:published_at;autoCreateTime;"`
	Keywords              string        `json:"keywords" gorm:"column:keywords;"`
	CategoryId            int           `json:"-" gorm:"column:category_id;"`
	UserId                int           `json:"-" gorm:"column:user_id;"`
	FakeCategoryId        *common.UID   `json:"category_id" gorm:"-"`
	FakeUserId            *common.UID   `json:"-" gorm:"-"`
}

func (PostCreate) TableName() string {
	return Post{}.TableName()
}

func (p *PostCreate) Fulfill() {
	if p.FakeCategoryId != nil {
		p.CategoryId = int(p.FakeCategoryId.GetLocalID())
	}

	if p.FakeUserId != nil {
		p.UserId = int(p.FakeUserId.GetLocalID())
	}
}

func (res *PostCreate) Validate() error {
	res.Title = strings.TrimSpace(res.Title)
	res.Slug = strings.TrimSpace(res.Slug)
	res.ShortDesc = strings.TrimSpace(res.ShortDesc)
	res.Body = strings.TrimSpace(res.Body)
	res.Keywords = strings.TrimSpace(res.Keywords)

	if len(res.Title) == 0 {
		return ErrPostTitleCannotBeEmpty
	}
	if len(res.Title) > 200 {
		return ErrPostTitleIsTooLong
	}
	if len(res.Slug) == 0 {
		return ErrPostSlugCannotBeEmpty
	}
	if len(res.Slug) > 255 {
		return ErrPostSlugIsTooLong
	}
	if !slug.IsSlug(res.Slug) {
		return ErrPostSlugIsInvalid
	}
	if len(res.ShortDesc) == 0 {
		return ErrPostShortDescCannotBeEmpty
	}
	if len(res.ShortDesc) > 255 {
		return ErrPostShortDescIsTooLong
	}
	if len(res.Body) == 0 {
		return ErrPostBodyCannotBeEmpty
	}
	if len(res.Body) > 20000 {
		return ErrPostBodyIsTooLong
	}
	if res.PublishedAt == nil {
		return ErrPostPublishAtCannotBeEmpty
	}
	if res.Image == nil {
		return ErrPostImageCannotBeEmpty
	}
	if res.FakeCategoryId == nil {
		return ErrPostCategoryCannotBeEmpty
	}
	if len(res.Keywords) > 255 {
		return ErrPostKeywordsIsTooLong
	}

	return nil
}

func (data *PostCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypePost)
}
