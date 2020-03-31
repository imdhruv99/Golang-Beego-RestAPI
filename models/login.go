package models

import (
	"fmt"
)

type Login struct {
	Username 	string		`json:"username"`
	Password	string		`json:"password"`
}

var login Login

func init() {

}

func VerifyLogin(object Login) (ObjectId bool, err error) {
	objlogin := loginRequest(object)
	if objlogin == (Login{}) {
		fmt.Printf("\n No user Found \n \n")
		return false, nil
	}	else {
		return true, nil
	}
}