package unit3

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tools"
	"models"
)

// BlogFrontHandler is the HTTP handler for displaying the most recent posts.
func BlogFrontHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Method: %v", r.Method)
	if r.Method == "GET" {
		posts := models.RecentPosts(c)
		c.Infof("cs253: Posts len: %v", len(posts))
		for i, _ := range posts {
			c.Infof("cs253: Post id: %v", posts[i].Id)
		}
		writeBlog(w, posts)
	}else{
		tools.Error404(w)
		return
	}
}

// writeBlog executes the blog template with a slice of Posts.
func writeBlog(w http.ResponseWriter, posts []*models.Post){
	tmpl, _ := template.ParseFiles("templates/blog.html","templates/post.html")
	tmpl.ExecuteTemplate(w,"blog",posts)
}

// NewPostHandler is the HTTP handler to create a new Post
func NewPostHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Method: %v", r.Method)
	
	if r.Method == "GET" {
		writeNewPostForm(w, &NewPostForm{})
	}else if r.Method == "POST"{
		postForm := NewPostForm{
			r.FormValue("subject"),
			r.FormValue("content"),
			"",
		}
		if !(tools.IsStringValid(postForm.Subject) && 
			tools.IsStringValid(postForm.Content)){
			
			postForm.Error = "We need to set both a subject and some content"
			writeNewPostForm(w, &postForm)
		}else{
			c.Infof("cs253: Blog new post:")
			
			postID, _, _ := datastore.AllocateIDs(c, "Post", nil, 1)
			key := datastore.NewKey(c, "Post", "", postID, nil)
			p := models.Post{
				postID,
				postForm.Subject,
				postForm.Content,
				time.Now(),
			}
			key, err := datastore.Put(c, key, &p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c.Infof("cs253: Blog Key: %v", key.IntID())

			// build url and redirect
			permalinkURL := "/blog/"+strconv.FormatInt(p.Id,10)
			http.Redirect(w, r, permalinkURL, http.StatusFound)
		}
	}else{
		tools.Error404(w)
		return
	}
}

// NewPostForm is the type used to hold the new post information.
type NewPostForm struct{
	Subject string 
	Content string
	Error string
}

// writeNewPostForm executes the newpost.html template with NewPostForm type as param.
func writeNewPostForm(w http.ResponseWriter, postForm *NewPostForm){
	tmpl, _ := template.ParseFiles("templates/newpost.html")
	tmpl.Execute(w,postForm)
}

// PermalinkHandler is the HTTP handler for displaying a single post.
// post information is retreive via the URL: /blog/postId
func PermalinkHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	if r.Method == "GET" {
		
		path := strings.Split(r.URL.String(), "/")
		intID, _ := strconv.ParseInt(path[2], 0, 64)
		c.Infof("cs253: PATH : %v", intID)
		// build post key
		c.Infof("cs253: postAndTimeByID call : %v", intID)
		postAndTime := models.PostAndTimeByID(c, intID)
		c.Infof("cs253: postAndTimeByID done ! : %v", intID)
		
		c.Infof("cs253: Post id: %v", postAndTime.Post.Id)
		c.Infof("cs253: Cache hit time: %v", postAndTime.Cache_hit_time)
		
		writePermalink(w, postAndTime)
	}else{
		tools.Error404(w)
		return
	}
}

// writePermalink executes the permalink.html template with a PostAndTime type as param.
func writePermalink(w http.ResponseWriter, p models.PostAndTime){
	tmpl, _ := template.ParseFiles("templates/permalink.html","templates/post.html")
	tmpl.ExecuteTemplate(w,"permalink",p)
}
