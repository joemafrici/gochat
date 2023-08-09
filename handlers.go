package main

import (
	"html/template"
	"log"
	"net/http"
)

// ***********************************************
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/signup-submit":
		submitHandler(w, r)
	case "/login-submit":
		loginHandler(w, r)
	case "/addfriend-submit":
		addFriend(w, r)
	case "/":
		if sessionValid(r) == true {
			loginSuccessHandler(w, r)
		} else {
			indexHandler(w, r)
		}
	default:
		indexHandler(w, r)
	}
}

// idiomatic http handler naming
func handleGetUserById() {

}

func handleResizeImage() {

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
	if userExists == true && user.Password == r.FormValue("password") {
		_, err := r.Cookie("sessionToken")
		if err == http.ErrNoCookie {
			user.SessionToken = generateSessionToken()
			cookie := &http.Cookie{
				Name:     "sessionToken",
				Value:    user.SessionToken,
				Path:     "/",
				HttpOnly: true,
				Secure:   false, // true for production
				SameSite: http.SameSiteStrictMode,
				MaxAge:   3600,
			}
			http.SetCookie(w, cookie)
		} else {
			// logged in and request has cookie
		}
		loginSuccessHandler(w, r)
	} else {
		w.WriteHeader(401)
		log.Printf("log in request failed")
	}
}

// ***********************************************
func submitHandler(w http.ResponseWriter, r *http.Request) {
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
func loginSuccessHandler(w http.ResponseWriter, r *http.Request) {
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
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var fileName = "index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error: Parsing %v\n", fileName)
		return
	}
	t.ExecuteTemplate(w, fileName, nil)
}
