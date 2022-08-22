package main

import "fmt"

func main() {

	age := 10

	switch age { // age 조건문 실행 👇

	case 10:  // age가 10 일 경우
		fmt.Println("초등생1") // 출력
		fmt.Println("초등생2") // 출력
	case 11, 13:	// age가 11 이나 13 일 경우
		fmt.Println("초등생3")
		
	default:  // 아무것도 해당안될경우
		fmt.Println("학생아님") // 출력
	}
//--------------------------------------------

	ages := 2

	switch { // switch에 ages 를 넣지않고 case마다 조건을 넣어줄수도 있다

	case ages > 1:  // age가 1보다 클 경우
		fmt.Println("초등생1") // 출력

	case ages < 0:	// age가 0보다 작을 경우
		fmt.Println("초등생3") // 출력
		
	default:  // 아무것도 해당안될경우
		fmt.Println("학생아님") // 출력
	}
}
