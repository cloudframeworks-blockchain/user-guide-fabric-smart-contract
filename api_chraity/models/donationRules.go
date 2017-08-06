package models

import (
	"encoding/json"
	"fmt"
)

type RcRules struct {
	CNotes []*CharityNote
	CUser  *CharityUser
}

func DonationRulesUser(args ...string) (*RcRules, error) {
	rr := &RcRules{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "donationRules")
	cexec.Args = append(cexec.Args, args[0])
	cexec.Args = append(cexec.Args, args[1])
	if len(args) == 3 {
		cexec.Args = append(cexec.Args, args[2])
	}
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		return rr, err
	}
	f, err := RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f)

	cUser := &CharityUser{
		DonationName: args[0],
		ALLMoney:     60000,
		LeftMoney:    80000,
		DealNumbers:  4,
	}
	rr.CUser = cUser
	i := 1
	for i <= 4 {
		cNote := &CharityNote{
			Direction:    "juankuan",
			CostMoney:    10000,
			DonationName: args[0],
		}
		rr.CNotes = append(rr.CNotes, cNote)
		i += 1
	}

	return rr, nil
}
