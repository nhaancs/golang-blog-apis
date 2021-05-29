package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	router := gin.New()
	socketServer := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	socketServer.OnConnect("/", func(socketConnection socketio.Conn) error {
		socketConnection.SetContext("")
		log.Println("connected: ", socketConnection.ID(), "ip: ", socketConnection.RemoteAddr())
		return nil
	})

	// listen to msg event of /chat namespace
	// process the message and return the data
	socketServer.OnEvent("/chat", "msg", func(socketConnection socketio.Conn, msg string) string {
		socketConnection.SetContext(msg)
		return "received " + msg
	})

	// listen to notice event of root namespace
	// process message and emit data to reply event of root namespace
	socketServer.OnEvent("/", "notice", func(socketConnection socketio.Conn, msg string) {
		log.Println("notice:", msg)
		socketConnection.Emit("reply", "have "+msg)
	})

	socketServer.OnEvent("/", "bye", func(socketConnection socketio.Conn) string {
		last := socketConnection.Context().(string)
		socketConnection.Emit("bye", last)
		socketConnection.Close()
		return last
	})

	socketServer.OnError("/", func(socketConnection socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	socketServer.OnDisconnect("/", func(socketConnection socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer socketServer.Close()

	router.GET("/socket.io/*any", gin.WrapH(socketServer))
	router.POST("/socket.io/*any", gin.WrapH(socketServer))
	router.StaticFS("/socket-example-1", http.Dir("./asset"))

	if err := router.Run(":8000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
