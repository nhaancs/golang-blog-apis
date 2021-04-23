package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/middleware"
	"nhaancs/modules/productcategory/productcategorytransport/ginproductcategory"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
	  log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	productCategories := r.Group("/product-categories")
	{
		productCategories.GET("", ginproductcategory.ListProductCategory(appCtx))
		productCategories.POST("", ginproductcategory.CreateProductCategory(appCtx))
		productCategories.PATCH("/:id", ginproductcategory.UpdateProductCategory(appCtx))
		productCategories.DELETE("/:id", ginproductcategory.DeleteProductCategory(appCtx))
		productCategories.GET("/:slug", ginproductcategory.GetProductCategoryBySlug(appCtx))
	}

	return r.Run()
}
