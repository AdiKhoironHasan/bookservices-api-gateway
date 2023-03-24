package handler

import (
	"net/http"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	*Handler
}

func NewHelloHandler(h *Handler) *HelloHandler {
	return &HelloHandler{h}
}

func (r *HelloHandler) Ping(c *gin.Context) {
	rsp, err := r.client.BooksPing(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success ping books services",
		Data:    rsp,
	})
}
