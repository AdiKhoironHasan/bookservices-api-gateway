package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	*Handler
}

type HelloRequest struct {
	text string
}

func NewHelloHandler(h *Handler) *HelloHandler {
	return &HelloHandler{h}
}

func (r *HelloHandler) Ping(c *gin.Context) {
	rsp, err := r.client.BooksPing(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": rsp,
	})
	return
}
