package main

import (
	"net/http"

	"crawler/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strings"
	"log"
	"io/ioutil"
	"encoding/json"
	"crawler/utils"
	"strconv"
)

const listPreUrl = "https://www.ximalaya.com/revision/category/queryCategoryPageAlbums?"
const cookie = `Hm_lvt_4a7d8ec50cfd6af753c4f8aee3425070=1539345894,1539347507,1539414843,1539441506; login_from=qq; nickname=%E6%93%8D%E7%A2%8E%E4%BA%86%E5%BF%83%E7%9A%84mio%E9%85%B1; 1&remember_me=y; 1&_token=133099983&E1388119AB444NdV439D0D4A93262DBD6D6DA588AD9EDABFADA8F40A1F59F5184504F80F11DA1A9B; 1_l_flag="133099983&E1388119AB444NdV439D0D4A93262DBD6D6DA588AD9EDABFADA8F40A1F59F5184504F80F11DA1A9B_2018-10-13 23:05:45"; Hm_lpvt_4a7d8ec50cfd6af753c4f8aee3425070=1539443148`
const httpPre = "https://www.ximalaya.com"
const regexp = "\\<[\\S\\s]+?\\>"
const totalCategories = "https://www.ximalaya.com/category/"

var (
	collectorCommonFunc = []func(*colly.Collector){
		colly.AllowedDomains("www.ximalaya.com")}
)

func main() {

	work()
}

func work() {
	//var wg sync.WaitGroup
	var count int = 0
	categories := getCategories(totalCategories)
	jobs := make(chan models.Category, 5)
	datas := make(chan []models.Course, 305)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				job := <-jobs
				log.Println("-->正在收集", job.FirstCategory, job.SecondCategory)
				courses := httpGetCourseList(job.Url)
				for i := 0; i < len(courses); i++ {
					courses[i].FirstCategory = job.FirstCategory
					courses[i].SecondCategory = job.SecondCategory
				}
				log.Println("-->收集完毕", job.FirstCategory, job.SecondCategory)
				datas <- courses
			}
		}()
	}
	for i := 0; i < len(categories); i++ {
		jobs <- categories[i]
	}
	for {
		if (count == 305) {
			return
		}
		courses := <-datas
		count++
		log.Printf("正在插入第%d", count)
		models.SaveCourses(courses)

	}

}

//获取所有目录
func getCategories(link string) []models.Category {
	var categories = make([]models.Category, 0)
	var categoryCollector = colly.NewCollector(collectorCommonFunc...)

	categoryCollector.OnHTML(".category_hotword.Kx", func(element1 *colly.HTMLElement) {

		element1.ForEach(".category_hotword-wrapper.Kx", func(i int, element2 *colly.HTMLElement) {

			selection := element2.DOM.Find(".category_hotword .hotword.Kx")
			firstCategory := selection.Find(".category_hotword .hotword .center.Kx").Text()

			element2.ForEach(".category_hotword .list .item.Kx", func(i int, element3 *colly.HTMLElement) {
				category := models.Category{
					FirstCategory:  firstCategory,
					SecondCategory: element3.Text,
					Url:            element3.Attr("href"),
				}
				categories = append(categories, category)
			})

		})

	})

	categoryCollector.OnHTML(".category_plate .body.Kx", func(element *colly.HTMLElement) {
		selections := element.DOM.Children()
		selections.Each(func(i int, selection1 *goquery.Selection) {
			firstCategory := selection1.Find(".category_plate .subject_wrapper .subject h2.Kx").Text()
			selection1.Find(".category_plate .subject_wrapper .list .item.Kx").Each(func(i int, selection2 *goquery.Selection) {
				secondCategory := selection2.Text()
				url, _ := selection2.Attr("href")
				category := models.Category{
					FirstCategory:  firstCategory,
					SecondCategory: secondCategory,
					Url:            url,
				}
				categories = append(categories, category)
			})
		})

	})

	categoryCollector.Visit(link)
	return categories
}

//func getCourseList(courseId string) string {
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

func httpGetCourseList(href string) ([]models.Course) {
	courses := make([]models.Course, 0)
	firstUrl, secondUrl := splitHref(href)
	url := "https://www.ximalaya.com/revision/category/queryCategoryPageAlbums?"
	url = url + "category=" + firstUrl + "&subcategory=" + secondUrl + "&meta=&sort=0&page=1&perPage=1000"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := httpGet(request)
	if err != nil {
		log.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	result := &utils.CourseList{}
	err = json.Unmarshal(body, result)

	for _, album := range result.Data.Albums {
		course := models.Course{
			Title:     album.Title,
			PlayCount: album.PlayCount,
			Author:    album.AnchorName,
			CourseId:  strconv.Itoa(album.AlbumID),
		}
		courses = append(courses, course)
	}

	return courses

}

func splitHref(href string) (firstCategory string, secondCategory string) {
	words := strings.Split(href, "/")
	return words[1], words[2]
}

func httpGet(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	return response, err
}
