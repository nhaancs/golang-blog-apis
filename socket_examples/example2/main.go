package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"nhaancs/component"
	"nhaancs/component/tokenprovider/jwt"
	userstorage "nhaancs/modules/user/store"
	"os"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load("../../.env")
	db, _ := gorm.Open(mysql.Open(os.Getenv("DSN")), nil)
	appCtx := component.NewAppContext(db, nil, "", nil)
	router := gin.New()
	socketServer := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	socketServer.OnConnect("/", func(socketConnection socketio.Conn) error {
		//socketConnection.SetContext("")
		log.Println("connected:", socketConnection.ID(), " IP:", socketConnection.RemoteAddr())
		//socketConnection.Join("Shipper")
		// socketServer.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
		return nil
	})

	socketServer.OnEvent("/", "greeting", func(socketConnection socketio.Conn, msg string) {
		log.Println("greeting from client:", msg)
	})

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	socketServer.OnEvent("/", "sendInfo", func(socketConnection socketio.Conn, p Person) {
		log.Println("info sent from client:", p.Name, p.Age)
		p.Age = 33
		socketConnection.Emit("sendInfo", p)
	})

	socketServer.OnEvent("/", "authenticate", func(socketConnection socketio.Conn, token string) {
		// Validate token
		// If false: socketConnection.Close(), and return
		// If true
		// => UserId
		// Fetch db find user by Id
		// Here: socketConnection belongs to who? (user_id)
		// We need a map[user_id][]socketio.Conn

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			socketConnection.Emit("authentication_failed", err.Error())
			socketConnection.Close()
			return
		}

		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			socketConnection.Emit("authentication_failed", "database error: " + err.Error())
			socketConnection.Close()
			return
		}
		if user.DeletedAt != nil {
			socketConnection.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			socketConnection.Close()
			return
		}
		user.Mask(false)

		socketConnection.Emit("your_profile", user)
	})

	socketServer.OnError("/", func(socketConnection socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	socketServer.OnDisconnect("/", func(socketConnection socketio.Conn, reason string) {
		log.Println("closed", reason)
		// Remove socket from socket router (from app context)
	})

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %socketConnection\n", err)
		}
	}()
	defer socketServer.Close()

	router.GET("/socket.io/*any", gin.WrapH(socketServer))
	router.POST("/socket.io/*any", gin.WrapH(socketServer))
	router.StaticFS("/socket-example-2", http.Dir("./asset"))
	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
