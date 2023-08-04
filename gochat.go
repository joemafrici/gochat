package main

import (
	"container/list"
	"log"
	"net/http"
	// "github.com/joemafrici/gochat/models"
)

// ***********************************************
// Globals
var (
	dbase map[string]User
)

// 401 - authentication (identity)
// 403 - authorization (permissions)
// ***********************************************
func main() {
	dbase = make(map[string]User)

	testuser := User{
		Username:       "deepwater",
		Password:       "passwd",
		SessionID:      generateSessionID(),
		FriendRequests: list.New(),
	}
	dbase[testuser.Username] = testuser
	testuser.FriendRequests.PushBack("Goribus")

	http.HandleFunc("/", handler)
	log.Println("gochat listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
