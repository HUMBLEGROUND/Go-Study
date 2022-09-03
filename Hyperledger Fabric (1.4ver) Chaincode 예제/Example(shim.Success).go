package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
// Init 함수
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 👇 이번엔 shim.Success(nil) 을 바로 리턴하지말고 확인하는 작업을 거친다
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters() // 👉 Func & Para
	// 관리자가 체인코드를 배포하면서 처음부터 key 값, value 값을 넣어서 배포작업 
	// 함수 이름과 파라미터를 잘라서 사용할수 있도록 처리
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
 
	// Initialize the chaincode
	A = args[0] // 데이터를 읽어와서 배열에 넣는다
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	// 배포할때 들어온 데이터를 가지고 초기 월드스테이트에 넣어준다

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	// 👉 월드스테이트에 A 값과 B 값을 각각 저장

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// shim 인터페이스를 사용할 경우 Init 과 같이 ⭐ Invoke 도 필수요소이다

	// 체인코드가 배포가 끝나고 어플, 웹서버가 동작을하게 되면
	// 사용자가 자신의 단말기(컴퓨터, 핸드폰, IOT 등)로 인터넷을 통해 트랜잭션을 일으키면 
	// 트랜잭션 내용이 패브릭 네트워크내에 peer 에게 전달
	// peer 는 트랜잭션을 받아서 
	// 자신이 보유하고 있는 체인코드 컨테이너쪽에 그 트랜잭션을 보낸다
	// 그리고 시뮬레이션을 돌림
	// 👉 그때 호출되는게 Invoke 함수이다
	
	// 사용자가 블록체인에 트랜잭션을 일으켰을때
	// 데이터를 써 넣는다거나 읽어오거나 요청하거나 데이터를 저장할때
	// 반드시 호출되는게 Invoke 함수이다
	// Invoke 함수가 호출되어 동작하게되면 peer는 트랜잭션을 받은것이다
	// 
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters() // 👉 Func & Para
	// 들어온 데이터를 스마트컨트랙트 함수 이름과 
	// 스마트컨트랙트가 동작하기 위한 파라미터들이 배열형식으로 들어온다.
	// 함수 이름과 파라미터를 잘라서 사용할수 있도록 처리
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
