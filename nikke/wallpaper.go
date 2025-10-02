package nikke

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BYT0723/go-tools/packer"
	"github.com/BYT0723/go-tools/transport/httpx"
)

type WallpaperType uint8

const (
	WallpaperTypeHorizontal WallpaperType = iota
	WallpaperTypeVertical1080p
	WallpaperTypeVertical2k
)

func GetStaticWallpaper(_type WallpaperType) (infos []*WallpaperInfo, err error) {
	var (
		offset, size = 0, 12
		total        int
		direct       string
	)
	switch _type {
	case WallpaperTypeHorizontal:
		direct = "2626"
	case WallpaperTypeVertical1080p:
		direct = "2627"
	case WallpaperTypeVertical2k:
		direct = "2650"
	default:
		return nil, errors.New("invalid type")
	}

	once := func() error {
		resp, result, err := httpx.PostAny[WallpaperResponse](
			"https://na-community.playerinfinite.com/api/gpts.information_feeds_svr.InformationFeedsSvr/GetContentByLabel",
			httpx.WithHeader(http.Header{
				"Referer":    []string{"https://nikke-en.com/"},
				"X-Areaid":   []string{"na"},
				"X-Gameid":   []string{"16"},
				"X-Language": []string{"en"},
				"X-Source":   []string{"pc_web"},
			}),
			httpx.WithPayload(map[string]any{
				"language":              []string{"en"},
				"gameid":                "16",
				"offset":                offset,
				"get_num":               size,
				"ext_info_type_list":    []int{0, 1, 2},
				"secondary_label_id":    1864,
				"content_class":         1,
				"content_label_id_list": []string{direct},
				"primary_label_id":      4114,
			}),
		)
		if err != nil {
			return err
		}
		if resp.Code != http.StatusOK {
			return fmt.Errorf("status code error: %d, %s", resp.Code, http.StatusText(resp.Code))
		}

		if result.Code != 0 {
			return fmt.Errorf("response code %d: %s", result.Code, result.Msg)
		}
		for _, pic := range result.Data.InfoContent {
			info := &WallpaperInfo{ContentID: pic.ContentID}
			if u := strings.TrimSpace(pic.JumpLinkInfo.JumpURL); u != "" {
				info.Urls = []string{u}
				continue
			}
			info.Urls = pic.PicUrls
			infos = append(infos, info)
		}
		offset += size
		if total == 0 {
			total = result.Data.TotalNum
		}

		return nil
	}
	for total == 0 || offset < total {
		if err = once(); err != nil {
			return infos, err
		}
	}
	return infos, err
}

func GetLiveWallpaperInfo() (infos []*WallpaperInfo, err error) {
	var (
		offset, size = 0, 12
		total        int
	)
	once := func() error {
		resp, result, err := httpx.PostAny[WallpaperResponse](
			"https://na-community.playerinfinite.com/api/gpts.information_feeds_svr.InformationFeedsSvr/GetContentByLabel",
			httpx.WithHeader(http.Header{
				"Referer":    []string{"https://nikke-en.com/"},
				"X-Areaid":   []string{"na"},
				"X-Gameid":   []string{"16"},
				"X-Language": []string{"en"},
				"X-Source":   []string{"pc_web"},
			}),
			httpx.WithPayload(map[string]any{
				"language":              []string{"en"},
				"gameid":                "16",
				"offset":                offset,
				"get_num":               size,
				"ext_info_type_list":    []int{0, 1, 2},
				"secondary_label_id":    1866,
				"content_class":         1,
				"content_label_id_list": []string{},
				"primary_label_id":      4114,
			}),
		)
		if err != nil {
			return err
		}
		if resp.Code != http.StatusOK {
			return fmt.Errorf("status code error: %d, %s", resp.Code, http.StatusText(resp.Code))
		}

		if result.Code != 0 {
			return fmt.Errorf("response code %d: %s", result.Code, result.Msg)
		}
		for _, live := range result.Data.InfoContent {
			if u := strings.TrimSpace(live.JumpLinkInfo.JumpURL); u != "" {
				infos = append(infos, &WallpaperInfo{
					ContentID: live.ContentID,
					Urls:      []string{u},
				})
			}
		}
		offset += size
		if total == 0 {
			total = result.Data.TotalNum
		}
		return nil
	}

	for total == 0 || offset < total {
		if err = once(); err != nil {
			return infos, err
		}
	}
	return infos, err
}

