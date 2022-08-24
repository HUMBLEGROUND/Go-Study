package main

import "fmt"

func test() {
	fmt.Println("hello")
}

func test2(x int) { // 매개변수를 지정할수도 있다
	fmt.Println(x)
}

func test3(x int, y int) { // 매개변수를 지정할수도 있다
	fmt.Println(x + y) // 계산
}

func test4(x , y int) int  { // 매개변수 뒤에 어떤 값(타입)을 반환해줄지 써준다
	// 하나만 타입을 지정해주면 하나는 생략이 가능하다
	return x + y // Println를 안쓰고 리턴으로도 가능하다
}

func test5(x , y int) (int, int)  { // 반환값을 여러개 지정할수도 있다
	return x + y, x - y // 2개의 반환값
}

func test6(x , y int) (w1, w2 int)  { // 반환값에 변수명을 지정할수도 있다
	w1 = x + y
	w2 = x - y
	return
}

func test7(x , y int) (w1, w2 int)  { // 반환값에 변수명을 지정할수도 있다
	defer fmt.Println("반환 성공! 출력값 직전에 실행")
	// ⭐ defer 는 반환순서를 조정할수 있다
	w1 = x + y
	w2 = x - y
	fmt.Println("반환 작업 중")
	return
}

//--------------------------------------------

func main() {
// 위에서 만들어놓은 함수를 컴포넌트처럼 부를수 있다
	test() // 👉 출력값 hello
	test2(5) // 👉 출력값 5
	test2(7) // 👉 출력값 7
	test3(10, 3) // 👉 출력값 13  / 10 + 3
	
//--------------------------------------------

	z := test4(10, 5) // 변수에 함수를 담고
	fmt.Println(z) // Println로 출력 👉 출력값 15  / 10 + 5

//--------------------------------------------

	z1, z2 := test5(10, 7) // 반환값 2개
	fmt.Println(z1, z2) // 👉 출력값 17 , 3  / 10 + 7, 10 - 7

	// result, result2 := test5(10, 6) // 👉 변수를 2개 생성함
	result, _ := test5(10, 6) // 👉 변수를 2개 생성했는데 하나는 안쓸경우 _ (언더바) 표시
	fmt.Println(result) // 👉 변수를 다 호출하지 않으면 에러가 뜬다

//--------------------------------------------

	z3, z4 := test6(10, 7) // 반환값 2개
	fmt.Println(z3, z4) // 👉 출력값 17 , 3  / 10 + 7, 10 - 7

//--------------------------------------------

	z5, z6 := test6(10, 7) // 반환값 2개
	fmt.Println(z5, z6) 
	// 👉 출력값 
	// 반환 작업 중 
	// 반환 성공! 출력값 직전에 실행
	// 17 , 3
}
