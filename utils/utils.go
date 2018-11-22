package utils

type CourseList struct {
	Data struct {
		Albums []struct {
			AlbumID    int    `json:"albumId"`
			AnchorName string `json:"anchorName"`
			CoverPath  string `json:"coverPath"`
			IsFinished int    `json:"isFinished"`
			IsPaid     bool   `json:"isPaid"`
			Link       string `json:"link"`
			PlayCount  int    `json:"playCount"`
			Title      string `json:"title"`
			UID        int    `json:"uid"`
		} `json:"albums"`
		Page       int `json:"page"`
		PageConfig struct {
			H1title string `json:"h1title"`
		} `json:"pageConfig"`
		PageSize int `json:"pageSize"`
		Total    int `json:"total"`
	} `json:"data"`
	Msg string `json:"msg"`
	Ret int    `json:"ret"`
}
