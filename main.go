/*
Copyright © 2024 Christopher Ritter Christopher.Ritter301@gmail.com
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fdaygon/rift/cmd/commands"
	"github.com/fdaygon/rift/pkg/spotify"
	"github.com/fdaygon/rift/pkg/terminal"
	"github.com/go-chi/chi/v5"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var htmlContent string

	spotify.AuthCode = r.URL.Query().Get("code")
	if spotify.AuthCode != "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else {

		welcomeContent := `<html><head><title>Rift</title></head><body>Welcome to rift. Login in <a href="http://localhost:3000/login">Here<a/> /body></html>`
		htmlContent = welcomeContent
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)
	}

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	shell := terminal.CurrentShell()
	terminal.CheckProfileFile(shell, spotify.AuthCode)

}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {
	//not sure how i want to handle the call back so for now the app will just close out.
	fmt.Println("Log in failed. Closing application")
}

func main() {
	router := chi.NewRouter()
	commands.Execute()
	router.Get("/", HandleAuth)
	router.Get("/login", HandleLogin)
	router.Get("/callback", HandleCallBack)

	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Println("Unable to start server")
		os.Exit(1)
	}

}
