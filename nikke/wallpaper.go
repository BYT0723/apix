package nikke

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/BYT0723/apix/utils"
	"github.com/BYT0723/go-tools/transport/httpx"
)

func GetWallpapers() (urls []string, err error) {
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
				"secondary_label_id":    1864,
				"content_class":         1,
				"content_label_id_list": []string{"2626"},
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
			if len(pic.PicUrls) == 0 {
				continue
			}
			urls = append(urls, pic.PicUrls[0])
		}
		offset += size
		if total == 0 {
			total = result.Data.TotalNum
		}

		return nil
	}
	for total == 0 || offset < total {
		if err = once(); err != nil {
			return
		}
	}
	return
}

type DownloadFailed struct {
	Index int
	Url   string
	Err   error
}

func DownloadStaticWallpapers(dir string, gcount int) (result []DownloadFailed) {
	if gcount == 0 {
		gcount = 5
	}
	var (
		ch = make(chan struct{}, gcount)
		wg sync.WaitGroup
	)
	os.MkdirAll(dir, os.ModePerm)

	urls, err := GetWallpapers()
	if err != nil {
		panic(err)
	}
	for i, url := range urls {
		wg.Add(1)

		ch <- struct{}{}

		go func(i int, url string) {
			defer func() {
				<-ch
				wg.Done()
			}()
			var (
				path = filepath.Join(dir, fmt.Sprintf("nikke-%d.jpeg", i))
				err  error
			)

			for range 5 {
				if err = utils.Download(url, path); err == nil {
					break
				}
			}
			if err != nil {
				result = append(result, DownloadFailed{Index: i, Url: url, Err: err})
			}
		}(i, url)
	}
	wg.Wait()
	return
}
