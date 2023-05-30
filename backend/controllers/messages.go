package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMessages(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	messagesByUser := map[int]UserMessages{}
	sql := `SELECT m.id,
	s.id, s.firstname, s.lastname, s.username, s.headline,
	r.id, r.firstname, r.lastname, r.username, r.headline,
	message, date
	FROM messages m
	JOIN users s ON m.sender_id = s.id
	JOIN users r ON m.reciever_id = r.id
	WHERE m.sender_id = $1 OR m.reciever_id = $1
	ORDER BY date DESC;`
	rows, err := database.DB.Query(sql, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var message Message
		// Scan the row into the Message struct
		err := rows.Scan(&message.ID,
			&message.Sender.ID, &message.Sender.FirstName, &message.Sender.LastName, &message.Sender.UserName, &message.Sender.Headline,
			&message.Receiver.ID, &message.Receiver.FirstName, &message.Receiver.LastName, &message.Receiver.UserName, &message.Receiver.Headline,
			&message.Message, &message.Date)
		if err != nil {
			// handle error
		}

		// Determine the ID and details of the other user in the conversation
		var otherUserID int
		var otherUser UserShort
		if message.Sender.ID == userObject.ID {
			otherUserID = message.Receiver.ID
			otherUser = message.Receiver
		} else {
			otherUserID = message.Sender.ID
			otherUser = message.Sender
		}

		shortMessage := MessageShort{ID: message.ID, Message: message.Message, Date: message.Date}

		// Check if we already have a group for the other user
		if _, ok := messagesByUser[otherUserID]; !ok {
			// If not, create a new group
			messagesByUser[otherUserID] = UserMessages{
				OtherUser: otherUser,
				Messages:  []MessageShort{shortMessage},
			}
		} else {
			// If the user group exists, retrieve it, append the new message, and reassign it
			group := messagesByUser[otherUserID]
			group.Messages = append(group.Messages, shortMessage)
			messagesByUser[otherUserID] = group
		}
	}
	groupedMessages := []UserMessages{}
	for _, userMessages := range messagesByUser {
		groupedMessages = append(groupedMessages, userMessages)
	}
	c.JSON(http.StatusOK, OkMessage(groupedMessages))
}

func GetSpecificMessage(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	targetUsername := c.Param("username")
	messagesByUser := map[int]UserMessages{}
	sql := `SELECT m.id,
	s.id, s.firstname, s.lastname, s.username, s.headline,
	r.id, r.firstname, r.lastname, r.username, r.headline,
	message, date
	FROM messages m
	JOIN users s ON m.sender_id = s.id
	JOIN users r ON m.reciever_id = r.id
	WHERE (m.sender_id = $1 AND m.reciever_id = (SELECT id FROM users WHERE username = $2)) OR (m.reciever_id = $1 AND m.sender_id = (SELECT id FROM users WHERE username = $2))
	ORDER BY date DESC;`
	rows, err := database.DB.Query(sql, userObject.ID, targetUsername)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	type MiniUser struct {
		ID        int    `json:"id"`
		UserName  string `json:"user_name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Headline  string `json:"headline"`
	}
	for rows.Next() {
		var message Message
		// Scan the row into the Message struct
		err := rows.Scan(&message.ID,
			&message.Sender.ID, &message.Sender.FirstName, &message.Sender.LastName, &message.Sender.UserName, &message.Sender.Headline,
			&message.Receiver.ID, &message.Receiver.FirstName, &message.Receiver.LastName, &message.Receiver.UserName, &message.Receiver.Headline,
			&message.Message, &message.Date)
		if err != nil {
			// handle error
		}

		// Determine the ID and details of the other user in the conversation
		var otherUserID int
		var otherUser UserShort
		if message.Sender.ID == userObject.ID {
			otherUserID = message.Receiver.ID
			otherUser = message.Receiver
		} else {
			otherUserID = message.Sender.ID
			otherUser = message.Sender
		}

		shortMessage := MessageShort{ID: message.ID, Message: message.Message, Date: message.Date}

		// Check if we already have a group for the other user
		if _, ok := messagesByUser[otherUserID]; !ok {
			// If not, create a new group
			messagesByUser[otherUserID] = UserMessages{
				OtherUser: otherUser,
				Messages:  []MessageShort{shortMessage},
			}
		} else {
			// If the user group exists, retrieve it, append the new message, and reassign it
			group := messagesByUser[otherUserID]
			group.Messages = append(group.Messages, shortMessage)
			messagesByUser[otherUserID] = group
		}
	}
	groupedMessages := []UserMessages{}
	for _, userMessages := range messagesByUser {
		groupedMessages = append(groupedMessages, userMessages)
	}
	c.JSON(http.StatusOK, OkMessage(groupedMessages))
}

func SendMessage(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	targetUsername := c.Param("username")
	var messageRequest MessageRequest
	err := c.ShouldBindJSON(&messageRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	var targetID int
	sql := `SELECT id FROM users WHERE username = $1`
	err = database.DB.QueryRow(sql, targetUsername).Scan(&targetID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql = `INSERT INTO messages(sender_id,reciever_id,message) VALUES($1,$2,$3);`
	_, err = database.DB.Exec(sql, userObject.ID, targetID, messageRequest.Message)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(messageRequest))
}
