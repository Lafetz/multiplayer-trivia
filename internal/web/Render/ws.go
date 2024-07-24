package render

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	webentities "github.com/Lafetz/showdown-trivia-game/internal/web/entity"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/components"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/a-h/templ"
)

func SendGamePage(w http.ResponseWriter, r *http.Request, gameConfig webentities.GameConfig) error {
	p := components.Game(gameConfig)
	err := layout.Base("Game", p).Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
func RenderPlayers(id string, players []string) (*bytes.Buffer, error) {
	component := components.Players(id, players)
	return returnBuf(component)
}
func RenderQuestion(q entities.Question, current int, total int, timer int, players []*game.Player) (*bytes.Buffer, error) {
	component := components.Question(q, current, total, timer, players)
	return returnBuf(component)
}
func RenderGameMessage(Info game.Info) (*bytes.Buffer, error) {
	component := components.GameMessage(Info.Text)
	return returnBuf(component)
}

func GameEnd(winners game.Winners) (*bytes.Buffer, error) {
	component := components.GameEndMessage(winners)
	return returnBuf(component)
}
func RenderUserAnswer(userAnswer string) (*bytes.Buffer, error) {
	component := components.Answer(userAnswer)
	return returnBuf(component)
}
func WsServerError() (*bytes.Buffer, error) {
	component := components.InfoToast("Internal server error", true)
	return returnBuf(component)
}
func returnBuf(component templ.Component) (*bytes.Buffer, error) {
	buffer := &bytes.Buffer{}
	err := component.Render(context.Background(), buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
