package models

import (
	"time"
	"appengine"
	"appengine/datastore"
)

type Post struct {
	Id int64 
	Subject string
	Content string
	Created time.Time
}

func RecentPosts(c appengine.Context)([]*Post){
	c.Infof("cs253: RecentPosts")
	q := datastore.NewQuery("Post").Limit(10)
	var posts []*Post
	if _, err := q.GetAll(c, &posts); err != nil {
		c.Infof("cs253: Error: %v",err)
		return nil
	}
	return posts
}
