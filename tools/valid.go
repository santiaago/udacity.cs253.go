
package tools

import (
	"regexp"
)

// IsUsernameValid returns true is username is a string between 3 and 20 characters.
var USERNAME_RE = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
func IsUsernameValid(username string) bool{

	return USERNAME_RE.MatchString(username) 
}

// IsPasswordValid returns true for all passwords that are between 3 and 20 characters.
var PASSWORD_RE = regexp.MustCompile(`^.{3,20}$`)
func IsPasswordValid(password string) bool{
	
	return PASSWORD_RE.MatchString(password)
}
// IsEmailValid returns true if string email has the form a@b.c
var EMAIL_RE = regexp.MustCompile(`^[\S]+@[\S]+\.[\S]+$`)
func IsEmailValid(email string) bool{

	return EMAIL_RE.MatchString(email)
}

// IsStringValid returns true if string is not empty.
func IsStringValid(s string) bool{
	return  len(s)>0
}
