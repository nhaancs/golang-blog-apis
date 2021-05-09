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
	// v1.GET("/favorited-posts", middleware.RequiredAuth(appCtx), ginfavorite.List(appCtx)) // todo: implement
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
		posts.GET("/:id/favorited-users", middleware.RequiredAuth(appCtx), ginfavorite.List(appCtx))
	}

	return r.Run()
}

/*
todo:
- A có thể implement mẫu một repository đc ko. Ví dụ như trong phần list restaurant biz có phần mapping like count cho từng nhà hàng, a nói có thể làm trong repository nhưng e  chưa hiểu lắm ạ?

- Mình implement searching và sorting như thế nào v a?

- E thấy trong code có biến ctx chưa dùng tới, tác dụng của nó là gì v a?

- E có một project bao gồm:
    + Web  : https://domain.com - server-side rendering
    + Portal: https://admin.commain.com - client-side rendering
    + API    : https://api.domain.com
    + DB     : mysql
Tất cả đc chứa trong cùng một monorepo, giờ e muốn deploy tất cả  lên 1 instance của amazon bằng docker thì phải làm sao a? E có thể viết dockerfile để chạy từng phần riêng lẻ dưới local đc, nhưng phần https với config domain thì e ko biết ạ. Repository của e là https://github.com/nhaancs/ecommerce-nx-nestjs-angular

- Server nên trả về cho client JSON có key dạng snack_case hay camelCase vậy a?

- Anh thấy thư viện https://github.com/kyleconroy/sqlc như thế nào ạ?

- Nếu ko dùng GORM thì thay thế bằng sqlx (xử lý nhanh) và query builder https://github.com/Masterminds/squirrel (đỡ phải gõ query, hạn chế lỗi typing) thì có ổn ko a?

- E thấy client khi gọi api hay bị lỗi CORS, và sever thường phải disable nó đi. Vậy CORS ở đây có tác dụng gì v a, tắt đi thì có sao k?

- Để làm một ứng dụng high traffic thì cần lưu ý những gì, thường thì mình sẽ cần phải optimize ở những phần nào v a?

- A phân tích giúp e một số JD:
    + https://itviec.com/it-jobs/middle-senior-backend-golang-python-shopee-5946
    + https://itviec.com/it-jobs/mid-senior-backend-java-python-go-tiki-corporation-1625
    + https://itviec.com/it-jobs/backend-developer-golang-python-java-grab-vietnam-ltd-3331
    + https://itviec.com/it-jobs/backend-developer-golang-nodejs-sendo-vn-3747

- Qui trình phỏng vấn Golang của các cty ở VN thường sẽ như thế nào vậy a?

- A review giúp e CV này ạ https://nhaancs.github.io
*/
