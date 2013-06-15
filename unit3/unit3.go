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
			p := models.Post{
				postForm.Subject,
				postForm.Content,
				time.Now(),
			}
			key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "post", nil), &p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c.Infof("cs253: Blog Key: %v", key)
			// build url and redirect
			permalinkURL := "/blog/"+strconv.FormatInt(key.IntID(),10)
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
		c.Infof("cs253: PATH ID : %v", path[2])

		// back to int64
		intID,_ := strconv.ParseInt(path[2], 0, 64)
		c.Infof("cs253: PATH : %v", intID)
		// build key
		key := datastore.NewKey(c, "post", "", intID, nil)
		c.Infof("cs253: PATH : %v", key)
		
		var p Post
		if err := datastore.Get(c, key, &p); err != nil {
			c.Infof("cs253: ERROR : %v", key)			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, p.Subject)
	}else{
		tools.Error404(w)
		return
	}
}










