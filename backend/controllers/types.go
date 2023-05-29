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
	User            User `json:"user"`
	ConnectionCount int  `json:"connection_count"`
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

type Company struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type Application struct {
	JobID  int       `json:"job_id"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
}

type Job struct {
	Company struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"company"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
}

type Education struct {
	ID           int       `json:"id"`
	SchoolName   string    `json:"school_name"`
	Degree       string    `json:"degree"`
	FieldOfStudy string    `json:"field_of_study"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}

type Experince struct {
	ID      int `json:"id"`
	Company struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"company"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"title"`
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
	Likes       int       `json:"likes"`
	Comments    []Comment `json:"comments"`
	LikedStatus bool      `json:"liked_status"`
	Date        time.Time `json:"date"`
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
	ID     int `json:"id"`
	Sender struct {
		ID        int    `json:"id"`
		UserName  string `json:"user_name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Headline  string `json:"headline"`
	} `json:"sender"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
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

type CommentRequest struct {
	Text string `json:"text" binding:"required"`
}

type EducationRequest struct {
	SchoolName   string    `json:"school_name" binding:"required"`
	Degree       string    `json:"degree" binding:"required"`
	FieldOfStudy string    `json:"field_of_study" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	StartDate    time.Time `json:"start_date" binding:"required"`
	EndDate      time.Time `json:"end_date" binding:"required"`
}

type ExperienceRequest struct {
	CompanyID   int       `json:"company_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

type CompanyRequest struct {
	ID          int    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Industry    string `json:"industry" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	Logo        string `json:"logo" binding:"required"`
}
