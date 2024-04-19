package web

import (
	"net/http"

	"github.com/Lafetz/showdown-trivia-game/internal/web/form"
	layout "github.com/Lafetz/showdown-trivia-game/internal/web/views/layouts"
	"github.com/Lafetz/showdown-trivia-game/internal/web/views/pages"
)

func (a *App) initAppRoutes() {
	fileServer := http.FileServer(http.Dir("./internal/web/static"))

	a.router.Handle("/static/", http.StripPrefix("/static/", fileServer))
	a.router.HandleFunc("GET /signup", func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signup(form.SignupUser{})
		err := layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})
	a.router.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			println("problem parsing form")
			return
		}
		form := form.SignupUser{
			Username: r.PostForm.Get("username"),
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		if !form.Valid() {
			p := pages.Signup(form)
			println("errrr")

			w.WriteHeader(422)
			err = layout.Base("Sign up", p).Render(r.Context(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			return
		}
		println(form.Email, form.Password)
		p := pages.SignupSuccess()
		w.WriteHeader(201)
		err = layout.Base("Sign up", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

	})
	a.router.HandleFunc("GET /signin", func(w http.ResponseWriter, r *http.Request) {
		p := pages.Signin(form.SigninUser{}, "")
		err := layout.Base("Sign in", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})
	a.router.HandleFunc("POST /signin", func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			println("problem parsing form")
			return
		}
		form := form.SigninUser{
			Email:    r.PostForm.Get("email"),
			Password: r.PostForm.Get("password"),
		}
		println(form.Email, form.Password)
		if !form.Valid() {
			println("not failed")
			p := pages.Signin(form, "")
			w.WriteHeader(422)
			err = layout.Base("Sign up", p).Render(r.Context(), w)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			return
		}

		w.Header().Set("HX-Redirect", "/home")
		w.WriteHeader(http.StatusOK)

	})

	a.router.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		p := pages.Home()
		err := layout.Base("Home", p).Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})
	a.router.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		sendGamePage(w, r, true)
	})
	a.router.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		sendGamePage(w, r, false)
	})
	//ws
	a.router.HandleFunc("/wscreate", func(w http.ResponseWriter, r *http.Request) {
		a.hub.CreateRoom(w, r)
	})
	a.router.HandleFunc("/wsjoin", func(w http.ResponseWriter, r *http.Request) {
		a.hub.CreateRoom(w, r)
	})
}
func sendGamePage(w http.ResponseWriter, r *http.Request, create bool) {
	p := pages.Game(create)
	err := layout.Base("Game", p).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
