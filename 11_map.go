package main

import "fmt"

func main() {

	var zoo map[string]int = map[string]int { // 맵의 value : key 값 타입설정
		"코끼리": 22,
		"사자": 13,
		"호랑이": 7,
	}

	// zoo := make(map[string]int)
	
	fmt.Println(zoo) // 👉 출력값 map[호랑이:7 사자:13 코끼리:22]

//--------------------------------------------

	zoo["코알라"] = 87 // 맵 추가
	fmt.Println(zoo) // 👉 출력값 map[호랑이:7 사자:13 코끼리:22 코알라:87]

//--------------------------------------------

	fmt.Println(zoo["사자"]) // 👉 출력값 13

	delete(zoo, "사자") // 배열안에 값을 지울수도 있다
	fmt.Println(zoo) // 👉 출력값 map[호랑이:7 코끼리:22 코알라:87]

//--------------------------------------------

	val, ok := zoo["호랑이"] // 배열안에 값이 있는지 확인할수도 있다 (개수, bool)
	// val, ok := zoo["고릴라"] // 배열안에 값이 있는지 확인할수도 있다
	// fmt.Println(val, ok)	// 👉 출력값 0 false 고릴라는 없음
	fmt.Println(val, ok)	// 👉 출력값 7 true

}
