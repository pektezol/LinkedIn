package middleware

import (
	"fmt"
	"linkedin/controllers"
	"linkedin/database"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {
	// Get auth cookie
	tokenString := c.GetHeader("Authorization")
	// Validate token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if token == nil {
		c.Next()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Next()
			return
		}
		// Get user from DB
		var user controllers.User
		sql := `SELECT id, username, firstname, lastname, email, dateofbirth, profilepicture, headline, industry, location, bio, cv, background FROM users WHERE id = $1;`
		database.DB.QueryRow(sql, claims["sub"]).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.ProfilePicture, &user.Headline, &user.Industry, &user.Location, &user.Bio, &user.CV)
		if user.ID == 0 {
			c.Next()
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.Next()
		return
	}
}
