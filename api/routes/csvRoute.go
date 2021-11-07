package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for question module
type CsvRoute struct {
	Controller controller.CsvController
	Handler    infrastructure.GinRouter
}

//NewPostRoute -> initializes new choice routes
func NewCsvRoute(
	controller controller.CsvController,
	handler infrastructure.GinRouter,

) CsvRoute {
	return CsvRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (c CsvRoute) Setup() {
	csv := c.Handler.Gin.Group("/csv") //Router group
	{
		csv.GET("/download-csv", c.Controller.WriteCsv)
	}
}
