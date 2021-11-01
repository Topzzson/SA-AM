package main

import (
	"github.com/Topzzson/SA-AM/controller"

	"github.com/Topzzson/SA-AM/entity"

	"github.com/Topzzson/SA-AM/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			// Path_status Routes
			protected.GET("/path_statuses", controller.ListPathStatus)
			protected.GET("/path_status/:id", controller.GetPathStatus)
			protected.POST("/path_statuses", controller.CreatePathStatus)
			protected.PATCH("/path_statuses", controller.UpdatePathStatus)
			protected.DELETE("/path_statuses/:id", controller.DeletePathStatus)

			// Checklist Routes
			protected.GET("/check_lists", controller.ListCheckList)
			protected.GET("/check_list/:id", controller.GetCheckList)
			protected.POST("/check_lists", controller.CreateCheckList)
			protected.PATCH("/check_lists", controller.UpdateCheckList)
			protected.DELETE("/check_lists/:id", controller.DeleteCheckList)

			// Car_path Routes
			protected.GET("/car_paths", controller.ListCarPath)
			protected.GET("/car_path/:id", controller.GetCarPath)
			protected.POST("/car_paths", controller.CreateCarPath)
			protected.PATCH("/car_paths", controller.UpdateCarPath)
			protected.DELETE("/car_paths/:id", controller.DeleteCarPath)

			// AmbulanceType Routes
			protected.GET("/ambulanceTypes", controller.ListAmbulanceType)
			protected.GET("/ambulanceType/:id", controller.GetAmbulanceType)
			protected.POST("/ambulanceTypes", controller.CreateAmbulanceType)
			protected.PATCH("/ambulanceTypes", controller.UpdateAmbulanceType)
			protected.DELETE("/ambulanceTypes/:id", controller.DeleteAmbulanceType)

			// Brand Routes
			protected.GET("/brands", controller.ListBrand)
			protected.GET("/brand/:id", controller.GetBrand)
			protected.POST("/brands", controller.CreateBrand)
			protected.PATCH("/brands", controller.UpdateBrand)
			protected.DELETE("/brands/:id", controller.DeleteBrand)

			// Status Routes
			protected.GET("/statuses", controller.ListStatuses)
			protected.GET("/status/:id", controller.GetStatus)
			protected.POST("/statuses", controller.CreateStatus)
			protected.PATCH("/statuses", controller.UpdateStatus)
			protected.DELETE("/statuses/:id", controller.DeleteStatus)

			//Ambulance Routes
			protected.GET("/ambulances", controller.ListAmbulance)
			protected.GET("/ambulance/:id", controller.GetAmbulance)
			protected.POST("/ambulances", controller.CreateAmbulance)
			protected.PATCH("/ambulances", controller.UpdateAmbulance)
			protected.DELETE("/ambulancers/:id", controller.DeleteAmbulance)

		}
	}

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
