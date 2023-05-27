package private

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while working on file: %w", err))

		return
	}
	csvFile, err := file.Open()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while working on file: %w", err))

		return
	}

	for line := range csvLines {
		fmt.Println(line)
	}
}
