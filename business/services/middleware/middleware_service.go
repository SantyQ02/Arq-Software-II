package middlewareService

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	authUtils "mvc-go/utils/auth"

	userClient "mvc-go/clients/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type middlewareService struct{}

type middlewareServiceInterface interface {
	DeserializeUser() gin.HandlerFunc
	CheckAdmin() gin.HandlerFunc
}

var (
	MiddlewareService middlewareServiceInterface
)

func init() {
	MiddlewareService = &middlewareService{}
}

var Db *gorm.DB

func (m *middlewareService) DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
			return
		}

		sub, err := authUtils.TokenClient.ValidateToken(token, os.Getenv("TOKEN_SECRET"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user := userClient.UserClient.GetUserById(fmt.Sprint(sub))
		if user.UserID == uuid.Nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "the user belonging to this token no logger exists"})
			return
		}
		if !user.Active {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Your account has been desactivated"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}

func (m *middlewareService) CheckAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var token string
		cookie, err := ctx.Cookie("token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
			return
		}

		sub, err := authUtils.TokenClient.ValidateToken(token, os.Getenv("TOKEN_SECRET"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user := userClient.UserClient.GetUserById(fmt.Sprint(sub))
		if user.UserID == uuid.Nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "the user belonging to this token no logger exists"})
			return
		}

		if user.Role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required"})
			return
		}
		if !user.Active {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Your account has been desactivated"})
			return
		}

		ctx.Next()
	}
}
