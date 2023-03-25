package dto

type UserReqDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthTokenDTO struct {
	Token string `json:"token"`
}
