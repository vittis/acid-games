package iogames

import "github.com/gorilla/sessions"

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

//SetupSessions .
func SetupSessions() {
	//30 minutes
	store.MaxAge(60 * 30)

}
