package entity

import "time"

type Training struct {
	ID               int       `json:"id"`
	Date             time.Time `json:"date"`
	PlayerID         int       `json:"player_id"`
	ShootingPower    int       `json:"shooting_power"`
	Time             int       `json:"time"`
	Distance         int       `json:"distance"`
	SuccessfulPasses int       `json:"successful_passes"`
}

type TotalValuesPerPlayerWeek struct {
	PlayerID              int    `json:"player_id"`
	Name                  string `json:"name"`
	TotalShootingPower    int    `json:"total_shooting_power"`
	TotalTime             int    `json:"total_time"`
	TotalDistance         int    `json:"total_distance"`
	TotalSuccessfulPasses int    `json:"total_successful_passes"`
}
