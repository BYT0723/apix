package qqmusic

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/BYT0723/go-tools/ds"
	"github.com/tidwall/gjson"
)

var ErrGetSongUrlFailed = errors.New("获取链接失败")

type SongType int

const (
	SongTypeM4A SongType = iota
	SongType128
	SongType320
	SongTypeAPE
	SongTypeFLAC
)

func (t SongType) String() string {
	switch t {
	case SongTypeM4A:
		return "m4a"
	case SongType128:
		return "128"
	case SongType320:
		return "320"
	case SongTypeAPE:
		return "ape"
	case SongTypeFLAC:
		return "flac"
	}
	return ""
}

// 如果指定类型无法获取链接则按照此顺序尝试
var songTypes = []SongType{
	SongTypeFLAC,
	SongTypeAPE,
	SongType320,
	SongType128,
	SongTypeM4A,
}

func (t SongType) Suffix() string {
	switch t {
	case SongTypeM4A:
		return ".m4a"
	case SongType128:
		return ".mp3"
	case SongType320:
		return ".mp3"
	case SongTypeAPE:
		return ".ape"
	case SongTypeFLAC:
		return ".flac"
	}
	return ""
}

func (t SongType) Prefix() string {
	switch t {
	case SongTypeM4A:
		return "C400"
	case SongType128:
		return "M500"
	case SongType320:
		return "M800"
	case SongTypeAPE:
		return "A000"
	case SongTypeFLAC:
		return "F000"
	}
	return ""
}

var rd = rand.New(rand.NewSource(time.Now().UnixNano()))

func (c *Client) GetSongUrl(mid, mediaId string, t SongType) (url string, rt SongType, err error) {
	tset := ds.NewHashSet[SongType]()

	url, err = c.getSongUrl(mid, mediaId, t)
	if err == nil {
		rt = t
		return
	}

	tset.Append(t)

	for _, t := range songTypes {
		if tset.Contains(t) {
			continue
		}
		tset.Append(t)

		// 随机睡眠1s - 3s
		time.Sleep(c.Interval + time.Duration(rd.Int63n(int64(c.IntervalInc))))

		if url, err = c.getSongUrl(mid, mediaId, t); err == nil {
			rt = t
			return
		}
	}
	err = ErrGetSongUrlFailed
	return
}

func (c *Client) getSongUrl(mid, mediaId string, t SongType) (url string, err error) {
	var resp SongUrlResponse

	if mediaId == "" {
		mediaId = mid
	}

	data, err := json.Marshal(map[string]any{
		"req_0": map[string]any{
			"module": "vkey.GetVkeyServer",
			"method": "CgiGetVkey",
			"param": map[string]any{
				"filename":  []string{fmt.Sprintf("%s%s%s%s", t.Prefix(), mid, mediaId, t.Suffix())},
				"guid":      c.guid,
				"songmid":   []string{mid},
				"songtype":  []int{0},
				"uin":       c.cookies["uin"],
				"loginflag": 1,
				"platform":  "20",
			},
		},
		"comm": map[string]any{
			"uin":    c.cookies["uin"],
			"format": "json",
			"ct":     19,
			"cv":     0,
			"authst": c.cookies["qqmusic_key"],
		},
	})
	if err != nil {
		return
	}

	for range c.Retry {
		if err = c.do(
			http.MethodGet,
			"https://u.y.qq.com/cgi-bin/musicu.fcg",
			nil,
			map[string]any{
				"-":           "getplaysongvkey",
				"g_tk":        5381,
				"loginUin":    c.cookies["uin"],
				"hostUin":     0,
				"format":      "json",
				"inCharset":   "utf8",
				"outCharset":  "utf-8¬ice=0",
				"platform":    "yqq.json",
				"needNewCode": 0,
				"data":        string(data),
			},
			&resp,
		); err != nil {
			return
		}

		if resp.Code != 0 {
			err = fmt.Errorf("response code %d", resp.Code)
			return
		}

		if resp.Req0.Code != 0 {
			err = fmt.Errorf("response code %d", resp.Req0.Code)
			return
		}
		if len(resp.Req0.Data.Midurlinfo) == 0 || resp.Req0.Data.Midurlinfo[0].Purl == "" {
			time.Sleep(c.Interval + time.Duration(rd.Int63n(int64(c.IntervalInc))))
			continue
		}

		index := slices.IndexFunc(resp.Req0.Data.Sip, func(s string) bool {
			return !strings.HasPrefix(s, "http://ws")
		})
		url = resp.Req0.Data.Sip[max(0, index)] + resp.Req0.Data.Midurlinfo[0].Purl
		return
	}
	err = ErrGetSongUrlFailed
	return
}

func (c *Client) GetSongLyric(mid string) (lyric, trans []byte, err error) {
	var resp SongLyricResponse
	if err = c.do(
		http.MethodGet,
		"http://c.y.qq.com/lyric/fcgi-bin/fcg_query_lyric_new.fcg",
		http.Header{
			"Referer": []string{"https://y.qq.com"},
		},
		map[string]any{
			"songmid":     mid,
			"pcachetime":  time.Now().UnixMilli(),
			"g_tk":        5381,
			"loginUin":    0,
			"hostUin":     0,
			"inCharset":   "utf8",
			"outCharset":  "utf-8",
			"notice":      0,
			"platform":    "yqq",
			"needNewCode": 0,
		},
		&resp,
	); err != nil {
		return
	}
	if resp.Code != 0 {
		err = fmt.Errorf("response code %d", resp.Code)
		return
	}

	lyric, _ = base64.StdEncoding.DecodeString(resp.Lyric)
	trans, _ = base64.StdEncoding.DecodeString(resp.Trans)
	return
}

func (c *Client) GetAlbumArtBySongMid(mid string) (url string, err error) {
	var resp []byte

	data, err := json.Marshal(map[string]any{
		"songinfo": map[string]any{
			"method": "get_song_detail_yqq",
			"module": "music.pf_song_detail_svr",
			"param": map[string]any{
				"song_mid": mid,
			},
		},
	})
	if err != nil {
		return
	}

	if err = c.do(
		http.MethodGet,
		"http://u.y.qq.com/cgi-bin/musicu.fcg",
		http.Header{
			"Referer": []string{"https://y.qq.com"},
		},
		map[string]any{
			"data": string(data),
		},
		&resp,
	); err != nil {
		return
	}

	albumMid := gjson.GetBytes(resp, "songinfo.data.track_info.album.mid")
	if !albumMid.Exists() {
		err = errors.New("获取专辑信息失败")
		return
	}

	url = fmt.Sprintf("https://y.gtimg.cn/music/photo_new/T002R300x300M000%s.jpg", albumMid.String())

	return
}
