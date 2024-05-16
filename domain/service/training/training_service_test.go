package training

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

func TestSaveTraining(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	trainingService := NewTrainingService(mockDB, mockLogger)

	players := []entity.Player{
		{
			ID:   1,
			Name: "Player1",
			Stats: entity.Stats{
				Power: "10",
				Speed: entity.Speed{
					Distance: "50",
					Time:     "5",
				},
				Passes: "20",
			},
		},
		{
			ID:   2,
			Name: "Player2",
			Stats: entity.Stats{
				Power: "15",
				Speed: entity.Speed{
					Distance: "60",
					Time:     "6",
				},
				Passes: "25",
			},
		},
	}

	err := trainingService.SaveTraining(players)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestGetTrainingCount(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	trainingService := NewTrainingService(mockDB, mockLogger)

	minTrainings := 50

	count, err := trainingService.GetTrainingCount(minTrainings)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if count != 0 {
		t.Errorf("Expected training count to be 0, but got %d", count)
	}
}

func TestGetTotalValuesPerPlayerWeek(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	trainingService := NewTrainingService(mockDB, mockLogger)

	totalValues, err := trainingService.GetTotalValuesPerPlayerWeek()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(totalValues) == 0 {
		t.Errorf("Expected non-empty list of total values, but got empty list")
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
