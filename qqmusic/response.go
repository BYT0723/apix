package qqmusic

type Response[T any] struct {
	Code    int    `json:"code,omitempty"`
	SubCode int    `json:"subcode,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type UserSongListResponse struct {
	EncryptUin string `json:"encrypt_uin,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	Totoal     int    `json:"totoal,omitempty"`
	DissList   []Diss `json:"disslist,omitempty"`
}
