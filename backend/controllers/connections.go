package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConnectionRequests(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	connections := ConnectionsResponse{Connections: []Connection{}}
	sql := `SELECT id, sender_id, reciever_id, status, date FROM connections WHERE sender_id = $1 OR reciever_id = $1;`
	rows, err := database.DB.Query(sql, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var connection Connection
		if err := rows.Scan(&connection.ID, &connection.SenderID, &connection.RecieverID, &connection.Status, &connection.Date); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		connections.Connections = append(connections.Connections, connection)
	}
	c.JSON(http.StatusOK, OkMessage(connections))
}

func SendConnectionRequest(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	targetUsername := c.Param("username")
	var count int
	sql := `SELECT COUNT(*) FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	if err := database.DB.QueryRow(sql, userObject.ID, targetUsername).Scan(&count); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if count != 0 {
		c.JSON(http.StatusOK, ErrorMessage("You already have sent a connection request to this person!"))
		return
	}
	sql = `INSERT INTO connections(sender_id,reciever_id,status) VALUES($1,(SELECT id FROM users WHERE username = $2),$3);`
	_, err := database.DB.Exec(sql, userObject.ID, targetUsername, false)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(targetUsername))
}

func AcceptConnectionRequest(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	targetUsername := c.Param("username")
	var count int
	sql := `SELECT COUNT(*) FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&count); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, ErrorMessage("You do not have a connection request from this person!"))
		return
	}
	var status bool
	sql = `SELECT status FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&status); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if status {
		c.JSON(http.StatusOK, ErrorMessage("You already have a connection to this person!"))
		return
	}
	sql = `UPDATE connections SET status = true;`
	_, err := database.DB.Exec(sql)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(targetUsername))
}

func RemoveConnection(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	targetUsername := c.Param("username")
	var count int
	sql := `SELECT COUNT(*) FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&count); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, ErrorMessage("You are not connected with this person!"))
		return
	}
	var status bool
	sql = `SELECT status FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&status); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if !status {
		c.JSON(http.StatusOK, ErrorMessage("You are not connected with this person!"))
		return
	}
	sql = `DELETE FROM connections WHERE sender_id = $1 AND reciever_id = (SELECT id FROM users WHERE username = $2);`
	_, err := database.DB.Exec(sql, targetUsername, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(targetUsername))
}
