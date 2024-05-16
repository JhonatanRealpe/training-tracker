package handler

import (
	"github.com/JhonatanRealpe/training-tracker/domain/usecase"
	"github.com/gin-gonic/gin"
)

type TeamHandler interface {
	GetTeam(c *gin.Context)
}

type teamHandler struct {
	tUseCase usecase.TeamUseCase
}

func NewTeamHandler(tUseCase usecase.TeamUseCase) TeamHandler {
	return &teamHandler{
		tUseCase: tUseCase,
	}
}

// GetTeam retrieves team information.
// @Summary Retrieve team information
// @Description Get information about a specific team
// @Tags Team
// @Accept json
// @Produce json
// @Success 200 {object} entity.Response
// @Router /team [get]
func (t teamHandler) GetTeam(c *gin.Context) {
	response := t.tUseCase.GetTeam(c)
	c.JSON(response.Status, response)
}
