package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"tildegit.org/andinus/perseus/account"
	"tildegit.org/andinus/perseus/storage"
)

// RegisterHandler handles registration.
func RegisterHandler(w http.ResponseWriter, r *http.Request, db *storage.DB) {
	p := Page{}
	var err error

	t, err := template.ParseFiles("web/templates/register.html")
	if err != nil {
		log.Printf("web/register.go: 500 Internal Server Error :: %s", err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	p.Notice = []string{
		"Only [a-z] & [0-9] allowed for username",
		"Password length must be greater than 8 characters",
	}

	switch r.Method {
	case http.MethodGet:
		t.Execute(w, p)

	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			log.Printf("web/register.go: 400 Bad Request :: %s", err.Error())
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		// Get form values
		u := account.User{}
		u.Username = r.FormValue("username")
		u.Password = r.FormValue("password")

		// Perform registration
		err = u.Register(db)

		if err != nil {
			log.Printf("web/register.go: %s :: %s",
				"registration failed",
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

		t.Execute(w, p)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("web/register.go: %v not allowed on %v", r.Method, r.URL)
	}

}
