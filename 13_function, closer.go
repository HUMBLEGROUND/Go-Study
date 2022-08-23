package main

import "fmt"

// Go 언어는 변수에 할당할 수 있다
// 변수에 함수를 담을 수 있다
func hamsu(x int) {
	fmt.Println("함수 출력", x)
}

func hamsu4(myHamsu func(int) int) { // 함수 안에 함수
	fmt.Println(myHamsu(8))
}

// ⭐ 클로저 함수
// 함수 안에서 함수를 선언 & 정의할 수 있다
// 바깥 함수에 선언된 변수에도 접근할 수 있다
func hamsuReturn(b string) func() { 
	// sum := 0 // 👉 함수 안에서 함수를 선언 & 정의할 경우
	return func() { 
		fmt.Println(b) 
		// fmt.Println(sum) // 👉 출력값 0
	}
}

//--------------------------------------------

func main() {
	x := hamsu // 함수를 변수에 담고 
	x(7) // 변수를 출력하면 = 함수를 출력하는 효과랑 같다
	// 👉 출력값 "함수 출력" 7

//--------------------------------------------

	hamsu2 := func (y int) { // ⭐ 익명의 함수도 가능하다
		fmt.Println(y) // 변수 안에 함수를 정의할 수도 있다
	}
	hamsu2(9) // 👉 출력값 9

//--------------------------------------------
	hamsu3 := func(z int) int { // 매개변수 뒤에 어떤 값(타입)을 반환해줄지 써준다
		return z * -1 // 
	}(5) // 중괄호 뒤에 매개변수 값을 설정

	fmt.Println(hamsu3) // 👉 출력값 -5

//--------------------------------------------

	hamsuCall := func(a int) int {
		return a * -1
	}
	hamsu4(hamsuCall) // 👉 출력값 -8
	// 불러온 함수 안에 함수를 계산

//--------------------------------------------

// 클로저 함수 👉 바깥 함수에 선언된 변수에도 접근할 경우

	hamsuReturn("이렇게도 가능")() // 👉 이렇게도 가능
	c := hamsuReturn("변수에 담기도 가능")
	d := hamsuReturn("변수에 담기 가능2")
	c() // 👉 변수에 담기도 가능
	d() // 👉 변수에 담기 가능2
}
