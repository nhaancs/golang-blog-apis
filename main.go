package main

import (
	"fmt"
	"log"
	"nhaancs/component"
	"nhaancs/component/uploadprovider"
	"nhaancs/middleware"
	gincategory "nhaancs/modules/category/transport/gin"
	ginfavorite "nhaancs/modules/favorite/transport/gin"
	ginpost "nhaancs/modules/post/transport/gin"
	ginupload "nhaancs/modules/upload/transport/gin"
	ginuser "nhaancs/modules/user/transport/gin"
	"nhaancs/pubsub/pblocal"
	"nhaancs/socketengine"
	"nhaancs/subscriber"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println(os.Getenv("DSN"))
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
	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub())
	r := gin.Default()

	rtEngine := socketengine.NewEngine()
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatalln(err)
	}

	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.RequiredAuthOrNot(appCtx))

	v1 := r.Group("v1")
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))
	v1.GET("/favorited-posts", middleware.RequiredAuth(appCtx), ginfavorite.ListFavoritedPostsOfAUser(appCtx))
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
		posts.GET("/:id/favorited-users", middleware.RequiredAuth(appCtx), ginfavorite.ListUsersFavoritedAPost(appCtx))
	}

	return r.Run()
}
