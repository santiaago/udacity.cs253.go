package unit3

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tools"
	"models"
)

func BlogFrontHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	if r.Method == "GET" {
		fmt.Fprint(w,"Blog front!")		
	}else{
		tools.Error404(w)
		return
	}
}

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
		if !(tools.ValidStr(postForm.Subject) && tools.ValidStr(postForm.Content)){
			postForm.Error = "We need to set both a subject and some content"
			writeNewPostForm(w, &postForm)
		}else{
			c.Infof("cs253: Blog new post:")
			p := models.Post{
				0,
				postForm.Subject,
				postForm.Content,
				time.Now(),
			}
			incKey := datastore.NewIncompleteKey(c,"post",nil)
			key, err := datastore.Put(c, incKey, &p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c.Infof("cs253: Blog Key: %v", key.IntID())
			// set post id
			p.Id = key.IntID()
			key, err = datastore.Put(c, key, &p)
			if err != nil {
			 	http.Error(w, err.Error(), http.StatusInternalServerError)
			 	return
			}
			c.Infof("cs253: Blog new post key: %v", key)
			c.Infof("cs253: Blog new post id: %v", p.Id)
			// build url and redirect
			permalinkURL := "/blog/"+strconv.FormatInt(p.Id,10)
			http.Redirect(w, r, permalinkURL, http.StatusFound)
		}
	}else{
		tools.Error404(w)
		return
	}
}

type NewPostForm struct{
	Subject string 
	Content string
	Error string
}


func writeNewPostForm(w http.ResponseWriter, postForm *NewPostForm){
	tmpl, _ := template.ParseFiles("templates/newpost.html")
	tmpl.Execute(w,postForm)
}

func PermalinkHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	if r.Method == "GET" {
		
		path := strings.Split(r.URL.String(), "/")
		
		// back to int64
		intID, _ := strconv.ParseInt(path[2], 0, 64)
		c.Infof("cs253: PATH : %v", intID)
		// build key
		key := datastore.NewKey(c, "post", "", intID, nil)
		c.Infof("cs253: PATH : %v", key)
		
		var p models.Post
		if err := datastore.Get(c, key, &p); err != nil {
			c.Infof("cs253: ERROR : %v", key)			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Infof("cs253: Key : %v", p.Id)
		writePermalink(w, &p)
	}else{
		tools.Error404(w)
		return
	}
}

func writePermalink(w http.ResponseWriter, p *models.Post){
	tmpl, _ := template.ParseFiles("templates/permalink.html","templates/post.html")
	tmpl.ExecuteTemplate(w,"permalink",p)
}


















