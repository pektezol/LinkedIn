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
	// Scan for each posts
	for rows.Next() {
		var post Post
		post.Comments = []Comment{}
		if err := rows.Scan(&post.ID, &post.Content.Text, &post.Content.Image, &post.Date, &post.User.ID, &post.User.FirstName, &post.User.LastName, &post.User.Headline); err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		posts.Posts = append(posts.Posts, post)
		sql = `SELECT c.id, c.comment, c.date, u.id, u.firstname, u.lastname, u.headline, (SELECT COUNT(*) FROM likes WHERE post_id = c.post_id) 
		FROM comments c INNER JOIN users u ON c.user_id=u.id
		INNER JOIN users u ON c.user_id=u.id WHERE c.post_id = $1;`
		rows, err = database.DB.Query(sql, post)
		if err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		// Scan for each comments and likes
		for rows.Next() {
			var comment Comment
			if err := rows.Scan(&comment.ID, &comment.Comment, &comment.Date, &comment.User.ID, &comment.User.FirstName, &comment.User.LastName, &comment.User.Headline, &post.Likes); err != nil {
				c.JSON(http.StatusOK, ErrorMessage(err.Error()))
				return
			}
			post.Comments = append(post.Comments, comment)
		}
	}
	c.JSON(http.StatusOK, posts)
}
