package controllers

import "time"

type OkResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ProfileResponse struct {
	User User `json:"user"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}

type ConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}

type User struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	DateOfBirth    string `json:"date_of_birth"`
	ProfilePicture string `json:"profile_picture"`
	Headline       string `json:"headline"`
	Industry       string `json:"industry"`
	Location       string `json:"location"`
	Bio            string `json:"bio"`
	CV             string `json:"cv"`
}

type Post struct {
	ID   int `json:"id"`
	User struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Headline  string `json:"headline"`
	} `json:"user"`
	Content struct {
		Text  string `json:"text"`
		Image string `json:"image_base64"`
	} `json:"content"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments"`
	Date     time.Time `json:"date"`
}

type Comment struct {
	ID   int `json:"id"`
	User struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Headline  string `json:"headline"`
	} `json:"user"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

type Connection struct {
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	RecieverID int       `json:"reciever_id"`
	Status     bool      `json:"status"`
	Date       time.Time `json:"date"`
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

type RegisterRequest struct {
	Username    string `json:"user_name" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}
type LoginRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostRequest struct {
	Text  string `json:"text" binding:"required"`
	Image string `json:"image_base64"`
}
