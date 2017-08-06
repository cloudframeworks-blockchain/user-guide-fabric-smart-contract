package models

import (
	"encoding/json"
	"fmt"
)

func GetUserInfo(username string) (*CharityUser, error) {
	cUser := &CharityUser{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "queryUserInfo")
	cexec.Args = append(cexec.Args, username)
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		return cUser, err
	}
	f, err := RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f)

	cUser = &CharityUser{
		DonationName: username,
		ALLMoney:     80000,
		LeftMoney:    16000,
		DealNumbers:  4,
	}
	return cUser, nil
}
