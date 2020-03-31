package models

import (
	"fmt"
	"time"
	"math/rand"
	"gopkg.in/mgo.v2"
	"restful_api/patterns"
)

const (
	user_collection = patterns.Users_collection
)

var session *mgo.Session
var errDial error


func accessDB(collection string)(*mgo.Collection) {
	session, errDial = patterns.GetSessErrMongoDBSession("Dial")
	verifyErr(errDial)
	session.SetMode(mgo.Monotonic, true)
	col := session.DB(patterns.Dbname).C(collection)
	return col
}

func verifyErr(err error) {
	if err != nil {
		fmt.Printf("\n Error: %s \n", err)
		panic(err)
	}
}

func generarPassword(longitud int) (cad string) {
	rand.Seed(time.Now().UTC().UnixNano())
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	password := " "
	for i := 0; i < longitud; i++ {
		ln := rand.Intn(len(characters))
		password += string(characters[ln])
	}
	return password
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Encrypt(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
