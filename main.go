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
	dbase    map[string]User
	sessions map[string]Session
)

// 401 - authentication (identity)
// 403 - authorization (permissions)
// ***********************************************
func main() {
	dbase = make(map[string]User)
	sessions = make(map[string]Session)

	testuser := User{
		Username:       "deepwater",
		Password:       "passwd",
		SessionToken:   generateSessionID(),
		FriendRequests: list.New(),
	}
	testSession := Session{
		Username:     testuser.Username,
		SessionToken: testuser.SessionToken,
		Active:       true,
	}

	sessions[testuser.SessionToken] = testSession
	dbase[testuser.Username] = testuser
	testuser.FriendRequests.PushBack("Goribus")
    // check if the above line changes the testuser in dbase
	http.HandleFunc("/", handler)
	log.Println("gochat listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
