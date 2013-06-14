package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)

var welcomeTemplate = template.Must(template.New("Welcome").Parse(welcomeHTML))

const welcomeHTML = `
<!DOCTYPE html><html>
<head>
<title>Unit 2 Signup</title>
</head>
<body>
<h2>Welcome, {{.}}!</h2>
</body>
</html>
`

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Http METHOD: %v",r.Method)
	if r.Method == "GET" {
		username := r.FormValue("username")
		if err := welcomeTemplate.Execute(w,username); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}else{
		tools.Error404(w)
		return
	}
}
