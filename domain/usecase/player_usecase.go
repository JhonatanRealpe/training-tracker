package usecase

import (
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"github.com/JhonatanRealpe/training-tracker/domain/service/player"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PlayerUseCase interface {
	SavePlayer(c *gin.Context) *entity.Response
	GetPlayer() *entity.Response
}

type playerUseCase struct {
	logger        *log.Logger
	playerService player.PlayerService
}

func NewPlayerUseCase(logger *log.Logger, playerService player.PlayerService) PlayerUseCase {
	return &playerUseCase{
		logger:        logger,
		playerService: playerService,
	}
}

func (p *playerUseCase) SavePlayer(c *gin.Context) *entity.Response {
	response := entity.NewResponse(http.StatusCreated, constants.TrainingSuccessfully, nil)
	var training struct {
		Players []entity.Player
	}

	err := c.BindJSON(&training)
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		p.logger.Println(err.Error())
		return response
	}
	for _, player := range training.Players {
		if p.playerService.PlayerExists(player.ID) {
			p.playerService.UpdatePlayer(player)
		} else {
			p.playerService.InsertPlayer(player)
		}
	}
	return response
}

func (p *playerUseCase) GetPlayer() *entity.Response {
	response := entity.NewResponse(http.StatusOK, constants.TrainingSuccessfully, nil)
	Players, err := p.playerService.GetPlayers()
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		p.logger.Println(err.Error())
		return response
	}
	response.SetData(http.StatusOK, "", Players)
	return response
}
