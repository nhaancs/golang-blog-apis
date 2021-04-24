package productbiz

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"

	"github.com/Machiel/slugify"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error

	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(
	ctx context.Context, 
	data *productmodel.ProductCreate,
) error {
	if err := data.ValidateCreate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	data.Slug = slugify.Slugify(data.Name)
	{
		_, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"slug": data.Slug});
		if err != common.ErrRecordNotFound {
			return common.ErrEntityExisted(productmodel.EntityName, nil)
		}
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return  common.ErrCannotCreateEntity(productmodel.EntityName, err)
	}

	return nil
}
