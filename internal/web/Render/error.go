package render

import (
	"net/http"

	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func InternalServer(w http.ResponseWriter, r *http.Request) error {
	p := pages.InternalError()
	err := layout.Base("error", p).Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
func NotFound(w http.ResponseWriter, r *http.Request) error {
	p := pages.NotFound()
	err := layout.Base("error", p).Render(r.Context(), w)
	if err != nil {
		return err
	}
	return nil
}
