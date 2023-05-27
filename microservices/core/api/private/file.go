package private

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while working on file: %w", err))

		return
	}

	xlsxFile, err := file.Open()
	excelFile, err := excelize.OpenReader(xlsxFile)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while working on file: %w", err))

		return
	}

	rows, err := excelFile.Rows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	if err = rows.Close(); err != nil {
		fmt.Println(err)
	}
}
