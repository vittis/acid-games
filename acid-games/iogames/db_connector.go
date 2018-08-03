package iogames

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

//CreateDBConnection .
func CreateDBConnection() (err error) {
	session, err = mgo.Dial(MONGO_HOST)

	if err != nil {
		log.Printf("[ERROR] failed to connect to database")
		return
	}
	return
}

//GetSession .
func GetSession() *mgo.Session {

	return session
}
