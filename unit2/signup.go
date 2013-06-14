package unit2

import (
	"html/template"
	"net/http"
	"appengine"
	"tools"
)

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
		s := Signup{}

		if err := signupTemplate.Execute(w,s); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
		if !(tools.IsUsernameValid(s.Username) && tools.IsPasswordValid(s.Password) && s.Password == s.Verify && tools.IsEmailValid(s.Email)){
			if ! tools.IsUsernameValid(s.Username){
				s.ErrorUser = "That's not a valid user name."
			}
			if ! tools.IsPasswordValid(s.Password){
				s.ErrorPassword = "That wasn't a valid password."
			}
			if s.Password != s.Verify{
				s.ErrorPasswordMatch = "Your passwords didn't match."
			}
			if !tools.IsEmailValid(s.Email){
				s.ErrorEmail = "That's not a valid email."
			}
			s.Password = ""
			s.Verify = ""
			if err := signupTemplate.Execute(w,s); err != nil{
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}else{
			http.Redirect(w,r, "/unit2/welcome?username="+s.Username, http.StatusFound)
		}
	}else{
		tools.Error404(w)
		return
	}
}
