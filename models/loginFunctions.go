package models

import (
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
)

func loginRequest(object Login) Login {
	var obj Login

	c := accessDB(user_collection)
	err := c.Find(bson.M{"username": object.Username, "password": object.Password}).One(&obj)
	if err != nil {
		fmt.Printf("\n Error: \t",err)
	}
	js, __ := json.Marshal(obj)
	verifyErr(__)
	fmt.Printf("\n User Logged in: \n %s \n ", js)
	return obj
}