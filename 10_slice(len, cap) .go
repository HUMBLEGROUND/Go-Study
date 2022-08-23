package main

import "fmt"

func main() {

	var x [5]int = [5]int{12, 23, 34, 45, 56}
	var s []int = x[0:4] // 0번째부터 4번째 까지 자르기
	// 4번째는 포함되지 않고 자른 후 3번째 까지만 보여준다

	fmt.Println(s) // 👉 출력값 [12, 23, 34, 45]
	fmt.Println(len(s)) // 👉 배열의 길이 출력값 4
	fmt.Println(cap(s)) // 👉 출력값 5 / cap 은 자르기 시작부분부터 배열의 끝까지 출력한다

//--------------------------------------------

	var ss []int = []int{12, 23, 34, 45, 56}

	fmt.Println(cap(ss[:3])) // 0 ~ 3번째 까지 슬라이스
	// 👉 출력값 5 / cap 은 자르기 시작부분부터 배열의 끝까지 출력한다
	// 0 1 2 3 4 총 개수를 계산 👉 5개

//--------------------------------------------

	var sss []int = []int{12, 23, 34, 45, 56}
	b := append(sss,77,93) // sss 배열에 77, 93 추가
	// c = append(sss,77,93) // sss 배열에 77, 93 추가

	fmt.Println(b) // 👉 출력값 [12, 23, 34, 45, 56, 77, 93]
	fmt.Println(cap(b)) // 👉 출력값 10
	// 왜 10이 나오는지 👉 b 변수에서 sss값을 한번더 생성하기때문에 5+5 
	// fmt.Println(cap(c)) // 👉 출력값 10

//--------------------------------------------

	d := make([]int,3,4) // len 3 / cap 4 인 slice 생성

	fmt.Println(d) // 👉 출력값 [0 0 0]
	fmt.Println(len(d)) // 👉 3
	fmt.Println(cap(d)) // 👉 4

	d[0] = 12
	d[1] = 23
	d[2] = 33

	fmt.Println(d) // 👉 출력값 [12 23 33]
	fmt.Printf("%T", d) // 타입 알아보기 👉 []int

}
