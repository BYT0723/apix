package qqmusic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"maps"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/BYT0723/go-tools/spider"
	"github.com/BYT0723/go-tools/transport/httpx"
)

type Client struct {
	cli     *httpx.Client
	cookie  string
	cookies map[string]string
	guid    string
}

func NewClient(cookie string) (*Client, error) {
	c, err := http.ParseCookie(cookie)
	if err != nil {
		return nil, err
	}

	cookies := make(map[string]string)
	for _, v := range c {
		cookies[v.Name] = v.Value
	}

	// 用当前时间做随机种子，不然每次都是一样的

	return &Client{
		cli:     httpx.NewClient(),
		cookie:  cookie,
		cookies: cookies,
		guid:    strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()),
	}, nil
}

func (c *Client) do(method, url string, header http.Header, payload map[string]any, response any) error {
	h := http.Header{}
	h.Set("Accept", "application/json")
	h.Set("Accept-Encoding", "gzip, deflate, br")
	h.Set("Cookie", c.cookie)
	h.Set("User-Agent", spider.GetRandUserAgent())

	if header != nil {
		maps.Copy(h, header)
	}

	resp, err := c.cli.Do(context.Background(), method, url, httpx.WithHeader(h), httpx.WithPayload(payload))
	if err != nil {
		return err
	}
	if resp.Code != http.StatusOK {
		return errors.New("resp code not OK")
	}
	switch v := response.(type) {
	case *[]byte:
		*v = resp.Body
	default:
		resp.Body = bytes.TrimSpace(resp.Body)
		switch {
		case bytes.HasPrefix(resp.Body, []byte("jsonCallback(")):
			resp.Body = resp.Body[13 : len(resp.Body)-1]
		case bytes.HasPrefix(resp.Body, []byte("MusicJsonCallback(")):
			resp.Body = resp.Body[18 : len(resp.Body)-1]
		}
		return json.Unmarshal(resp.Body, response)
	}
	return nil
}
