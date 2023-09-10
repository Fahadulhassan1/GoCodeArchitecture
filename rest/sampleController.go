package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goPractice/service"
)

type SampelController interface {
	SampleMessage(ctx *gin.Context)
}
type sampleController struct {
	sampleService service.SampleService
}

func NewSampleController(sampleService service.SampleService) SampelController {
	return &sampleController{
		sampleService: sampleService,
	}
}

func (s *sampleController) SampleMessage(ctx *gin.Context) {
	message, err := s.sampleService.SampleFunc()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewStandardResponse(false, 400, err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, NewStandardResponse(true, 200, "got response", message))
}
