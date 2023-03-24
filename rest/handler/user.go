package handler

import (
	"net/http"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	*Handler
}

func NewUserHandler(h *Handler) *UserHandler {
	return &UserHandler{h}
}

func (r *UserHandler) Ping(c *gin.Context) {
	// res, err := r.client.UsersPing(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
	// 		Code:    http.StatusInternalServerError,
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success ping Users services",
		// Data:    res.Message,
	})
}
