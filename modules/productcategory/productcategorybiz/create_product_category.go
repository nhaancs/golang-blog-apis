package productcategorybiz

import (
	"context"
	"errors"
	"nhaancs/modules/productcategory/productcategorymodel"

	"github.com/Machiel/slugify"
)

type CreateProductCategoryStore interface {
	Create(ctx context.Context, data *productcategorymodel.ProductCategoryCreate) error

	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productcategorymodel.ProductCategory, error)
}

type createProductCategoryBiz struct {
	store CreateProductCategoryStore
}

func NewCreateProductCategoryBiz(store CreateProductCategoryStore) *createProductCategoryBiz {
	return &createProductCategoryBiz{store: store}
}

func (biz *createProductCategoryBiz) CreateProductCategory(
	ctx context.Context, 
	data *productcategorymodel.ProductCategoryCreate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	data.Slug = slugify.Slugify(data.Name)
	{
		res, _ := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug});

		if res != nil {
			return errors.New("the product category already exists")
		}
	}

	return biz.store.Create(ctx, data)
}
