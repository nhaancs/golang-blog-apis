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
	if err := data.ValidateCreate(); err != nil {
		return err
	}

	data.Slug = slugify.Slugify(data.Name)
	{
		res, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug});

		//todo: check err instead
		// if res != nil {
		// 	return errors.New("the product category already exists")
		// }
	}

	return biz.store.Create(ctx, data)
}
