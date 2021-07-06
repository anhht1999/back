package crawl

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	"github.com/anhht1999/back/lec-04/model"
)


func Crawl(link string) []model.Movie {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) { 
		log.Println("Something went wrong:", err)
	})

	var listMovies []model.Movie

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) { //Bóc tách dữ liệu từ HTML lấy được
		data := model.Movie{}
		
		data.Title = e.ChildText(".titleColumn") //Tìm đến thẻ con của thẻ tr có class titleColumn và lấy nội dung
		title := regexp.MustCompile(`\s+`).ReplaceAllString(data.Title, " ")
		listTitleInfo := strings.Split(title,"(")
		data.Title = listTitleInfo[0]
		data.Year,_ = strconv.Atoi(listTitleInfo[1][0:len(listTitleInfo[1])-1])
		data.Rating = e.ChildText(".imdbRating") //tìm đến thẻ con có class imdbRating và lấy nội dung
		data.Url = e.ChildAttr("a", "href")      //Tìm đến thẻ con a và lấy nội dung thuộc tính href
		data.Url = "https://www.imdb.com/"+data.Url

		listMovies = append(listMovies, data)
	})

	c.OnScraped(func(r *colly.Response) { //Hoàn thành job craw
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(link) //Trình thu thập truy cập URL đó
	return listMovies
}
