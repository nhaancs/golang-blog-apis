package productstore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

func (s *sqlStore) ListDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *productmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]productmodel.Product, error) {
	var result []productmodel.Product
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(productmodel.Product{}.TableName()).
		Where(conditions).
		Where("deleted_at is null")

	// custom filters here
	// if v := filter; v != nil {
		// if v.CityId > 0 {
		// 	db = db.Where("city_id = ?", v.CityId)
		// }
	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//todo: implement order
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}