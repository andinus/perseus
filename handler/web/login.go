package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"tildegit.org/andinus/perseus/account"
	"tildegit.org/andinus/perseus/storage"
)

// LoginHandler handles login.
func LoginHandler(w http.ResponseWriter, r *http.Request, db *storage.DB) {
	p := Page{}
	var err error

	t, err := template.ParseFiles("web/templates/login.html")
	if err != nil {
		log.Printf("web/login.go: 500 Internal Server Error :: %s", err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		t.Execute(w, p)

	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			log.Printf("web/login.go: 400 Bad Request :: %s", err.Error())
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		// Get form values
		u := account.User{}
		u.Username = r.FormValue("username")
		u.Password = r.FormValue("password")

		// Perform login
		err = u.Login(db)

		if err != nil {
			log.Printf("web/login.go: %s :: %s",
				"login failed",
				err.Error())

			error := []string{}
			error = append(error,
				fmt.Sprintf("Login failed"))

			p.Error = error
			t.Execute(w, p)
			return
		}

		// Login successful, set token
		cookie := http.Cookie{
			Name:  "token",
			Value: u.Token,
			// Expire the cookie after 16 days from
			// current UTC time.
			Expires:  time.Now().UTC().Add(16 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		success := []string{}
		success = append(success,
			fmt.Sprintf("Login successful"))
		p.Success = success
		t.Execute(w, p)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("web/login.go: %v not allowed on %v", r.Method, r.URL)
	}

}
