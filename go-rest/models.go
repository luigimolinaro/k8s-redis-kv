package main

import (
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Post struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Name	  string    `json:"name"`
	Surname   string    `json:"surname"`
	//Topic     string    `json:"topic"`
	//Text      string    `json:"text"`
	//Comment   Comment   `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
}

type Posts []Post

type Users []User
