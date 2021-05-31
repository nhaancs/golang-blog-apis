package socketengine

import (
	"context"
	"errors"
	"fmt"
	"nhaancs/component"
	"nhaancs/component/tokenprovider/jwt"
	userstorage "nhaancs/modules/user/store"
	"nhaancs/modules/user/transport/socket"
	"sync"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type RealtimeEngine interface {
	UserSockets(userId int) []AppSocket
	EmitToRoom(room string, eventName string, data interface{}) error
	EmitToUser(userId int, eventName string, data interface{}) error
	// Run(ctx component.AppContext, engine *gin.Engine) error
}

type rtEngine struct {
	server  *socketio.Server
	storage map[int][]AppSocket
	locker  *sync.RWMutex
}

func NewEngine() *rtEngine {
	return &rtEngine{
		storage: make(map[int][]AppSocket),
		locker:  new(sync.RWMutex),
	}
}

func (engine *rtEngine) UserSockets(userId int) []AppSocket {
	engine.locker.RLock()
	defer engine.locker.RUnlock()
	return engine.storage[userId]
}

func (engine *rtEngine) EmitToRoom(room string, eventName string, data interface{}) error {
	engine.server.BroadcastToRoom("/", room, eventName, data)
	return nil
}

func (engine *rtEngine) EmitToUser(userId int, eventName string, data interface{}) error {
	sockets := engine.UserSockets(userId)
	for _, socketConnection := range sockets {
		socketConnection.Emit(eventName, data)
	}
	return nil
}

func (engine *rtEngine) saveAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	if v, ok := engine.storage[userId]; ok {
		engine.storage[userId] = append(v, appSck)
	} else {
		engine.storage[userId] = []AppSocket{appSck}
	}

	engine.locker.Unlock()
}

func (engine *rtEngine) removeAppSocket(userId int, appSck AppSocket) {
	engine.locker.Lock()
	defer engine.locker.Unlock()
	if v, ok := engine.storage[userId]; ok {
		for i := range v {
			if v[i] == appSck {
				engine.storage[userId] = append(v[:i], v[i+1:]...)
				break
			}
		}
	}
}

// starts socket server and
// listens to important events
func (engine *rtEngine) Run(appCtx component.AppContext, r *gin.Engine) error {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})
	engine.server = server

	server.OnConnect("/", func(socketConnection socketio.Conn) error {
		socketConnection.SetContext("")
		fmt.Println("connected:", socketConnection.ID(), " IP:", socketConnection.RemoteAddr(), socketConnection.ID())
		return nil
	})

	server.OnError("/", func(socketConnection socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(socketConnection socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	// Setup
	server.OnEvent("/", "authenticate", func(socketConnection socketio.Conn, token string) {
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
			socketConnection.Emit("authentication_failed", err.Error())
			socketConnection.Close()
			return
		}
		if user.DeletedAt != nil {
			socketConnection.Emit("authentication_failed", errors.New("you has been banned/deleted"))
			socketConnection.Close()
			return
		}
		user.Mask(false)

		appSck := NewAppSocket(socketConnection, user)
		engine.saveAppSocket(user.Id, appSck)
		socketConnection.Emit("authenticated", user)

		//appSck.Join(user.GetRole()) // the same
		//if user.GetRole() == "admin" {
		//	appSck.Join("admin")
		//}

		server.OnEvent("/", "UserUpdateLocation", socketuser.OnUserUpdateLocation(appCtx, user))
	})

	go server.Serve()
	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	return nil
}
