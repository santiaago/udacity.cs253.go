package main

import (
	"fmt"
	"net/http"
	"appengine"
	"unit1"
	"unit2"
	"unit3"
	"tools"
)

func init(){
	h := new( tools.RegexpHandler)
	
	h.HandleFunc("/",mainHandler)
	h.HandleFunc("/unit1/date", unit1.DateHandler)
	h.HandleFunc("/unit1/thanks", unit1.ThanksHandler)
	h.HandleFunc("/unit2/rot13", unit2.Rot13Handler)
	h.HandleFunc("/unit2/signup/?", unit2.SignupHandler)
	h.HandleFunc("/unit2/welcome/?", unit2.WelcomeHandler)
	h.HandleFunc("/blog/?", unit3.BlogFrontHandler)
	h.HandleFunc("/blog/newpost", unit3.NewPostHandler)
	h.HandleFunc("/blog/[0-9]+/?",unit3.PermalinkHandler)
	
	http.Handle("/",h)
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	
	fmt.Fprint(w,"hello Udacity with Go!")
}
