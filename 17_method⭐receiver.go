package main

import "fmt"

// ⭐ 함수(func) - 처리의 묶음
// ⭐ 객체(obj) = 변수의 묶음
// ⭐ 구조체(struct) - 객체를 선언하는 방법
// ⭐ 리시버(receiver) - 객체와 함수를 연결하는 매개체
// ⭐ 메소드(method) - 객체에 연결된 함수

type Eating struct { // struct 선언
	x, y int
}

// func (리시버 인자) 메소드이름 (리턴타입) { }
// - value receiver - 객체를 value로 그대로 가져오는 리시버
// - pointer receiver - 객체를 포인터(역참조) 로 가져오는 리시버

func (e *Eating) eat_count(n int) { // pointer receiver
	// Eating의 이름을 e 라고 지정해줄수도 있다 (약어)
	// eat_count 👉 함수이름
	e.x += n
	e.y += n
}

func (e *Eating) eat_exit(n int) {
	e.x -= n
	e.y -= n
}

func main() {

	e := Eating{3, 2}
	fmt.Println("오늘 먹은 횟수", e) // 👉 오늘 먹은 횟수 {3, 2}

	e.eat_count(3) // 👉 pointer receiver 👉 역참조
	fmt.Println("더 먹은 횟수", e) // 👉 더 먹은 횟수 {6, 5}

	e.eat_exit(2) // 👉 pointer receiver 👉 역참조
	fmt.Println("먹다가 버린 횟수", e) // 👉 먹다가 버린 횟수 {4, 3}

}
