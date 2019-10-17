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

	//if Store != nil {
	//		return Store
	//	}
	Pool = newPool(host)
	SessionTimeOut = timeout
	Domain = Domain
	Store, err = redistore.NewRediStoreWithPool(Pool, []byte(secret))
	Store.DefaultMaxAge = SessionTimeOut
	return err
	//return Store
}

func GetSession(r *http.Request, sessName, sessKey string) (*AuthSession, bool) {

	if Store == nil {
		log.Fatal(fmt.Errorf("redis: redistore is null"))
	}

	session, err := Store.Get(r, sessName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("sess key : %s \n", sessKey)
	fmt.Printf("sess name: %s \n", sessName)
	fmt.Printf("store options %v \n ", Store.Options)
	fmt.Printf("debug sess %v \n ", session.Values[sessKey])

	obj, ok := session.Values[sessKey].(*AuthSession)
	fmt.Printf("debug obj %v \n ", obj)

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

	err = session.Save(r, w)
	fmt.Printf("save sess key : %s \n", sessKey)
	fmt.Printf("save sess name: %s \n", sessName)

	fmt.Printf("save session %v \n", session.Values[sessKey])
	fmt.Printf("saved obj %v \n", sessObj)
	return err
}
