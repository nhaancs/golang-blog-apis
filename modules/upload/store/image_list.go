package uploadstore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/upload/model"
)

func (s *sqlStore) List(ctx context.Context,
	conditions map[string]interface{},
	filter *uploadmodel.ImageFilter,
	paging *common.Paging,
	moreKeys ...string,
) ([]uploadmodel.Image, error) {
	var result []uploadmodel.Image
	db := s.db

	db = db.Table(uploadmodel.Image{}.TableName()).
		Where(conditions).
		Where("deleted_at IS NULL")

	// if v := filter; v != nil {
	// 	if v.CityId > 0 {
	// 		db = db.Where("city_id = ?", v.CityId)
	// 	}
	// }

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
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
