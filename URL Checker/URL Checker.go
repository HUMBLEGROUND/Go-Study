package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("req error") // 에러핸들러 메세지

func main() {
	// var result = map[string]string{}
	var results = make(map[string]string) // 비어있는 map 만들기
	urls := []string{
		"https://naver.com/",
		"https://google.com/",
		"https://daum.net/",
		"https://amazon.com/",
		"https://zum.com/",
		"https://facebook.com/",
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

		result := "OK" // url 체크 결과 멘트출력

		err := hitURL(url) // hitURL의 에러가 있다면 받음

		if err != nil { // 에러 발생시
			result = "FAILED"
		}
		results[url] = result // map url 리스트의 결과 출력
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)

	resp, err := http.Get(url) // http 주소 get 요청 (index 순)
	
	if err != nil || resp.StatusCode >= 400 { // 400 이상부터는 에러 ex)404
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}