package usecase

import (
	"database/sql"
	"fmt"
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"github.com/JhonatanRealpe/training-tracker/domain/service/team"
	"github.com/JhonatanRealpe/training-tracker/domain/service/training"
	"github.com/JhonatanRealpe/training-tracker/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TeamUseCase interface {
	GetTeam(c *gin.Context) *entity.Response
}

type teamUseCase struct {
	db              *sql.DB
	logger          *log.Logger
	configuration   ConfigurationUseCase
	trainingService training.TrainingService
}

func NewTeamUseCase(db *sql.DB, logger *log.Logger, configuration ConfigurationUseCase, trainingService training.TrainingService) TeamUseCase {
	return &teamUseCase{
		db:              db,
		logger:          logger,
		configuration:   configuration,
		trainingService: trainingService,
	}
}

func (t *teamUseCase) GetTeam(c *gin.Context) *entity.Response {
	response := entity.NewResponse(http.StatusOK, constants.TeamSuccessfully, nil)
	config, err := t.configuration.GetConfiguration()
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}

	trainingCount, err := t.trainingService.GetTrainingCount(config.MinTrainings)
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}

	if trainingCount < config.StartingPlayers {
		response.SetMessageWithError(http.StatusBadRequest, err, fmt.Sprintf(constants.MessageConfig,
			util.IntToStrin(config.MinTrainings), util.IntToStrin(config.StartingPlayers)))
		t.logger.Println(response.Message)
		return response
	}

	totalValuesPerPlayerWeek, err := t.trainingService.GetTotalValuesPerPlayerWeek()
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}

	titularPlayers, err := team.GetTitularPlayers(totalValuesPerPlayerWeek, config)
	if err != nil {
		response.SetMessageWithError(http.StatusInternalServerError, err, "")
		t.logger.Println(err.Error())
		return response
	}
	response.SetData(http.StatusOK, "", titularPlayers)
	return response
}
