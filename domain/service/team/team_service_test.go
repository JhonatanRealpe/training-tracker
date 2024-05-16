package team

import (
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"testing"
)

func TestGetTitularPlayers(t *testing.T) {
	totalValues := []entity.TotalValuesPerPlayerWeek{
		{
			PlayerID:              1,
			Name:                  "Player 1",
			TotalTime:             10,
			TotalDistance:         200,
			TotalShootingPower:    80,
			TotalSuccessfulPasses: 50,
		},
		{
			PlayerID:              2,
			Name:                  "Player 2",
			TotalTime:             12,
			TotalDistance:         120,
			TotalShootingPower:    90,
			TotalSuccessfulPasses: 40,
		},
		{
			PlayerID:              3,
			Name:                  "Player 3",
			TotalTime:             10,
			TotalDistance:         150,
			TotalShootingPower:    80,
			TotalSuccessfulPasses: 10,
		},
		{
			PlayerID:              4,
			Name:                  "Player 4",
			TotalTime:             12,
			TotalDistance:         250,
			TotalShootingPower:    90,
			TotalSuccessfulPasses: 50,
		},
		{
			PlayerID:              5,
			Name:                  "Player 5",
			TotalTime:             10,
			TotalDistance:         150,
			TotalShootingPower:    80,
			TotalSuccessfulPasses: 60,
		},
	}

	config := entity.Configuration{
		StartingPlayers:            5,
		ShootingPowerPercentage:    70,
		SpeedPercentage:            80,
		SuccessfulPassesPercentage: 90,
		MinTrainings:               3,
	}

	titularPlayers, err := GetTitularPlayers(totalValues, config)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedTitularPlayers := []entity.PlayerResult{
		{ID: 4, Name: "Player 4", Score: 124},
		{ID: 5, Name: "Player 5", Score: 122},
		{ID: 1, Name: "Player 1", Score: 117},
		{ID: 2, Name: "Player 2", Score: 107},
		{ID: 3, Name: "Player 3", Score: 77},
	}

	if len(titularPlayers) != len(expectedTitularPlayers) {
		t.Errorf("expected %d titular players, but got %d", len(expectedTitularPlayers), len(titularPlayers))
	}

	for i, player := range titularPlayers {
		expectedPlayer := expectedTitularPlayers[i]
		if player.ID != expectedPlayer.ID {
			t.Errorf("expected player ID %d, but got %d", expectedPlayer.ID, player.ID)
		}
		if player.Name != expectedPlayer.Name {
			t.Errorf("expected player name %s, but got %s", expectedPlayer.Name, player.Name)
		}
		if player.Score != expectedPlayer.Score {
			t.Errorf("expected player score %f, but got %f", expectedPlayer.Score, player.Score)
		}
	}
}

func TestCalculateSpeed(t *testing.T) {
	tabla := []struct {
		timeInSeconds    int
		distanceInMeters int
		speed            int
	}{
		{2, 4, 2},
		{0, 0, 0},
	}
	for _, item := range tabla {
		speed := CalculateSpeed(item.timeInSeconds, item.distanceInMeters)

		if speed != item.speed {
			t.Errorf("expected speed %v, but got %v ", item.speed, speed)
		}
	}
}
