package main

import (
	"log"
	"nhaancs/component"
	"nhaancs/middleware"
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
	return r.Run()
}
