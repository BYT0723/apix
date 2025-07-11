package qqmusic

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/BYT0723/go-tools/transport/httpx"
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
			if len(v.Songlist) == 0 {
				continue
			}
			var (
				s           = v.Songlist[0]
				nameBuilder strings.Builder
				n           = len(s.Singer)
				st          = SongType320
				sdir        = "./songs"
				ldir        = "./lyrics"
			)

			os.MkdirAll(sdir, os.ModePerm)
			os.MkdirAll(ldir, os.ModePerm)

			nameBuilder.WriteString(s.Songname)
			nameBuilder.WriteString(" - ")

			for i, s := range s.Singer {
				nameBuilder.WriteString(s.Name)
				if i < n-1 {
					nameBuilder.WriteString(",")
				}
			}
			addr, err := c.GetSongUrl(s.Songmid, s.StrMediaMid, st)
			if err != nil {
				t.Fatal(err)
			}

			if err := httpx.Download(addr, filepath.Join(sdir, nameBuilder.String()+st.Suffix())); err != nil {
				t.Fatal(err)
			}

			lyric, trans, err := c.GetSongLyric(s.Songmid)
			if err != nil {
				t.Fatal(err)
			}

			os.WriteFile(filepath.Join(ldir, nameBuilder.String()+".lrc"), lyric, os.ModePerm)
			os.WriteFile(filepath.Join(ldir, nameBuilder.String()+"_trans.lrc"), trans, os.ModePerm)
			fmt.Println(nameBuilder.String(), "Downloaded.")
		}
	}
}
