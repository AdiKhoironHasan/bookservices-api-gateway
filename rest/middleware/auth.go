package middleware

import (
	"net/http"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/utils"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResDTO{
				Code:    http.StatusUnauthorized,
				Message: "Token is required",
			})
			return
		}

		expired := utils.IsTokenExpired(token)

		if expired {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResDTO{
				Code:    http.StatusUnauthorized,
				Message: "Token is expired",
			})

			return
		}

		c.Next()
	}
}
