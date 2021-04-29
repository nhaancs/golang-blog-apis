package component

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *sqlx.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}
