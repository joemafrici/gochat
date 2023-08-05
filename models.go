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

// for if you do the data hiding thing
//func (usr *User) Username() string { return usr.username }
