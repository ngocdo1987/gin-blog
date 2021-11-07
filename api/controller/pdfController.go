package controller

import (
	"blog/api/service"
	"blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PdfController struct {
	service service.PdfService
}

func NewPdfController(s service.PdfService) PdfController {
	return PdfController{
		service: s,
	}
}

func (p PdfController) Index(ctx *gin.Context) {
	err := p.service.GeneratePdf()

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Generate PDF")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Generate PDF",
		Data:    "",
	})
}
