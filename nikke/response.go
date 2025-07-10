package nikke

type WallpaperResponse struct {
	Code int `json:"code,omitempty"`
	Data struct {
		ContentClass int    `json:"content_class,omitempty"`
		Errmsg       string `json:"errmsg,omitempty"`
		InfoContent  []struct {
			ContentClass         int    `json:"content_class,omitempty"`
			ContentDesc          string `json:"content_desc,omitempty"`
			ContentID            string `json:"content_id,omitempty"`
			ContentLabelInfoList []struct {
				LabelID   string `json:"label_id,omitempty"`
				LabelName string `json:"label_name,omitempty"`
				Order     int    `json:"order,omitempty"`
			} `json:"content_label_info_list,omitempty"`
			ContentType  int    `json:"content_type,omitempty"`
			ExtInfo      string `json:"ext_info,omitempty"`
			JumpLinkInfo struct {
				JumpType int    `json:"jump_type,omitempty"`
				JumpURL  string `json:"jump_url,omitempty"`
			} `json:"jump_link_info,omitempty"`
			JumpScheme        string   `json:"jump_scheme,omitempty"`
			PicUrls           []string `json:"pic_urls,omitempty"`
			PrimaryLabelIds   []int    `json:"primary_label_ids,omitempty"`
			PubTimestamp      string   `json:"pub_timestamp,omitempty"`
			SecondaryLabelIds []int    `json:"secondary_label_ids,omitempty"`
			SpiderScore       string   `json:"spider_score,omitempty"`
			Title             string   `json:"title,omitempty"`
		} `json:"info_content,omitempty"`
		NextOffset int `json:"next_offset,omitempty"`
		Result     int `json:"result,omitempty"`
		TotalNum   int `json:"total_num,omitempty"`
	} `json:"data,omitempty"`
	Msg       string `json:"msg,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}
