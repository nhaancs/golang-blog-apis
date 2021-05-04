package postmodel

import (
	"nhaancs/common"
	"time"
)

const EntityName = "Post"

// todo: how to get user and category object, what if separate as microservice
type Post struct {
	common.SQLModel `json:",inline"`
	Title           string        `json:"title" gorm:"column:title;"`
	Slug            string        `json:"slug" gorm:"column:slug;"`
	ShortDesc       string        `json:"short_desc" gorm:"column:short_desc;"`
	Body            string        `json:"body" gorm:"column:body;"`
	Image           *common.Image `json:"image" gorm:"column:image;"`
	PublishedAt     *time.Time    `json:"published_at" gorm:"column:published_at;"`
	Keywords        string        `json:"keywords" gorm:"column:keywords;"`
	CategoryId      int           `json:"-" gorm:"column:category_id;"`
	FakeCategoryId  *common.UID   `json:"category_id" gorm:"-"`
	UserId          int           `json:"-" gorm:"column:user_id;"`
	FakeUserId      *common.UID   `json:"user_id" gorm:"-"`
	FavoriteCount   int           `json:"favorite_count" gorm:"-"`
}

func (data *Post) Mask(isAdmin bool) {
	data.GenUID(common.DbTypePost)
	data.FakeCategoryId = common.NewUID(uint32(data.CategoryId), common.DbTypeCategory, 1)
	data.FakeUserId = common.NewUID(uint32(data.UserId), common.DbTypeUser, 1)
}

func (Post) TableName() string {
	return "posts"
}

var (
	ErrPostTitleCannotBeEmpty     = common.NewCustomError(nil, "post title can't be blank", "ErrPostTitleCannotBeEmpty")
	ErrPostTitleIsTooLong         = common.NewCustomError(nil, "post title is too long", "ErrPostTitleIsTooLong")
	ErrPostSlugCannotBeEmpty      = common.NewCustomError(nil, "slug can't be blank", "ErrPostTitleCannotBeEmpty")
	ErrPostSlugIsTooLong          = common.NewCustomError(nil, "slug is too long", "ErrPostTitleIsTooLong")
	ErrPostSlugIsInvalid          = common.NewCustomError(nil, "slug is invalid", "ErrPostSlugIsInvalid")
	ErrPostShortDescCannotBeEmpty = common.NewCustomError(nil, "short description can't be blank", "ErrPostShortDescCannotBeEmpty")
	ErrPostShortDescIsTooLong     = common.NewCustomError(nil, "short description is too long", "ErrPostShortDescIsTooLong")
	ErrPostBodyCannotBeEmpty      = common.NewCustomError(nil, "body can't be blank", "ErrPostBodyCannotBeEmpty")
	ErrPostBodyIsTooLong          = common.NewCustomError(nil, "body is too long", "ErrPostBodyIsTooLong")
	ErrPostKeywordsIsTooLong      = common.NewCustomError(nil, "keywords is too long", "ErrPostKeywordsIsTooLong")
	ErrPostPublishAtCannotBeEmpty = common.NewCustomError(nil, "Published date can't be blank", "ErrPostPublishAtCannotBeEmpty")
	ErrPostImageCannotBeEmpty     = common.NewCustomError(nil, "Image can't be blank", "ErrPostImageCannotBeEmpty")
	ErrPostCategoryCannotBeEmpty  = common.NewCustomError(nil, "category can't be blank", "ErrPostCategoryCannotBeEmpty")
	// ErrPostCategoryIsInvalid      = common.NewCustomError(nil, "category is invalid", "ErrPostCategoryIsInvalid")
	// ErrPostAuthorIsInvalid        = common.NewCustomError(nil, "post author is invalid", "ErrPostAuthorIsInvalid")
)
