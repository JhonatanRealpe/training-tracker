package entity

type Configuration struct {
	ID                         int `json:"id"`
	ShootingPowerPercentage    int `json:"shooting_power_percentage"`
	SpeedPercentage            int `json:"speed_percentage"`
	SuccessfulPassesPercentage int `json:"successful_passes_percentage"`
	StartingPlayers            int `json:"starting_players"`
	MinTrainings               int `json:"min_trainings"`
}
