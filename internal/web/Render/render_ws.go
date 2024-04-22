package render

import (
	"bytes"
	"context"

	"github.com/Lafetz/showdown-trivia-game/internal/core/game"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func RenderPlayers() *bytes.Buffer {
	component := pages.Players()
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
