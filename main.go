package main

import (
	"appengine"
	// "appengine/datastore"
	// "appengine/user"
	"html/template"
	"net/http"
	"fmt"
	"io"
	// "time"
)

type Date struct{
	Month string
	Day string
	Year string
	Error string
}

func init(){
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/unit1/date", dateHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("Requested URL: %v", r.URL)
	
	fmt.Fprint(w,"hello Udacity with Go!")
}

var dateTemplate = template.Must(template.New("MyDate").Parse(dateHTML))

const dateHTML = `
<html>
<body>
<form method="post">
   What is your birthday?
   <br>
	<label>Month
   	<input name="month" value="{{.Month}}">
   </label>
	<label>Day
		<input name="day" value="{{.Day}}">
   </label>
	<label>Year
		<input name="year" value="{{.Year}}">
   </label>
	<div style="color:red">{{.Error}}</div>
	<br>
   <br>
   <input type="submit">
</form>
</body>
</html>
`

func dateHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Http METHOD: %v",r.Method)
	if r.Method == "GET" {
		date := Date{
			Month: "",
			Day: "",
			Year: "",
		}
		
		if err := dateTemplate.Execute(w,date); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST"{
		d := Date{
			Month: r.FormValue("month"),
			Day: r.FormValue("day"),
			Year: r.FormValue("year"),			
		}
		fmt.Fprint(w, "Month:",d.Month," Day:",d.Day," Year:",d.Year)
	}else{
		error404(w)
		return
	}
}

func error404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "404: Not Found")
}
 
















