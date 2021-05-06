package categorymodel

import (
	"nhaancs/common"
	"strings"

	"github.com/gosimple/slug"
)

type CategoryCreate struct {
	common.SQLCreateModel `json:",inline"`
	Name                  string `json:"name" gorm:"column:name;"`
	Slug                  string `json:"slug" gorm:"column:slug;"`
	Desc                  string `json:"desc" gorm:"column:desc;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (res *CategoryCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	res.Slug = strings.TrimSpace(res.Slug)

	if len(res.Name) == 0 {
		return ErrCategoryNameCannotBeEmpty
	}
	if len(res.Name) > 200 {
		return ErrCategoryNameIsTooLong
	}
	if len(res.Slug) == 0 {
		return ErrCategorySlugCannotBeEmpty
	}
	if len(res.Slug) > 255 {
		return ErrCategorySlugIsTooLong
	}
	if !slug.IsSlug(res.Slug) {
		return ErrCategorySlugIsInvalid
	}

	return nil
}

func (data *CategoryCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCategory)
}
