package helper

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boj/redistore"
	"github.com/gomodule/redigo/redis"

	"github.com/gorilla/sessions"
)

var (
	Pool   *redis.Pool
	Store  *redistore.RediStore
	Domain string
	//Secret         string
	SessionTimeOut int
)

type AuthSession struct {
	Session map[string]interface{}
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func InitSessionStore(host, domain, secret string, timeout int) error {
	var err error

	Pool = newPool(host)
	SessionTimeOut = timeout
	Domain = Domain
	Store, err = redistore.NewRediStoreWithPool(Pool, []byte(secret))
	Store.DefaultMaxAge = SessionTimeOut
	return err
}

func GetSession(r *http.Request, sessName, sessKey string) (*AuthSession, bool) {

	if Store == nil {
		log.Fatal(fmt.Errorf("redis: redistore is null"))
	}

	fmt.Printf("request: %v", r.Cookies())
	session, err := Store.Get(r, sessName)
	if err != nil {
		log.Fatal(err.Error())
	}

	obj, ok := session.Values[sessKey].(*AuthSession)
	return obj, ok
}
func SaveSession(r *http.Request, w http.ResponseWriter, sessName, sessKey string, sessObj *AuthSession) error {
	if Store == nil {
		log.Fatal(fmt.Errorf("redis: redistore is null"))
	}
	session, err := Store.Get(r, sessName)
	if err != nil {
		log.Fatal(err.Error())
	}
	session.Values[sessKey] = sessObj
	session.Options = &sessions.Options{
		Domain:   Domain,
		Path:     "/",
		MaxAge:   SessionTimeOut,
		HttpOnly: true,
	}

	err = sessions.Save(r, w)
	return err
}
