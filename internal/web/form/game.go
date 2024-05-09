package form

import "strconv"

type NewGame struct {
	Category string `json:"category"`
	Timer    string `json:"timer"`
	Amount   string `json:"amount"`
	Errors   map[string]string
}

func (f *NewGame) Valid() bool {

	f.Errors = make(map[string]string)
	category, err := strconv.Atoi(f.Category)
	if err != nil {
		f.Errors["category"] = "needs to be number"
	}
	time, err := strconv.Atoi(f.Timer)
	if err != nil {
		f.Errors["timer"] = "needs to be number"
	}
	amount, err := strconv.Atoi(f.Amount)
	if err != nil {
		f.Errors["amount"] = "needs to be number"

	}

	if category < 0 || category > 32 {
		f.Errors["category"] = "category can't be less 0 or greater than 10"
	}
	if amount < 1 || amount > 50 {
		f.Errors["amount"] = "number of questions can't be less than 1 or greater than 50"
	}
	if time < 2 || time > 20 {
		f.Errors["timer"] = "timer can't be less than 2 or greater than 20"
	}
	return len(f.Errors) == 0
}
