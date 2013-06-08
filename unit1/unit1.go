package unit1

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"appengine"
	"tools"
)

type Date struct{
	Month string
	Day string
	Year string
	Error string
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


func ThanksHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Http METHOD: %v",r.Method)
	if r.Method == "GET" {
		fmt.Fprint(w, "Thanks! That's a totally valid date!")
	}else{
		tools.Error404(w)
		return
	}
}

func DateHandler(w http.ResponseWriter, r *http.Request){
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
			Month: validMonth(r.FormValue("month")),
			Day: validDay(r.FormValue("day")),
			Year: validYear(r.FormValue("year")),			
		}
		if d.Day == "" || d.Month == "" || d.Year == ""{
			d.Error = "That's an error!"
			if err := dateTemplate.Execute(w,d); err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w,r, "/unit1/thanks", http.StatusFound)
	}else{
		tools.Error404(w)
		return
	}
}

var months = []string{"JANUARY","FEBRUARY","MARCH","APRIL","MAY","JUNE","JULY","AUGUST","SEPTEMBER","OCTOBER","NOVEMBER","DECEMBER"}

func validMonth(month string) string{
	if len(month) >0{
		if up := strings.ToUpper(month); tools.Contains(months,up){
			return month
		}
	}
	return ""
}

func validDay(day string) string{
	if d, err := strconv.Atoi(day); err == nil{
		if d >0 && d < 31{
			return day
		}
	}
	return ""
}

func validYear(year string) string{
	if y, err := strconv.Atoi(year); err == nil{
		if y > 1900 && y < 2020{
			return year
		}
	}
	return ""
}
