package main

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

// ***********************************************
// Classes
type User struct {
	Username       string
	Password       string
	SessionID      string
	FriendRequests *list.List
}

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

	testuser := User{"deepwater", "passwd", generateSessionID(), list.New()}
	dbase[testuser.Username] = testuser
	testuser.FriendRequests.PushBack("Goribus")

	http.HandleFunc("/", handler)
	log.Println("gochat listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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

	}
	return
}

// ***********************************************
func handler(w http.ResponseWriter, r *http.Request) {
	//log.Println("in handler")
	//log.Println(r.URL.Path)
	switch r.URL.Path {
	case "/signup-submit":
		signupSubmit(w, r)
	case "/login-submit":
		loginHandler(w, r)
	case "/addfriend-submit":
		addFriend(w, r)
	case "/":
		if sessionValid(r) == true {
			serveLoginSuccess(w, r)
		} else {
			serveIndex(w, r)
		}
	default:
		serveIndex(w, r)
	}
}

// ***********************************************
func addFriend(w http.ResponseWriter, r *http.Request) {
	// need to keep track of who is logged in
}

// ***********************************************
func signupSubmit(w http.ResponseWriter, r *http.Request) {
	user, ok := dbase[r.FormValue("username")]
	if ok == true {
		// user does exist
		// handle user already exists
	} else {
		// user does not exist
		user.Username = r.FormValue("username")
		user.Password = r.FormValue("password")
		dbase[r.FormValue("username")] = user
		log.Printf("Created user %v\n", user.Username)
	}
}

// ***********************************************
// verify username/password with database
// create a temporary user session
// issue cookie with session ID
// user sends cookie with each request
// validate cookie against session store
// cookie is a header
// server signs cookie. checks hash of cookie received
// from client to make sure the client has not changed it
// see HMAC algorithm to sign
// url encode cookie (for compatibility)
// or could use json web tokens
func loginHandler(w http.ResponseWriter, r *http.Request) {
	user, userExists := dbase[r.FormValue("username")]
	if userExists == true {
		if r.FormValue("password") == user.Password {
			// need to check if they already have a valid sessionID
			cookie, err := r.Cookie("sessionID")
			if err == http.ErrNoCookie {
				user.SessionID = generateSessionID()
				cookie := http.Cookie{
					Name:     "sessionID",
					Value:    user.SessionID,
					Path:     "/",
					HttpOnly: true,
					Secure:   false, // true for production
					SameSite: http.SameSiteStrictMode,
					MaxAge:   3600,
				}

			}
			http.SetCookie(w, &cookie)
			serveLoginSuccess(w, r)
		} else {
			// incorrect password
			w.WriteHeader(401)
			log.Printf("%v log in request. Incorrect password\n", user.Username)
		}
	} else {
		// user does not exist
		w.WriteHeader(401)
		log.Printf("%v log in request. Does not exist\n",
			r.FormValue("username"))
	}
}

// ***********************************************
func serveLoginSuccess(w http.ResponseWriter, r *http.Request) {
	var fileName = "loginSuccess.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error: Parsing %v\n", fileName)
		return
	}
	user, _ := dbase[r.FormValue("username")]
	t.ExecuteTemplate(w, fileName, user)
}

// ***********************************************
func serveIndex(w http.ResponseWriter, r *http.Request) {
	var fileName = "index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error: Parsing %v\n", fileName)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
}
