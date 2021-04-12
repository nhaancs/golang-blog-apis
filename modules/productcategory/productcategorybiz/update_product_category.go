package productcategorybiz

import (
	"context"
	"nhaancs/modules/productcategory/productcategorymodel"
)

// todo: update category dont update slug

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productcategorymodel.ProductCategory, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *productcategorymodel.ProductCategoryUpdate,
	)
}

type updateProductCategoryBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateProductCategoryBiz {
	return &updateProductCategoryBiz{store: store}
}

func (biz *updateProductCategoryBiz) UpdateProductCategory() {

}
