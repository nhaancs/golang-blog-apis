package favoritestore

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

func (s *sqlStore) Delete(
	ctx context.Context,
	userId int,
	postId int,
) error {
	db := s.db
	err := db.Table(favoritemodel.Favorite{}.TableName()).
		Where("user_id = ? ", userId).
		Where("post_id = ? ", postId).
		Delete(nil).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
