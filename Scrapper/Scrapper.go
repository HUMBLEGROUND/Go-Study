package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	var jobs []extractedJob // 빈배열 struct
	c := make(chan []extractedJob)

	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c // 채널에서 전송받은 메세지를 변수에 저장

		// extractedJobs := getPage(i) 
		// getPage함수에서 나오는 배열들을 반복문으로 돌려서 extractedJobs변수에 담는다
		jobs = append(jobs, extractedJobs...) // 위에서 정의한 빈배열에 추가시킨다
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv") // 데이터를 csv 파일로 저장하는 라이브러리
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush() // 작성된 모든것들을 csv파일에 입력

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers) // 배열에 쓴다
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice) // 쓰기
		checkErr(jwErr)
	}
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob // 빈배열 struct

	c := make(chan extractedJob) // 채널 생성 (struct 를 보냄)

	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	// baseURL 뒤에 &start=0 ~ 50 / 100 / 150 ... 50 단위로 넣는다는 뜻
	// 숫자(int)를 string 으로 바꿔주는 라이브러리 strconv.Itoa
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c // 채널에서 전송받은 메세지를 jop 변수에 저장
		jobs = append(jobs, job) // 위에서 정의한 빈배열에 추가시킨다
	}
	// searchCards.Each(func(i int, card *goquery.Selection) {
	// 	job := extractJob(card) // extractJob 함수를 변수에 담고 
	// 	jobs = append(jobs, job) // 위에서 정의한 빈배열에 추가시킨다
	// })
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) { 
	id, _ := card.Attr("data-jk")
	// id가 data-jk 들 다 찾기
	title := cleanString(card.Find(".title>a").Text()) 
	// title 클래스 안에 a 태그 에 들어있는 텍스트들 찾기
	location := cleanString(card.Find(".sjcl").Text())
	// sjcl 클래스 안에 텍스트들 찾기
	salary := cleanString(card.Find(".salaryText").Text())
	// salaryText 클래스 안에 텍스트들 찾기
	summary := cleanString(card.Find(".summary").Text())
	// summary 클래스 안에 텍스트들 찾기
	c <- extractedJob{ // 채널을 통해 전달
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
	// Fields의 기능은 
	// str  is  good 👉 간격이 긴 문자들을 "str","is","good" 으로 만들어준다
	// Join 은 "str","is","good" 을 풀어서 str is good 으로 만들어준다
}

func getPages() int {
	pages := 0 // 기본값 0 에 a 태그들 대입하기
	res, err := http.Get(baseURL) // http 주소 get 요청 (index 순)
	checkErr(err)
	checkCode(res) // res에 ttp.Response 주입 (아래 checkCode 함수에)

	defer res.Body.Close() // ⭐ defer 는 반환순서를 조정할수 있다
	// res.Body 를 닫아줌 👉 메모리가 새어나가는걸 막아준다

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
	// 클래스이름 pagination 를 찾는다
		pages = s.Find("a").Length() // a 태그 대입하기
		// 클래스 pagination 안에 a 태그들 다 찾기
	})
	// goquery에 내장된 기능 Find / Each / Selection
	
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) { // http.Response 역참조
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}