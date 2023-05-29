package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("q")
	response := SearchResponse{Users: []struct {
		ID        int    "json:\"id\""
		FirstName string "json:\"first_name\""
		LastName  string "json:\"last_name\""
		UserName  string "json:\"user_name\""
		Headline  string "json:\"headline\""
	}{}, Company: []struct {
		ID   int    "json:\"id\""
		Name string "json:\"name\""
		Logo string "json:\"logo\""
	}{}}
	sql := `SELECT id, firstname, lastname, username, headline FROM users WHERE LOWER(username) LIKE $1 OR LOWER(firstname) LIKE $1 OR LOWER(lastname) LIKE $1;`
	rows, err := database.DB.Query(sql, "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		user := struct {
			ID        int    "json:\"id\""
			FirstName string "json:\"first_name\""
			LastName  string "json:\"last_name\""
			UserName  string "json:\"user_name\""
			Headline  string "json:\"headline\""
		}{}
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Headline); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		response.Users = append(response.Users, user)
	}
	c.JSON(http.StatusOK, OkMessage(response))
}
