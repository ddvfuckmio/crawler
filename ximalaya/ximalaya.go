package main

import (
	"github.com/gocolly/colly"
	"log"
)

const listPreUrl = "https://www.ximalaya.com/revision/category/queryCategoryPageAlbums?"
const cookie = `Hm_lvt_4a7d8ec50cfd6af753c4f8aee3425070=1539345894,1539347507,1539414843,1539441506; login_from=qq; nickname=%E6%93%8D%E7%A2%8E%E4%BA%86%E5%BF%83%E7%9A%84mio%E9%85%B1; 1&remember_me=y; 1&_token=133099983&E1388119AB444NdV439D0D4A93262DBD6D6DA588AD9EDABFADA8F40A1F59F5184504F80F11DA1A9B; 1_l_flag="133099983&E1388119AB444NdV439D0D4A93262DBD6D6DA588AD9EDABFADA8F40A1F59F5184504F80F11DA1A9B_2018-10-13 23:05:45"; Hm_lpvt_4a7d8ec50cfd6af753c4f8aee3425070=1539443148`

var (
	collectorCommonFunc = []func(*colly.Collector){
		colly.AllowedDomains("www.ximalaya.com")}
)

func main() {
	_, categories := getCategories("https://www.ximalaya.com/category/")
	for index, value := range categories {
		log.Printf("%d %s", index, value)
	}

}

//总目录
func getCategories(link string) ([]string, []string) {

	var categoryCollector = colly.NewCollector(collectorCommonFunc...)
	var categoryUrls = make([]string, 0, 400)
	var categorys = make([]string, 0, 400)
	categoryCollector.OnHTML(".category_hotword.Kx", func(element *colly.HTMLElement) {
		categoryCollector.MaxDepth = 1

		element.ForEach(".category_hotword .list .item.Kx", func(i int, element *colly.HTMLElement) {
			url := element.Attr("href")
			category := element.Text
			categoryUrls = append(categoryUrls, url)
			categorys = append(categorys, category)
		})

	})

	categoryCollector.OnHTML(".category_plate .body.Kx", func(element *colly.HTMLElement) {
		element.ForEach(".category_plate .subject_wrapper .list .item.Kx", func(i int, element *colly.HTMLElement) {
			url := element.Attr("href")
			category := element.Text
			categoryUrls = append(categoryUrls, url)
			categorys = append(categorys, category)
		})
	})

	categoryCollector.Visit(link)
	return categoryUrls, categorys
}

//func getDescription(courseId string) string {
//	s, err := httpGetDescription(courseId)
//	if err != nil {
//		log.Println(err)
//	}
//	reg, err := regexp.Compile("\\<[\\S\\s]+?\\>")
//	if err != nil {
//		log.Println(err)
//	}
//	s = reg.ReplaceAllString(s, "")
//	return s
//}

//func httpGetDescription(courseId string) (string, error) {
//	url := "https://www.ximalaya.com/revision/album?albumId=" + courseId
//	client := &http.Client{}
//	var description = ""
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return description, err
//	}
//
//	req.Header.Set("Accept", "*/*")
//	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
//	req.Header.Set("Connection", "keep-alive")
//	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
//	req.Header.Set("Cookie", cookie)
//	req.Header.Set("Host", "www.ximalaya.com")
//	//req.Header.Set("Referer", "https://www.ximalaya.com/youshengshu/wenxue/")
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")
//	req.Close = true
//	response, err := client.Do(req)
//	if err != nil {
//		log.Println("-->获取内容介绍失败:", err)
//		return description, nil
//	}
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		return description, err
//	}
//	result := &db.ToolStruct{}
//	err = json.Unmarshal(body, result)
//	if err != nil {
//		log.Println("-->获取内容介绍失败:", err)
//		return description, err
//	}
//	description = result.Data.MainInfo.RichIntro
//	return description, nil
//}