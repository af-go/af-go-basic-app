package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"

	"github.com/af-go/basic-app/pkg/model"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// NewRestfulController create new controller
func NewRestfulController(options model.RestServiceOptions) *Controller {
	return &Controller{options: options}
}

// Controller api controller
type Controller struct {
	options model.RestServiceOptions
	server  *http.Server
}

// Start start http controller
func (c *Controller) Start(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	port := c.options.Port
	if port == 0 {
		port = 8080
	}
	r := BuildEngine()
	r.GET("/healthz", c.Healthz)
	pprof.Register(r)
	c.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	go func() {
		if err := c.server.ListenAndServe(); err != http.ErrServerClosed {
			logging.WithError(err).Error("Failed to start API controller")
		}
	}()
	logging.Infof("API controller is listening on port %d", port)
}

// Stop stop API controller
func (c *Controller) Stop(ctx context.Context) {
	logging.Infof("shutting down API controller at %v", time.Now())
	_, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer func() {
		cancel()
	}()
	if err := c.server.Shutdown(ctx); err != nil {
		logging.WithError(err).Error("failed to shut down API controller gracefully")
	}
	logging.Infof("API controller is shutdown at %v", time.Now())
}

// Healthz health check api
// @Produce json
// @Summary health check
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /healthz [get]
func (c *Controller) Healthz(gc *gin.Context) {
	statusCode := 200
	resp := model.StatusResponse{}
	gc.JSON(statusCode, &resp)
}

func NewError(gc *gin.Context, status int, err error) {
	er := model.HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	gc.JSON(status, er)
}
