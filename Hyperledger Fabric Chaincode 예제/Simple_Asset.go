// 패키지 정의
package main
	
// 1. 외부모듈 포함
import (
	"fmt"
	"encoding/json" // marshal 을 사용하기위해 import
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
// 2. 체인코드 클래스-구조체정의 SimpleAsset
type SimpleAsset struct {
}

type Asset struct {
	key string `json:key`
	Value string `json:value`
}

// 3. Init 함수
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success([]byte("init success"))
}
// 4. Invoke 함수
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	fn, args := stub.GetFunctionAndParameters()
	
	if fn == "set" {
		return t.Set(stub, args)
	} else if fn == "get" {
		return t.Get(stub, args)
	} 

	return shim.Error("Not supported function name")
}
// 5. Set 함수
func (t *SimpleAsset) Set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
		
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and value")
	}
	// 오류체크 중복 키 검사 -> 덮어쓰기로 해결

	asset := Asset{Key:args[0], Value:args[1]}
	// 구조체 Key - Value 세팅

	assetAsBytes, err := json.Marshal(asset)
	// Marshal 잘됐을때와 에러가 났을때

	if err != nil {
		// 에러가 null 이 아닐경우 👉 에러가 발생했을경우
		return shim.Error("Failed to Marshal args: " + args[0]+" "+args[1])
	}

	err = stub.PutState(args[0], assetAsBytes) // Marshal값 대입
	
	if err != nil {
		return shim.Error("Failed to set asset: " + args[0])
	}

	return shim.Success(assetAsBytes)
}
// 6. Get 함수
func (t *SimpleAsset) Get(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	
	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Filed to get asset: " + args[0] + " with error: " + err.Error())
	}
	if value == nil {
		return shim.Error("Asset not found: " + args[0])
	}

	return shim.Success([]byte(value))
}

// 6-1. Del 함수
func (t *SimpleAsset) Del(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	if len(args) != 1 { // 값이 있는지 확인 (args 값은 1개만 있으면 된다)
		return shim.Error("Incorrect arguments. Expecting a key")
	}

	value, err := stub.GetState(args[0]) // 조회 (args 값이 있는지 확인)

	if err != nil { // 에러가 null 이 아닐경우 👉 에러가 발생했을경우
		return shim.Error("Filed to get asset: " + args[0] + " with error: " + err.Error())
	}
	if value == nil { // value 값이 없을 경우
		return shim.Error("Asset not found: " + args[0])
	}

	err = stub.DelState(args[0]) // 삭제(args[0] 키 값 삭제 )

	if err != nil { // 에러가 null 이 아닐경우 👉 에러가 발생했을경우
		return shim.Error("Filed to get asset: " + args[0] + " with error: " + err.Error())
	}
	return shim.Success([]byte(args[0])) // key 값 (삭제한 값을 그대로 반환해본다)
}

// 6-2. Transfer 함수
func (t *SimpleAsset) Transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response{
	if len(args) != 3 { // 인자값 체크
		return shim.Error("Incorrect arguments. Expecting a from_key, to_key and amount")
	} // 보내는사람 / 받는사람 / 금액

	// args[0] : from_key / 보내는사람
	// args[1] : to_key / 받는사람
	// args[2] : amount / 금액
//-------------------------------------
	from_asset, err := stub.GetState(args[0]) // 보내는사람 값 조회 (args 값이 있는지 확인)

	if err != nil { // 에러가 null 이 아닐경우 👉 에러가 발생했을경우
		return shim.Error("Filed to get asset: " + args[0] + " with error: " + err.Error())
	}
	if from_asset == nil { // value 값이 없을 경우
		return shim.Error("Asset not found: " + args[0])
	}
//-------------------------------------
	to_asset, err := stub.GetState(args[1]) // 받는사람 값 조회 (args 값이 있는지 확인)

	if err != nil { // 에러가 null 이 아닐경우 👉 에러가 발생했을경우
		return shim.Error("Filed to get asset: " + args[1] + " with error: " + err.Error())
	}
	if to_asset == nil { // value 값이 없을 경우
		return shim.Error("Asset not found: " + args[1])
	}
//-------------------------------------
	// from_asset / to_asset 👉 언마샬 작업
	from := Asset{} // struct 값
	to := Asset{} // struct 값

	json.Unmarshal(from_asset.&from) // from_asset 을 언마샬해서 from 주소에 넣는다
	json.Unmarshal(to_asset.&to) // to_asset 을 언마샬해서 to 주소에 넣는다
//-------------------------------------

	// from_asset / to_asset / amount 정수형으로 변환

	from_amount, _ := strconv.Atoi(from.Value) // from 의 struct 값 (from_asset 을 언마샬해서 from 주소값)
	to_amount, _ := strconv.Atoi(to.Value) // to 의 struct 값 (to_asset 을 언마샬해서나온 to 주소값)
	amount, _ := strconv.Atoi(args[2])

	//-------------------------------------

	// 잔액 검증
	if (from_amount < amount) { // 보내는 값이 잔액보다 작은경우
		return shim.Error("Not enough asset value: "+args[0])
	}

	//-------------------------------------

	// 송금된 결과값 계산
	from.Value = strconv.Itoa(from_amount - amount) // 보냈으니 값을 그만큼 빼준다
	// amount값을 string 으로 바꿔서 넣어준다
	to.Value = strconv.Itoa(to_amount + amount) // 받았으니 값을 그만큼 더해준다

//-------------------------------------
// from_asset / to_asset 👉 마샬 작업
	from_asset, _ = json.Marshal(from)
	to_asset, _ = json.Marshal(to)
//-------------------------------------
	// PutState (값 추가작업)
	stub.PutState(args[0], from_asset)
	stub.PutState(args[1], to_asset)

	return shim.Success([]byte("transfer done!"))
}

// 7. Main 함수
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode : %s", err)
	}
}