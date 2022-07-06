package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"
	"unicode"
)

type LoginForm struct {
	Username string
	Password string
	Token    string
}

const SET_USERNAME string = "c137@onecause.com"
const SET_PASSWORD string = "#th@nH@rm#y#r!$100%D0p#"

func responseMsg(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
	io.WriteString(w, msg)
}

func handler(w http.ResponseWriter, r *http.Request) {
	SET_TOKEN := time.Now().Format("1504") // format: hhmm

	if r.Body != nil {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responseMsg(w, http.StatusBadRequest, `{"error": "bad request"}`)
			return
		}

		var body LoginForm
		err = json.Unmarshal(bytes, &body)
		if err != nil {
			responseMsg(w, http.StatusBadRequest, `{"error": "invalid json"}`)
			return
		}

		_, err = mail.ParseAddress(body.Username)
		if err != nil {
			responseMsg(w, http.StatusBadRequest, `{"error": "invalid email"}`)
			return
		}

		for i := 0; i < len(body.Password); i++ {
			if body.Password[i] > unicode.MaxASCII {
				responseMsg(w, http.StatusBadRequest, `{"error": "invalid password"}`)
				return
			}
		}

		_, err = strconv.Atoi(body.Token)
		if err != nil || len(body.Token) > 4 {
			responseMsg(w, http.StatusBadRequest, `{"error": "invalid token"}`)
			return
		}

		if body.Username == SET_USERNAME && body.Password == SET_PASSWORD && body.Token == SET_TOKEN {
			responseMsg(w, http.StatusOK, `{"error": null}`)
		} else {
			responseMsg(w, http.StatusUnauthorized, `{"error": "incorrect credentials"}`)
		}
	}
}

func main() {
	http.HandleFunc("/auth/login", handler)
	log.Fatal(http.ListenAndServeTLS(":8080", "../ssl/localhost.ctr", "../ssl/localhost.key", nil))
}
