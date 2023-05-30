package controllers

import (
	"linkedin/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	var request CompanyRequest
	err := c.ShouldBindJSON(&request)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql := `INSERT INTO companies(name,industry,logo,location,description,employer_id) VALUES($1,$2,$3,$4,$5,$6);`
	_, err = database.DB.Exec(sql, request.Name, request.Industry, request.Logo, request.Location, request.Description, userObject.ID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(request))
}

func CreateJobOpening(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	companyID := c.Param("id")
	var employerID int
	sql := `SELECT employer_id FROM companies WHERE id = $1;`
	database.DB.QueryRow(sql, companyID).Scan(&employerID)
	if employerID != userObject.ID {
		c.JSON(http.StatusOK, ErrorMessage("You are not the employer of this company."))
		return
	}
	var request JobOpeningRequest
	err := c.ShouldBindJSON(&request)
	// Error on Invalid JSON
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql = `INSERT INTO jobs(title,company_id,description,location,"type") VALUES($1,$2,$3,$4,$5);`
	_, err = database.DB.Exec(sql, request.Title, companyID, request.Description, request.Location, request.Type)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(request))
}

func GetJobOpenings(c *gin.Context) {
	response := JobOpeningsResponse{Openings: []JobOpening{}}
	sql := `SELECT c.id, c.name, c.logo, j.id, j.title, j.location, j.description, j.type, j.date FROM jobs j INNER JOIN companies c ON j.company_id=c.id WHERE j.filled = false`
	rows, err := database.DB.Query(sql)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var job JobOpening
		rows.Scan(&job.Company.ID, &job.Company.Name, &job.Company.Logo, &job.ID, &job.Title, &job.Location, &job.Description, &job.Type, &job.Date)
		response.Openings = append(response.Openings, job)
	}
	c.JSON(http.StatusOK, OkMessage(response))
}

func SendJobApplication(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	jobID := c.Param("id")
	var count int
	var filled bool
	sql := `SELECT filled FROM jobs WHERE id = $1`
	database.DB.QueryRow(sql, userObject.ID).Scan(&filled)
	if filled {
		c.JSON(http.StatusOK, ErrorMessage("This job opening is filled."))
		return
	}
	sql = `SELECT COUNT(*) FROM applications WHERE user_id = $1`
	database.DB.QueryRow(sql, userObject.ID).Scan(&count)
	if count != 0 {
		c.JSON(http.StatusOK, ErrorMessage("You have already applied for this job application."))
		return
	}
	sql = `INSERT INTO applications(user_id,job_id,status) VALUES($1,$2,false);`
	_, err := database.DB.Exec(sql, userObject.ID, jobID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(nil))
}

func AcceptJobApplication(c *gin.Context) {
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	applicationID := c.Param("applicationid")
	var employerID int
	var jobID int
	sql := `SELECT j.id, employer_id FROM companies c INNER JOIN jobs j ON c.id=j.company_id INNER JOIN applications a ON j.id=a.job_id WHERE a.id = $1;`
	database.DB.QueryRow(sql, applicationID).Scan(&jobID, &employerID)
	if employerID != userObject.ID {
		c.JSON(http.StatusOK, ErrorMessage("You are not the employer of this company."))
		return
	}
	sql = `UPDATE applications SET status = true WHERE id = $1`
	_, err := database.DB.Exec(sql, applicationID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql = `DELETE FROM applications WHERE status = false AND id != $1`
	_, err = database.DB.Exec(sql, applicationID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	sql = `UPDATE jobs SET filled = true WHERE id = $1`
	_, err = database.DB.Exec(sql, jobID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	c.JSON(http.StatusOK, OkMessage(nil))
}

func GetJobApplications(c *gin.Context) {
	// employer, show user profile
	sessionUser, exists := c.Get("user")
	if !exists {
		// User not logged in
		c.JSON(http.StatusOK, ErrorMessage("User not logged in."))
		return
	}
	userObject := sessionUser.(User)
	companyID := c.Param("id")
	var employerID int
	sql := `SELECT employer_id FROM companies WHERE id = $1;`
	database.DB.QueryRow(sql, companyID).Scan(&employerID)
	if employerID != userObject.ID {
		c.JSON(http.StatusOK, ErrorMessage("You are not the employer of this company."))
		return
	}
	response := JobApplicationsResponse{Applications: []Application{}}
	sql = `SELECT u.id, u.username, u.firstname, u.lastname, u.headline, u.cv, a.job_id, a.date FROM applications a INNER JOIN users u ON a.user_id=u.id INNER JOIN jobs j ON a.job_id=j.id INNER JOIN companies c ON j.company_id=c.id WHERE c.id = $1 AND j.filled = false`
	rows, err := database.DB.Query(sql, companyID)
	if err != nil {
		c.JSON(http.StatusOK, ErrorMessage(err.Error()))
		return
	}
	for rows.Next() {
		var application Application
		rows.Scan(&application.User.ID, &application.User.UserName, &application.User.FirstName, &application.User.LastName, &application.User.Headline, &application.User.CV, &application.JobID, &application.Date)
		response.Applications = append(response.Applications, application)
	}
	c.JSON(http.StatusOK, OkMessage(response))
}
