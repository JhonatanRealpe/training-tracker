package usecase

import (
	"database/sql"
	"fmt"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	configurationUseCase := NewConfigurationUseCase(mockDB, mockLogger)

	expectedConfig := entity.Configuration{
		ID:                         1,
		ShootingPowerPercentage:    20,
		SpeedPercentage:            30,
		SuccessfulPassesPercentage: 50,
		StartingPlayers:            4,
		MinTrainings:               2,
	}

	config, err := configurationUseCase.GetConfiguration()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if config.ID != expectedConfig.ID || config.ShootingPowerPercentage != expectedConfig.ShootingPowerPercentage || config.SpeedPercentage != expectedConfig.SpeedPercentage || config.SuccessfulPassesPercentage != expectedConfig.SuccessfulPassesPercentage || config.StartingPlayers != expectedConfig.StartingPlayers || config.MinTrainings != expectedConfig.MinTrainings {
		t.Errorf("Config does not match expected configuration")
	}
}

func setupLogger(t *testing.T) *log.Logger {
	if testing.Verbose() {
		return log.New(os.Stdout, "test-logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	// Si no está en modo verbose, se devuelve un logger que descarta los logs
	return log.New(ioutil.Discard, "", 0)
}

func setupDB() *sql.DB {
	dsn := fmt.Sprintf("trainingtracker:123456@tcp(localhost:3306)/training_tracker")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	if err := db.Ping(); err != nil {
		return nil
	}

	return db
}
