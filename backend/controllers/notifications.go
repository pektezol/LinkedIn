package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotifications(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var notifications []Notification
	sql := `SELECT id, notification, status, date FROM notifications WHERE user_id = $1`
	rows, err := database.DB.Query(sql, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var notification Notification
		rows.Scan(&notification.ID, &notification.Notification, &notification.Status, &notification.Date)
		notifications = append(notifications, notification)
	}
	c.JSON(http.StatusOK, OkMessage(notifications))
}

func SendNotification(userID int, notification string) error {
	sql := `INSERT INTO notifications(user_id,notification) VALUES($1, $2);`
	_, err := database.DB.Exec(sql, userID, notification)
	if err != nil {
		return err
	}
	return nil
}
