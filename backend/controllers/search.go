package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("q")
	response := SearchResponse{Users: []UserShort{}, Company: []CompanyShort{}}
	sql := `SELECT id, firstname, lastname, username, headline FROM users WHERE LOWER(username) LIKE $1 OR LOWER(firstname) LIKE $1 OR LOWER(lastname) LIKE $1;`
	rows, err := database.DB.Query(sql, "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var user UserShort
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Headline); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		response.Users = append(response.Users, user)
	}
	sql = `SELECT id, name, logo FROM companies WHERE LOWER(name) LIKE $1;`
	rows, err = database.DB.Query(sql, "%"+query+"%")
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var company CompanyShort
		if err := rows.Scan(&company.ID, &company.Name, &company.Logo); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		response.Company = append(response.Company, company)
	}
	c.JSON(http.StatusOK, OkMessage(response))
}
