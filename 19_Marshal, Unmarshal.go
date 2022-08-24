package main

import (
	"encoding/json" // marshal 을 사용하기위해 import
	"fmt"
	"log"
)

// 데이터를 json화 시키면 일렬로 바뀌는데 이것을 바이트 코드라고 한다
// 👉 직렬화 = 마샬(marshal) 이라고 한다

// json을 데이터(객체)로 변환시키면 이것을 👉 역직렬화 = 언마샬(unmarshal) 이라고 한다


type Task struct {
	Title string
	Status int
	}

func main() {
ExampleTask_marshalJSON()
ExampleTask_unmarshalJSON()
}

func ExampleTask_marshalJSON() {
	t := Task{
		"Laundry", 
		1,
	} // 👉 marsha 시킨다 (한줄로 만들어줌)
	b, err := json.Marshal(t) // t 변수를 Marshal함수에 넣고 돌린다
	// 에러가 날 경우도 변수로

	if err != nil { // 에러 발생시
		log.Println(err)
		return
	}
	
	fmt.Println(string(b))
	// 출력값 👉 {"Title": "Laundry", "Status": 1}
}

func ExampleTask_unmarshalJSON() {
	b := []byte(`{"Title":"Buy Milk","Status":2}`)
	t := Task{} // b 를 unmarsha시켜서 t 에 넣어주기 위한 공간
	err := json.Unmarshal(b, &t) // 👉 unmarsha 시킨다 (json을 풀어준다(해체))
	// b 를 unmarsha 시켜서 t 주소로 넣어준다
	
	if err != nil {
		log.Println(err)
		return
	}
	
	fmt.Println(t.Title) // 👉 Buy Milk
	fmt.Println(t.Status) // 👉 2
}

// 구조체(데이터) 👉 json 👉 json에서 필요한 값을 구조체(데이터)화 로 다시 변환

// 모바일앱 데이터 👉 데이터베이스(json) 👉 웹브라우저에서 읽음

// 스마트컨트랙트 데이터 👉 월드스테이트 데이터베이스 👉 데이터베이스자체 / 익스플로어 모니터링 / 스마트컨트랙트



