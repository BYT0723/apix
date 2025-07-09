package qqmusic

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	b, err := os.ReadFile("cookie.txt")
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewClient(string(b))
	if err != nil {
		t.Fatal(err)
	}
	c.GetUserSongList()
}
