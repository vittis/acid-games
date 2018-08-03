package iogames

import "gopkg.in/mgo.v2/bson"

//InsertCategory .
func InsertCategory(c Category) (err error) {
	collection := GetSession().DB("goDB").C("categories")

	err = collection.Insert(c)

	return
}

//GetAllCategories .
func GetAllCategories() (allCategories []Category, err error) {

	collection := GetSession().DB("goDB").C("categories")

	var categories []Category

	err = collection.Find(nil).All(&categories)

	return categories, err
}

//GetCategoryByName .
func GetCategoryByName(name string) (c Category, err error) {
	collection := GetSession().DB("goDB").C("categories")

	var category Category

	err = collection.Find(bson.M{"name": name}).One(&category)

	return category, err
}
