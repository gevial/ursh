package ursh

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Storage struct
type Storage struct {
	Client redis.UniversalClient
}

// Connect to a Storage
func Connect(host string, port int, user string, password string) (*Storage, error) {
	rclient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
	})
	_, err := rclient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &Storage{Client: rclient}, nil
}

// Put URL struct to Storage
func (s *Storage) Put(u *URL, e time.Duration) error {
	err := s.Client.Set(u.ShortURL, u.URL, e).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get long URL using short URL
func (s *Storage) Get(shortURL string) (*URL, error) {
	longURL, err := s.Client.Get(shortURL).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &URL{ShortURL: shortURL, URL: longURL}, nil
	}
}
