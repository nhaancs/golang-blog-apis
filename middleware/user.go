package middleware

import (
	"errors"
	"nhaancs/common"
	"nhaancs/component"

	"github.com/gin-gonic/gin"
)

func RequiredUser(appCtx component.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var requester = c.MustGet(common.CurrentUser).(common.Requester)
		if requester.GetRole() != "user" {
			panic(common.ErrNoPermission(errors.New("you have no permission")))
		}
		c.Next()
	}
}
