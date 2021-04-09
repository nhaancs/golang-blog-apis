package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/modules/productcategory/productcategorymodel"
	"nhaancs/modules/productcategory/productcategorytransport/ginproductcategory"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// todo: new package for database
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&productcategorymodel.ProductCategory{}); err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db)

	productCategories := r.Group("/product-categories")
	{
		productCategories.POST("", ginproductcategory.CreateProductCategory(appCtx))
	}

	return r.Run()
}