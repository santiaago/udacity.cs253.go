
package tools

import (
	"regexp"
)

// IsUsernameValid returns true is username is a string between 3 and 20 characters.
func IsUsernameValid(username string) bool{
	r := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	
	if r.MatchString(username) == true{
		return true
	}
	return false
}

// IsPasswordValid returns true for all passwords that are between 3 and 20 characters.
func IsPasswordValid(password string) bool{
	r := regexp.MustCompile(`^.{3,20}$`)
	
	if r.MatchString(password) == true{
		return true
	}
	return false
}
// IsEmailValid returns true if string email has the form a@b.c
func IsEmailValid(email string) bool{

	r := regexp.MustCompile(`^[\S]+@[\S]+\.[\S]+$`)
	
	if r.MatchString(email) == true{
		return true
	}
	return false
}

// IsStringValid returns true if string is not empty.
func IsStringValid(s string) bool{
	return  len(s)>0
}
