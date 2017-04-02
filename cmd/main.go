package main

import (
	"fmt"

	"github.com/gevial/ursh"
)

func main() {
	c := ursh.Parse("./ursh.conf.yaml")
	fmt.Printf("%#v\n", c)
	// s, err := ursh.Connect(
	// 	c.Storage.Host, c.Storage.Port, c.Storage.User, c.Storage.Password)
}
