package favoritestore

import "context"

func (s *sqlStore) GetFavoriteCountsOfPosts(
	ctx context.Context,
	postIds []int,
) (map[int]int, error) {
	return nil, nil
}