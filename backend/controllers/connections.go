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
	sql := `SELECT c.id, u.id, u.username, u.firstname, u.lastname, u.headline, c.status, c.date FROM connections c INNER JOIN users u ON c.sender_id=u.id WHERE reciever_id = $1;`
	rows, err := database.DB.Query(sql, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var connection Connection
		if err := rows.Scan(&connection.ID, &connection.Sender.ID, &connection.Sender.UserName, &connection.Sender.FirstName, &connection.Sender.LastName, &connection.Sender.Headline, &connection.Status, &connection.Date); err != nil {
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
	sql := `SELECT COUNT(*) FROM connections WHERE reciever_id = $2 AND sender_id = (SELECT id FROM users WHERE username = $1);`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&count); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, ErrorMessage("You do not have a connection request from this person!"))
		return
	}
	var status bool
	sql = `SELECT status FROM connections WHERE reciever_id = $2 AND sender_id = (SELECT id FROM users WHERE username = $1);`
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
	var connectionID int
	var status bool
	sql := `SELECT id, status FROM connections WHERE (reciever_id = $2 AND sender_id = (SELECT id FROM users WHERE username = $1)) OR (sender_id = $2 AND reciever_id = (SELECT id FROM users WHERE username = $1));`
	if err := database.DB.QueryRow(sql, targetUsername, userObject.ID).Scan(&connectionID, &status); err != nil {
		c.JSON(http.StatusOK, ErrorMessage("You are not connected to this person!"))
		return
	}
	sql = `DELETE FROM connections WHERE id = $1;`
	_, err := database.DB.Exec(sql, connectionID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(targetUsername))
}
