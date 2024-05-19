package render

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/entities"
	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/components"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/a-h/templ"
)

func SendGamePage(w http.ResponseWriter, r *http.Request, wsUrl string, create bool, id string, catagory int, timer int, amount int) error {
	p := components.Game(wsUrl, create, id, catagory, timer, amount)
	err := layout.Base("Game", p).Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
func RenderPlayers(id string, players []string) *bytes.Buffer {
	component := components.Players(id, players)
	return returnBuf(component)
}
func RenderQuestion(q entities.Question, current int, total int, timer int, players []*game.Player) *bytes.Buffer {
	component := components.Question(q, current, total, timer, players)
	return returnBuf(component)
}
func RenderGameMessage(Info game.Info) *bytes.Buffer {
	component := components.GameMessage(Info.Text)
	return returnBuf(component)
}

func GameEnd(winners game.Winners) *bytes.Buffer {
	component := components.GameEndMessage(winners)
	return returnBuf(component)
}
func RenderUserAnswer(userAnswer string) *bytes.Buffer {
	component := components.Answer(userAnswer)
	return returnBuf(component)
}
func WsServerError() *bytes.Buffer {
	component := components.InfoToast("Internal server error", true)
	return returnBuf(component)
}
func returnBuf(component templ.Component) *bytes.Buffer {
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
