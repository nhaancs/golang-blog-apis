package poststore

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
	// userStore UserStore // can get url here
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
