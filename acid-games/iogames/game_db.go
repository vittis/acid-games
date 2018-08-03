package iogames

import "gopkg.in/mgo.v2/bson"

//InsertGame .
func InsertGame(g Game) (err error) {
	collection := GetSession().DB("goDB").C("games")

	err = collection.Insert(g)

	return
}

//GetGameByTitle .
func GetGameByTitle(title string) (g Game, err error) {
	collection := GetSession().DB("goDB").C("games")

	var game Game

	err = collection.Find(bson.M{"title": title}).One(&game)

	return game, err
}

//GetAllGames .
func GetAllGames() (allGames []Game, err error) {

	collection := GetSession().DB("goDB").C("games")

	var games []Game

	err = collection.Find(nil).All(&games)

	return games, err
}

//GetAllGamesOfDeveloper .
func GetAllGamesOfDeveloper(username string) (gamesOfDev []Game, err error) {

	collection := GetSession().DB("goDB").C("games")

	var games []Game

	err = collection.Find(bson.M{"owner": username}).All(&games)

	return games, err
}

//GetAllGamesOfCategory .
func GetAllGamesOfCategory(category string) (gamesOfCategory []Game, err error) {

	collection := GetSession().DB("goDB").C("games")

	var games []Game

	err = collection.Find(bson.M{"category": category}).All(&games)

	return games, err
}

//UpvoteGameByTitle .
func UpvoteGameByTitle(title string) (err error) {
	collection := GetSession().DB("goDB").C("games")

	g, _ := GetGameByTitle(title)

	err = collection.Update(bson.M{"title": title}, bson.M{"$set": bson.M{"upvotes": g.Upvotes + 1, "totalvotes": g.TotalVotes + 1}})

	return
}

//DownvoteGameByTitle .
func DownvoteGameByTitle(title string) (err error) {
	collection := GetSession().DB("goDB").C("games")

	g, _ := GetGameByTitle(title)

	err = collection.Update(bson.M{"title": title}, bson.M{"$set": bson.M{"totalvotes": g.TotalVotes + 1}})

	return
}
