package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type CharityNote struct {
	Direction    string `json:"direction"`
	CostMoney    int32  `json:"costMoney"`
	DonationName string `json:"donationName"`
}

type CharityUser struct {
	DonationName string `json:"donationName"`
	ALLMoney     int32  `json:"allMoney"`
	LeftMoney    int32  `json:"leftMoney"`
	DealNumbers  int    `json:"dealNumbers"`
}

func Sha8(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs[:4])
}

func Skey(name string, num int) string {
	return fmt.Sprintf("%s_%d", Sha8(name), num)
}

const (
	D0 string = "D0"
	// origin
	D1 string = "D1"
	D2 string = "D2"
	D3 string = "D3"
)

func (s *SmartContract) Init(api shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(api shim.ChaincodeStubInterface) peer.Response {

	function, args := api.GetFunctionAndParameters()

	switch function {
	case "donation":
		return s.donation(api, args)
	case "queryDealOnce":
		return s.queryDealOnce(api, args)
	case "queryDealALL":
		return s.queryDealALL(api, args)
	case "queryUserInfo":
		return s.queryUserInfo(api, args)
	case "donationRules":
		return s.donationRules(api, args)
	}

	return shim.Error("Invalid function name.")
}

func (s *SmartContract) set(api shim.ChaincodeStubInterface, key string, value []byte) error {
	err := api.PutState(key, value)
	if err != nil {
		return fmt.Errorf("Failed to set asset: %s", key)
	}
	return nil
}

func (s *SmartContract) get(api shim.ChaincodeStubInterface, key string) ([]byte, error) {
	value, err := api.GetState(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to get asset: %s with error: %s", key, err)
	}
	if value == nil {
		return nil, fmt.Errorf("Asset not found: %s", key)
	}
	return value, nil
}

/*
func (s *SmartContract) getRange(api shim.ChaincodeStubInterface, keyStart string, keyEnd string) (shim.StateQueryIteratorInterface, error) {
	resultsIterator, err := api.GetStateByRange(keyStart, keyEnd)
	if err != nil {
		return shim.StateQueryIteratorInterface, fmt.Errorf("get range error, %s", err)
	}
	defer resultsIterator.Close()
	return resultsIterator, nil
}
*/

func (s *SmartContract) donation(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("need your name and deal numbers")
	}
	moneyCount, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error("strconv money error.")
	}
	cUser := CharityUser{}
	cNote := CharityNote{}
	cUserValues, err := api.GetState(args[0])
	if err != nil {
		return shim.Error("get cuer info error, " + err.Error())
	}

	if cUserValues == nil {
		cUser = CharityUser{
			DonationName: args[0],
			ALLMoney:     int32(moneyCount),
			LeftMoney:    int32(moneyCount),
			DealNumbers:  0,
		}
		cNote = CharityNote{
			Direction:    "nochange",
			CostMoney:    0,
			DonationName: args[0],
		}
		cNoteKey := Skey(cNote.DonationName, cUser.DealNumbers)
		cNoteBytes, _ := json.Marshal(cNote)
		err = s.set(api, cNoteKey, cNoteBytes)
		if err != nil {
			return shim.Error("set cnote error" + err.Error())
		}
		cUserKey := cNote.DonationName
		cUserBytes, _ := json.Marshal(cUser)
		err = s.set(api, cUserKey, cUserBytes)
		if err != nil {
			return shim.Error("set cuser error" + err.Error())
		}
	} else {
		json.Unmarshal(cUserValues, &cUser)
		cUser.ALLMoney = cUser.ALLMoney + int32(moneyCount)
		cUser.LeftMoney = cUser.LeftMoney + int32(moneyCount)
		cUserKey := args[0]
		cUserBytes, _ := json.Marshal(cUser)
		err = s.set(api, cUserKey, cUserBytes)

		if err != nil {
			return shim.Error("update cuser error" + err.Error())
		}
	}
	return shim.Success(nil)
}

func (s *SmartContract) queryDealOnce(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("need your name and deal numbers")
	}
	if args[1] == "0" {
		return shim.Error("cant query 0 nums deal.")
	}
	nums, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
	cNoteKey := Skey(args[0], nums)
	cNoteValue, err := s.get(api, cNoteKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(cNoteValue)
}

