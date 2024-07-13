package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Lafetz/showdown-trivia-game/internal/core/question"
	render "github.com/Lafetz/showdown-trivia-game/internal/web/Render"
	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	"github.com/Lafetz/showdown-trivia-game/internal/web/ws"
)

type key string

const (
	UsernameKey key = "username"
)

func Home(logger *slog.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		username, ok := r.Context().Value(UsernameKey).(string)
		if !ok {
			ServerError(w, r, errors.New("username couldn't"), logger)
			return
		}

		err := render.Home(w, r, username)
		if err != nil {
			ServerError(w, r, err, logger)
		}

	}
}
func sendGamePage(w http.ResponseWriter, r *http.Request, wsUrl string, create bool, id string, catagory int, timer int, amount int) error {
	return render.SendGamePage(w, r, wsUrl, create, id, catagory, timer, amount)

}

func CreateFormGet(logger *slog.Logger, questionService question.QuestionServiceApi) http.HandlerFunc { // sends create game form
	return func(w http.ResponseWriter, r *http.Request) {
		cat, err := questionService.GetCategories()
		if err != nil {
			ServerError(w, r, err, logger)
		}
		err = render.CreateGameForm(w, r, form.NewGame{}, cat)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func CreateFormPost(logger *slog.Logger, questionService question.QuestionServiceApi, wsUrl string) http.HandlerFunc { // sends ws component with /wscreate
	return func(w http.ResponseWriter, r *http.Request) {
		cat, err := questionService.GetCategories()
		if err != nil {
			ServerError(w, r, err, logger)
		}
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err = r.ParseForm()
		if err != nil {
			ServerError(w, r, err, logger)
			return
		}
		form := form.NewGame{
			Category: r.PostForm.Get("category"),
			Timer:    r.PostForm.Get("timer"),
			Amount:   r.PostForm.Get("amount"),
		}
		if !form.Valid() {
			if err := render.InvliadCreateGameForm(w, r, form, cat); err != nil {
				ServerError(w, r, err, logger)
			}
			return
		}
		category, _ := strconv.Atoi(form.Category)
		time, _ := strconv.Atoi(form.Timer)
		amount, _ := strconv.Atoi(form.Amount)
		gameOwner := true
		id := ""
		err = sendGamePage(w, r, wsUrl, gameOwner, id, category, time, amount)
		if err != nil {
			ServerError(w, r, err, logger)
		}

	}
}

func ActiveGames(hub *ws.Hub, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms := hub.ListRooms()
		err := render.ActiveGames(w, r, rooms)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func Join(logger *slog.Logger, wsUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameOwner := false
		id := r.PathValue("id")
		err := sendGamePage(w, r, wsUrl, gameOwner, id, 0, 0, 0)
		if err != nil {
			ServerError(w, r, err, logger)
		}
	}
}
func CreateWs(hub *ws.Hub, questionService question.QuestionServiceApi) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category, err := strconv.Atoi(r.URL.Query().Get("category"))
		if err != nil {
			ServerError(w, r, err, hub.Logger)
			return
		}
		timer, err := strconv.Atoi(r.URL.Query().Get("timer"))
		if err != nil {
			ServerError(w, r, err, hub.Logger)
			return
		}
		amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
		if err != nil {
			ServerError(w, r, err, hub.Logger)
			return
		}
		username, ok := r.Context().Value(UsernameKey).(string)
		if !ok {
			ServerError(w, r, errors.New("username couldn't be found"), hub.Logger)
			return
		}
		questions, err := questionService.GetQuestions(amount, category)
		if err != nil {
			ServerError(w, r, err, hub.Logger)
			return
		}
		hub.CreateRoom(w, r, username, timer, questions)
	}
}
func JoinWs(hub *ws.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(UsernameKey).(string)
		if !ok {
			ServerError(w, r, errors.New("username couldn't be found"), hub.Logger)
			return
		}
		id := r.PathValue("id")
		hub.JoinRoom(w, r, id, username)
	}
}
