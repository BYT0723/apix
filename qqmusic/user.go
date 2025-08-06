package qqmusic

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

// GetUserSongList 获取用户歌单
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

	var favExist bool
	for _, d := range resp.Data.DissList {
		if d.DirId == 201 {
			favExist = true
		}
	}
	if !favExist {
		udr, err := c.GetUserDetail()
		if err == nil && len(udr.Data.Mymusic) > 0 {
			fav := udr.Data.Mymusic[0]
			id, err := strconv.Atoi(fav.ID)
			if err == nil {
				resp.Data.DissList = append(resp.Data.DissList, Diss{
					DissName:  fav.Title,
					DissCover: "http://y.gtimg.cn/mediastyle/y/img/cover_qzone_130.jpg",
					SongCnt:   fav.Num0,
					ListenNum: 0,
					Tid:       id,
					DirId:     201,
					DirShow:   1,
				})
			}
		}
	}
	return &resp.Data, nil
}

func (c *Client) GetUserDetail() (*UserDetailResponse, error) {
	var resp UserDetailResponse

	if err := c.do(
		http.MethodGet,
		"https://c6.y.qq.com/rsc/fcgi-bin/fcg_get_profile_homepage.fcg",
		http.Header{},
		map[string]any{
			"cid":     "205360838",
			"userid":  c.cookies["uin"],
			"reqfrom": 1,
		},
		&resp,
	); err != nil {
		return nil, err
	}

	return &resp, nil
}
