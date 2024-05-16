package usecase

import (
	"database/sql"
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"log"
)

type ConfigurationUseCase interface {
	GetConfiguration() (entity.Configuration, error)
}

type configurationUseCase struct {
	db     *sql.DB
	logger *log.Logger
}

func NewConfigurationUseCase(db *sql.DB, logger *log.Logger) ConfigurationUseCase {
	return &configurationUseCase{
		db:     db,
		logger: logger,
	}
}

func (c configurationUseCase) GetConfiguration() (entity.Configuration, error) {
	stmt, err := c.db.Prepare(constants.QueryGetConfiguration)
	if err != nil {
		return entity.Configuration{}, err
	}
	defer stmt.Close()

	var config entity.Configuration
	err = stmt.QueryRow().Scan(&config.ID, &config.ShootingPowerPercentage, &config.SpeedPercentage, &config.SuccessfulPassesPercentage, &config.StartingPlayers, &config.MinTrainings)
	if err != nil {
		return entity.Configuration{}, err
		c.logger.Println(err.Error())
	}

	return config, nil
}
