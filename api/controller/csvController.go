package controller

import (
	"bytes"
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

type CsvController struct {
}

func NewCsvController() CsvController {
	return CsvController{}
}

func (c CsvController) WriteCsv(ctx *gin.Context) {
	//var csvStruct [][]string

	var csvStruct = [][]string{
		{"名前", "住所", "電話"},
		{"Ram", "Tokyo", "1236524"},
		{"Shay", "Beijing", "8575675484"},
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(csvStruct)
	ctx.Writer.Write(b.Bytes())
}
