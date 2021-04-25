package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/middleware"
	"nhaancs/modules/product/producttransport/ginproduct"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	db, err := sqlx.Connect("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalln("Error connecting database: ", err)
	}

	if err := doMigrations(db); err != nil {
		log.Fatalln("Error doing migrations: ", err)
	}

	if err := runService(db); err != nil {
		log.Fatalln("Error running service: ", err)
	}
}

func doMigrations(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", "mysql", driver)
	if err != nil {
		return err
	}
	err = m.Steps(1)
	if err != nil {
		return err
	}

	return nil
}

func runService(db *sqlx.DB) error {
	appCtx := component.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	products := r.Group("/products")
	{
		products.GET("", ginproduct.ListProduct(appCtx))
		products.POST("", ginproduct.CreateProduct(appCtx))
		products.PATCH("/:id", ginproduct.UpdateProduct(appCtx))
		products.DELETE("/:id", ginproduct.DeleteProduct(appCtx))
		products.GET("/:slug", ginproduct.GetProductBySlug(appCtx))
	}

	return r.Run()
}
