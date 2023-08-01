package main

import (
	"html/template"
	"log"
	"net/http"
    "container/list"
)

//***********************************************
// Classes
type User struct {
    Username string
    Password string 
    FriendRequests *list.List
}
//***********************************************
// Globals
var (
    dbase map[string]User
)
//***********************************************
func main() {
    dbase = make(map[string]User)
    testuser := User{"deepwater", "passwd", list.New()}
    dbase[testuser.Username] = testuser
    testuser.FriendRequests.PushBack("newFriend")

	http.HandleFunc("/", handler)
    log.Println("gochat listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//***********************************************
func handler(w http.ResponseWriter, r *http.Request) {
	//log.Println("in handler")
	//log.Println(r.URL.Path)
	switch r.URL.Path {
    case "/signup-submit":
		signupSubmit(w, r)
	case "/login-submit":
		loginSubmit(w, r)
    case "/addfriend-submit":
        addFriend(w, r)    
	case "/":
		serveIndex(w, r)
	default:
		serveIndex(w, r)
	}
}
//***********************************************
func addFriend(w http.ResponseWriter, r *http.Request) {
    // need to keep track of who is logged in
}
//***********************************************
func signupSubmit(w http.ResponseWriter, r *http.Request) {
    user, ok := dbase[r.FormValue("username")]
    if ok == true {
        // user does exist
        // handle user already exists
    } else {
        // user does not exist
        user.Username = r.FormValue("username")
        user.Password= r.FormValue("password")
        dbase[r.FormValue("username")] = user
        log.Printf("Created user %v\n", user.Username)
    }
}
//***********************************************
func loginSubmit(w http.ResponseWriter, r *http.Request) {
    user, ok := dbase[r.FormValue("username")]
    if ok == true {
        // user does exist
        if r.FormValue("password") == user.Password {
            serveLoginSuccess(w, r)
        } else {
            // incorrect password
            w.WriteHeader(404)
            log.Printf("%v log in request. Incorrect password\n", user.Username)
        }
    } else {
        // user does not exist
        w.WriteHeader(404)
        log.Printf("%v log in request. Does not exist\n",
            r.FormValue("username"))
    }
}
//***********************************************
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
//***********************************************
func serveIndex(w http.ResponseWriter, r *http.Request) {
	var fileName = "index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error: Parsing %v\n", fileName)
        return
	}
	t.ExecuteTemplate(w, fileName, nil)
}
