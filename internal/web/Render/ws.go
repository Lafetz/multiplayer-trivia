package render

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/components"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/a-h/templ"
)

func SendGamePage(w http.ResponseWriter, r *http.Request, create bool, id string) {
	p := components.Game(create, id)
	err := layout.Base("Game", p).Render(r.Context(), w)
	if err != nil {
		InternalServer(w, r)
	}
}
func RenderPlayers(id string, players []string) *bytes.Buffer {
	component := components.Players(id, players)
	return returnBuf(component)
}
func RenderQuestion(q game.Question, current int, total int) *bytes.Buffer {
	component := components.Question(q, current, total)
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
func returnBuf(component templ.Component) *bytes.Buffer {
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
