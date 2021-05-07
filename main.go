package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/component/uploadprovider"
	"nhaancs/middleware"
	gincategory "nhaancs/modules/category/transport/gin"
	ginfavorite "nhaancs/modules/favorite/transport/gin"
	ginpost "nhaancs/modules/post/transport/gin"
	ginupload "nhaancs/modules/upload/transport/gin"
	ginuser "nhaancs/modules/user/transport/gin"
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
	secretKey := os.Getenv("AUTH_SECRET")

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln("Error running service: ", err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.RequiredAuthOrNot(appCtx))

	v1 := r.Group("v1")
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))
	// v1.GET("/liked-posts", middleware.RequiredAuth(appCtx), middleware.RequiredUser(appCtx), ginfavorite.List(appCtx)) // todo: implement
	v1.POST("/upload-image", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), ginupload.UploadImage(appCtx))
	categories := v1.Group("/categories")
	{
		categories.POST("", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), gincategory.Create(appCtx))
		categories.GET("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), gincategory.Get(appCtx))
		categories.GET("/slug/:slug", gincategory.Get(appCtx))
		categories.GET("", gincategory.List(appCtx))
		categories.PATCH("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), gincategory.Update(appCtx))
		categories.DELETE("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), gincategory.Delete(appCtx))
	}
	posts := v1.Group("/posts")
	{
		posts.POST("", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), ginpost.Create(appCtx))
		posts.GET("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), ginpost.Get(appCtx))
		posts.GET("/slug/:slug", ginpost.Get(appCtx))
		posts.GET("", ginpost.List(appCtx))
		posts.PATCH("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), ginpost.Update(appCtx))
		posts.DELETE("/:id", middleware.RequiredAuth(appCtx), middleware.RequiredAdmin(appCtx), ginpost.Delete(appCtx))
		posts.POST("/:id/favorite", middleware.RequiredAuth(appCtx), middleware.RequiredUser(appCtx), ginfavorite.Favorite(appCtx))
		posts.DELETE("/:id/unfavorite", middleware.RequiredAuth(appCtx), middleware.RequiredUser(appCtx), ginfavorite.Unfavorite(appCtx))
		// posts.GET("/:id/liked-users", middleware.RequiredAuth(appCtx), middleware.RequiredUser(appCtx), ginfavorite.List(appCtx)) // todo: implement
	}

	return r.Run()
}
