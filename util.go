package gobsc

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
	"strings"
)

// parseBscResponseString
func parseBscResponseString(data string) string {

	if strings.HasPrefix(data, "0x") {
		data = data[2:]
	}
	if len(data) > 128 {
		// todo
	} else if len(data) == 64 {
		// todo
	} else {
		value, ok := new(big.Int).SetString(data, 16)
		if ok {
			return value.String()
		}
	}

	return ""
}

// strToHex 字符转十六进制
func strToHex(s string) string {
	a, _ := strconv.Atoi(s)
	return strconv.FormatInt(int64(a), 16)
}

// splicingStr 拼接字符串
func splicingStr(args ...string) string {
	var builder strings.Builder
	for i := range args {
		builder.WriteString(args[i])
	}
	return builder.String()
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["input"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	if msg.GasFeeCap != nil {
		arg["maxFeePerGas"] = (*hexutil.Big)(msg.GasFeeCap)
	}
	if msg.GasTipCap != nil {
		arg["maxPriorityFeePerGas"] = (*hexutil.Big)(msg.GasTipCap)
	}
	if msg.AccessList != nil {
		arg["accessList"] = msg.AccessList
	}
	if msg.BlobGasFeeCap != nil {
		arg["maxFeePerBlobGas"] = (*hexutil.Big)(msg.BlobGasFeeCap)
	}
	if msg.BlobHashes != nil {
		arg["blobVersionedHashes"] = msg.BlobHashes
	}
	return arg
}
