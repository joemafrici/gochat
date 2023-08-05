package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
)

// ***********************************************
func generateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("rand error: %v\n", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// ***********************************************
func sessionValid(r *http.Request) (valid bool) {
	cookie, err := r.Cookie("sessionID")
	if err == http.ErrNoCookie {
		valid = false
	} else {
		thing, ok := sessions[cookie.Value]
		if thing.Active == true {
			valid = true
			return
		}
	}
	return
}

// ***********************************************
func addFriend(w http.ResponseWriter, r *http.Request) {
	// need to keep track of who is logged in
}
