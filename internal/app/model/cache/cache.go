package cache

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

//Service cache object
type Service struct {
	Client     *redis.Client
	DefaultKey string
}

//InitClient instantiates redis client
func (s *Service) InitClient(address, key string) {
	s.Client = redis.NewClient(&redis.Options{
		Addr:         address,
		IdleTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		MaxRetries:   2,
	})

	s.DefaultKey = key
	return
}

//Get retrieves from cache
func (s *Service) Get(key string) (result string, err error) {
	s.Ping()

	res, err := s.Client.Get(key).Result()
	return res, err
}

//Set saves key:pair in cache
func (s *Service) Set(key string, keyValues string, expires time.Duration) error {
	s.Ping()
	err := s.Client.Set(key, keyValues, expires).Err()
	return err
}

//Expire sets expiration of key
func (s *Service) Expire(key string, expires time.Duration) error {
	s.Ping()
	err := s.Client.Expire(key, expires).Err()
	return err
}

//Delete key in cache
func (s *Service) Delete(key string) (int64, error) {
	s.Ping()
	res, err := s.Client.Del(key).Result()

	if err != nil {
		log.Printf("error in deleting in redis: %s", key)
		return 0, err
	}

	return res, nil
}

//Ping cache server
func (s *Service) Ping() {
	err := s.Client.Ping().Err()

	if err != nil {
		log.Println("unable to ping redis connection")
	}
}