func GetAvatarInfo() (infos []*WallpaperInfo, err error) {
	var (
		offset, size = 0, 12
		total        int
	)

	once := func() error {
		resp, result, err := httpx.PostAny[WallpaperResponse](
			"https://na-community.playerinfinite.com/api/gpts.information_feeds_svr.InformationFeedsSvr/GetContentByLabel",
			httpx.WithHeader(http.Header{
				"Referer":    []string{"https://nikke-en.com/"},
				"X-Areaid":   []string{"na"},
				"X-Gameid":   []string{"16"},
				"X-Language": []string{"en"},
				"X-Source":   []string{"pc_web"},
			}),
			httpx.WithPayload(map[string]any{
				"language":              []string{"en"},
				"gameid":                "16",
				"offset":                offset,
				"get_num":               size,
				"ext_info_type_list":    []int{0, 1, 2},
				"secondary_label_id":    1865,
				"content_class":         1,
				"content_label_id_list": []string{},
				"primary_label_id":      4114,
			}),
		)
		if err != nil {
			return err
		}
		if resp.Code != http.StatusOK {
			return fmt.Errorf("status code error: %d, %s", resp.Code, http.StatusText(resp.Code))
		}

		if result.Code != 0 {
			return fmt.Errorf("response code %d: %s", result.Code, result.Msg)
		}
		for _, pic := range result.Data.InfoContent {
			infos = append(infos, &WallpaperInfo{
				ContentID: pic.ContentID,
				Urls:      pic.PicUrls,
			})
		}
		offset += size
		if total == 0 {
			total = result.Data.TotalNum
		}

		return nil
	}
	for total == 0 || offset < total {
		if err = once(); err != nil {
			return infos, err
		}
	}
	return infos, err
}

type DownloadFailed struct {
	ID  string
	URL string
	Err error
}

func DownloadStaticWallpapers(_type WallpaperType, dir string, gcount int) (result []DownloadFailed) {
	if gcount == 0 {
		gcount = 5
	}
	var (
		ch    = make(chan struct{}, gcount)
		wg    sync.WaitGroup
		retry = 5
	)
	os.MkdirAll(dir, os.ModePerm)

	infos, err := GetStaticWallpaper(_type)
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		if len(info.Urls) == 0 {
			result = append(result, DownloadFailed{ID: info.ContentID, Err: errors.New("no url")})
			continue
		}

		wg.Add(1)

		ch <- struct{}{}

		go func(info *WallpaperInfo) {
			defer func() {
				<-ch
				wg.Done()
			}()
			var (
				path = filepath.Join(dir, fmt.Sprintf("nikke-%s.jpeg", info.ContentID))
				err  error
			)
			if _, err := os.Stat(path); err == nil {
				return
			}

			for range retry {
				if err = httpx.Download(info.Urls[0], path); err == nil {
					break
				}
			}
			if err != nil {
				result = append(result, DownloadFailed{ID: info.ContentID, URL: info.Urls[0], Err: err})
			}
		}(info)
	}
	wg.Wait()
	return result
}

func DownloadLiveWallpapers(dir string, gcount int, uncompress bool) (result []DownloadFailed) {
	if gcount == 0 {
		gcount = 5
	}
	var (
		ch    = make(chan struct{}, gcount)
		wg    sync.WaitGroup
		retry = 5
	)
	os.MkdirAll(dir, os.ModePerm)

	infos, err := GetLiveWallpaperInfo()
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		if len(info.Urls) == 0 {
			result = append(result, DownloadFailed{ID: info.ContentID, Err: errors.New("no url")})
			continue
		}
		wg.Add(1)

		ch <- struct{}{}

		go func(info *WallpaperInfo) {
			defer func() {
				<-ch
				wg.Done()
			}()
			var (
				path = filepath.Join(dir, fmt.Sprintf("nikke-live-%s.zip", info.ContentID))
				err  error
			)
			if _, err := os.Stat(path); err == nil {
				return
			}

			for range retry {
				if err = httpx.Download(info.Urls[0], path); err == nil {
					break
				}
			}
			if err != nil {
				result = append(result, DownloadFailed{ID: info.ContentID, URL: info.Urls[0], Err: err})
				return
			}
			if uncompress {
				packer.Unzip(path, strings.TrimSuffix(path, ".zip"))
			}
		}(info)
	}
	wg.Wait()
	return result
}

func DownloadAvatars(dir string, gcount int) (result []DownloadFailed) {
	if gcount == 0 {
		gcount = 5
	}
	var (
		ch    = make(chan struct{}, gcount)
		wg    sync.WaitGroup
		retry = 5
	)
	os.MkdirAll(dir, os.ModePerm)

	infos, err := GetAvatarInfo()
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		if len(info.Urls) == 0 {
			result = append(result, DownloadFailed{ID: info.ContentID, Err: errors.New("no url")})
			continue
		}

		wg.Add(1)

		ch <- struct{}{}

		go func(info *WallpaperInfo) {
			defer func() {
				<-ch
				wg.Done()
			}()
			var (
				path = filepath.Join(dir, fmt.Sprintf("nikke-avatar-%s.jpeg", info.ContentID))
				err  error
			)
			if _, err := os.Stat(path); err == nil {
				return
			}

			for range retry {
				if err = httpx.Download(info.Urls[0], path); err == nil {
					break
				}
			}
			if err != nil {
				result = append(result, DownloadFailed{ID: info.ContentID, URL: info.Urls[0], Err: err})
			}
		}(info)
	}
	wg.Wait()
	return result
}
