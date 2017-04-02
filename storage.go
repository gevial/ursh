package ursh

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Storage struct
type Storage struct {
}

// Connect to a storage
func Connect(host string, port int, user string, password string) (*Storage, error) {
	return &Storage{}, nil
}

func exampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
