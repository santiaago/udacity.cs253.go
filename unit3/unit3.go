package unit3

import (
	"html/template"
	"net/http"
	"appengine"
	"fmt"
	"tools"
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
		writeNewPostForm(w,"","","")
	}else if r.Method == "POST"{
		p := Post{
			r.FormValue("subject"),
			r.FormValue("content"),
			"",
		}
		if !(tools.ValidStr(p.Subject) && tools.ValidStr(p.Content)){
			writeNewPostForm(w,
				p.Subject,
				p.Content,
				"We need to set both a subject and some content")
		}else{
			fmt.Fprint(w,"good!")
		}
	}else{
		tools.Error404(w)
		return
	}
}

type Post struct{
	Subject string 
	Content string
	Error string
}


func writeNewPostForm(w http.ResponseWriter, subject string, content string, error string){
	tmpl, _ := template.ParseFiles("templates/newpost.html")
	p := Post{subject,content,error}
	tmpl.Execute(w,p) 
}

func PermalinkHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	if r.Method == "GET" {
		fmt.Fprint(w,"Permalink!")		
	}else{
		tools.Error404(w)
		return
	}
}









