package favoritestore

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

func (s *sqlStore) List(ctx context.Context,
	conditions map[string]interface{},
	filter *favoritemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]favoritemodel.Favorite, error) {
	var result []favoritemodel.Favorite
	db := s.db

	db = db.Table(favoritemodel.Favorite{}.TableName()).
		Where(conditions)

	if v := filter; v != nil {
		if v.UserId > 0 {
			db = db.Where("user_id = ?", v.UserId)
		}
		if v.PostId > 0 {
			db = db.Where("post_id = ?", v.PostId)
		}
	}

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
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
