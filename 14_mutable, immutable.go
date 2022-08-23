package main

import "fmt"

// mutable (변하는) 데이터타입
// immutable (변하지않는) 데이터타입
func main() {

	var x int = 5
	z := x
	z = 7
	fmt.Println(x, z) // 👉 출력값 5 7

	// x := 5
	// z := x
	// z = 9
	// fmt.Println(x, z)
	// 👉 출력값 5 9
//--------------------------------------------

	var a []int = []int{1,2,3} // 슬라이스 생성
	b := a // b 변수에 대입
	a[2] = 33 // 배열의 2번째 값을 변경
	fmt.Println(a, b) // 👉 출력값 [1 2 33]


 // ⭐ 👉 '포인트' 의 개념이다
}
