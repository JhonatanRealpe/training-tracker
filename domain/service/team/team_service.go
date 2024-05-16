package team

import (
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"sort"
)

func GetTitularPlayers(totalValues []entity.TotalValuesPerPlayerWeek, config entity.Configuration) ([]entity.PlayerResult, error) {
	var results []struct {
		totalValue entity.TotalValuesPerPlayerWeek
		result     float64
	}
	for _, totalValue := range totalValues {
		totalResult := CalculateScore(totalValue, config)
		results = append(results, struct {
			totalValue entity.TotalValuesPerPlayerWeek
			result     float64
		}{
			totalValue: totalValue,
			result:     totalResult,
		})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].result > results[j].result
	})

	titularPlayers := []entity.PlayerResult{}
	for i := 0; i < config.StartingPlayers && i < len(results); i++ {
		playerResult := entity.PlayerResult{
			ID:    results[i].totalValue.PlayerID,
			Name:  results[i].totalValue.Name,
			Score: results[i].result,
		}
		titularPlayers = append(titularPlayers, playerResult)
	}

	return titularPlayers, nil
}

func CalculateScore(t entity.TotalValuesPerPlayerWeek, config entity.Configuration) float64 {
	speedPercentage := CalculateSpeed(t.TotalTime, t.TotalDistance)

	powerResult := float64(config.ShootingPowerPercentage) * float64(t.TotalShootingPower) / 100
	speedResult := float64(config.SpeedPercentage) * float64(speedPercentage) / 100
	passesResult := float64(config.SuccessfulPassesPercentage * t.TotalSuccessfulPasses / 100)

	totalResult := powerResult + speedResult + passesResult

	return totalResult
}

func CalculateSpeed(timeInSeconds int, distanceInMeters int) int {
	if timeInSeconds <= 0 || distanceInMeters <= 0 {
		return 0
	}
	speed := distanceInMeters / timeInSeconds
	return speed
}
