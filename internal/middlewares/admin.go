package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly(ctx *gin.Context) {
	role, exists := ctx.Get("role")
	if !exists || role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "Admin access required",
		})
		return
	}
	ctx.Next()
}
