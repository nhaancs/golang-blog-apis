package uploadstore

import (
	"context"
	"nhaancs/common"
	uploadmodel "nhaancs/modules/upload/model"
)

func (store *sqlStore) CreateImage(context context.Context, data *uploadmodel.UploadedImage) error {
	db := store.db
	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
