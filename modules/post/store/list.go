package poststore

import (
	"context"
	"nhaancs/common"
	"nhaancs/modules/post/model"
)

func (s *sqlStore) List(ctx context.Context,
	conditions map[string]interface{},
	filter *postmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	var result []postmodel.Post
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(postmodel.Post{}.TableName()).
		Where(conditions).
		Where("deleted_at IS NULL")

	if v := filter; v != nil {
		if userId, err := common.FromBase58(v.UserId); len(v.UserId) > 0 && err != nil {
			db = db.Where("user_id = ?", userId.GetLocalID())
		}
		if categoryId, err := common.FromBase58(v.CategoryId); len(v.CategoryId) > 0 && err != nil {
			db = db.Where("category_id = ?", categoryId.GetLocalID())
		}
	}

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
