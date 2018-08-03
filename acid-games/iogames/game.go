package iogames

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
)

//Game .
type Game struct {
	Title       string `json:"title" validate:"required,max-length-20,unique-game-title"`
	URL         string `json:"url" validate:"required,max-length-50,valid-URL"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required,valid-category"`
	Thumbnail   string `json:"thumbnail" validate:"required"`
	Owner       string `json:"owner" validate:"required,max-length-50"`
	Featured    bool   `json:"featured"`
	Upvotes     int    `json:"upvotes"`
	TotalVotes  int    `json:"totalvotes"`
	DateAdded   string `json:"dateadded"`
}

//AllGamesAndCategories .
type AllGamesAndCategories struct {
	Games      []Game
	Categories []Category
}

//AllCategoriesAndLoggedUsername .
type AllCategoriesAndLoggedUsername struct {
	Categories []Category
	Username   string
}

//OneGameAndCategories .
type OneGameAndCategories struct {
	Game       Game
	Categories []Category
}

//RegisterGame .
func RegisterGame(w http.ResponseWriter, r *http.Request) {
	log.Print("RegisterGame requested")

	var g Game

	err := DecodeJSON(r.Body, &g)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(responses["bad-json"]))
		return
	}

	log.Printf("Game populado: %v", g)

	dateFormated := time.Now().Format("2006-01-02 15:04:05")
	g.DateAdded = dateFormated

	err = validate.Struct(g)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(responses["bad-validation"]))
		return
	}

	err = InsertGame(g)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["insert-db-error"]))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responses["sucess"]))

	return
}

//ListAllGames .
func ListAllGames(w http.ResponseWriter, r *http.Request) {
	log.Print("AllGames requested")

	absPath, _ := filepath.Abs("./iogames/templates/all_games.html")
	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	allGames, err := GetAllGames()

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	allCategories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	allGamesAndCategories := AllGamesAndCategories{Games: allGames, Categories: allCategories}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, allGamesAndCategories)
}

//ShowGame .
func ShowGame(w http.ResponseWriter, r *http.Request) {
	log.Print("ShowGame requested")

	vars := mux.Vars(r)
	gameTitle := vars["gameTitle"]

	game, err := GetGameByTitle(gameTitle)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(responses["pageNotFound"]))
		return
	}

	allCategories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	absPath, _ := filepath.Abs("./iogames/templates/game_page.html")
	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	gameAndCategories := OneGameAndCategories{Game: game, Categories: allCategories}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, gameAndCategories)
}

//ShowNewGameForm .
func ShowNewGameForm(w http.ResponseWriter, r *http.Request) {
	log.Print("ShowNewGameForm requested")

	absPath, _ := filepath.Abs("./iogames/templates/new_game.html")

	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	allCategories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	session, _ := store.Get(r, "cookie-login")
	//log.Printf("username: %s  \n", session.Values["user"])

	username, isString := session.Values["user"].(string)

	if !isString {
		log.Print("Could not get username session value")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	categoriesAndUsername := AllCategoriesAndLoggedUsername{Categories: allCategories, Username: username}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, categoriesAndUsername)
	return
}

//ListAllGamesByCategory .
func ListAllGamesByCategory(w http.ResponseWriter, r *http.Request) {
	log.Print("ListAllGamesByCategory requested")

	vars := mux.Vars(r)
	categoryName := vars["categoryName"]

	absPath, _ := filepath.Abs("./iogames/templates/category_games.html")
	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	games, err := GetAllGamesOfCategory(categoryName)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	allCategories, err := GetAllCategories()

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["get-from-db-error"]))
		return
	}

	allGamesAndCategories := AllGamesAndCategories{Games: games, Categories: allCategories}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, allGamesAndCategories)

}

//UpvoteGame .
func UpvoteGame(w http.ResponseWriter, r *http.Request) {
	log.Print("UpvoteGame requested")
	vars := mux.Vars(r)
	gameTitle := vars["gameTitle"]

	err := UpvoteGameByTitle(gameTitle)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	game, _ := GetGameByTitle(gameTitle)
	jdata, _ := json.Marshal(game)
	w.Write(jdata)
}

//DownvoteGame .
func DownvoteGame(w http.ResponseWriter, r *http.Request) {
	log.Print("UpvoteGame requested")
	vars := mux.Vars(r)
	gameTitle := vars["gameTitle"]

	err := DownvoteGameByTitle(gameTitle)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	game, _ := GetGameByTitle(gameTitle)
	jdata, _ := json.Marshal(game)
	w.Write(jdata)
}
