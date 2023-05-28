package private

import (
	"fmt"
	"medica/sdk/db"
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

	var Records []*db.Record

	rows, err := excelFile.Rows("Sheet1")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while reading file: %w", err))

		return
	}

	first := true

	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while reading row: %w", err))

			return
		}
		if !first {
			Records = append(Records, db.NewRecord(row))
		} else {
			first = false
		}

	}

	ok, err := db.NewCopy(Records)
	if err != nil {
		if !ok {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("while copying data: %w", err))

			return
		}
	}

	if err = rows.Close(); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while closing file: %w", err))

		return
	}
}
