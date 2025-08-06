package qqmusic

import (
	"fmt"
	"net/http"
)

// GetSongList 获取歌单详情
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
