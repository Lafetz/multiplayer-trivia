package webentities

type GameConfig struct {
	Host     string
	Owner    bool
	Id       string
	Catagory int
	Timer    int
	Amount   int
}

func NewConfig(host string, owner bool, id string, catagory int, timer int, amount int) GameConfig {
	return GameConfig{
		Host:     host,
		Owner:    owner,
		Id:       id,
		Catagory: catagory,
		Timer:    timer,
		Amount:   amount,
	}
}
