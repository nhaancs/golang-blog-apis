package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Get(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*postmodel.Post, error) {
	var result postmodel.Post
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
