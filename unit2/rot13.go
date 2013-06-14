package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)

type Rot13 struct{
	Str string
}

func rot13(b byte) byte{
	var first, second byte
	switch{
	case 'a' <= b && b <= 'z':
		first, second = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		first, second = 'A', 'Z'
	default:
		return b
	}
	return (b - first + 13)%(second - first + 1) + first
}

func (r Rot13) Encode() string{
	n := len(r.Str)
	t:= []byte(r.Str)
	for i := 0; i < n; i++{
		t[i] = rot13(t[i])
	}
	return string(t)
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
		r13 := Rot13{}
		if err := rot13Template.Execute(w,r13); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST"{
		r13 := Rot13{r.FormValue("text"),}
		r13.Str = r13.Encode()
		c.Infof("cs253: Rot13 %v",r13.Str)
		if err := rot13Template.Execute(w,r13); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}else{
		tools.Error404(w)
		return
	}
}
