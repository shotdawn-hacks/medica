package private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"medica/sdk/db"
	"medica/sdk/destination"
	"medica/sdk/run"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
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

	analyzerDst, ok := ctx.Get("dst-analyzer")
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("no core destination"))

		return
	}

	analyzer := analyzerDst.(*destination.Destination)

	go SendRecordsWaitAndFlush(analyzer, Records)

	if err = rows.Close(); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while closing file: %w", err))

		return
	}
}

type AnalyzerRecord struct {
	ID           string `json:"_id"`
	ICD          string `json:"icd"`
	Prescription string `json:"prescription"`
}

func RecordsToAnalyzerRecords(rec []*db.Record) []*AnalyzerRecord {
	var AnalyzerRecords []*AnalyzerRecord

	for _, r := range rec {
		AnalyzerRecords = append(AnalyzerRecords,
			&AnalyzerRecord{
				ID:           r.ID,
				ICD:          r.ICD,
				Prescription: r.Prescription})
	}

	return AnalyzerRecords
}

func SendRecordsWaitAndFlush(analyzer *destination.Destination, records []*db.Record) {
	jsonBody, err := json.Marshal(RecordsToAnalyzerRecords(records))
	if err != nil {
		run.Logger.Error("while encoding records_to_analyzer_records:", zap.Error(err))
	}

	analyzerURL := fmt.Sprintf("http://%s:%s/classify", analyzer.Config.Address, analyzer.Config.Port)
	fmt.Println(analyzerURL)
	req, err := http.NewRequest("POST", analyzerURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	cli := http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		run.Logger.Error("", zap.Error(err))
	}

	fmt.Println(resp.StatusCode)

	var destCfg string

	json.NewDecoder(resp.Body).Decode(&destCfg)

	fmt.Println(destCfg)
}
