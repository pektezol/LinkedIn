package controllers

import (
	"fmt"
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts PostsResponse
	posts.Posts = []Post{}
	sql := `SELECT p.id, p.text, p.image, p.date, u.id, u.firstname, u.lastname, u.headline FROM posts p INNER JOIN users u ON p.user_id=u.id ORDER BY date DESC;;`
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
		sql = `SELECT c.id, c.comment, c.date, u.id, u.firstname, u.lastname, u.headline, (SELECT COUNT(*) FROM likes WHERE post_id = c.post_id) 
		FROM comments c INNER JOIN users u ON c.user_id=u.id
		WHERE c.post_id = $1;`
		commentRows, err := database.DB.Query(sql, post.ID)
		if err != nil {
			c.JSON(http.StatusOK, ErrorMessage(err.Error()))
			return
		}
		// Scan for each comments and likes
		for commentRows.Next() {
			var comment Comment
			if err := commentRows.Scan(&comment.ID, &comment.Comment, &comment.Date, &comment.User.ID, &comment.User.FirstName, &comment.User.LastName, &comment.User.Headline, &post.Likes); err != nil {
				c.JSON(http.StatusOK, ErrorMessage(err.Error()))
				return
			}
			post.Comments = append(post.Comments, comment)
		}
		fmt.Println(post)
		posts.Posts = append(posts.Posts, post)
	}
	c.JSON(http.StatusOK, OkMessage(posts))
}

func CreatePost(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var postRequest PostRequest
	err := c.ShouldBindJSON(&postRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO posts(user_id, text, image) VALUES($1, $2, $3);`
	_, err = database.DB.Exec(sql, userObject.ID, postRequest.Text, postRequest.Image)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(postRequest))
}

func Like(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var likeID int
	postID := c.Param("post")
	sql := `SELECT id FROM likes WHERE user_id = $1 AND post_id = $2;`
	database.DB.QueryRow(sql, userObject.ID, postID).Scan(&likeID)
	if likeID == 0 {
		// Didn't like the post
		sql := `INSERT INTO likes(post_id,user_id) VALUES($1,$2);`
		database.DB.Exec(sql, postID, userObject.ID)
		c.JSON(http.StatusOK, OkMessage(nil))
		return
	}
	// Already liked the post
	sql = `DELETE FROM likes WHERE id = $1;`
	database.DB.Exec(sql, likeID)
	c.JSON(http.StatusOK, OkMessage(nil))
}
