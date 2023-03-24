package handler

import (
	"fmt"
	"net/http"

	"github.com/AdiKhoironHasan/bookservices/domain/dto"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	*Handler
}

func NewBookHandler(h *Handler) *BookHandler {
	return &BookHandler{h}
}

func (r *BookHandler) List(c *gin.Context) {
	data, err := r.client.BookList(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"data":    data.Books,
		"message": "Success get all books",
	})
}

func (r *BookHandler) Store(c *gin.Context) {
	bookReq := dto.BookReqDTO{}

	err := c.Bind(&bookReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("%v", err),
		})
		return
	}

	data, err := r.client.BookStore(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Data:    data.Books,
		Message: "Success store book",
	},
	)
}