func (s *SmartContract) queryDealALL(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("need your name")
	}
	cUserValue, err := s.get(api, args[0])
	if err != nil {
		return shim.Error("get cuer key error.")
	}
	cUser := &CharityUser{}
	err = json.Unmarshal(cUserValue, &cUser)
	if err != nil {
		return shim.Error("json cuservalue error.")
	}
	totalDealNums := cUser.DealNumbers

	cNoteKeyEnd := Skey(args[0], totalDealNums+1)
	cNoteKeyStart := Skey(args[0], 0)
	if cNoteKeyEnd == cNoteKeyStart {
		var buffer bytes.Buffer
		cNoteValue, err := s.get(api, cNoteKeyEnd)
		if err != nil {
			return shim.Error(err.Error())
		}
		buffer.WriteString("[[[")
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(cNoteKeyEnd)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(cNoteValue))
		buffer.WriteString("}")
		buffer.WriteString("]]]")
		return shim.Success(buffer.Bytes())
	}
	resultsIter, err := api.GetStateByRange(cNoteKeyStart, cNoteKeyEnd)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[[[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		fmt.Printf("requset is %s", queryResponse.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]]]")
	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryUserInfo(api shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("need your name")
	}
	cUserValue, err := s.get(api, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(cUserValue)
}

//########################
// donation ruls contract
//########################

func (s *SmartContract) modelAssign(api shim.ChaincodeStubInterface, name string, id string) error {
	cUserValue, err := s.get(api, name)
	if err != nil {
		return fmt.Errorf("modelassign get name error")
	}
	cUser := &CharityUser{}
	err = json.Unmarshal(cUserValue, &cUser)
	if err != nil {
		return fmt.Errorf("get cuser error")
	}
	cNote := &CharityNote{
		Direction:    id,
		CostMoney:    cUser.LeftMoney,
		DonationName: name,
	}

	cNoteKey := Skey(name, cUser.DealNumbers+1)
	cNoteBytes, _ := json.Marshal(cNote)
	err = s.set(api, cNoteKey, cNoteBytes)
	if err != nil {
		return fmt.Errorf("set cnote error")
	}
	cUser.LeftMoney = 0
	cUser.DealNumbers = cUser.DealNumbers + 1
	cUserBytes, _ := json.Marshal(cUser)
	err = s.set(api, name, cUserBytes)
	if err != nil {
		return fmt.Errorf("set cuser error")
	}
	return nil
}

func (s *SmartContract) modelRandom(api shim.ChaincodeStubInterface, name string) error {
	cUserValue, err := s.get(api, name)
	if err != nil {
		return fmt.Errorf("modelassign get name error")
	}
	cUser := &CharityUser{}
	err = json.Unmarshal(cUserValue, &cUser)
	if err != nil {
		return fmt.Errorf("get cuser error")
	}
	var cost int32
	cost = 10000
	if cUser.LeftMoney < 10000 {
		fmt.Print("left money less than 10000. can't do.")
		cost = cUser.LeftMoney
	}
	randomList := []string{"xiwanggongcheng", "zaiqu", "guojijiuyuan", "kunnanjiating"}
	t := rand.Intn(3)
	cNote := &CharityNote{
		Direction:    randomList[t],
		CostMoney:    cost,
		DonationName: name,
	}
	cNoteKey := Skey(name, cUser.DealNumbers+1)
	cNoteBytes, _ := json.Marshal(cNote)
	err = s.set(api, cNoteKey, cNoteBytes)
	if err != nil {
		return fmt.Errorf("set cnote error")
	}
	// 更新charityNote

	cUser.LeftMoney = cUser.LeftMoney - cost
	cUser.DealNumbers = cUser.DealNumbers + 1
	cUserBytes, _ := json.Marshal(cUser)
	err = s.set(api, name, cUserBytes)
	if err != nil {
		return fmt.Errorf("set cuser error")
	}
	// 更新charityUser

	//###########################
	// 随机4个机构其中的一个，单次最大捐款额度为1w
	return nil
}

func (s *SmartContract) donationRules(api shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 3 {
		// assignModel
		if args[1] == "assign" {
			//'{"Args":["username", "assing", "xiwangxiaoxue"]}'
			err := s.modelAssign(api, args[0], args[2])
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	} else if len(args) == 2 {
		if args[1] == "random" {
			err := s.modelRandom(api, args[0])
			if err != nil {
				return shim.Error(err.Error())
			}
		}
	} else {
		return shim.Error("donationRules args Error.")
	}

	return shim.Success(nil)
}

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
