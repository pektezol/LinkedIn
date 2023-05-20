package controllers

type OkResponse struct {
	Status string `json:"status"`
	Logged bool   `json:"logged"`
	Perm   string `json:"perm"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func OkMessage(data any) OkResponse {
	return OkResponse{
		Status: "ok",
		Data:   data,
	}
}

func ErrorMessage(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
	}
}

type UserRegister struct {
	Username    string `json:"user_name" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}
type UserLogin struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
