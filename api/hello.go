package api

import (
	"fmt"
	"os"

	"github.com/af-go/basic-app/pkg/model"
	"github.com/gin-gonic/gin"
)

type HelloAPIManager struct {
}

func (h *HelloAPIManager) Build(engine *gin.Engine) {
	engine.GET("/ping", h.OnPing)
}

// OnPing ping
// @Produce json
// @Summary ping
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 500 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /ping [get]
func (h *HelloAPIManager) OnPing(gc *gin.Context) {
	statusCode := 200
	name, err := os.Hostname()
	var resp model.StatusResponse
	if err != nil {
		statusCode = 500
		resp = model.StatusResponse{Message: fmt.Sprintf("internal server error %v", err)}
	} else {
		resp = model.StatusResponse{Message: fmt.Sprintf("pong from %s", name)}
	}
	gc.JSON(statusCode, &resp)
}
