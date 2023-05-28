package private

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

const hospitals = `{"hospitals":[{"name":"Городская клиническая больница № 1 им. Н.И. Пирогова","departments":["Хирургическое","Приемное","Реанимационное"]},{"name":"Городская больница № 3","departments":["Хирургическое","Приемное","Реанимационное"]},{"name":"Городская клиническая больница № 4.","departments":["Хирургическое","Приемное","Реанимационное"]}]}`

func Hospitals(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, json.RawMessage(hospitals))
}
