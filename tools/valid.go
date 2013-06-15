
package tools

import (
	"regexp"
)

func IsUsernameValid(username string) bool{
	r := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	
	if r.MatchString(username) == true{
		return true
	}
	return false
}

func IsPasswordValid(password string) bool{
	r := regexp.MustCompile(`^.{3,20}$`)
	
	if r.MatchString(password) == true{
		return true
	}
	return false
}

func IsEmailValid(email string) bool{
	r := regexp.MustCompile(`^[\S]+@[\S]+\.[\S]+$`)
	
	if r.MatchString(email) == true{
		return true
	}
	return false
}

func ValidStr(s string) bool{
	return  len(s)>0
}
