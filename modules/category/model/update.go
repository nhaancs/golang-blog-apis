package categorymodel

import (
	"nhaancs/common"
	"strings"

	"github.com/gosimple/slug"
)

type CategoryUpdate struct {
	common.SQLUpdateModel `json:",inline"`
	Name                  string  `json:"name" gorm:"column:name;"`
	Slug                  string  `json:"slug" gorm:"column:slug;"`
	Desc                  *string `json:"desc" gorm:"column:desc;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}
// todo: we only need to validate a few column get updated
func (res *CategoryUpdate) Validate() error {
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
	if len(res.Slug) > 200 {
		return ErrCategorySlugIsTooLong
	}
	if !slug.IsSlug(res.Slug) {
		return ErrCategorySlugIsInvalid
	}

	return nil
}
