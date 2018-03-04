package session

import (
	"crypto/rand"
	"fmt"
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

func NewSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	sess, err := store.New(r, SessionName)
	if err != nil {
		return nil, err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 2,
		HttpOnly: true,
	}
	return sess, nil
}

func SessionExists(r *http.Request, name string) (bool, error) {
	sess, err := store.Get(r, name)
	if err != nil {
		return false, err
	}
	return !sess.IsNew, nil
}

func GetStrValue(r *http.Request, value string) (string, error) {
	sess, err := store.Get(r, SessionName)
	if err != nil {
		return "", err
	}

	// Perform type assertion
	if s, ok := sess.Values[value].(string); !ok {
		return "", fmt.Errorf("error performing type assertion on %v", sess.Values[value])
	} else {
		if s == "" {
			return "", fmt.Errorf("value was either nil or never set")
		}
		return s, nil
	}
}
