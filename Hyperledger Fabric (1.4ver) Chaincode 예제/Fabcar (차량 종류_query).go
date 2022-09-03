 package main

 import (
	 "bytes"
	 "encoding/json" // Marshal 사용
	 "fmt"
	 "strconv"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 ) // shim, peer 설치 라이브러리 인터페이스
 // shim 👉 GetState / PutState 👉 원장의 데이터를 쓰거나 읽어올때 사용한다
 // GetFunctionAndParameters 👉 트랜잭션이 들어왔을때 트랜잭션 내용에 대해서 스마트컨트랙트의 이름과 스마트컨트랙트를 동작시키는데 필요한 파라미터들을 잘라주는 함수
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 type Car struct { // 자동차에 대한 구조체
	 Make   string `json:"make"`
	 Model  string `json:"model"`
	 Colour string `json:"colour"`
	 Owner  string `json:"owner"`
 } // 월드스테이트의 value로 저장된다
 
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 } 
 
 // shim 과 peer 를 불러와서 체인코드를 작성한다면 
 // 👉 Init 과 Invoke 를 필수로 써줘야한다 ⭐
 
 // 개발자가 체인코드를 작성하면 관리자에게 보내준다
 // 관리자는 블록체인(패브릭)에 자신의 peer 에게 체인코드를 설치 (copy)
 // 패브릭 네트워크에 체인코드를 활성화 시킨다 (배포)
 // 체인코드를 배포하게되면 
 // peer 는 체인코드 컨테이너를 생성
 // 해당 체인코드를 체인코드 컨테이너에 보내고 컴파일
 // 그럼 트랜잭션을 받을 준비한다 
 // 체인코드가 활성화 된 후에 제일먼저 실행되는게 Init 함수다 ⭐
 
 // shim.Success(nil) 👉 체인코드를 배포를 시킨 관리자에게
 // 관리자에 요청한 체인코드의 작업이 성공적으로 끝났다는 뜻
 // 배포를 했을때 문제가 있다면 Init 함수가 실행되지 않음
 
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 

	 function, args := APIstub.GetFunctionAndParameters()

	 if function == "queryCar" {
		 return s.queryCar(APIstub, args)
	 } else if function == "initLedger" {
		 return s.initLedger(APIstub)
	 } else if function == "createCar" {
		 return s.createCar(APIstub, args)
	 } else if function == "queryAllCars" {
		 return s.queryAllCars(APIstub)
	 } else if function == "changeCarOwner" {
		 return s.changeCarOwner(APIstub, args)
	 }
 
	 return shim.Error("Invalid Smart Contract function name.")
 }
 
 func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 return shim.Success(carAsBytes)
 }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 cars := []Car{
		 Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		 Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		 Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		 Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		 Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		 Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		 Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		 Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		 Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		 Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	 }
 
	 i := 0
	 for i < len(cars) {
		 fmt.Println("i is ", i)
		 carAsBytes, _ := json.Marshal(cars[i])
		 APIstub.PutState("CAR"+strconv.Itoa(i), carAsBytes)
		 fmt.Println("Added", cars[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 5 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
 
	 var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
 
	 carAsBytes, _ := json.Marshal(car)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "CAR0"
	 endKey := "CAR999"
 
	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
		 }
		 // Add a comma before array members, suppress it for the first array member
		 if bArrayMemberAlreadyWritten == true {
			 buffer.WriteString(",")
		 }
		 buffer.WriteString("{\"Key\":")
		 buffer.WriteString("\"")
		 buffer.WriteString(queryResponse.Key)
		 buffer.WriteString("\"")
 
		 buffer.WriteString(", \"Record\":")
		 // Record is a JSON object, so we write as-is
		 buffer.WriteString(string(queryResponse.Value))
		 buffer.WriteString("}")
		 bArrayMemberAlreadyWritten = true
	 }
	 buffer.WriteString("]")
 
	 fmt.Printf("- queryAllCars:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }
 
 func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 carAsBytes, _ := APIstub.GetState(args[0])
	 car := Car{}
 
	 json.Unmarshal(carAsBytes, &car)
	 car.Owner = args[1]
 
	 carAsBytes, _ = json.Marshal(car)
	 APIstub.PutState(args[0], carAsBytes)
 
	 return shim.Success(nil)
 }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 