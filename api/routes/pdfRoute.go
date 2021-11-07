package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for question module
type PdfRoute struct {
	Controller controller.PdfController
	Handler    infrastructure.GinRouter
}

//NewPostRoute -> initializes new choice rouets
func NewPdfRoute(
	controller controller.PdfController,
	handler infrastructure.GinRouter,

) PdfRoute {
	return PdfRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p PdfRoute) Setup() {
	pdf := p.Handler.Gin.Group("/pdf") //Router group
	{
		pdf.GET("/", p.Controller.Index)
	}
}
