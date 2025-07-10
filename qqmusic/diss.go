package qqmusic

import (
	"fmt"
	"net/http"
	"slices"
)

func (c *Client) GetUserSongList() (*UserSongListResponse, error) {
	var resp Response[UserSongListResponse]
	if err := c.do(
		http.MethodGet,
		"https://c6.y.qq.com/rsc/fcgi-bin/fcg_user_created_diss",
		http.Header{
			"Referer": []string{"https://y.qq.com/"},
		},
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
	resp.Data.DissList = slices.DeleteFunc(resp.Data.DissList, func(e Diss) bool { return e.DirShow == 0 })
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
		return nil, fmt.Errorf("response code %d, %s", resp.Code, resp.Msg)
	}
	return &resp, nil
}
