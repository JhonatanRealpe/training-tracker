package handler

import (
	_ "github.com/JhonatanRealpe/training-tracker/cmd/api/docs"
	"github.com/JhonatanRealpe/training-tracker/domain/usecase"
	"github.com/gin-gonic/gin"
)

type TrainingHandler interface {
	SaveTraining(c *gin.Context)
}

type trainingHandler struct {
	tUseCase usecase.TrainingUseCase
}

func NewTrainingHandler(tUseCase usecase.TrainingUseCase) TrainingHandler {
	return &trainingHandler{
		tUseCase: tUseCase,
	}
}

// SaveTraining creates a new training.
// @Summary Create a new training
// @Description Create a new training in the system
// @Tags Training
// @Accept json
// @Param players body entity.Players true "Players object"
// @Produce json
// @Success 201 {object} entity.Response
// @Router /training [post]
func (t trainingHandler) SaveTraining(c *gin.Context) {
	response := t.tUseCase.SaveTraining(c)
	c.JSON(response.Status, response)
}
