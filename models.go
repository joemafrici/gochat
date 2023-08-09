package main

import "container/list"

type User struct {
	Username       string
	Password       string
	SessionToken   string
	FriendRequests *list.List
}
type Session struct {
	Username     string
	SessionToken string
	Active       bool
}

func NewSession(username string, sessionToken string, active bool) *Session {
	return &Session{
		Username:     username,
		SessionToken: sessionToken,
		Active:       active,
	}
}

// for if you do the data hiding thing
//func (usr *User) Username() string { return usr.username }
