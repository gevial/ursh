package ursh_test

import (
	"testing"
	"time"

	"github.com/gevial/ursh"
)

func TestStorage(t *testing.T) {
	// Run Redis in Docker locally first
	s, _ := ursh.Connect("localhost", 6379, "", "")
	s.Put(&url, e)
}

var url = ursh.URL{
	URL:      "https://vk.com/a/b/c/d185728238/",
	ShortURL: "https://localhost:10101/vkco1a",
}
var e = time.Duration(0)
