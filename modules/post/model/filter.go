package postmodel

import (
	"nhaancs/common"
)

type Filter struct {
	FakeCategoryId string `json:"category_id,omitempty" form:"category_id"`
	FakeUserId     string `json:"user_id,omitempty" form:"user_id"`
	CategoryId     int    `json:"-"`
	UserId         int    `json:"-"`
}

func (f *Filter) Fulfill() {
	if len(f.FakeCategoryId) > 0 {
		if categoryId, err := common.FromBase58(f.FakeCategoryId); err == nil {
			f.CategoryId = int(categoryId.GetLocalID())
		}
	}
	if len(f.FakeUserId) > 0 {
		if userId, err := common.FromBase58(f.FakeUserId); err == nil {
			f.UserId = int(userId.GetLocalID())
		}
	}
}
