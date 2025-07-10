package qqmusic

type (
	Response[T any] struct {
		Code    int    `json:"code"`
		SubCode int    `json:"sub_code"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}
	UserSongListResponse struct {
		EncryptUin string `json:"encrypt_uin,omitempty"`
		Hostname   string `json:"hostname,omitempty"`
		Totoal     int    `json:"totoal,omitempty"`
		DissList   []Diss `json:"disslist,omitempty"`
	}
	SongListResponse struct {
		Code               int    `json:"code,omitempty"`
		Subcode            int    `json:"subcode,omitempty"`
		Msg                string `json:"msg,omitempty"`
		AccessedPlazaCache int    `json:"accessed_plaza_cache,omitempty"`
		AccessedFavbase    int    `json:"accessed_favbase,omitempty"`
		Login              string `json:"login,omitempty"`
		Cdnum              int    `json:"cdnum,omitempty"`
		Cdlist             []struct {
			AlbumPicMid    string `json:"album_pic_mid,omitempty"`
			Buynum         int    `json:"buynum,omitempty"`
			Cmtnum         int    `json:"cmtnum,omitempty"`
			Coveradurl     string `json:"coveradurl,omitempty"`
			Ctime          int    `json:"ctime,omitempty"`
			CurSongNum     int    `json:"cur_song_num,omitempty"`
			Desc           string `json:"desc,omitempty"`
			DirPicURL2     string `json:"dir_pic_url2,omitempty"`
			DirShow        int    `json:"dir_show,omitempty"`
			Dirid          int    `json:"dirid,omitempty"`
			Dissid         int    `json:"dissid,omitempty"`
			Dissname       string `json:"dissname,omitempty"`
			Disstid        string `json:"disstid,omitempty"`
			Disstype       int    `json:"disstype,omitempty"`
			EncryptUin     string `json:"encrypt_uin,omitempty"`
			Headurl        string `json:"headurl,omitempty"`
			Ifpicurl       string `json:"ifpicurl,omitempty"`
			IsAd           int    `json:"isAd,omitempty"`
			Isdj           int    `json:"isdj,omitempty"`
			Isvip          int    `json:"isvip,omitempty"`
			Login          string `json:"login,omitempty"`
			Logo           string `json:"logo,omitempty"`
			Mtime          int    `json:"mtime,omitempty"`
			Nick           string `json:"nick,omitempty"`
			Nickname       string `json:"nickname,omitempty"`
			Owndir         int    `json:"owndir,omitempty"`
			PicDpi         int    `json:"pic_dpi,omitempty"`
			PicMid         string `json:"pic_mid,omitempty"`
			Scoreavage     string `json:"scoreavage,omitempty"`
			Scoreusercount int    `json:"scoreusercount,omitempty"`
			Singerid       int    `json:"singerid,omitempty"`
			Singermid      string `json:"singermid,omitempty"`
			SongBegin      int    `json:"song_begin,omitempty"`
			SongUpdateNum  int    `json:"song_update_num,omitempty"`
			SongUpdateTime int    `json:"song_update_time,omitempty"`
			Songids        string `json:"songids,omitempty"`
			Songlist       []Song `json:"songlist,omitempty"`
			Songnum        int    `json:"songnum,omitempty"`
			Songtypes      string `json:"songtypes,omitempty"`
			// Tags           []string `json:"tags,omitempty"`
			TotalSongNum int    `json:"total_song_num,omitempty"`
			Type         int    `json:"type,omitempty"`
			Uin          string `json:"uin,omitempty"`
			Visitnum     int    `json:"visitnum,omitempty"`
		} `json:"cdlist,omitempty"`
		Realcdnum int `json:"realcdnum,omitempty"`
	}
	SongUrlResponse struct {
		Code int `json:"code,omitempty"`
		Req0 struct {
			Code int `json:"code,omitempty"`
			Data struct {
				DeviceResult string `json:"deviceResult,omitempty"`
				Expiration   int    `json:"expiration,omitempty"`
				LoginKey     string `json:"login_key,omitempty"`
				Midurlinfo   []struct {
					AuthSwitch        int    `json:"auth_switch,omitempty"`
					AuthSwitch2       int    `json:"auth_switch2,omitempty"`
					CommonDownfromtag int    `json:"common_downfromtag,omitempty"`
					Ekey              string `json:"ekey,omitempty"`
					Errtype           string `json:"errtype,omitempty"`
					Filename          string `json:"filename,omitempty"`
					Flowfromtag       string `json:"flowfromtag,omitempty"`
					Flowurl           string `json:"flowurl,omitempty"`
					Hisbuy            int    `json:"hisbuy,omitempty"`
					Hisdown           int    `json:"hisdown,omitempty"`
					Isbuy             int    `json:"isbuy,omitempty"`
					Isonly            int    `json:"isonly,omitempty"`
					Onecan            int    `json:"onecan,omitempty"`
					Opi128kurl        string `json:"opi128kurl,omitempty"`
					Opi192koggurl     string `json:"opi192koggurl,omitempty"`
					Opi192kurl        string `json:"opi192kurl,omitempty"`
					Opi30surl         string `json:"opi30surl,omitempty"`
					Opi48kurl         string `json:"opi48kurl,omitempty"`
					Opi96koggurl      string `json:"opi96koggurl,omitempty"`
					Opi96kurl         string `json:"opi96kurl,omitempty"`
					Opiflackurl       string `json:"opiflackurl,omitempty"`
					P2pfromtag        int    `json:"p2pfromtag,omitempty"`
					Pdl               int    `json:"pdl,omitempty"`
					Pneed             int    `json:"pneed,omitempty"`
					Pneedbuy          int    `json:"pneedbuy,omitempty"`
					Premain           int    `json:"premain,omitempty"`
					Purl              string `json:"purl,omitempty"`
					Qmdlfromtag       int    `json:"qmdlfromtag,omitempty"`
					Result            int    `json:"result,omitempty"`
					Songmid           string `json:"songmid,omitempty"`
					Subcode           int    `json:"subcode,omitempty"`
					Tips              string `json:"tips,omitempty"`
					UiAlert           int    `json:"uiAlert,omitempty"`
					VipDownfromtag    int    `json:"vip_downfromtag,omitempty"`
					Vkey              string `json:"vkey,omitempty"`
					Wififromtag       string `json:"wififromtag,omitempty"`
					Wifiurl           string `json:"wifiurl,omitempty"`
					Xcdnurl           string `json:"xcdnurl,omitempty"`
				} `json:"midurlinfo,omitempty"`
				Msg          string   `json:"msg,omitempty"`
				Retcode      int      `json:"retcode,omitempty"`
				Servercheck  string   `json:"servercheck,omitempty"`
				Sip          []string `json:"sip,omitempty"`
				Testfile2g   string   `json:"testfile2g,omitempty"`
				Testfilewifi string   `json:"testfilewifi,omitempty"`
				Thirdip      []string `json:"thirdip,omitempty"`
				Uin          string   `json:"uin,omitempty"`
				VerifyType   int      `json:"verify_type,omitempty"`
			} `json:"data"`
		} `json:"req_0"`
		StartTs int    `json:"start_ts,omitempty"`
		Traceid string `json:"traceid,omitempty"`
		Ts      int    `json:"ts,omitempty"`
	}
	SongLyricResponse struct {
		Code    int    `json:"code,omitempty"`
		Subcode int    `json:"subcode,omitempty"`
		Lyric   string `json:"lyric,omitempty"`
		Trans   string `json:"trans,omitempty"`
	}
)

type (
	Diss struct {
		DissName  string `json:"diss_name,omitempty"`
		DissCover string `json:"diss_cover,omitempty"`
		SongCnt   int    `json:"song_cnt,omitempty"`
		ListenNum int    `json:"listen_num,omitempty"`
		DirId     int    `json:"dirid,omitempty"`
		Tid       int    `json:"tid,omitempty"`
		DirShow   int    `json:"dir_show,omitempty"`
	}
	Song struct {
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
	Singer struct {
		ID   int    `json:"id,omitempty"`
		Mid  string `json:"mid,omitempty"`
		Name string `json:"name,omitempty"`
	}
)
