package models

import "os/exec"

const (
	CmdExec string = "peer chaincode invoke -n charity -c '%s' -C myc"
)

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

type ExecArgs struct {
	Args []string
}

func RunCommand(arg string) (string, error) {

	f, err := exec.Command("peer", "chaincode", "invoke", "-n", "charity", "-c", arg, "-C", "mycs").Output()
	//f, err := exec.Command("echo", arg).Output()
	if err != nil {
		return "", err
	}
	for i := 0; i < len(f); i++ {
		if f[i] == 0 {
			return string(f[0:i]), nil
		}
	}
	return string(f), nil
}
