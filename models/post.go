package models

import (
	"time"
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"strconv"
	//"net/http"
	//"encoding/binary"
	"encoding/gob"
	"bytes"
)

type Post struct {
	Id int64 
	Subject string
	Content string
	Created time.Time
}

type PostAndTime struct{
	P Post
	T time.Time
}

func init(){
	gob.Register(PostAndTime{})	
}
func RecentPosts(c appengine.Context)([]*Post){
	c.Infof("cs253: RecentPosts")
	q := datastore.NewQuery("Post").Limit(20).Order("-Created")
	var posts []*Post
	if _, err := q.GetAll(c, &posts); err != nil {
		c.Infof("cs253: Error: %v",err)
		return nil
	}
	return posts
}

func PostAndTimeByID(c appengine.Context, id int64)( Post,  time.Time){
	memcacheKey := "posts_and_time"+strconv.FormatInt(id, 10)
	c.Infof("cs253: Post and time by id memcache key is: %v ",memcacheKey)

	var p Post
	var t time.Time

	c.Infof("cs253: query cache first with memcache key")
	if item, err := memcache.Get(c, memcacheKey); err == memcache.ErrCacheMiss {
		c.Infof("cs253: item not in the cache :%v will perform query instead",err)

		key := datastore.NewKey(c, "Post", "", id, nil)
		c.Infof("cs253: key to use: v%", key)

		if err := datastore.Get(c, key, &p); err != nil {
			c.Errorf("cs253: post not found : %v", err)
		}
		c.Infof("cs253: get time:")
		t = time.Now()
		c.Infof("cs253: time is: %v", t)
		// record information in cache for next time
		postAndTime := PostAndTime{P: p, T: t,}
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		c.Infof("cs253: New Encoder done", t)
		encCache.Encode(postAndTime)
		c.Infof("cs253: Encode done", t)

		postItem := &memcache.Item{
			Key:   memcacheKey,
			Value: mCache.Bytes(),
		}
		c.Infof("cs253: memcache Item ready")
		if err := memcache.Add(c, postItem); err == memcache.ErrNotStored {
			c.Infof("cs253: postAndTime with key %q already exists", item.Key)
		} else if err != nil {
			c.Errorf("error adding item: %v", err)
		}
		c.Infof("cs253: item read to be returned")

	} else if err != nil {
		c.Errorf("cs253: Memcache error getting item: %v",err)
	} else {
		c.Infof("cs253: Memcache item found")
		var postAndTime PostAndTime
		
		pCache := bytes.NewBuffer(item.Value)//mCache.Bytes())
		c.Infof("cs253: New Buffer ready")
		
		decCache := gob.NewDecoder(pCache)
		c.Infof("cs253: decoder ready")
		
		decCache.Decode(&postAndTime)
		c.Infof("cs253: god decoder of item done")
		
		p = postAndTime.P
		t = postAndTime.T
		
	}	
	return p, t
}
