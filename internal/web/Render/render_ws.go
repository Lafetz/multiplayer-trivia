package render

import (
	"bytes"
	"context"

	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func RenderPlayers(id string, players []string) *bytes.Buffer {
	component := pages.Players(id, players)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
func RenderQuestion(q game.Question) *bytes.Buffer {
	component := pages.Question(q)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
func RenderGameMessage(Info game.Info) *bytes.Buffer {
	component := pages.GameMessage(Info.Text)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}

func GameEnd(winners game.Winners) *bytes.Buffer {
	component := pages.GameEndMessage(winners)
	buffer := &bytes.Buffer{}
	component.Render(context.Background(), buffer)
	return buffer
}
