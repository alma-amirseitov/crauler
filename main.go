package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)


type Article struct{
	title string `json:"hading"`
	time string `json:"time"`
	href string `json:"href"`
	content string `json:"content"`
}

var articleCollection = []Article{}

func main()  {
	res,err := http.Get("https://www.zakon.kz/news/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("#dle-content .cat_news_item").Each(func(index int, element *goquery.Selection) {
		title := element.Find("a").Text()
		href,_ := element.Find("a").Attr("href")
		time := element.Find(".date").Text()
		content := scrap("https://www.zakon.kz"+href)
		articleCollection = append(articleCollection,Article{title,time,href,content})
	})
	writeResultXLS()
}
func scrap(url string) string{
	res,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	var content string
	document.Find("#initial_news_story").Each(func(index int, element *goquery.Selection) {
		p := element.Find("p").Text()
		content += p
	})
	return content
}
func writeResultXLS(){
	xlsx := excelize.NewFile()
	// Create a new sheet.
	xlsx.NewSheet("Sheet1")
	j := 1
	for _,article := range articleCollection {
		xlsx.SetCellValue("Sheet1",fmt.Sprintf("A%v",j),article.time)
		xlsx.SetCellValue("Sheet1",fmt.Sprintf("B%v",j),article.title)
		xlsx.SetCellValue("Sheet1",fmt.Sprintf("C%v",j+1),article.content)
		j+=2
	}
	err := xlsx.SaveAs("./news.xlsx")
	if err != nil{
		fmt.Println(err)
	}
}
