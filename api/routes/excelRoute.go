package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for question module
type ExcelRoute struct {
	Controller controller.ExcelController
	Handler    infrastructure.GinRouter
}

//NewPostRoute -> initializes new choice routes
func NewExcelRoute(
	controller controller.ExcelController,
	handler infrastructure.GinRouter,

) ExcelRoute {
	return ExcelRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (e ExcelRoute) Setup() {
	excel := e.Handler.Gin.Group("/excel") //Router group
	{
		excel.GET("/write-excel", e.Controller.WriteExcel)
	}
}
