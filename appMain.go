package main

import (
	"employee-management-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/greetMe", controller.GreetMe)
	router.GET("/listAllEmployees", controller.ListEmployeesHandler)
	router.POST("/addEmployee", controller.AddEmployeeHandler)
	router.Run(":8080")
}
