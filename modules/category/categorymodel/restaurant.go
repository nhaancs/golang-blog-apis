package categorymodel

import (
	"nhaancs/common"
	"strings"
)

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
	LikedCount      int            `json:"liked_count" gorm:"-"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (res *CategoryCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "category name can't be blank", "ErrNameCannotBeEmpty")
)

func (data *Category) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeCategory)
}
