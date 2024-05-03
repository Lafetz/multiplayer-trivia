package render

import (
	"bytes"
	"context"
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/components"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
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
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
func RenderQuestion(q game.Question, current int, total int) *bytes.Buffer {
	component := components.Question(q, current, total)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
func RenderGameMessage(Info game.Info) *bytes.Buffer {
	component := components.GameMessage(Info.Text)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}

func GameEnd(winners game.Winners) *bytes.Buffer {
	component := components.GameEndMessage(winners)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
func RenderUserAnswer(userAnswer string) *bytes.Buffer {
	component := components.Answer(userAnswer)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
