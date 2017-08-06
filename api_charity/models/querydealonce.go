package models

import (
	"encoding/json"
	"fmt"
)

func QueryOnce(username, nums string) (*CharityNote, error) {
	cNote := &CharityNote{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "queryDealOnce")
	cexec.Args = append(cexec.Args, username)
	cexec.Args = append(cexec.Args, nums)
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		return cNote, err
	}
	_, err = RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}
	cNote = &CharityNote{
		Direction:    "xiwangxiaoxue",
		CostMoney:    30000,
		DonationName: username,
	}
	return cNote, nil
}
