package redisSession

import (
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	session *RedisSession
	prefix  string
	pool    *redis.Pool
)

func init() {
	var e error
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   50,
		Wait:        true,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "117.50.95.182:6379")
			if err != nil {
				e = err
				return nil, err
			}

			if _, err := c.Do("AUTH", "dev123"); err != nil {
				c.Close()
				e = err
				return nil, err
			}

			return c, nil
		},
	}
	prefix = "testing"

	if e != nil {
		str := fmt.Sprintf("create redis pool has error: %v", e)
		panic(str)
	}

	session = NewRedisSessionWithPool(pool)
}

func TestGet(t *testing.T) {
	session.SetPrefix(prefix)
	session.AddPrefix("get_test")
	//session.Set("1", "1")
	//session.Set("2", "2")
	//session.Set("3", "3")
	//session.Set("4", "4")
	//session.Set("5", "5")
	//session.Set("6", "6")

	s, e := session.Get("11")
	if e != nil {
		t.Logf("%v\n", e)
	} else {
		t.Logf("%s\n", s)
	}
}

func TestMGet(t *testing.T) {
	session.SetPrefix(prefix)
	session.AddPrefix("mget_test")
	//session.Set("1", "1")
	//session.Set("2", "2")
	//session.Set("3", "3")
	//session.Set("4", "4")
	//session.Set("5", "5")
	//session.Set("6", "6")

	r, e := session.MGet([]string{"1", "2", "3", "4", "5", "6", "7", "8"})
	if e != nil {
		t.Logf("%v\n", e)
	} else {
		for _, v := range r {
			t.Logf("%s\n", v)
		}
	}
}

func TestPrefix(t *testing.T) {
	session.SetPrefix(prefix)
	key := session.AddPrefix("muppets")
	if key != "testing:muppets" {
		t.Errorf("Key must be formed correctly")
	}
}
