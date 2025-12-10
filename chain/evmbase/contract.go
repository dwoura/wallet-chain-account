package evmbase

import (
	"encoding/hex"
	"strings"

	"github.com/dapplink-labs/wallet-chain-account/rpc/account"
	"github.com/dapplink-labs/wallet-chain-account/rpc/common"
)

// CallContract executes a read-only contract call using eth_call
func CallContract(client EthClient, req *account.CallContractRequest) (*account.CallContractResponse, error) {
	// Remove "0x" prefix if present
	dataHex := strings.TrimPrefix(req.Data, "0x")

	// Decode hex to bytes
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return &account.CallContractResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "invalid calldata hex: " + err.Error(),
		}, nil
	}

	// Call the contract using existing EthCallContract
	result, err := client.EthCallContract(req.ContractAddress, data)
	if err != nil {
		return &account.CallContractResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "contract call failed: " + err.Error(),
		}, nil
	}

	return &account.CallContractResponse{
		Code:   common.ReturnCode_SUCCESS,
		Msg:    "success",
		Result: result,
	}, nil
}
