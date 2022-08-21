package main

import "fmt"

func main() {

	var x, y, z int // 변수 이름을 정하고 / int 자료형 (숫자)
	x = 32
	y = 43
	z = 55
	// var 변수로 지정한 값은 변할 수 있는 값이다

	const Pi = 3.14
	// const 변수로 지정한 값은 변할 수 없다

	fmt.Println(x, y, z)
	fmt.Println(Pi)

	x = 72 // 👉 위에서 지정한 값을 재지정해도 바뀔수 있다
	// Pi = 9.13 // 👉 지정된 값을 바꿀수 없다
}
