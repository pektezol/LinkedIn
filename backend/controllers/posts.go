package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts PostsResponse
	posts.Posts = []Post{}
	sql := `SELECT p.id, p.text, p.image, p.date, u.id, u.firstname, u.lastname, u.headline FROM posts p INNER JOIN users u ON p.user_id=u.id;`
	rows, err := database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Content.Text, &post.Content.Image, &post.Date, &post.User.ID, &post.User.FirstName, &post.User.LastName, &post.User.Headline); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		posts.Posts = append(posts.Posts, post)
	}
	c.JSON(http.StatusOK, posts)
}
