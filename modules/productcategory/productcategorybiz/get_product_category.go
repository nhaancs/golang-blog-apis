package productcategorybiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

type GetProductCategoryStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productcategorymodel.ProductCategory, error)
}

type getProductCategoryBiz struct {
	store GetProductCategoryStore
}

func NewGetProductCategoryBiz(store GetProductCategoryStore) *getProductCategoryBiz {
	return &getProductCategoryBiz{store: store}
}

func (biz *getProductCategoryBiz) GetProductCategoryBySlug(ctx context.Context, slug string) (*productcategorymodel.ProductCategory, error) {
	data, err := biz.store.FindDataByCondition(ctx, 
		map[string]interface{}{"slug": slug})
	if err != nil {
		return nil, common.ErrCannotGetEntity(productcategorymodel.EntityName, err)
	}
	if data.DeletedAt != nil {
		return nil, common.ErrEntityDeleted(productcategorymodel.EntityName, nil)
	}
	
	return data, nil
}