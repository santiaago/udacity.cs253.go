package main

import (
	// "appengine"
	// "appengine/datastore"
	// "appengine/user"
	// "html/template"
	"net/http"
	"fmt"
	// "time"

)


func init(){
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hello world!")
}






