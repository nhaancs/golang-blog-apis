package productbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

type GetProductStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type getProductBiz struct {
	store GetProductStore
}

func NewGetProductBiz(store GetProductStore) *getProductBiz {
	return &getProductBiz{store: store}
}

func (biz *getProductBiz) GetProductBySlug(ctx context.Context, slug string) (*productmodel.Product, error) {
	data, err := biz.store.FindDataByCondition(ctx, 
		map[string]interface{}{"slug": slug})
	if err != nil {
		return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
	}
	if data.DeletedAt != nil {
		return nil, common.ErrEntityDeleted(productmodel.EntityName, nil)
	}
	
	return data, nil
}