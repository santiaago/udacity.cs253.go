package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)

// Rot13 is the type used to hold the string to encode.
type Rot13 struct{
	Str string
}

// rot13 returns the rot13 substitution of single byte.
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

// Rot13 implement Encode function to perform ROT13 substitution.
// this is a slight modification of Go tour 60. 
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
            <textarea name="text" style="height: 100px; width: 400px;">{{.Str}}</textarea><br>
            <input type="submit">
        </form>
    </body>
</html>
`

// writeFormRot13 executes the rot13Template with a given Rot13 variable
func writeFormRot13(w http.ResponseWriter, r13 Rot13){
	if err := rot13Template.Execute(w,r13); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Rot13Handler is the HTTP handler for encoding and decoding (rot13(rot13(x)) = x ) a string
func Rot13Handler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: HTTP METHOD: %v",r.Method)
	if r.Method == "GET" {
		r13 := Rot13{}
		writeFormRot13(w, r13)
	} else if r.Method == "POST"{
		r13 := Rot13{r.FormValue("text"),}
		r13.Str = r13.Encode()
		c.Infof("cs253: Rot13 %v",r13.Str)
		writeFormRot13(w, r13)
	}else{
		tools.Error404(w)
		return
	}
}
