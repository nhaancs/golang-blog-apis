package categorystore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/category/categorymodel"
)

func (s *sqlStore) List(ctx context.Context,
	conditions map[string]interface{},
	filter *categorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(categorymodel.Category{}.TableName()).Where(conditions).Where("status in (1)")

	// if v := filter; v != nil {
	// 	if v.CityId > 0 {
	// 		db = db.Where("city_id = ?", v.CityId)
	// 	}
	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
