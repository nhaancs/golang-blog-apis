package productcategorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/productcategory/productcategorymodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *productcategorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]productcategorymodel.ProductCategory, error) {
	var result []productcategorymodel.ProductCategory
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(productcategorymodel.ProductCategory{}.TableName()).Where(conditions)

	// custom filters here
	// if v := filter; v != nil {
		// if v.CityId > 0 {
		// 	db = db.Where("city_id = ?", v.CityId)
		// }
	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	//todo: implement order
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}