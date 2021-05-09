package favoritestore

import (
	"context"
	"fmt"
	"nhaancs/common"
	favoritemodel "nhaancs/modules/favorite/model"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

const timeLayout = "2006-01-02T15:04:05.999999"

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
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))
		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i, item := range result {
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return result, nil
}
