package main

import (
	"net/http"
	"myGo/src/database"
	"errors"
)

func session(w http.ResponseWriter, r *http.Request) (sess database.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = database.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
