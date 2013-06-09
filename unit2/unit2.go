package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)

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
		r13 := Rot13{Rot13: r.FormValue("text"),}
		r13.Rot13 = r13.Encode()
		c.Infof("cs253: Rot13 %v",r13.Rot13)
		if err := rot13Template.Execute(w,r13); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}else{
		tools.Error404(w)
		return
	}
}

type Rot13 struct{
	Rot13 string
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
	n := len(r.Rot13)
	t:= []byte(r.Rot13)
	for i := 0; i < n; i++{
		t[i] = rot13(t[i])
	}
	return string(t)
}

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

var signupTemplate = template.Must(template.New("Signup").Parse(signupHTML))

const signupHTML = `
<!DOCTYPE html><html>
  <head>
    <title>Sign Up</title>
    <style type="text/css">.label {text-align: right}.error {color: red}</style>
  </head>
  <body>
	<h2>Signup</h2>
	<form method="post">
	  <table>
	    <tr>
	      <td class="label">Username</td>
	      <td><input type="text" name="username" value="{{.Username}}"></td>
	      <td class="error">{{.ErrorUser}}</td>
	    </tr>
	    <tr>
	      <td class="label">Password</td>
	      <td><input type="password" name="password" value="{{.Password}}"></td>
	      <td class="error">{{.ErrorPassword}}</td>
	    </tr>
	    <tr>
	      <td class="label">Verify Password</td>
	      <td><input type="password" name="verify" value="{{.Verify}}"></td>
	      <td class="error">{{.ErrorPasswordMatch}}</td>
	    </tr>
	    <tr>
	      <td class="label">Email (optional)</td>
	      <td><input type="text" name="email" value="{{.Email}}"></td>
	      <td class="error">{{.ErrorEmail}}</td>
	    </tr>
	  </table>
	  <input type="submit">
	</form>
  </body>
</html>
`

type Signup struct{
	Username string
	Password string
	Email string
	Verify string
	ErrorUser string
	ErrorPassword string
	ErrorPasswordMatch string
	ErrorEmail string
}

func SignupHandler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	c.Infof("cs253: Requested URL: %v", r.URL)
	c.Infof("cs253: Http METHOD: %v",r.Method)
	if r.Method == "GET" {
		s := Signup{
			Username: "",
			Password: "",
			Email: "",
			Verify: "",
			ErrorUser: "",
			ErrorPassword: "",
			ErrorPasswordMatch: "",
			ErrorEmail: "",
		}

		if err := signupTemplate.Execute(w,s); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//todo
	} else if r.Method == "POST"{
		s := Signup{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email: r.FormValue("email"),
			Verify: r.FormValue("verify"),
			ErrorUser: "",
			ErrorPassword: "",
			ErrorPasswordMatch: "",
			ErrorEmail: "",
		}
		
		// if not (valid_username(username) and valid_password(password) and (password == verify) and valid_email(email)):
		//     if not valid_username(username):
		//         errorUser = 'That\'s not a valid username.'
		//     if not valid_password(password):
		//         errorPassword = 'That wasn\'t a valid password.'
		//     if password != verify:
		//         errorVerify = 'Your passwords didn\'t match.'
		//     if not valid_email(email):
		//         errorEmail = 'That\'s not a valid email.'
                
		//     password = ''
		//     verify = ''
		if err := signupTemplate.Execute(w,s); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}else{
		tools.Error404(w)
		return
	}
}

func isUsernameValid(username string) bool{
	return false
}

func isPasswordValid(password string) bool{
	return false
}

func isEmailValid(email string) bool{
	return false
}










