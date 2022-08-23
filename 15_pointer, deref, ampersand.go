package main

import "fmt"

// 변수 앞에 '&'(Ampersand) 를 붙이면 해당 변수의 메모리 주소를 뜻한다
// 👉 포인터형 변수에 대입을 할 수 있음

// '*' 는 포인터 변수를 역참조(dereference) 해서
// 👉 포인터형 변수에 값을 대입하거나, 가져올 수 있다

func star(starz *string) { // 역참조하는 함수
	*starz = "Look! Star!"
}

// func moon(moond string) { // 그냥 함수 
// 	moond = "Moon!"
// 	return moond
// }

func main() {

	x := 7
	z := &x // x 가 저장된 주소값
	fmt.Println(x, z) // 👉 7 0xc0000100a0

	*z = 9 // x 가 저장된 주소값에 9를 주입
	fmt.Println(x, z) // 👉 9 0xc0000100a0
	// x 값은 변하고 주소는 그대로 나오는걸 볼 수 있다

//--------------------------------------------

	look := "Sun!"
	fmt.Println(look) // 👉 Sun! 출력

	star(&look) // star 함수가 look 의 저장된 주소값을 역참조함
	// 👉 look 이 저장된 주소에 starz 주입
	fmt.Println(look) // 👉 Look! Star! 출력

//--------------------------------------------

	real := "good!"
	var move *string = &real // real 이 저장된 주소값을 역참조

	fmt.Println(move) // 👉 0xc00003e1f0 저장된 주소값 출력
	fmt.Println(*move) // 👉 주소값을 다시 역참조하면? 👉 good! 출력

	fmt.Println(move, &move) 
	// 👉 move의 저장된 주소값(real의 주소)과 / move 의 변수 주소값은 같을까?
	// 👉 0xc00003e1f0  0xc000006028  다르다
	// real 의 주소값 , 저장된 move 변수의 주소값
}
