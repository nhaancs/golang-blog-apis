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


func (data *Category) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCategory)
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "category name can't be blank", "ErrNameCannotBeEmpty")
	ErrNameIsTooLong = common.NewCustomError(nil, "category name is too long", "ErrNameIsTooLong")
	ErrSlugCannotBeEmpty = common.NewCustomError(nil, "slug can't be blank", "ErrNameCannotBeEmpty")
	ErrSlugIsTooLong = common.NewCustomError(nil, "slug is too long", "ErrNameIsTooLong")
	ErrSlugIsInvalid = common.NewCustomError(nil, "slug is invalid", "ErrSlugIsInvalid")
)