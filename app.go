package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)


func main() {



	router := gin.Default()

	//fmt.println("imported gin")

	router.GET("/poli", controllers.GetPoli)
	router.GET("/no_patient", controllers.GetNoPatient)
	router.GET("/job", controllers.GetJob)
	router.GET("/education", controllers.GetEducation)
	router.GET("/poli_schedule", controllers.GetSchedule)

	router.Run(":1000")
}
