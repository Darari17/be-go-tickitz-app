package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Darari17/be-go-tickitz-app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Silahkan login terlebih dahulu",
		})
		return
	}

	token := strings.TrimPrefix(bearerToken, "Bearer ")

	claims := &pkg.Claims{}

	if err := claims.VerifyToken(token); err != nil {
		if strings.Contains(err.Error(), jwt.ErrTokenInvalidIssuer.Error()) {
			log.Println("JWT Error.\nCause: ", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Silahkan login kembali",
			})
			return
		}
		if strings.Contains(err.Error(), jwt.ErrTokenExpired.Error()) {
			log.Println("JWT Error.\nCause: ", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Silahkan login kembali",
			})
			return
		}
		fmt.Println(jwt.ErrTokenExpired)
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal Server Error",
		})
		return
	}

	ctx.Set("claims", claims)
	ctx.Next()
}
