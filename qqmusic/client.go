package qqmusic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/BYT0723/go-tools/transport/httpx"
)

type Client struct {
	cli     *httpx.Client
	cookie  string
	cookies map[string]string
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

	return &Client{
		cli:     httpx.NewClient(),
		cookie:  cookie,
		cookies: cookies,
	}, nil
}

func (c *Client) GetUserSongList() (*UserSongListResponse, error) {
	var resp Response[UserSongListResponse]
	if err := c.do(
		http.MethodGet,
		"https://c6.y.qq.com/rsc/fcgi-bin/fcg_user_created_diss",
		http.Header{},
		map[string]any{
			"hostuin":     c.cookies["uin"],
			"uin":         c.cookies["uin"],
			"sin":         0,
			"size":        200,
			"g_tk":        5381,
			"loginUin":    0,
			"format":      "json",
			"inCharset":   "utf-8",
			"outCharset":  "utf-8",
			"notice":      0,
			"platform":    "yqq.json",
			"needNewCode": 0,
		},
		&resp,
	); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("response code %d: %s", resp.Code, resp.Message)
	}
	return &resp.Data, nil
}

func (c *Client) GetSongList(tid int) (*SongListResponse, error) {
	var resp SongListResponse
	if err := c.do(
		http.MethodGet,
		"http://c.y.qq.com/qzone/fcg-bin/fcg_ucc_getcdinfo_byids_cp.fcg",
		http.Header{"Referer": []string{"https://y.qq.com/yqq/playlist"}},
		map[string]any{
			"type":     1,
			"utf8":     1,
			"disstid":  tid,
			"loginUin": 0,
		},
		&resp,
	); err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("response code %d", resp.Code)
	}
	return &resp, nil
}

func (c *Client) do(method, url string, header http.Header, payload map[string]any, response any) error {
	header.Set("Accept", "application/json")
	header.Set("Accept-Encoding", "gzip, deflate, br")
	header.Set("Cookie", c.cookie)
	header.Set("Referer", "https://y.qq.com/")
	header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:140.0) Gecko/20100101 Firefox/140.0")

	resp, err := c.cli.Do(context.Background(), method, url, httpx.WithHeader(header), httpx.WithPayload(payload))
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
		if bytes.HasPrefix(resp.Body, []byte("jsonCallback(")) {
			resp.Body = resp.Body[13 : len(resp.Body)-1]
		}
		fmt.Printf("resp.Body: %s\n", resp.Body)
		return json.Unmarshal(resp.Body, response)
	}
	return nil
}
