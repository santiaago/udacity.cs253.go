package models

import (
	"bytes"
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"encoding/gob"
	"strconv"
	"time"
)

// Post is the type used to hold the Post information.
type Post struct {
	Id int64 
	Subject string
	Content string
	Created time.Time
}

// PostAndTime is the type used to hold a Post and it's "cache hit time" information.
type PostAndTime struct{
	Post Post
	Cache_hit_time time.Time
}

func init(){
	// Registering a Type more than once causes a panic.
	// so registration should go in an init function.
	gob.Register(PostAndTime{})	
}

// RecentPosts returns a pointer to a slice of Posts.
// orderBy creation time, limit = 20
func RecentPosts(c appengine.Context)([]*Post){
	c.Infof("cs253: RecentPosts")
	q := datastore.NewQuery("Post").Limit(20).Order("-Created")
	var posts []*Post
	if _, err := q.GetAll(c, &posts); err != nil {
		c.Errorf("cs253: Error: %v",err)
		return nil
	}
	return posts
}

// PostAndTimeByID returns a PostAndTime for the requested id
func PostAndTimeByID(c appengine.Context, id int64)( PostAndTime){
	memcacheKey := "posts_and_time"+strconv.FormatInt(id, 10)
	c.Infof("cs253: Post and time by id memcache key is: %v ",memcacheKey)

	var postAndTime PostAndTime
	c.Infof("cs253: query cache first with memcache key")
	if item, err := memcache.Get(c, memcacheKey); err == memcache.ErrCacheMiss {
		c.Infof("cs253: item not in the cache :%v will perform query instead",err)

		key := datastore.NewKey(c, "Post", "", id, nil)
		c.Infof("cs253: key to use: v%", key)

		if err := datastore.Get(c, key, &postAndTime.Post); err != nil {
			c.Errorf("cs253: post not found : %v", err)
		}
		c.Infof("cs253: get time:")
		postAndTime.Cache_hit_time = time.Now()
		c.Infof("cs253: time is: %v", postAndTime.Cache_hit_time)
		// record information in cache for next time
		mCache := new(bytes.Buffer)
		encCache := gob.NewEncoder(mCache)
		c.Infof("cs253: New Encoder done")
		encCache.Encode(postAndTime)
		c.Infof("cs253: Encode done")

		postItem := &memcache.Item{
			Key:   memcacheKey,
			Value: mCache.Bytes(),
		}
		c.Infof("cs253: memcache Item ready")
		if err := memcache.Add(c, postItem); err == memcache.ErrNotStored {
			c.Errorf("cs253: postAndTime with key %q already exists", item.Key)
		} else if err != nil {
			c.Errorf("error adding item: %v", err)
		}
		c.Infof("cs253: item read to be returned")

	} else if err != nil {
		c.Errorf("cs253: Memcache error getting item: %v",err)
	} else {
		c.Infof("cs253: Memcache item found")

		pCache := bytes.NewBuffer(item.Value)
		c.Infof("cs253: New Buffer ready")
		decCache := gob.NewDecoder(pCache)
		c.Infof("cs253: decoder ready")
		decCache.Decode(&postAndTime)
		c.Infof("cs253: gob decoder of item done")
	}	
	return postAndTime
}
