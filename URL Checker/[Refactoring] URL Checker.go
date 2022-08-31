package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string // ok 아니면 failed
}

var errRequestFailed = errors.New("Request failed") // 에러핸들러 메세지

func main() {
	results := make(map[string]string) // 비어있는 map 만들기
	c := make(chan requestResult) // 채널 생성 (struct 를 보냄)
	urls := []string{ // 빈배열
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}
	for _, url := range urls { // url 👉 index / urls 전체를 순회함
		go hitURL(url, c) // 앞에 go를 붙여서 비동기로 동시에 출력
	}

	for i := 0; i < len(urls); i++ { // urls 전체를 반복
		result := <-c
		results[result.url] = result.status // map 유형에 맞게 대입
	}

	for url, status := range results { 
		// 반복문으로 빈 map에 담긴 배열 전체를 순회
		fmt.Println(url, status)
	}
}
// 👉 출력값
// https://www.amazon.com/ OK
// https://www.airbnb.com/ OK    
// https://soundcloud.com/ OK    
// https://www.facebook.com/ OK  
// https://www.instagram.com/ OK 
// https://www.reddit.com/ FAILED
// https://www.google.com/ OK

func hitURL(url string, c chan<- requestResult) {

	resp, err := http.Get(url) // http 주소 get 요청 (index 순)

	status := "OK" // status 기본값은 "OK"

	if err != nil || resp.StatusCode >= 400 { // 400 이상부터는 에러 ex)404
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status} // 타입은 struct를 따른다
}