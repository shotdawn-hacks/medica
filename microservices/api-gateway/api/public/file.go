package public

import (
	"bytes"
	"fmt"
	"io"
	"medica/sdk/destination"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("fetching file: %w", err))

		return
	}

	coreDst, ok := ctx.Get("dst-core")
	if !ok {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("no core destination"))

		return
	}
	core := coreDst.(*destination.Destination)

	// New multipart writer.
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while creating fileform: %w", err))

		return
	}

	fileToUpload, err := file.Open()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while opening file: %w", err))

		return
	}

	_, err = io.Copy(fw, fileToUpload)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while copying %w", err))

		return
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s:%s/upload", core.Config.Address, core.Config.Port), bytes.NewReader(body.Bytes()))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := core.Base.Post(req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("while posting: %w", err))

		return
	}

	ctx.Status(resp.StatusCode)
}
