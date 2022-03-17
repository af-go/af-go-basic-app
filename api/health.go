package api

import (
	"github.com/af-go/basic-app/pkg/model"
	"github.com/gin-gonic/gin"
)

type HealthAPIManager struct {
	inservice bool
}

func (h *HealthAPIManager) Build(engine *gin.Engine) {
	engine.GET("/liveness", h.OnLivenessProbe)
	engine.GET("/readiness", h.OnReadinessProbe)
	engine.GET("/inservice", h.OnInServiceProbe)
	engine.POST("/inservice", h.OnInService)
	engine.POST("/outofservice", h.OnOutOfService)
}

// OnLivenessProbe liveness probe
// @Produce json
// @Summary liveness probe
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /liveness [get]
func (h *HealthAPIManager) OnLivenessProbe(gc *gin.Context) {
	statusCode := 200
	resp := model.StatusResponse{Message: "liveness probe success"}
	gc.JSON(statusCode, &resp)
}

// TODO add logic handle dependency health check
// OnReadinessProbe readiness probe
// @Produce json
// @Summary rediness probe
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /readiness [get]
func (h *HealthAPIManager) OnReadinessProbe(gc *gin.Context) {
	statusCode := 200
	resp := model.StatusResponse{Message: "readiness probe success"}
	gc.JSON(statusCode, &resp)
}

// OnInServiceProbe inservice probe
// @Produce text
// @Summary inservice probe
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /inservice [get]
func (h *HealthAPIManager) OnInServiceProbe(gc *gin.Context) {
	statusCode := 200
	resp := "OKOKOK"
	if !h.inservice {
		statusCode = 404
		resp = "NONONO"
	}
	gc.String(statusCode, resp)
}

// OnInService inservice
// @Produce json
// @Summary inservice
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /inservice [post]
func (h *HealthAPIManager) OnInService(gc *gin.Context) {
	h.inservice = true
	statusCode := 200
	resp := model.StatusResponse{Message: "in service now"}
	gc.JSON(statusCode, &resp)
}

// OnOutOfService out of service
// @Produce json
// @Summary out of service
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /outofservice [post]
func (h *HealthAPIManager) OnOutOfService(gc *gin.Context) {
	h.inservice = false
	statusCode := 200
	resp := model.StatusResponse{Message: "out of service now"}
	gc.JSON(statusCode, &resp)
}
