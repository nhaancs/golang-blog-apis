package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/middleware"
	"nhaancs/modules/category/categorytransport/gincategory"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting database: ", err)
	}

	if err := runService(db); err != nil {
		log.Fatalln("Error running service: ", err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")
	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.Create(appCtx))
		categories.GET("/:id", gincategory.Get(appCtx))
		categories.GET("", gincategory.List(appCtx))
		categories.PATCH("/:id", gincategory.Update(appCtx))
		categories.DELETE("/:id", gincategory.Delete(appCtx))
	}

	return r.Run()
}
