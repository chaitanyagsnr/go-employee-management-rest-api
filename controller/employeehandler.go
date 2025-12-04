package controller

import (
	"employee-management-service/dto"
	"employee-management-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetMe(context *gin.Context) {
	context.String(http.StatusOK, "Hello, Employee Management Service!")
}

func ListEmployeesHandler(context *gin.Context) {
	employees := service.ListEmployees()
	context.JSON(http.StatusOK, employees)
}

func AddEmployeeHandler(context *gin.Context) {
	var employee dto.Employee
	context.BindJSON(&employee)
	service.AddEmployee(employee)
	context.Status(http.StatusCreated)
}
