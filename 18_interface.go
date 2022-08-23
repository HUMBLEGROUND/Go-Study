package main

import "fmt"

// 인터페이스 - 메소드(기능)들의 집합
// type 인터페이스명 interface { 메소드 반환형 }

type Pizza struct{}

func (p Pizza) eat() {
	fmt.Println("피자 한조각 먹기")
}
// func (리시버 인자) 메소드이름 (리턴타입) { }
func (p Pizza) taste() {
	fmt.Println("피자는 역시 맛있어")
}

//--------------------------------------------

type Coke struct {}

func (c Coke) eat() {
	fmt.Println("느끼할땐 콜라")
}

func (c Coke) taste() {
	fmt.Println("콜라는 달콤해")
}

//--------------------------------------------

type fastFood interface { // 👉 interface는 메소드의 집합 / fastFood로 묶어줌
	eat() // 메소드
	taste() // 메소드
} 
// 👇 interface 명 을 가져감

func westFood (f fastFood) {
	f.eat()
	f.taste()
}

//--------------------------------------------

func main() {
	var NewYork Pizza // Pizza struct 를 변수에
	var Coca Coke // Coke struct 를 변수에

	westFood(NewYork) // interface를 담은 함수에 struct 대입
	westFood(Coca)

	// interface로 메소드들을 묶어줘서 한번에 출력이 가능하다

	// 👉 출력값
	// 피자 한조각 먹기
	// 피자는 역시 맛있어
	// 느끼할땐 콜라
	// 콜라는 달콤해

}
