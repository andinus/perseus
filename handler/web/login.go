package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"tildegit.org/andinus/perseus/auth"
	"tildegit.org/andinus/perseus/auth/token"
	"tildegit.org/andinus/perseus/core"
	"tildegit.org/andinus/perseus/storage/sqlite3"
)

// HandleLogin handles /login pages.
func HandleLogin(w http.ResponseWriter, r *http.Request, db *sqlite3.DB) {
	p := Page{Version: core.Version()}
	error := []string{}
	success := []string{}

	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("web/login.html")
		t.Execute(w, p)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			log.Printf("web/login.go: 400 Bad Request :: %s", err.Error())
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		// Get form values
		uInfo := make(map[string]string)
		uInfo["username"] = r.FormValue("username")
		uInfo["password"] = r.FormValue("password")

		// Perform authentication
		err := auth.Login(db, uInfo)

		if err != nil {
			log.Printf("web/login.go: %s :: %s :: %s",
				"login failed",
				uInfo["username"],
				err.Error())

			error = append(error,
				fmt.Sprintf("Login failed"))

			p.Error = error
		} else {
			success = append(success,
				fmt.Sprintf("Login successful"))
			p.Success = success

			// Set token if login was successful.
			token, err := token.AddToken(db, uInfo)
			if err != nil {
				log.Printf("web/login.go: %s :: %s :: %s",
					"token generation failed",
					uInfo["username"],
					err.Error())

				error = append(error,
					fmt.Sprintf("Token generation failed"))
			}
			// If token was generated then ask browser to
			// set it as cookie.
			expiration := time.Now().Add(1 * 24 * time.Hour)
			cookie := http.Cookie{Name: "token", Value: token, Expires: expiration}
			http.SetCookie(w, &cookie)
		}

		t, _ := template.ParseFiles("web/login.html")
		t.Execute(w, p)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("web/login.go: %v not allowed on %v", r.Method, r.URL)
	}

}
