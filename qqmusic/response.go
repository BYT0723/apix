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

type Song struct {
	Albumdesc string `json:"albumdesc,omitempty"`
	Albumid   int    `json:"albumid,omitempty"`
	Albummid  string `json:"albummid,omitempty"`
	Albumname string `json:"albumname,omitempty"`
	Alertid   int    `json:"alertid,omitempty"`
	BelongCD  int    `json:"belongCD,omitempty"`
	CdIdx     int    `json:"cdIdx,omitempty"`
	Interval  int    `json:"interval,omitempty"`
	Isonly    int    `json:"isonly,omitempty"`
	Label     string `json:"label,omitempty"`
	Msgid     int    `json:"msgid,omitempty"`
	Pay       struct {
		Payalbum      int `json:"payalbum,omitempty"`
		Payalbumprice int `json:"payalbumprice,omitempty"`
		Paydownload   int `json:"paydownload,omitempty"`
		Payinfo       int `json:"payinfo,omitempty"`
		Payplay       int `json:"payplay,omitempty"`
		Paytrackmouth int `json:"paytrackmouth,omitempty"`
		Paytrackprice int `json:"paytrackprice,omitempty"`
		Timefree      int `json:"timefree,omitempty"`
	} `json:"pay"`
	Preview struct {
		Trybegin int `json:"trybegin,omitempty"`
		Tryend   int `json:"tryend,omitempty"`
		Trysize  int `json:"trysize,omitempty"`
	} `json:"preview"`
	Rate        int      `json:"rate,omitempty"`
	Singer      []Singer `json:"singer,omitempty"`
	Size128     int      `json:"size128,omitempty"`
	Size320     int      `json:"size320,omitempty"`
	Size51      int      `json:"size5_1,omitempty"`
	Sizeape     int      `json:"sizeape,omitempty"`
	Sizeflac    int      `json:"sizeflac,omitempty"`
	Sizeogg     int      `json:"sizeogg,omitempty"`
	Songid      int      `json:"songid,omitempty"`
	Songmid     string   `json:"songmid,omitempty"`
	Songname    string   `json:"songname,omitempty"`
	Songorig    string   `json:"songorig,omitempty"`
	Songtype    int      `json:"songtype,omitempty"`
	StrMediaMid string   `json:"strMediaMid,omitempty"`
	Stream      int      `json:"stream,omitempty"`
	Switch      int      `json:"switch,omitempty"`
	Type        int      `json:"type,omitempty"`
	Vid         string   `json:"vid,omitempty"`
}

type Singer struct {
	ID   int    `json:"id,omitempty"`
	Mid  string `json:"mid,omitempty"`
	Name string `json:"name,omitempty"`
}

type SongListResponse struct {
	Code               int    `json:"code,omitempty"`
	Subcode            int    `json:"subcode,omitempty"`
	AccessedPlazaCache int    `json:"accessed_plaza_cache,omitempty"`
	AccessedFavbase    int    `json:"accessed_favbase,omitempty"`
	Login              string `json:"login,omitempty"`
	Cdnum              int    `json:"cdnum,omitempty"`
	Cdlist             []struct {
		AlbumPicMid    string   `json:"album_pic_mid,omitempty"`
		Buynum         int      `json:"buynum,omitempty"`
		Cmtnum         int      `json:"cmtnum,omitempty"`
		Coveradurl     string   `json:"coveradurl,omitempty"`
		Ctime          int      `json:"ctime,omitempty"`
		CurSongNum     int      `json:"cur_song_num,omitempty"`
		Desc           string   `json:"desc,omitempty"`
		DirPicURL2     string   `json:"dir_pic_url2,omitempty"`
		DirShow        int      `json:"dir_show,omitempty"`
		Dirid          int      `json:"dirid,omitempty"`
		Dissid         int      `json:"dissid,omitempty"`
		Dissname       string   `json:"dissname,omitempty"`
		Disstid        string   `json:"disstid,omitempty"`
		Disstype       int      `json:"disstype,omitempty"`
		EncryptUin     string   `json:"encrypt_uin,omitempty"`
		Headurl        string   `json:"headurl,omitempty"`
		Ifpicurl       string   `json:"ifpicurl,omitempty"`
		IsAd           int      `json:"isAd,omitempty"`
		Isdj           int      `json:"isdj,omitempty"`
		Isvip          int      `json:"isvip,omitempty"`
		Login          string   `json:"login,omitempty"`
		Logo           string   `json:"logo,omitempty"`
		Mtime          int      `json:"mtime,omitempty"`
		Nick           string   `json:"nick,omitempty"`
		Nickname       string   `json:"nickname,omitempty"`
		Owndir         int      `json:"owndir,omitempty"`
		PicDpi         int      `json:"pic_dpi,omitempty"`
		PicMid         string   `json:"pic_mid,omitempty"`
		Scoreavage     string   `json:"scoreavage,omitempty"`
		Scoreusercount int      `json:"scoreusercount,omitempty"`
		Singerid       int      `json:"singerid,omitempty"`
		Singermid      string   `json:"singermid,omitempty"`
		SongBegin      int      `json:"song_begin,omitempty"`
		SongUpdateNum  int      `json:"song_update_num,omitempty"`
		SongUpdateTime int      `json:"song_update_time,omitempty"`
		Songids        string   `json:"songids,omitempty"`
		Songlist       []Song   `json:"songlist,omitempty"`
		Songnum        int      `json:"songnum,omitempty"`
		Songtypes      string   `json:"songtypes,omitempty"`
		Tags           []string `json:"tags,omitempty"`
		TotalSongNum   int      `json:"total_song_num,omitempty"`
		Type           int      `json:"type,omitempty"`
		Uin            string   `json:"uin,omitempty"`
		Visitnum       int      `json:"visitnum,omitempty"`
	} `json:"cdlist,omitempty"`
	Realcdnum int `json:"realcdnum,omitempty"`
}
