package main

import (
	"fmt"
	"net/http"
	"appengine"
	"unit1"
	"unit2"
)

func init(){
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/unit1/date", unit1.DateHandler)
	http.HandleFunc("/unit1/thanks", unit1.ThanksHandler)
	http.HandleFunc("/unit2/rot13", unit2.Rot13Handler)
	http.HandleFunc("/unit2/signup", unit2.SignupHandler)
	http.HandleFunc("/unit2/welcome", unit2.WelcomeHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	
	fmt.Fprint(w,"hello Udacity with Go!")
} 
















