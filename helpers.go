package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
)

// ***********************************************
func generateSessionToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("rand error: %v\n", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ***********************************************
func sessionValid(r *http.Request) (valid bool) {
	valid = false
	c, err := r.Cookie("sessionToken")
	if err == http.ErrNoCookie {
		valid = false
	} else {
		s , _ := sessions[c.Value]
		if s.Active == true {
			valid = true
		}
	}
	return
}

// ***********************************************
func findUser(r *http.Request) (user User, userExists bool) {
	user, userExists = dbase[r.FormValue("username")]
	return 
}

// ***********************************************
func addFriend(w http.ResponseWriter, r *http.Request) {
	// need to keep track of who is logged in
	if sessionValid(r) {
		//user, userExists := findUser(r)
		_, userExists := findUser(r)
		if userExists {
			//cookie, err := r.Cookie("sessionToken")
			_, err := r.Cookie("sessionToken")
			if err == http.ErrNoCookie {

			} else {
				//user, exists := sessions[cookie.Value]
			}

		}
	} else {
		log.Println("addFriend session invalid")
	}
}
