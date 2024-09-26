package routes

import (
	"bai2/controllers"

	"github.com/gin-gonic/gin"
)

func Registeroutes(routes *gin.Engine) {
	api := routes.Group("/api/students")
	{
		api.GET("/", controllers.GetStudents)
		api.GET("/:id", controllers.GetStudentById)
		api.POST("/", controllers.AddStudent)
		api.PUT("/:id", controllers.UpdateStudent)
		api.DELETE("/:id", controllers.DeleteStudent)
		api.GET("/search", controllers.SearchStudent)
	}
}
