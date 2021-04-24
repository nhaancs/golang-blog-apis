package productbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

type ListProductStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *productmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.Product, error) {
	res, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(productmodel.EntityName, err)
	}

	return res, nil
}

