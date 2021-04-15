package productcategorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productcategorymodel.ProductCategory, error) {
	var result productcategorymodel.ProductCategory
	db := s.db

	for i := range moreKeys {
		// todo: can have error here .Error
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
