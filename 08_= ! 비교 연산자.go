package main

import "fmt"

func main() {

	a := 7
	b := 3

	answer := a < b 

	fmt.Printf("%t", answer) // Printf %t 👉 bool 타입 출력
				 // 👉 false 출력
//--------------------------------------------

	c := 7
	d := 3.5

	// answers := c < d  	   // 👉 같은타입이 아니기때문에 에러 int(정수), float64(실수)
	answer2 := float64(c) > d  // 👉 한쪽의 타입을 재지정해줘야 에러가 발생안함

	fmt.Printf("%t", answer2) // Printf %t 👉 bool 타입 출력
				  // 👉 true 출력 
//--------------------------------------------	

	e := "string2"
	f := "string2"
	// f := "string2 "  // 👉 한칸이라도 공백이 생겨도 같지않다
	// f := "String2"   // 👉 한글자라도 대문자라면 같지않다

	answer3 := e == f  // 👉 string 타입끼리도 비교연산을 할 수 있다

	fmt.Printf("%t", answer3) // Printf %t 👉 bool 타입 출력
				  // 👉 true 출력
//--------------------------------------------	

	x := "a" // 👉 26 0x61 아스키코드
	y := "b" // 👉 27 0x62 아스키코드


	answer4 := x < y  // 👉 string 타입끼리도 비교연산을 할 수 있다
			  // 문자끼리 계산은 아스키코드로 해석해서 계산된다

	fmt.Printf("%t", answer4) // Printf %t 👉 bool 타입 출력
				  // 👉 true 출력
//--------------------------------------------	

}
