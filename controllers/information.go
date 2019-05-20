package controllers

import (
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Job struct {
	JobName string
	JobCode string
}

type Education struct {
	EducationCode string
	EducationName string
}

func GetJob(c *gin.Context)  {
	var (
		jobs Job
		arrJob []Job
	)

	tsql := fmt.Sprintf("SELECT KDPKRJAAN, NMPKRJAAN FROM dbo.PEKERJAAN;")

	rows, err := Db.Query(tsql)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&jobs.JobCode, &jobs.JobName); err != nil {
			log.Println(err.Error())
		} else {
			arrJob = append(arrJob, jobs)
		}
	}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"code": 200,
		"data" : arrJob,

	})
}

func GetEducation(c *gin.Context){
	var (
		education Education
		arrEducation []Education
	)

	tsql := fmt.Sprintf("SELECT KODEDIDIK, NAMADIDIK FROM dbo.PENDDKAN;")

	rows, err := Db.Query(tsql)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&education.EducationCode, &education.EducationName); err != nil {
			log.Fatalln(err.Error())
		} else {
			arrEducation = append(arrEducation, education)
		}
	}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "",
		"data": arrEducation,
	})
}