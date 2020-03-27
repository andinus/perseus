package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"tildegit.org/andinus/perseus/auth"
	"tildegit.org/andinus/perseus/core"
	"tildegit.org/andinus/perseus/storage/sqlite3"
)

// HandleRegister handles /register pages.
func HandleRegister(w http.ResponseWriter, r *http.Request, db *sqlite3.DB) {
	p := Page{Version: core.Version()}
	p.Notice = []string{
		"Only [a-z] & [0-9] allowed for username",
		"Password length must be greater than 8 characters",
	}
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("web/register.html")
		t.Execute(w, p)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			log.Printf("web/register.go: 400 Bad Request :: %s", err.Error())
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		// Get form values
		uInfo := make(map[string]string)
		uInfo["username"] = r.FormValue("username")
		uInfo["password"] = r.FormValue("password")

		// Perform registration
		err := auth.Register(db, uInfo)

		if err != nil {
			log.Printf("web/register.go: %s :: %s :: %s",
				"registration failed",
				uInfo["username"],
				err.Error())

			error := []string{}
			error = append(error,
				fmt.Sprintf("Registration failed"))

			// Check if the error was because of username
			// not being unique.
			if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
				error = append(error,
					fmt.Sprintf("Username not unique"))
			}
			p.Error = error
		} else {
			success := []string{}
			success = append(success,
				fmt.Sprintf("Registration successful"))
			p.Success = success
		}

		t, _ := template.ParseFiles("web/register.html")
		t.Execute(w, p)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("web/register.go: %v not allowed on %v", r.Method, r.URL)
	}

}
