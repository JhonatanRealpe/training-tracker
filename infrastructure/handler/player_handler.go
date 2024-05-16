package handler

import (
	"github.com/JhonatanRealpe/training-tracker/domain/usecase"
	"github.com/gin-gonic/gin"
)

type PlayerHandler interface {
	SavePlayers(c *gin.Context)
	GetPlayers(c *gin.Context)
}

type playerHandler struct {
	pUseCase usecase.PlayerUseCase
}

func NewPLayerHandler(pUseCase usecase.PlayerUseCase) PlayerHandler {
	return &playerHandler{
		pUseCase: pUseCase,
	}
}

// SavePlayers creates new player.
// @Summary Create a new player
// @Description Create a new player in the system
// @Tags Player
// @Accept json
// @Param players body entity.Players true "Players object"
// @Produce json
// @Success 201 {object} entity.Response
// @Router /players [post]
func (p playerHandler) SavePlayers(c *gin.Context) {
	response := p.pUseCase.SavePlayer(c)
	c.JSON(response.Status, response)
}

// GetPlayers obtiene todos los jugadores.
// @Summary Obtener jugadores
// @Description Obtiene todos los jugadores del sistema
// @Tags Player
// @Produce json
// @Success 200 {array} entity.Player
// @Router /players [get]
func (p playerHandler) GetPlayers(c *gin.Context) {
	response := p.pUseCase.GetPlayer()
	c.JSON(response.Status, response)
}
