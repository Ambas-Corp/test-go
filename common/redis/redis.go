package redis

import (
	"github.com/bns-engineering/platformbanking-card/common/logging"
	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)

//ConnectRedis connect to redis server
func ConnectRedis(host string, port string) *redis.Pool {
	logging.Infof("Connect to Redis : %s:%s", host, port)
	pool = &redis.Pool{
		MaxIdle:   50,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", host+port)

			// Connection error handling
			if err != nil {
				logging.Error(err, "fail initializing the redis pool")
				// os.Exit(1)
				return conn, err
			}
			return conn, err
		},
	}

	return pool
}

// getConn pool based on settings
func getConn() (redis.Conn, error) {
	conn := pool.Get()

	return conn, conn.Err()
}

//Clear - clear value based on single key
func Clear(key string) (err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = redis.Int64(conn.Do("DEL", key))
	return
}

//Get - get value from single key
func Get(key string) (res string, err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	res, err = redis.String(conn.Do("GET", key))
	return
}

//GetHash - get specific hash value from single key
func GetHash(key string, hash string) (res string, err error) {
	conn, err := getConn()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	res, err = redis.String(conn.Do("HGET", key, hash))
	return
}

//GetHashAll - get all hash value
func GetHashAll(key string) (res map[string]string, err error) {
	conn, err := getConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	res, err = redis.StringMap(conn.Do("HGETALL", key))
	return
}

//Set - set single value
func Set(key string, value string) (err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Send("SET", key, value)
	return conn.Flush()
}

//SetWithTTL - set single value with ttl
func SetWithTTL(key string, value string, ttl int) (err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	conn.Send("SET", key, value)
	if ttl > 0 {
		conn.Send("EXPIRE", key, ttl)
	}
	return conn.Flush()
}

//SetHash - set multiple value
func SetHash(key string, input map[string]string) (err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	values := []interface{}{key}
	for key, value := range input {
		values = append(values, key, value)
	}

	conn.Send("HMSET", values...)
	return conn.Flush()
}

//SetHashWithTTL - set multiple value with ttl
func SetHashWithTTL(key string, input map[string]string, ttl int) (err error) {
	conn, err := getConn()
	if err != nil {
		return
	}
	defer conn.Close()

	values := []interface{}{key}
	for key, value := range input {
		values = append(values, key, value)
	}

	conn.Send("HMSET", values...)
	if ttl > 0 {
		conn.Send("EXPIRE", key, ttl)
	}
	return conn.Flush()
}

//IsErrNoData to determine is error no data
func IsErrNoData(err error) bool {
	return err == redis.ErrNil
}
