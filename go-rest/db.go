package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	currentPostID int
	currentUserID int
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "172.18.0.2:30000")
	HandleError(err)
	return c
}

func FindAll() Posts {

	c := RedisConnect()
	defer c.Close()

	keys, err := c.Do("KEYS", "post:*")
	HandleError(err)

	var posts Posts

	for _, k := range keys.([]interface{}) {
		var post Post

		reply, err := c.Do("GET", k.([]byte))
		HandleError(err)
		if err := json.Unmarshal(reply.([]byte), &post); err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}
	return posts
}

func FindPost(id int) Post {
	var post Post

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "post:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")
	if err = json.Unmarshal(reply.([]byte), &post); err != nil {
		panic(err)
	}
	return post
}

func CreatePost(p Post) {
	currentPostID++
	currentUserID++

	p.ID = currentPostID
	p.User.ID = currentUserID
	p.Timestamp = time.Now()

	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(p)
	HandleError(err)

	// Save JSON blob to Redis
	reply, err := c.Do("SET", "post:"+strconv.Itoa(p.ID), b)
	HandleError(err)

	fmt.Println("GET ", reply)
}

func DeletePost(id int) {

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("DEL", "post:"+strconv.Itoa(id))
	HandleError(err)

	if reply.(int) != 1 {
		fmt.Println("No post removed")
	} else {
		fmt.Println("Post removed")
	}
}
