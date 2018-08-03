package iogames

import (
	"github.com/gorilla/mux"
)

//SetupRoutes .
func SetupRoutes(r *mux.Router) {
	//Developers login and register operations
	r.HandleFunc("/developers/register", RegisterDev).Methods("POST")
	r.HandleFunc("/developers/sign_up", SignUpDevForm).Methods("GET")
	r.HandleFunc("/developers/login", LoginDev).Methods("POST")
	r.HandleFunc("/developers/sign_in", SignInDevForm).Methods("GET")
	r.HandleFunc("/developers/logout", LogoutDev).Methods("POST")

	//Logged developer operations
	r.HandleFunc("/developer/dashboard", ShowDeveloperDashBoard).Methods("GET")
	r.HandleFunc("/developer/games/new", ShowNewGameForm).Methods("GET")
	r.HandleFunc("/developer/games/new_game", RegisterGame).Methods("POST")

	//Game operations
	r.HandleFunc("/{gameTitle}", ShowGame).Methods("GET")
	r.HandleFunc("/", ListAllGames).Methods("GET")
	r.HandleFunc("/{gameTitle}/upvote", UpvoteGame).Methods("POST")
	r.HandleFunc("/{gameTitle}/downvote", DownvoteGame).Methods("POST")

	//Category operations
	r.HandleFunc("/category/{categoryName}", ListAllGamesByCategory).Methods("GET")

}
