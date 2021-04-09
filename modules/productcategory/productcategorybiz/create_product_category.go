package productcategorybiz

import (
	"context"
	"nhaancs/modules/productcategory/productcategorymodel"
	"github.com/Machiel/slugify"
)

type CreateProductCategoryStore interface {
	Create(ctx context.Context, data *productcategorymodel.ProductCategoryCreate) error
}

type createProductCategoryBiz struct {
	store CreateProductCategoryStore
}

func NewCreateProductCategoryBiz(store CreateProductCategoryStore) *createProductCategoryBiz {
	return &createProductCategoryBiz{store: store}
}

func (biz *createProductCategoryBiz) CreateProductCategory(ctx context.Context, data *productcategorymodel.ProductCategoryCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.Slug = slugify.Slugify(data.Name)

	return biz.store.Create(ctx, data)
}