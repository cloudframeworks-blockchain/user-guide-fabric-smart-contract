package models

import (
	"encoding/json"
	"fmt"
)

type ALLRecords struct {
	Records []*CharityNote
}

func QueryALLRecords(username string) (*ALLRecords, error) {
	rs := &ALLRecords{}
	cexec := &ExecArgs{}
	cexec.Args = append(cexec.Args, "queryDealALL")
	cexec.Args = append(cexec.Args, username)
	cexecBytes, err := json.Marshal(cexec)
	if err != nil {
		return rs, err
	}
	_, err = RunCommand(string(cexecBytes[:]))
	if err != nil {
		fmt.Println(err.Error())
	}

	i := 1
	for i <= 3 {
		cNote := &CharityNote{
			Direction:    "juankuan",
			CostMoney:    10000,
			DonationName: username,
		}
		rs.Records = append(rs.Records, cNote)
		i += 1
	}
	return rs, nil
}
