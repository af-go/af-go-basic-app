package api

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/af-go/basic-app/pkg/model"
	"github.com/gin-gonic/gin"
)

var memoryBlocks = []string{}

func BuildEngine() *gin.Engine {
	engine := gin.Default()
	engine.POST("/memory", onMemoryBlockCreate)
	engine.DELETE("/memory", onMemoryBlockDeleted)
	return engine
}

// Memory create memory blocks
// @Produce json
// @Summary create memoey blocks
// @Description check status
// @Success 201 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /memory [post]
func onMemoryBlockCreate(gc *gin.Context) {
	memoryBlocks = append(memoryBlocks, strings.Repeat("HelloWorld", 10000000))
	runtime.GC()
	statusCode := 201
	resp := model.StatusResponse{Message: fmt.Sprintf("memory blocks %d", len(memoryBlocks))}
	gc.JSON(statusCode, &resp)
}

// Memory delete memory blocks
// @Produce json
// @Summary delete memoey blocks
// @Description check status
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} model.HTTPError
// @Failure 503 {object} model.HTTPError
// @Router /memory [delete]
func onMemoryBlockDeleted(gc *gin.Context) {
	memoryBlocks = []string{}
	statusCode := 200
	runtime.GC()
	resp := model.StatusResponse{Message: fmt.Sprintf("memory blocks %d", len(memoryBlocks))}
	gc.JSON(statusCode, &resp)
}
