package player

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

func TestInsertPlayer(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	player := entity.Player{
		ID:   199,
		Name: "Player1",
		Stats: entity.Stats{
			Power: "10",
			Speed: entity.Speed{
				Distance: "50",
				Time:     "5",
			},
			Passes: "20",
		},
	}

	err := playerService.InsertPlayer(player)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestUpdatePlayer(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	player := entity.Player{
		ID:   199,
		Name: "Player1",
		Stats: entity.Stats{
			Power: "15",
			Speed: entity.Speed{
				Distance: "60",
				Time:     "6",
			},
			Passes: "25",
		},
	}

	err := playerService.UpdatePlayer(player)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestDeletePlayer(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	playerID := 199

	err := playerService.DeletePlayer(playerID)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestValidatePlayersData(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	player := entity.Player{
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
	}

	err := playerService.ValidatePlayersData(player)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestPlayerExists(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	exists := playerService.PlayerExists(1)

	if !exists {
		t.Errorf("Expected player to exist, but it does not")
	}
}

func TestGetPlayers(t *testing.T) {
	mockDB := setupDB()
	defer mockDB.Close()
	mockLogger := setupLogger(t)
	playerService := NewPlayerService(mockDB, mockLogger)

	players, err := playerService.GetPlayers()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(players) == 0 {
		t.Errorf("Expected non-empty list of players, but got empty list")
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
