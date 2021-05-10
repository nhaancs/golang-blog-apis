package postbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

type ListRepo interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		filter *postmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type listBiz struct {
	repo ListRepo
}

func NewListBiz(repo ListRepo) *listBiz {
	return &listBiz{repo: repo}
}

func (biz *listBiz) List(
	ctx context.Context,
	filter *postmodel.Filter,
	paging *common.Paging,
	isAdmin bool,
) ([]postmodel.Post, error) {
	conditions := map[string]interface{}{}
	if !isAdmin {
		conditions["is_enabled"] = true
	}
	result, err := biz.repo.List(ctx, conditions, filter, paging, "User", "Category")
	if err != nil {
		return nil, err
	}
	return result, nil
}
