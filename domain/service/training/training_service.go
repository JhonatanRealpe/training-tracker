package training

import (
	"database/sql"
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"log"
	"time"
)

type TrainingService interface {
	SaveTraining(players []entity.Player) error
	GetTrainingCount(minTrainings int) (int, error)
	GetTotalValuesPerPlayerWeek() ([]entity.TotalValuesPerPlayerWeek, error)
}

type trainingService struct {
	db     *sql.DB
	logger *log.Logger
}

func NewTrainingService(db *sql.DB, logger *log.Logger) TrainingService {
	return &trainingService{
		db:     db,
		logger: logger,
	}
}

func (t *trainingService) SaveTraining(players []entity.Player) error {
	date := time.Now()
	for _, player := range players {
		power := player.Stats.Power
		time := player.Stats.Speed.Time
		distance := player.Stats.Speed.Distance
		passes := player.Stats.Passes

		stmt, err := t.db.Prepare(constants.QueryInsertTraining)
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		_, err = stmt.Exec(date, player.ID, power, time, distance, passes)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *trainingService) GetTrainingCount(minTrainings int) (int, error) {
	stmt, err := t.db.Prepare(constants.QueryCountWeekPlayers)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var trainingCount int
	err = stmt.QueryRow(minTrainings).Scan(&trainingCount)
	if err != nil {
		return 0, err
	}

	return trainingCount, nil
}

func (t *trainingService) GetTotalValuesPerPlayerWeek() ([]entity.TotalValuesPerPlayerWeek, error) {
	stmt, err := t.db.Prepare(constants.QueryTotalValuesPerPlayerWeek)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalValues := []entity.TotalValuesPerPlayerWeek{}
	for rows.Next() {
		var totalValue entity.TotalValuesPerPlayerWeek
		err := rows.Scan(&totalValue.PlayerID, &totalValue.Name, &totalValue.TotalShootingPower, &totalValue.TotalTime, &totalValue.TotalDistance, &totalValue.TotalSuccessfulPasses)
		if err != nil {
			return nil, err
		}
		totalValues = append(totalValues, totalValue)
	}

	return totalValues, nil
}
