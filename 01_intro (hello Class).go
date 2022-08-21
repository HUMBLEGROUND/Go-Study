package main

// 패키지 명 선언
// 모든 Go 프로그램은 패키지로 구성되어있다
// F5키를 눌러서 build 를 하면 main 을 제일먼저 찾는다
// main 을 우선순위로 실행시킴
// 👉 ex) 리액트로 따지면 App.js

import "fmt"

// 패키지 불러오기
// 모든 Go 프로그램은 패키지로 구성되어있다
// 다른 파일의 함수를 부른다

func main() {
	fmt.Println("Hello My Class !")
}
// 함수 안에 프린트 명령
// 모든 Go 프로그램은 main함수를 포함한다