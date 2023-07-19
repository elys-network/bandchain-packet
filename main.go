package main

import (
	"encoding/hex"
	"fmt"

	"github.com/bandprotocol/bandchain-packet/packet"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	calldata, _ := hex.DecodeString("030000004254436400000000000000")
	req := packet.OracleRequestPacketData{
		ClientID:       "",
		OracleScriptID: 1,
		Calldata:       calldata,
		AskCount:       1,
		MinCount:       1,
		FeeLimit:       sdk.NewCoins(sdk.NewCoin("uband", sdk.NewInt(10000))),
		PrepareGas:     100,
		ExecuteGas:     100,
	}
	// This can be data of packet
	fmt.Println(string(req.GetBytes()))
}
