package controller

import (
	"blog/api/service"
	"blog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExcelController struct {
	service service.ExcelService
}

func NewExcelController(s service.ExcelService) ExcelController {
	return ExcelController{
		service: s,
	}
}

func (e ExcelController) WriteExcel(ctx *gin.Context) {
	err := e.service.WriteExcel()

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Generate Excel")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Generate Excel",
		Data:    "",
	})
}
