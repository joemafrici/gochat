package main

import "container/list"

type User struct {
	Username       string
	Password       string
	SessionID      string
	FriendRequests *list.List
}

// for if you do the data hiding thing
//func (usr *User) Username() string { return usr.username }
