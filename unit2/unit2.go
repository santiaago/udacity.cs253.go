package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)


type Rot13 struct{
	Rot13 string
}

var rot13Template = template.Must(template.New("Rot13").Parse(rot13HTML))

const rot13HTML = `
<!DOCTYPE html><html>
    <head><title>Unit 2 Rot 13</title></head>
    <body>
        <h2>Enter some text to ROT13:</h2>
        <form method="post">
            <textarea name="text" style="height: 100px; width: 400px;">{{.Rot13}}</textarea><br>
            <input type="submit">
        </form>
    </body>
</html>
`

func Rot13Handler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Http METHOD: %v",r.Method)
	if r.Method == "GET" {
		r13 := Rot13{Rot13: "",}
		if err := rot13Template.Execute(w,r13); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST"{
		if err := rot13Template.Execute(w,"hola"); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w,r, "/unit1/thanks", http.StatusFound)
	}else{
		tools.Error404(w)
		return
	}

}

// class Rot13Handler(webapp2.RequestHandler):
    
//     def write_form(self,rot13=""):
//         self.response.out.write(htmlRot13%{"rot13":escape_html(rot13)})
        
//     def get(self):
//         self.write_form()
        
//     def post(self):
//         user_input = self.request.get('text')
//         input_changed = user_input.encode('rot13')
//         self.write_form(input_changed)








