package uploadstore

import (
	"context"
	uploadmodel "nhaancs/modules/upload/model"
	"time"
)

func (store *sqlStore) SoftDeleteImages(ctx context.Context, ids []int) error {
	db := store.db
	if err := db.Table(uploadmodel.Image{}.TableName()).
		Where("id in (?)", ids).
		Updates(map[string]interface{}{"deleted_at": time.Now()}).Error; err != nil {
		return err
	}

	return nil
}

func (store *sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db
	if err := db.Table(uploadmodel.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
