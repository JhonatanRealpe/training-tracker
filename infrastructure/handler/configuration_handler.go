package handler

import (
	"database/sql"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
)

func GetConfiguration(db *sql.DB) (entity.Configuration, error) {
	stmt, err := db.Prepare("SELECT * FROM Configuration")
	if err != nil {
		return entity.Configuration{}, err
	}
	defer stmt.Close()

	var config entity.Configuration
	err = stmt.QueryRow().Scan(&config.ID, &config.ShootingPowerPercentage, &config.SpeedPercentage, &config.SuccessfulPassesPercentage, &config.StartingPlayers, &config.MinTrainings)
	if err != nil {
		return entity.Configuration{}, err
	}

	return config, nil
}
