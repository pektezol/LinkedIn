package controllers

import (
	"fmt"
	"linkedin/database"
	"net/http"
	"net/mail"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var registerRequest RegisterRequest
	err := c.ShouldBindJSON(&registerRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	// Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 14)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(fmt.Sprintf("Error on hashing password: %s", err)))
		return
	}
	// Validation check
	// Email
	_, err = mail.ParseAddress(registerRequest.Email)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage("Invalid email."))
		return
	}
	// Password
	if len(registerRequest.Password) < 6 {
		c.JSON(http.StatusOK, ErrorMessage("Password cannot be shorter than 6 characters."))
		return
	}
	if len(registerRequest.Password) > 32 {
		c.JSON(http.StatusOK, ErrorMessage("Password cannot be longer than 32 characters."))
		return
	}
	// Date of birth
	_, err = time.Parse("2006-01-02", registerRequest.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage("Invalid date of birth format. Use YYYY-MM-DD."))
		return
	}
	sql := `INSERT INTO users (username, firstname, lastname, email, password, dateofbirth, profilepicture, headline, industry, location, bio, cv) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	_, err = database.DB.Exec(sql, registerRequest.Username, registerRequest.FirstName, registerRequest.LastName, registerRequest.Email, string(hashedPassword), registerRequest.DateOfBirth, "", "", "", "", "", "")
	// Error on SQL
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	// Everything is ok
	c.JSON(http.StatusOK, OkMessage(registerRequest))
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(fmt.Sprintf("Invalid JSON body: %s", err)))
		return
	}
	// Check if username exists
	var id string
	var username string
	var password string
	sql := `SELECT id, username, password FROM users WHERE username = $1;`
	database.DB.QueryRow(sql, loginRequest.Username).Scan(&id, &username, &password)
	if username == "" {
		// User not found
		c.JSON(http.StatusOK, ErrorMessage("Invalid credentials."))
		return
	}
	// If it exists, check password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginRequest.Password))
	if err != nil {
		// Password incorrect
		c.JSON(http.StatusOK, ErrorMessage("Invalid credentials."))
		return
	}
	// If password is correct, generate and return JWT token
	token, err := createTokenJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorMessage(fmt.Sprintf("Could not generate JWT token: %s", err)))
		return
	}
	c.JSON(http.StatusOK, OkMessage(token))
}

func createTokenJWT(id string) (string, error) {
	// 1 - Start creating a new token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24 * 30) // 30 Days
	// 2 - Edit token details
	token.Claims = &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(exp),
		Subject:   id,
	}
	// 3 - Sign token with secret key in env
	signed, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	// Return signed token
	return signed, nil
}
