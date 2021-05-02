package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/postmodel"
	"time"
)

func (s *sqlStore) SoftDelete(
	ctx context.Context,
	id int,
) error {
	db := s.db
	err := db.Table(postmodel.Post{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error
	if err != nil {
		return common.ErrDB(err)
	}

	return nil
}
