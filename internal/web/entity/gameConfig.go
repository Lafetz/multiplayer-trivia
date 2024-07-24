package webentities

type GameConfig struct {
	Host     string
	Owner    bool
	Id       string
	Category int
	Timer    int
	Amount   int
}

func NewConfig(host string, owner bool, id string, category int, timer int, amount int) GameConfig {
	return GameConfig{
		Host:     host,
		Owner:    owner,
		Id:       id,
		Category: category,
		Timer:    timer,
		Amount:   amount,
	}
}
