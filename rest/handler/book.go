package handler

import (
	"net/http"
	"strconv"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type BookHandler struct {
	*Handler
}

func NewBookHandler(h *Handler) *BookHandler {
	return &BookHandler{h}
}

func (r *BookHandler) List(c *gin.Context) {
	bookReq := dto.BookReqDTO{
		Title: c.Query("title"),
	}

	data, err := r.client.BookList(c, &bookReq)
	if status, ok := status.FromError(err); ok {
		c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
			Code:    int(status.Code()),
			Message: status.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Data:    data,
		Message: "Success get list book",
	})
}

func (r *BookHandler) Store(c *gin.Context) {
	bookReq := dto.BookReqDTO{}

	err := c.Bind(&bookReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.BookStore(c, &bookReq)
	if status, ok := status.FromError(err); ok {
		c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
			Code:    int(status.Code()),
			Message: status.Message(),
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
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	data, err := r.client.BookDetail(c, Id)
	if status, ok := status.FromError(err); ok {
		c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
			Code:    int(status.Code()),
			Message: status.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success get detail book",
		Data:    data.Book,
	})
}

func (r *BookHandler) Update(c *gin.Context) {
	bookReq := dto.BookReqDTO{}

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	err = c.Bind(&bookReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.BookUpdate(c, &bookReq, Id)
	if status, ok := status.FromError(err); ok {
		c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
			Code:    int(status.Code()),
			Message: status.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success store book",
	},
	)
}

func (r *BookHandler) Delete(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.BookDelete(c, Id)
	if status, ok := status.FromError(err); ok {
		c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
			Code:    int(status.Code()),
			Message: status.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success delete book",
	},
	)
}
