package handler

import (
	"net/http"
	"strconv"

	"github.com/AdiKhoironHasan/bookservices-api-gateway/domain/dto"
	"github.com/AdiKhoironHasan/bookservices-api-gateway/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	*Handler
}

func NewUserHandler(h *Handler) *UserHandler {
	return &UserHandler{h}
}

func (r *UserHandler) Ping(c *gin.Context) {
	res, err := r.client.UserPing(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success ping user services",
		Data:    res.Message,
	})
}

func (r *UserHandler) List(c *gin.Context) {
	userReq := dto.UserReqDTO{
		Name: c.Query("name"),
		Role: c.Query("role"),
	}

	data, err := r.client.UserList(c, &userReq)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Data:    data.Users,
		Message: "Success get list user",
	})
}

func (r *UserHandler) Store(c *gin.Context) {
	userReq := dto.UserReqDTO{}

	err := c.Bind(&userReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.Usertore(c, &userReq)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, dto.ApiResDTO{
		Code:    http.StatusCreated,
		Message: "Success store user",
	},
	)
}

func (r *UserHandler) Detail(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	data, err := r.client.UserDetail(c, Id)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success get detail user",
		Data:    data.User,
	})
}

func (r *UserHandler) Update(c *gin.Context) {
	userReq := dto.UserReqDTO{}

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	err = c.Bind(&userReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.UserUpdate(c, &userReq, Id)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success update user",
	},
	)
}

func (r *UserHandler) Delete(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	_, err = r.client.UserDelete(c, Id)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success delete user",
	},
	)
}

func (r *UserHandler) Login(c *gin.Context) {
	loginReq := dto.LoginReqDTO{}

	err := c.Bind(&loginReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, dto.ApiResDTO{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
		return
	}

	data, err := r.client.UserLogin(c, &loginReq)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
				Code:    int(status.Code()),
				Message: status.Message(),
			})
			return
		}
	}

	token, err := utils.GenerateToken(data.User.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ApiResDTO{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.ApiResDTO{
		Code:    http.StatusOK,
		Message: "Success login user",
		Data:    token,
	})
}

// func (r *UserHandler) Logout(c *gin.Context) {
// 	token := c.Request.Header.Get("Authorization")
// 	tokenReq := dto.AuthTokenDTO{
// 		Token: token,
// 	}

// 	_, err := r.client.UserLogout(c, &tokenReq)
// 	if err != nil {
// 		if status, ok := status.FromError(err); ok {
// 			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
// 				Code:    int(status.Code()),
// 				Message: status.Message(),
// 			})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, dto.ApiResDTO{
// 		Code:    http.StatusOK,
// 		Message: "Success logout user",
// 	})
// }

// func (r *UserHandler) Verify(c *gin.Context) {
// 	token := c.Request.Header.Get("Authorization")
// 	tokenReq := dto.AuthTokenDTO{
// 		Token: token,
// 	}

// 	_, err := r.client.UserVerify(c, &tokenReq)
// 	if err != nil {
// 		if status, ok := status.FromError(err); ok {
// 			c.AbortWithStatusJSON(int(status.Code()), dto.ApiResDTO{
// 				Code:    int(status.Code()),
// 				Message: status.Message(),
// 			})
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, dto.ApiResDTO{
// 		Code:    http.StatusOK,
// 		Message: "Token is valid",
// 	})
// }
