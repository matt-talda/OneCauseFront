package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LoginForm struct {
	Username string
	Password string
	Token    string
}

const SET_USERNAME string = "c137@onecause.com"
const SET_PASSWORD string = "#th@nH@rm#y#r!$100%D0p#"

func handler(w http.ResponseWriter, r *http.Request) {
	SET_TOKEN := time.Now().Format("1504")

	if r.Body != nil {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ERROR: empty body")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		var body LoginForm
		err = json.Unmarshal(bytes, &body)
		if err != nil {
			log.Println("ERROR: Invalid JSON")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		if body.Username == SET_USERNAME && body.Password == SET_PASSWORD && body.Token == SET_TOKEN {
			w.Header().Set("status", "200")
			w.Header().Set("content-type", "application/json")
			io.WriteString(w, `{"name":"matt"}`)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}

	//http.Redirect(w, r, "http://onecause.com", http.StatusSeeOther)

	//w.Header().Set("content-type", "application/json")
	//w.Header().Set("status", http.StatusUnauthorized)
	//io.WriteString(w, `{"name":"matt"}`)
	//http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func main() {
	http.HandleFunc("/auth/login", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
