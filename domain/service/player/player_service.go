package player

import (
	"database/sql"
	"fmt"
	"github.com/JhonatanRealpe/training-tracker/domain/constants"
	"github.com/JhonatanRealpe/training-tracker/domain/entity"
	"github.com/JhonatanRealpe/training-tracker/util"
	"log"
)

type PlayerService interface {
	InsertPlayer(player entity.Player) error
	UpdatePlayer(player entity.Player) error
	DeletePlayer(playerID int) error
	GetPlayers() ([]entity.Player, error)
	ValidatePlayersData(player entity.Player) error
	PlayerExists(playerID int) bool
}

type playerService struct {
	db     *sql.DB
	logger *log.Logger
}

func NewPlayerService(db *sql.DB, logger *log.Logger) PlayerService {
	return &playerService{
		db:     db,
		logger: logger,
	}
}

func (p playerService) InsertPlayer(player entity.Player) error {
	player.GetPosition()
	query := constants.QueryInsertPlayer
	_, err := p.db.Exec(query, player.ID, player.Name, player.Stats.Power, player.Stats.Speed.Distance, player.Stats.Speed.Time, player.Stats.Passes, player.GetPosition())
	if err != nil {
		return err
	}
	return nil
}

func (p playerService) UpdatePlayer(player entity.Player) error {
	query := constants.QueryUpdatePlayer
	_, err := p.db.Exec(query, player.Name, player.Stats.Power, player.Stats.Speed.Distance, player.Stats.Speed.Time, player.Stats.Passes, player.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p playerService) DeletePlayer(playerID int) error {
	query := constants.QueryDeletePlayerByID
	_, err := p.db.Exec(query, playerID)
	if err != nil {
		return err
	}
	return nil
}

func (p playerService) GetPlayers() ([]entity.Player, error) {
	query := constants.QueryGetAllPlayers
	var players []entity.Player

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var player entity.Player
		err := rows.Scan(&player.ID, &player.Name, &player.Stats.Power, &player.Stats.Speed.Distance, &player.Stats.Speed.Time, &player.Stats.Passes, &player.Position)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

func (p playerService) ValidatePlayersData(player entity.Player) error {
	if player.ID == 0 || player.Name == "" || player.Stats.Power == "" || player.Stats.Speed.Time == "" || player.Stats.Speed.Distance == "" || player.Stats.Passes == "" {
		return fmt.Errorf(constants.InvalidPlayer, player.Name)
	}
	passes, _ := util.StrinToInt(player.Stats.Passes)
	power, _ := util.StrinToInt(player.Stats.Power)
	distance, _ := util.StrinToInt(player.Stats.Speed.Distance)
	time, _ := util.StrinToInt(player.Stats.Speed.Time)
	if !p.PlayerExists(player.ID) || passes <= 0 || power <= 0 || time <= 0 || distance < 0 {
		return fmt.Errorf(constants.InvalidPlayer, player.Name)
	}
	return nil
}

func (p playerService) PlayerExists(playerID int) bool {
	query := constants.QueryCountPlayerByID
	var count int

	err := p.db.QueryRow(query, playerID).Scan(&count)
	if err != nil {
		p.logger.Println(err.Error())
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
