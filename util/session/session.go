package session

import (
	"crypto/rand"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "user-session"
)

var store = newStore()

func newStore() *sessions.CookieStore {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal("error generating random byte array: ", err)
		return nil
	}
	store := sessions.NewCookieStore([]byte("foo"))
	return store
}

func NewSession(w http.ResponseWriter, r *http.Request) error {
	sess, err := store.New(r, SessionName)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 2,
		HttpOnly: true,
	}
	sess.Save(r, w)
	return nil
}

func SessionExists(r *http.Request, name string) (bool, error) {
	sess, err := store.Get(r, name)
	if err != nil {
		return false, err
	}
	return !sess.IsNew, nil
}
