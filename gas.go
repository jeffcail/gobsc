package gobsc

import (
	"github.com/ethereum/go-ethereum"
)

// BacEstimateGas
// 执行并估算一个交易需要的gas用量。该次交易不会写入区块链。注意，由于多种原因，例如EVM的机制 及节点旳性能，估算的数值可能比实际用量大的多。
func (r *rpcClient) BacEstimateGas(callMsg ethereum.CallMsg) (string, error) {
	var response string
	err := r.c.Call(&response, "eth_estimateGas", toCallArg(callMsg))
	if err != nil {
		return "", err
	}
	data := parseBscResponseString(response)
	return data, nil
}
