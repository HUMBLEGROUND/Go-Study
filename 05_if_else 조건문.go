package main

import "fmt"

func main() {

	name := "hello" // string

	fmt.Println("if 이전") // 👇

	if name == "hello" { 	// 👇 조건 name 이 "hello" 일 경우 실행
		fmt.Println("if 안") // "if 안" 실행
	}

	// if name != "hello" { 	// 👇 조건 name 이 "hello" 아닐 경우 실행
	// 	fmt.Println("if 안")
	// }

	// if name != "hello" || name == "hello" { 	// 👇 조건 name 이 "hello" 아니거나 맞을 경우 실행
	// 	fmt.Println("if 안")
	// }

	fmt.Println("if 밖") // 조건문을 빠져나왔으니 "if 밖" 실행

	// 👉 반환값 "if 이전", "if 안", "if 밖"

//--------------------------------------------

	age := 12 // number

	if age >= 18 { // 나이가 18세보다 크거나 같으면
		fmt.Println("군대 갈 수 있어요") // age >= 18 조건에 맞는다면 출력

	} else if age >= 14 { // 아니면 나이가 14세보다 크거나 같으면
		fmt.Println("학도병에 지원하세요")

	}	else {
		fmt.Println("군대 못가요") // age >= 18 에 맞지않는다면 출력
		fmt.Printf("학도병은 %d 년만 기다려 주세요", 14-age) // Printf %d 👉 10진수(정수) 출력
		fmt.Printf("군대는 %d 년만 기다려 주세요", 18-age) // Printf %d 👉 10진수(정수) 출력
	}
	// 👉 반환값 "군대 못가요" / "학도병은 2 년만 기다려 주세요" / "군대는 6 년만 기다려 주세요"
}
