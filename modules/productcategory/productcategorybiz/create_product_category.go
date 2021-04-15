package productcategorybiz

import (
	"context"
	"nhaancs/common"
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
		return common.ErrInvalidRequest(err)
	}

	data.Slug = slugify.Slugify(data.Name)
	{
		_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug});
		if err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(productcategorymodel.EntityName, nil)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return  common.ErrCannotCreateEntity(productcategorymodel.EntityName, err)
	}

	return nil
}
