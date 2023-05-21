package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var connectionCount int
	sql := `SELECT COUNT(*) FROM connections WHERE (sender_id = $1 OR reciever_id = $1) AND status = true;`
	if err := database.DB.QueryRow(sql, userObject.ID).Scan(&connectionCount); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	profileResponse := ProfileResponse{User: userObject, ConnectionCount: connectionCount}
	c.JSON(http.StatusOK, OkMessage(profileResponse))
}

func GetUser(c *gin.Context) {
	var user User
	username := c.Param("username")
	sql := `SELECT id, username, firstname, lastname, email, dateofbirth, profilepicture, headline, industry, location, bio, cv FROM users WHERE username = $1;`
	database.DB.QueryRow(sql, username).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.ProfilePicture, &user.Headline, &user.Industry, &user.Location, &user.Bio, &user.CV)
	if user.ID == 0 {
		// User does not exist
		c.JSON(http.StatusOK, ErrorMessage("User does not exist."))
		return
	}
	var connectionCount int
	sql = `SELECT COUNT(*) FROM connections WHERE (sender_id = $1 OR reciever_id = $1) AND status = true;`
	if err := database.DB.QueryRow(sql, user.ID).Scan(&connectionCount); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	profileResponse := ProfileResponse{User: user, ConnectionCount: connectionCount}
	c.JSON(http.StatusOK, OkMessage(profileResponse))
}

func AddEducation(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var educationRequest EducationRequest
	err := c.ShouldBindJSON(&educationRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO education(user_id,school_name,degree,field_of_study,description,start_date,end_date) 
	VALUES($1,$2,$3,$4,$5,$6,$7);`
	database.DB.Exec(sql, userObject.ID, educationRequest.SchoolName, educationRequest.Degree, educationRequest.FieldOfStudy, educationRequest.Description, educationRequest.StartDate, educationRequest.EndDate)
	c.JSON(http.StatusOK, OkMessage(educationRequest))
}

func DeleteEducation(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	educationID := c.Param("id")
	sql := `DELETE FROM education WHERE id = $1 AND user_id = $2;`
	database.DB.Exec(sql, educationID, userObject.ID)
	c.JSON(http.StatusOK, OkMessage(nil))
}

func AddExperience(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var experienceRequest ExperienceRequest
	err := c.ShouldBindJSON(&experienceRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO experience(user_id,company_id,title,location,description,start_date,end_date)
	VALUES($1,$2,$3,$4,$5,$6,$7);`
	database.DB.Exec(sql, userObject.ID, experienceRequest.CompanyID, experienceRequest.Title, experienceRequest.Location, experienceRequest.Description, experienceRequest.StartDate, experienceRequest.EndDate)
	c.JSON(http.StatusOK, OkMessage(experienceRequest))
}

func DeleteExperience(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	experienceID := c.Param("id")
	sql := `DELETE FROM experience WHERE id = $1 AND user_id = $2;`
	database.DB.Exec(sql, experienceID, userObject.ID)
	c.JSON(http.StatusOK, OkMessage(nil))
}
