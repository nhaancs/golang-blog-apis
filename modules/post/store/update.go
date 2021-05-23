package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *postmodel.PostUpdate,
) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) IncreaseFavoriteCount(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
		// use expression you dont need to query for old data, and can prevent race condition
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

// todo: what if favorite_count = 0, favorite_count - 1 = -1?
func (s *sqlStore) DecreaseFavoriteCount(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
		// use expression you dont need to query for old data, and can prevent race condition
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
