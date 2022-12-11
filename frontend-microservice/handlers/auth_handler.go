package handlers

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}

func NewAuthHandler() http.Handler {
	return &authHandler{}
}

type authHandler struct{}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("not allowed"))
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	fmt.Println(username, password)

	http.Redirect(w, r, "/app/", http.StatusFound)

	// w.Write([]byte("auth handler"))
}
