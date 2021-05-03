package categorymodel

import (
	"nhaancs/common"
)

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Slug            string `json:"slug" gorm:"column:slug;"`
	Desc            string `json:"desc" gorm:"column:desc;"`
}

func (Category) TableName() string {
	return "categories"
}

func (data *Category) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCategory)
}

var (
	ErrCategoryNameCannotBeEmpty = common.NewCustomError(nil, "category name can't be blank", "ErrCategoryNameCannotBeEmpty")
	ErrCategoryNameIsTooLong     = common.NewCustomError(nil, "category name is too long", "ErrCategoryNameIsTooLong")
	ErrCategorySlugCannotBeEmpty = common.NewCustomError(nil, "slug can't be blank", "ErrCategoryNameCannotBeEmpty")
	ErrCategorySlugIsTooLong     = common.NewCustomError(nil, "slug is too long", "ErrCategoryNameIsTooLong")
	ErrCategorySlugIsInvalid     = common.NewCustomError(nil, "slug is invalid", "ErrCategorySlugIsInvalid")
)
