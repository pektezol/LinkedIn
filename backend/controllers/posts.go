package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	sql := `SELECT p.id, p.text, p.image, p.date, u.id, u.firstname, u.lastname, u.headline FROM posts p INNER JOIN users u ON p.user_id=u.id;`
	fmt.Println(sql)
	// database.DB.QueryRow(sql, username).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.ProfilePicture, &user.Headline, &user.Industry, &user.Location, &user.Bio, &user.CV)
}
