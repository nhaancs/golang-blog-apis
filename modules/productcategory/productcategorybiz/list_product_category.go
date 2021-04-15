package productcategorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

type ListProductCategoryStore interface {
	ListDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *productcategorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productcategorymodel.ProductCategory, error)
}

type listProductCategoryBiz struct {
	store ListProductCategoryStore
}

func NewListProductCategoryBiz(store ListProductCategoryStore) *listProductCategoryBiz {
	return &listProductCategoryBiz{store: store}
}

func (biz *listProductCategoryBiz) ListProductCategory(
	ctx context.Context,
	filter *productcategorymodel.Filter,
	paging *common.Paging,
) ([]productcategorymodel.ProductCategory, error) {
	res, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(productcategorymodel.EntityName, err)
	}

	return res, nil
}

