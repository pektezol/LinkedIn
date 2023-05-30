package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var connectionCount int
	sql := `SELECT COUNT(*) FROM connections WHERE (sender_id = $1 OR reciever_id = $1) AND status = true;`
	if err := database.DB.QueryRow(sql, userObject.ID).Scan(&connectionCount); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	profileResponse := ProfileResponse{User: userObject, ConnectionCount: connectionCount}
	profileResponse.Educations = getEducations(userObject.ID)
	profileResponse.Experiences = getExperiences(userObject.ID)
	profileResponse.Skills = getSkills(userObject.ID)
	c.JSON(http.StatusOK, OkMessage(profileResponse))
}

func GetUser(c *gin.Context) {
	var user User
	username := c.Param("username")
	sql := `SELECT id, username, firstname, lastname, email, dateofbirth, profilepicture, headline, industry, location, bio, cv FROM users WHERE username = $1;`
	database.DB.QueryRow(sql, username).Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth, &user.ProfilePicture, &user.Headline, &user.Industry, &user.Location, &user.Bio, &user.CV)
	if user.ID == 0 {
		// User does not exist
		c.JSON(http.StatusOK, ErrorMessage("User does not exist."))
		return
	}
	var connectionCount int
	sql = `SELECT COUNT(*) FROM connections WHERE (sender_id = $1 OR reciever_id = $1) AND status = true;`
	if err := database.DB.QueryRow(sql, user.ID).Scan(&connectionCount); err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	profileResponse := ProfileResponse{User: user, ConnectionCount: connectionCount}
	profileResponse.Educations = getEducations(user.ID)
	profileResponse.Experiences = getExperiences(user.ID)
	profileResponse.Skills = getSkills(user.ID)
	c.JSON(http.StatusOK, OkMessage(profileResponse))
}

func AddCV(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var request CVRequest
	err := c.ShouldBindJSON(&request)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO cv(user_id,data_base64) VALUES($1, $2);`
	_, err = database.DB.Exec(sql, userObject.ID, request.Data)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(request))
}

func AddEducation(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var educationRequest EducationRequest
	err := c.ShouldBindJSON(&educationRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO education(user_id,school_name,degree,field_of_study,description,start_date,end_date) 
	VALUES($1,$2,$3,$4,$5,$6,$7);`
	database.DB.Exec(sql, userObject.ID, educationRequest.SchoolName, educationRequest.Degree, educationRequest.FieldOfStudy, educationRequest.Description, educationRequest.StartDate, educationRequest.EndDate)
	c.JSON(http.StatusOK, OkMessage(educationRequest))
}

func DeleteEducation(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	educationID := c.Param("id")
	sql := `DELETE FROM education WHERE id = $1 AND user_id = $2;`
	database.DB.Exec(sql, educationID, userObject.ID)
	c.JSON(http.StatusOK, OkMessage(nil))
}

func AddExperience(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var experienceRequest ExperienceRequest
	err := c.ShouldBindJSON(&experienceRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO experience(user_id,company_id,title,location,description,start_date,end_date)
	VALUES($1,$2,$3,$4,$5,$6,$7);`
	database.DB.Exec(sql, userObject.ID, experienceRequest.CompanyID, experienceRequest.Title, experienceRequest.Location, experienceRequest.Description, experienceRequest.StartDate, experienceRequest.EndDate)
	c.JSON(http.StatusOK, OkMessage(experienceRequest))
}

func DeleteExperience(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	experienceID := c.Param("id")
	sql := `DELETE FROM experience WHERE id = $1 AND user_id = $2;`
	database.DB.Exec(sql, experienceID, userObject.ID)
	c.JSON(http.StatusOK, OkMessage(nil))
}

func AddSkill(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var skillRequest SkillRequest
	err := c.ShouldBindJSON(&skillRequest)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO skill(user_id,name) VALUES($1,$2);`
	database.DB.Exec(sql, userObject.ID, skillRequest.Name)
	c.JSON(http.StatusOK, OkMessage(skillRequest))
}

func DeleteSkill(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	skillID := c.Param("id")
	sql := `DELETE FROM skill WHERE id = $1 AND user_id = $2;`
	database.DB.Exec(sql, skillID, userObject.ID)
	c.JSON(http.StatusOK, OkMessage(nil))
}

func getEducations(id int) []Education {
	var educations []Education
	sql := `SELECT id, school_name, degree, field_of_study, description, start_date, end_date FROM education WHERE user_id = $1;`
	rows, _ := database.DB.Query(sql, id)
	for rows.Next() {
		var education Education
		rows.Scan(&education.ID, &education.SchoolName, &education.Degree, &education.FieldOfStudy, &education.Description, &education.StartDate, &education.EndDate)
		educations = append(educations, education)
	}
	return educations
}

func getExperiences(id int) []Experince {
	var experiences []Experince
	sql := `SELECT c.id, c.name, c.logo, e.id, e.title, e.description, e.location, e.start_date, e.end_date
	FROM experience e INNER JOIN companies c ON e.company_id=c.id WHERE e.user_id = $1;`
	rows, _ := database.DB.Query(sql, id)
	for rows.Next() {
		var experience Experince
		rows.Scan(&experience.Company.ID, &experience.Company.Name, &experience.Company.Logo, &experience.ID, &experience.Title, &experience.Description, &experience.Location, &experience.StartDate, &experience.EndDate)
		experiences = append(experiences, experience)
	}
	return experiences
}

func getSkills(id int) []Skill {
	var skills []Skill
	sql := `SELECT id, name FROM skill WHERE user_id = $1;`
	rows, _ := database.DB.Query(sql, id)
	for rows.Next() {
		var skill Skill
		rows.Scan(&skill.ID, &skill.Name)
		skills = append(skills, skill)
	}
	return skills
}
