package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/infrastructure"
	"blog/models"
	"os"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter() //router has been initialized and configured
	db := infrastructure.NewDatabase()      // database has been initialized and configured

	// Post APIs setup
	postRepository := repository.NewPostRepository(db)          // repository are being setup
	postService := service.NewPostService(postRepository)       // service are being setup
	postController := controller.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router)    // post routes are initialized
	postRoute.Setup()                                           // post routes are being setup

	// PDF APIs setup
	pdfService := service.NewPdfService()
	pdfController := controller.NewPdfController(pdfService)
	pdfRoute := routes.NewPdfRoute(pdfController, router)
	pdfRoute.Setup()

	// Excel APIs setup
	excelService := service.NewExcelService()
	excelController := controller.NewExcelController(excelService)
	excelRoute := routes.NewExcelRoute(excelController, router)
	excelRoute.Setup()

	db.DB.AutoMigrate(&models.Post{})              // migrating Post model to datbase table
	router.Gin.Run(":" + os.Getenv("SERVER_PORT")) //server started on 8000 port
}
