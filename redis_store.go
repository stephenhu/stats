package stats

import (

	"github.com/gomodule/redigo/redis"
)

var Red redis.Conn


func ConnectRedis(protocol string, addr string) {

	c, err := redis.Dial(protocol, addr)

	if err != nil {
		logf("ConnectRedis", err.Error())
	} else {

		defer c.Close()

		Red = c

	}

} // ConnectRedis
