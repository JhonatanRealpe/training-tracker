package entity

type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Stats    Stats  `json:"stats"`
	Position string `json:"position"`
}

type Stats struct {
	Power  string `json:"power"`
	Speed  Speed  `json:"speed"`
	Passes string `json:"passes"`
}

type Speed struct {
	Distance string `json:"distance"`
	Time     string `json:"time"`
}

type PlayerResult struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type Players struct {
	Players *[]Player `json:"players"`
}

func (p Player) GetPosition() string {
	if p.Position == "" {
		return "0"
	}
	return p.Position
}
