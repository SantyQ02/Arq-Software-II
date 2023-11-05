package middlewareService

import (
	"net/http"
	// "encoding/json"
	"strings"
	// "bytes"

	"github.com/gin-gonic/gin"
	"fmt"
)

type middlewareService struct{}

type middlewareServiceInterface interface {
	CheckAdmin() gin.HandlerFunc
}

var (
	MiddlewareService middlewareServiceInterface
)

func init() {
	MiddlewareService = &middlewareService{}
}

// type Bodystruct struct {
//     IsAdmin bool `json:"isAdmin,omitempty"`
// 	Token  string `json:"token,omitempty"`
// }


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

		businessURL := "http://business:8080/api/business/checkadmin"
		// JSON body
		// body := Bodystruct{
		// 	Token: token,
		// }

		// jsonData, _ := json.Marshal(body)

		// Create a HTTP post request
		r, err := http.NewRequest("GET", businessURL, nil)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to make request!"})
		}

		// defer res.Body.Close()

		// response := &Bodystruct{}
		// derr := json.NewDecoder(res.Body).Decode(response)

		// if derr != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response!"})
		// }

		if res.StatusCode != http.StatusOK {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal api error!"})
			return
		}else {
			ctx.Next()
		}

		// if response.IsAdmin{
		// 	ctx.Next()
		// }else {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Admin privileges required!"})
		// 	return
		// }
	}
}
