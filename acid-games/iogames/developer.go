package iogames

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//DBDeveloper .
type DBDeveloper struct {
	Username string `json:"username" validate:"required,max-length-20,unique-dev-username"`
	Password string `json:"password" validate:"required,max-length-20"`
}

//ShowDeveloperDashBoard GET /developer/dashboard
func ShowDeveloperDashBoard(w http.ResponseWriter, r *http.Request) {
	log.Print("ShowDeveloperDashBoard requested")

	session, _ := store.Get(r, "cookie-login")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(responses["forbidden"]))
		return
	}

	absPath, _ := filepath.Abs("./iogames/templates/dev_dashboard.html")

	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(responses["internalError"]))
		return
	}
	username, isString := session.Values["user"].(string)

	if !isString {
		log.Print("Could not get username session value")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}
	gamesOfDev, err := GetAllGamesOfDeveloper(username)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["insert-db-error"]))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, gamesOfDev)
	return
}

//LoginDev .
func LoginDev(w http.ResponseWriter, r *http.Request) {
	log.Print("LoginDev requested")

	session, err := store.Get(r, "cookie-login")

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)

		w.Write([]byte(responses["internalError"]))
		return
	}

	var d DBDeveloper

	err = DecodeJSON(r.Body, &d)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		w.Write([]byte(responses["bad-json"]))
		return
	}

	isValid, err := CheckLoginIsValid(d)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusNotFound)

		w.Write([]byte(responses["userNotFound"]))
		return
	}

	if isValid {
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Values["user"] = d.Username
		session.Save(r, w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responses["sucess"]))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(responses["internalError"]))
	return
}

//LogoutDev .
func LogoutDev(w http.ResponseWriter, r *http.Request) {
	log.Print("LogoutDev requested")

	session, _ := store.Get(r, "cookie-login")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Values["user"] = ""
	session.Save(r, w)

	w.WriteHeader(http.StatusFound)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

//RegisterDev POST /developers/register
func RegisterDev(w http.ResponseWriter, r *http.Request) {
	log.Print("RegisterDev requested")

	var d DBDeveloper

	log.Printf("Body: %v", r.Body)

	err := DecodeJSON(r.Body, &d)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(responses["bad-json"]))
		return
	}

	log.Printf("Dev populado: %v", d)

	err = validate.Struct(d)
	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(responses["bad-validation"]))
		return
	}

	err = InsertDeveloper(d)
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

//SignUpDevForm GET /developers/sign_up
func SignUpDevForm(w http.ResponseWriter, r *http.Request) {
	log.Print("SignUpDevForm requested")

	absPath, _ := filepath.Abs("./iogames/templates/sign_up_dev.html")

	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
	return
}

//SignInDevForm .
func SignInDevForm(w http.ResponseWriter, r *http.Request) {
	log.Print("SignInDevForm requested")

	session, _ := store.Get(r, "cookie-login")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, "/developer/dashboard", http.StatusFound)
		return
	}

	absPath, _ := filepath.Abs("./iogames/templates/sign_in_dev.html")

	tmpl, err := template.ParseFiles(absPath)

	if err != nil {
		log.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(responses["internalError"]))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)

	return
}
