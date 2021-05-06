package common

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func ExtractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func IsRequesterAdmin(c *gin.Context) bool {
	requester, exists := c.Get(CurrentUser)
	if exists {
		return requester.(Requester).GetRole() == AdminRole
	} 

	return false
}