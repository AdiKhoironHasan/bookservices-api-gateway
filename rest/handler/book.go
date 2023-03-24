package handler

import (
	"net/http"
	"strconv"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
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
			"error": err.Error(),
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
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.BookStore(c, &bookReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success store book",
	},
	)
}

func (r *BookHandler) Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	data, err := r.client.BookDetail(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success get book",
		Data:    data.Book,
	})
}
