package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/component/uploadprovider"
	"nhaancs/middleware"
	"nhaancs/modules/category/categorytransport/gincategory"
	"nhaancs/modules/post/posttransport/ginpost"
	"nhaancs/modules/upload/uploadtransport/ginupload"
	"os"

	"github.com/gin-gonic/gin"
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

	s3Provider := uploadprovider.NewS3Provider(
		os.Getenv("S3_BUCKET_NAME"),
		os.Getenv("S3_REGION"),
		os.Getenv("S3_API_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("S3_DOMAIN"),
	)

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln("Error running service: ", err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")
	v1.POST("/upload", ginupload.Upload(appCtx))

	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.Create(appCtx))
		categories.GET("/:id", gincategory.Get(appCtx))
		categories.GET("", gincategory.List(appCtx))
		categories.PATCH("/:id", gincategory.Update(appCtx))
		categories.DELETE("/:id", gincategory.Delete(appCtx))
	}
	posts := v1.Group("/posts")
	{
		posts.POST("", ginpost.Create(appCtx))
		posts.GET("/:id", ginpost.Get(appCtx))
		posts.GET("", ginpost.List(appCtx))
		posts.PATCH("/:id", ginpost.Update(appCtx))
		posts.DELETE("/:id", ginpost.Delete(appCtx))
	}

	return r.Run()
}
