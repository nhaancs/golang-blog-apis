package favoritestore

import (
	"context"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
)

func (s *sqlStore) GetFavoriteCountsOfPosts(
	ctx context.Context,
	postIds []int,
) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		PostId    int `gorm:"column:post_id;"`
		LikeCount int `gorm:"column:like_count;"`
	}

	var listFavorite []sqlData
	if err := s.db.Table(favoritemodel.Favorite{}.TableName()).
		Select("post_id, count(post_id) as like_count").
		Where("post_id in (?)", postIds).
		Group("post_id").Find(&listFavorite).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listFavorite {
		result[item.PostId] = item.LikeCount
	}

	return result, nil
}
