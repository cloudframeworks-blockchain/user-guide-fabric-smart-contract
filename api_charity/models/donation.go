package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Donation struct {
	Username string
	Money    int32
}

func DonationUser(username, money string) (*CharityUser, error) {
	cUser := &CharityUser{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "donation")
	cexec.Args = append(cexec.Args, username)
	cexec.Args = append(cexec.Args, money)
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		return cUser, err
	}
	f, err := RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f)

	iMoney, err := strconv.ParseInt(money, 10, 64)
	if err != nil {
		return cUser, err
	}
	cUser = &CharityUser{
		DonationName: username,
		ALLMoney:     int32(iMoney),
		LeftMoney:    int32(iMoney),
		DealNumbers:  0,
	}
	return cUser, nil
}
