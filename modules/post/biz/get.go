package postbiz

import (
	"context"
	postmodel "nhaancs/modules/post/model"
)

type GetRepo interface {
	Get(ctx context.Context, conditions map[string]interface{}, isAdmin bool) (*postmodel.Post, error)
}

type getBiz struct {
	repo GetRepo
}

func NewGetBiz(repo GetRepo) *getBiz {
	return &getBiz{repo: repo}
}

func (biz *getBiz) Get(ctx context.Context, conditions map[string]interface{}, isAdmin bool) (*postmodel.Post, error) {
	if conditions != nil && !isAdmin {
		conditions["is_enabled"] = true
	}
	data, err := biz.repo.Get(ctx, conditions, isAdmin)
	if err != nil {
		return nil, err
	}

	return data, nil
}
