package application

import (
	"database/sql"
	"github.com/JhonatanRealpe/training-tracker/database"
	"github.com/JhonatanRealpe/training-tracker/domain/service/player"
	"github.com/JhonatanRealpe/training-tracker/domain/service/training"
	"github.com/JhonatanRealpe/training-tracker/domain/usecase"
	"github.com/JhonatanRealpe/training-tracker/infrastructure/handler"
	"log"
	"os"
)

type appDependencies struct {
	db              *sql.DB
	trainingHandler handler.TrainingHandler
	teamHandler     handler.TeamHandler
	playersHandler  handler.PlayerHandler
}

func loadAppDependencies() *appDependencies {
	logger := log.New(os.Stdout, "ERROR", log.LstdFlags|log.Lshortfile)
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	pService := player.NewPlayerService(db, logger)
	pUseCase := usecase.NewPlayerUseCase(logger, pService)
	pHandler := handler.NewPLayerHandler(pUseCase)

	tService := training.NewTrainingService(db, logger)
	tUseCase := usecase.NewTrainingUseCase(logger, pService, tService)
	tHandler := handler.NewTrainingHandler(tUseCase)

	cUseCase := usecase.NewConfigurationUseCase(db, logger)

	teamUseCase := usecase.NewTeamUseCase(db, logger, cUseCase, tService)
	teamHandler := handler.NewTeamHandler(teamUseCase)

	return &appDependencies{
		db:              db,
		trainingHandler: tHandler,
		teamHandler:     teamHandler,
		playersHandler:  pHandler,
	}
}
