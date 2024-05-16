package usecase

import (
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"github.com/JhonatanRealpe/training-tracker/domain/service/player"
	"github.com/JhonatanRealpe/training-tracker/domain/service/training"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TrainingUseCase interface {
	SaveTraining(c *gin.Context) *entity.Response
}

type trainingUseCase struct {
	logger          *log.Logger
	playerService   player.PlayerService
	trainingService training.TrainingService
}

func NewTrainingUseCase(logger *log.Logger, playerService player.PlayerService, trainingService training.TrainingService) TrainingUseCase {
	return &trainingUseCase{
		logger:          logger,
		playerService:   playerService,
		trainingService: trainingService,
	}
}

func (t *trainingUseCase) SaveTraining(c *gin.Context) *entity.Response {
	response := entity.NewResponse(http.StatusCreated, constants.PLayersSuccessfully, nil)
	var training struct {
		Players []entity.Player
	}

	err := c.BindJSON(&training)
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}

	for _, player := range training.Players {
		err := t.playerService.ValidatePlayersData(player)
		if err != nil {
			response.SetMessageWithError(http.StatusBadRequest, err, "")
			t.logger.Println(err.Error())
			return response
		}
	}

	err = t.trainingService.SaveTraining(training.Players)
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}

	return response
}
