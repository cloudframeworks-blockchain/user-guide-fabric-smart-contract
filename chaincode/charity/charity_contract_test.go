package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("0", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkDonation(t *testing.T, stub *shim.MockStub, name string) {
	res := stub.MockInvoke("0", [][]byte{[]byte("queryUserInfo"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	fmt.Println("Query value", "name is ", name, "value is ", value)
	t.FailNow()
}

func TestCharityContract_Donation(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("charity", scc)

	checkDonation(t, stub, "mm")
}
