package favoritestore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/favorite/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Get(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*favoritemodel.Favorite, error) {
	var result favoritemodel.Favorite
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
