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
	c.JSON(http.StatusOK, ProfileResponse{
		User: userObject,
	})
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
	c.JSON(http.StatusOK, ProfileResponse{
		User: user,
	})
}
