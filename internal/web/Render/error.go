package render

import (
	"net/http"

	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func InternalServer(w http.ResponseWriter, r *http.Request) {
	p := pages.InternalError()
	err := layout.Base("error", p).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		// return the error
		return
	}
}
func NotFound(w http.ResponseWriter, r *http.Request) {
	p := pages.NotFound()
	err := layout.Base("error", p).Render(r.Context(), w)
	if err != nil {
		InternalServer(w, r)
	}
}
