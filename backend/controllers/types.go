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
	User            User        `json:"user"`
	ConnectionCount int         `json:"connection_count"`
	Experiences     []Experince `json:"experiences"`
	Educations      []Education `json:"educations"`
	Skills          []Skill     `json:"skills"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}

type ConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}

type SearchResponse struct {
	Users   []UserShort    `json:"users"`
	Company []CompanyShort `json:"companies"`
}

type JobOpeningsResponse struct {
	Openings []JobOpening `json:"openings"`
}

type JobApplicationsResponse struct {
	Applications []Application `json:"applications"`
}

type JobOpening struct {
	ID          int          `json:"id"`
	Company     CompanyShort `json:"company"`
	Title       string       `json:"title"`
	Location    string       `json:"location"`
	Description string       `json:"description"`
	Type        string       `json:"type"`
	Date        time.Time    `json:"date"`
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

type UserShort struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Headline  string `json:"headline"`
}

type UserShortWithCV struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Headline  string `json:"headline"`
	CV        string `json:"cv"`
}

type Company struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type CompanyShort struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Application struct {
	ID    int             `json:"id"`
	User  UserShortWithCV `json:"user"`
	JobID int             `json:"job_id"`
	Date  time.Time       `json:"date"`
}

type Job struct {
	Company     CompanyShort `json:"company"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	Type        string       `json:"type"`
	Date        time.Time    `json:"date"`
}

type Education struct {
	ID           int       `json:"id"`
	SchoolName   string    `json:"school_name"`
	Degree       string    `json:"degree"`
	FieldOfStudy string    `json:"field_of_study"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date" time_format:"2006-01-02"`
	EndDate      time.Time `json:"end_date" time_format:"2006-01-02"`
}

type Experince struct {
	ID          int          `json:"id"`
	Company     CompanyShort `json:"company"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Location    string       `json:"location"`
	StartDate   time.Time    `json:"start_date" time_format:"2006-01-02"`
	EndDate     time.Time    `json:"end_date" time_format:"2006-01-02"`
}

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"title"`
}

type Post struct {
	ID      int       `json:"id"`
	User    UserShort `json:"user"`
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
	ID      int       `json:"id"`
	User    UserShort `json:"user"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

type Connection struct {
	ID     int       `json:"id"`
	Sender UserShort `json:"sender"`
	Status bool      `json:"status"`
	Date   time.Time `json:"date"`
}

type Notification struct {
	ID           int       `json:"id"`
	Notification string    `json:"notification"`
	Status       bool      `json:"status"`
	Date         time.Time `json:"date"`
}

type Message struct {
	ID       int       `json:"id"`
	Sender   UserShort `json:"sender"`
	Receiver UserShort `json:"receiver"`
	Message  string    `json:"message"`
	Date     time.Time `json:"date"`
}

type MessageShort struct {
	ID      int       `json:"id"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

type UserMessages struct {
	OtherUser UserShort      `json:"other_user"`
	Messages  []MessageShort `json:"messages"`
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
	StartDate    time.Time `json:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate      time.Time `json:"end_date" binding:"required" time_format:"2006-01-02"`
}

type ExperienceRequest struct {
	CompanyID   int       `json:"company_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate     time.Time `json:"end_date" binding:"required" time_format:"2006-01-02"`
}

type SkillRequest struct {
	Name string `json:"name" binding:"required"`
}

type CompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	Industry    string `json:"industry" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	Logo        string `json:"logo" binding:"required"`
}

type JobOpeningRequest struct {
	Title       string `json:"title" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type JobApplicationRequest struct {
	Title       string `json:"title" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}

type CVRequest struct {
	Data string `json:"data_base64" binding:"required"`
}

type UserUpdateRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	DateOfBirth    string `json:"date_of_birth"`
	ProfilePicture string `json:"profile_picture"`
	Headline       string `json:"headline"`
	Industry       string `json:"industry"`
	Location       string `json:"location"`
	Bio            string `json:"bio"`
}
