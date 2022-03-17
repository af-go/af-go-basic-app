package api

import (
	"fmt"
	"os"

	"github.com/af-go/basic-app/pkg/model"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

type HelloAPIManager struct {
}

func (h *HelloAPIManager) Build(engine *gin.Engine) {
	logging.Info("build hello")
	engine.GET("/ping", h.OnPing)
}

// OnPing ping
// @Produce json
// @Summary ping
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /ping [get]
func (h *HelloAPIManager) OnPing(gc *gin.Context) {
	podIP, present := os.LookupEnv("POD_ID")
	if !present {
		podIP = "127.0.0.1"
	}

	statusCode := 200
	resp := model.StatusResponse{Message: fmt.Sprintf("pong from %s", podIP)}
	gc.JSON(statusCode, &resp)
}
