package main

import "fmt"

func main() {

	var test1 string = "test string" 
	// 변수 이름을 정하고 / string 문자열 / 변수에 들어갈 내용
	
	//--------------------------------------------
	var test2 string // 변수이름과 타입만 지정하고
	test2 = "test string" // 변수내용은 따로 지정 할 수도 있다
	test2 = "test string2" // 👉 위에서 지정한 값을 재지정해도 바뀔수 있다
  //--------------------------------------------

	fmt.Println(test1)
	fmt.Println(test2)
	// 함수 안에 프린트 명령
}
