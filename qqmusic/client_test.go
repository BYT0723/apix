package qqmusic

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	b, err := os.ReadFile("cookie.txt")
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewClient(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatal(err)
	}

	sl, err := c.GetUserSongList()
	if err != nil {
		t.Fatal(err)
	}

	for _, d := range sl.DissList {
		l, err := c.GetSongList(d.Tid)
		if err != nil {
			t.Fatal(err)
		}
		if len(l.Cdlist) == 0 {
			continue
		}
		for _, v := range l.Cdlist {
			for _, s := range v.Songlist {
				fmt.Printf("s: %v\n", s)
			}
		}
	}
}
