package middleware

import (
	"net/http"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/utils"
	"github.com/gin-gonic/gin"
)

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {

		statusCode := c.Writer.Status()
		if !utils.IsValidCode(statusCode) {
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}

		c.Next()
	}
}
