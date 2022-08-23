package main

import "fmt"

// Go 언어에서는 클래스 대신 struct 라는 구조체를 사용한다
// 👉 type 구조체명 struct { }

// struct 필드를 엑세스하기 위해서는 쩜 . (닷 오퍼레이터) = (닷 연산자)를 사용한다

type Hunting struct {
	name string
	age int
	// good bool 👉 정의된 타입에 항목을 추가해줄수도 있다
}

func main() {
	dog1 := Hunting{} // 👉 위에서 지정한 타입을 변수에 불러오기
	dog2 := Hunting{}
	dog1.name = "Bab" 
	// 변수 뒤에 쩜 을 찍으면(닷 연산자) 타입안에 내용을 선택할수 있다 / 지정
	dog1.age = 12
	dog2.name = "Alex"

	fmt.Println(dog1, dog2) // 👉 {Bab 12} {Alex 0} dog2.age 를 지정안해주면 값은 0 이다
	
//--------------------------------------------

	var dog3 = Hunting{"Elen",22} // 이렇게 지정해줄수도 있다 
	dog4 := Hunting{age:21,name:"Big"} // 이렇게도 가능 name, age 적는 순서는 상관없다

	fmt.Println(dog3, dog4) // 👉 출력값 {Elen 22} {Big 21}
	fmt.Println(dog3.age) // 👉 출력값 22

	dog3.name = "Gail" // 지정된 name을 바꿔줄수도 있다

	fmt.Println(dog3) // 👉 출력값 {Gail 22}

	// dog4.good = false 
}
