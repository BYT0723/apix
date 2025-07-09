package qqmusic

type Diss struct {
	DissName  string `json:"diss_name,omitempty"`
	DissCover string `json:"diss_cover,omitempty"`
	SongCnt   int    `json:"song_cnt,omitempty"`
	ListenNum int    `json:"listen_num,omitempty"`
	DirId     int    `json:"dirid,omitempty"`
	Tid       int    `json:"tid,omitempty"`
	DirShow   int    `json:"dir_show,omitempty"`
}
