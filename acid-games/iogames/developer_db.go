package iogames

import (
	"log"

	"gopkg.in/mgo.v2/bson"
)

//InsertDeveloper .
func InsertDeveloper(d DBDeveloper) (err error) {
	collection := GetSession().DB("goDB").C("developers")

	err = collection.Insert(d)

	return
}

//GetDeveloperByUserName .
func GetDeveloperByUserName(username string) (d DBDeveloper, err error) {
	collection := GetSession().DB("goDB").C("developers")

	var dev DBDeveloper

	err = collection.Find(bson.M{"username": username}).One(&dev)

	return dev, err
}

//CheckLoginIsValid .
func CheckLoginIsValid(d DBDeveloper) (isValid bool, err error) {

	collection := GetSession().DB("goDB").C("developers")

	var dev DBDeveloper

	err = collection.Find(bson.M{"username": d.Username, "password": d.Password}).One(&dev)

	if err != nil {
		log.Printf("%v\n", err)
		return false, err
	}

	return true, err
}
