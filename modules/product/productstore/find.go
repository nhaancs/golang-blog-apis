package productstore

import (
	"context"
	"database/sql"
	"fmt"
	"nhaancs/common"
	"nhaancs/modules/product/productmodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productmodel.Product, error) {
	var result productmodel.Product
	db := s.db

	// todo: implement below code for sqlx
	// for i := range moreKeys {
	// 	// todo: can have error here .Error
	// 	db = db.Preload(moreKeys[i])
	// }

	for key, value := range conditions {
		
	}

	query := fmt.Sprintf("SELECT * FROM %s", productmodel.TableName)
	if err := db.Get(&result, query); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
