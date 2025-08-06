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

type UserDetailResponse struct {
	Code int `json:"code,omitempty"`
	Data struct {
		Creator struct {
			Backpic struct {
				Picurl string `json:"picurl,omitempty"`
				Title  string `json:"title,omitempty"`
				Type   int    `json:"type,omitempty"`
			} `json:"backpic,omitempty"`
			BuyLock int `json:"buy_lock,omitempty"`
			Cfinfo  struct {
				CfinfoBykey struct {
					UrlKey    string `json:"url_key,omitempty"`
					UrlParams string `json:"url_params,omitempty"`
				} `json:"cfinfo_bykey,omitempty"`
				Jumpkey string `json:"jumpkey,omitempty"`
				Jumpurl string `json:"jumpurl,omitempty"`
				Similar int    `json:"similar,omitempty"`
				Title   string `json:"title,omitempty"`
			} `json:"cfinfo,omitempty"`
			DissLock    int    `json:"diss_lock,omitempty"`
			EncryptUin  string `json:"encrypt_uin,omitempty"`
			Extra       string `json:"extra,omitempty"`
			FavLock     int    `json:"fav_lock,omitempty"`
			Forbidden   int    `json:"forbidden,omitempty"`
			Headpic     string `json:"headpic,omitempty"`
			Ifpic       string `json:"ifpic,omitempty"`
			IsBindWeibo int    `json:"is_bind_weibo,omitempty"`
			Isfollow    int    `json:"isfollow,omitempty"`
			Ishost      int    `json:"ishost,omitempty"`
			Islock      int    `json:"islock,omitempty"`
			Jumpkey     string `json:"jumpkey,omitempty"`
			Listeninfo  struct {
				Iconurl     string `json:"iconurl,omitempty"`
				Jumpkey     string `json:"jumpkey,omitempty"`
				Jumpurl     string `json:"jumpurl,omitempty"`
				ListenBykey struct {
					UrlKey    string `json:"url_key,omitempty"`
					UrlParams string `json:"url_params,omitempty"`
				} `json:"listen_bykey,omitempty"`
			} `json:"listeninfo,omitempty"`
			Lvinfo []struct {
				Iconurl     string `json:"iconurl,omitempty"`
				Jumpkey     string `json:"jumpkey,omitempty"`
				Jumpurl     string `json:"jumpurl,omitempty"`
				LvinfoBykey struct {
					UrlKey    string `json:"url_key,omitempty"`
					UrlParams string `json:"url_params,omitempty"`
				} `json:"lvinfo_bykey,omitempty"`
			} `json:"lvinfo,omitempty"`
			Medal struct {
				Flag       int    `json:"flag,omitempty"`
				Iconurl    string `json:"iconurl,omitempty"`
				Jumpkey    string `json:"jumpkey,omitempty"`
				Jumpurl    string `json:"jumpurl,omitempty"`
				MedalBykey struct {
					UrlKey    string `json:"url_key,omitempty"`
					UrlParams string `json:"url_params,omitempty"`
				} `json:"medal_bykey,omitempty"`
			} `json:"medal,omitempty"`
			Nick string `json:"nick,omitempty"`
			Nums struct {
				Fansnum         int `json:"fansnum,omitempty"`
				Follownum       int `json:"follownum,omitempty"`
				Followsingernum int `json:"followsingernum,omitempty"`
				Followusernum   int `json:"followusernum,omitempty"`
				Frdnum          int `json:"frdnum,omitempty"`
				Visitornum      int `json:"visitornum,omitempty"`
			} `json:"nums,omitempty"`
			ShareBykey struct {
				UrlKey    string `json:"url_key,omitempty"`
				UrlParams string `json:"url_params,omitempty"`
			} `json:"share_bykey,omitempty"`
			Shareurl   string `json:"shareurl,omitempty"`
			Singerinfo struct {
				Singerid int `json:"singerid,omitempty"`
			} `json:"singerinfo,omitempty"`
			Typeinfo struct {
				CfinfoBykey struct {
					UrlKey    string `json:"url_key,omitempty"`
					UrlParams string `json:"url_params,omitempty"`
				} `json:"cfinfo_bykey,omitempty"`
				Iconurl string `json:"iconurl,omitempty"`
				Jumpkey string `json:"jumpkey,omitempty"`
				Jumpurl string `json:"jumpurl,omitempty"`
				Type    int    `json:"type,omitempty"`
			} `json:"typeinfo,omitempty"`
			Uin        int    `json:"uin,omitempty"`
			UinWeb     string `json:"uin_web,omitempty"`
			UserInfoUI struct {
				Iconlist []struct {
					Desc   string `json:"desc,omitempty"`
					Ext    string `json:"ext,omitempty"`
					Height int    `json:"height,omitempty"`
					SrcUrl string `json:"srcUrl,omitempty"`
					Style  string `json:"style,omitempty"`
					Width  int    `json:"width,omitempty"`
				} `json:"iconlist,omitempty"`
				Nickname struct {
					DarkColor  string `json:"darkColor,omitempty"`
					LightColor string `json:"lightColor,omitempty"`
				} `json:"nickname,omitempty"`
			} `json:"userInfoUI,omitempty"`
			WeiboNick string `json:"weibo_nick,omitempty"`
			WeiboUid  string `json:"weibo_uid,omitempty"`
		} `json:"creator,omitempty"`
		Myarticle struct {
			Jumpkey  string        `json:"jumpkey,omitempty"`
			Jumpurl  string        `json:"jumpurl,omitempty"`
			Laypic   string        `json:"laypic,omitempty"`
			List     []interface{} `json:"list,omitempty"`
			Title    string        `json:"title,omitempty"`
			Totalcnt int           `json:"totalcnt,omitempty"`
		} `json:"myarticle,omitempty"`
		Mydiss struct {
			Jumpurl string `json:"jumpurl,omitempty"`
			Laypic  string `json:"laypic,omitempty"`
			List    []struct {
				DirShow  int    `json:"dir_show,omitempty"`
				Dirid    int    `json:"dirid,omitempty"`
				Dissid   int    `json:"dissid,omitempty"`
				Icontype int    `json:"icontype,omitempty"`
				Iconurl  string `json:"iconurl,omitempty"`
				Isshow   int    `json:"isshow,omitempty"`
				Picurl   string `json:"picurl,omitempty"`
				Subtitle string `json:"subtitle,omitempty"`
				Title    string `json:"title,omitempty"`
			} `json:"list,omitempty"`
			Num   int    `json:"num,omitempty"`
			Title string `json:"title,omitempty"`
		} `json:"mydiss,omitempty"`
		Mymusic []struct {
			ID         string `json:"id,omitempty"`
			Jumpkey    string `json:"jumpkey,omitempty"`
			Jumptype   int    `json:"jumptype,omitempty"`
			Jumpurl    string `json:"jumpurl,omitempty"`
			Laypic     string `json:"laypic,omitempty"`
			MusicBykey struct {
				UrlKey    string `json:"url_key,omitempty"`
				UrlParams string `json:"url_params,omitempty"`
			} `json:"music_bykey,omitempty"`
			Num0     int    `json:"num0,omitempty"`
			Num1     int    `json:"num1,omitempty"`
			Num2     int    `json:"num2,omitempty"`
			Picurl   string `json:"picurl,omitempty"`
			Subtitle string `json:"subtitle,omitempty"`
			Title    string `json:"title,omitempty"`
			Type     int    `json:"type,omitempty"`
		} `json:"mymusic,omitempty"`
		Mymusictype string `json:"mymusictype,omitempty"`
		Myradio     struct {
			Jumpkey  string        `json:"jumpkey,omitempty"`
			Jumpurl  string        `json:"jumpurl,omitempty"`
			Laypic   string        `json:"laypic,omitempty"`
			List     []interface{} `json:"list,omitempty"`
			Title    string        `json:"title,omitempty"`
			Totalcnt int           `json:"totalcnt,omitempty"`
		} `json:"myradio,omitempty"`
		Video struct {
			Jumpkey string        `json:"jumpkey,omitempty"`
			Jumpurl string        `json:"jumpurl,omitempty"`
			List    []interface{} `json:"list,omitempty"`
			Num     int           `json:"num,omitempty"`
			Title   string        `json:"title,omitempty"`
		} `json:"video,omitempty"`
	} `json:"data,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Subcode int    `json:"subcode,omitempty"`
}
