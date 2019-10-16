package helper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/boj/redistore"

	"github.com/gorilla/sessions"
)

type AuthSession struct {
	Session map[string]interface{}
}

func GetSession(rediStore *redistore.RediStore, r *http.Request, sessName, sessKey string) (*AuthSession, bool) {

	session, err := rediStore.Get(r, sessName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("debug sess %v", session.Values[sessKey])

	obj, ok := session.Values[sessKey].(*AuthSession)
	return obj, ok
}
func SaveSession(rediStore *redistore.RediStore, r *http.Request, w http.ResponseWriter, domain, sessName, sessKey string, sessionTimeout int, sessObj interface{}) error {
	session, err := rediStore.Get(r, sessName)
	if err != nil {
		log.Fatal(err.Error())
	}

	session.Values[sessKey] = sessObj
	session.Options = &sessions.Options{
		Domain:   domain,
		Path:     "/",
		MaxAge:   sessionTimeout,
		HttpOnly: true,
	}

	err = session.Save(r, w)
	return err
}
