package gobsc

import (
	"github.com/shopspring/decimal"
)

type BlockResponse struct {
	Difficulty      string         `json:"difficulty"`
	ExtraData       string         `json:"extraData"`
	GasLimit        string         `json:"gas_limit"`
	GasUsed         string         `json:"gasUsed"`
	Hash            string         `json:"hash"`
	LogsBloom       string         `json:"logsBloom"`
	Miner           string         `json:"miner"`
	MixHash         string         `json:"mixHash"`
	Nonce           string         `json:"nonce"`
	ParentHash      string         `json:"parentHash"`
	ReceiptsRoot    string         `json:"receiptsRoot"`
	Sha3Uncles      string         `json:"sha3Uncles"`
	Size            string         `json:"size"`
	StateRoot       string         `json:"stateRoot"`
	Timestamp       string         `json:"timestamp"`
	TotalDifficulty string         `json:"totalDifficulty"`
	Transactions    []*Transaction `json:"transactions"`
	TransactionRoot string         `json:"transactionRoot"`
	Uncles          []string       `json:"uncles"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
	V                string `json:"v"`
	Value            string `json:"value"`
}

// BscBlockLatestNumber 获取最新区块高度
func (r *rpcClient) BscBlockLatestNumber() (string, error) {
	var response string
	err := r.c.Call(&response, "eth_blockNumber")
	if err != nil {
		return "", nil
	}
	s := parseBscResponseString(response)
	newFromString, err := decimal.NewFromString(s)
	if err != nil {
		return "", err
	}
	return newFromString.String(), nil
}

// BscGetBlockByNumber 通过区块高度获取区块内容和元数据
func (r *rpcClient) BscGetBlockByNumber(blockNumber string) (*BlockResponse, error) {
	strHex := strToHex(blockNumber)
	strHex = splicingStr("0x", strHex)
	response := new(BlockResponse)
	if err := r.c.Call(response, "eth_getBlockByNumber", strHex, true); err != nil {
		return nil, err
	}
	return response, nil
}

// BscGetBlockByHash 通过区块hash获取区块内容和元数据
func (r *rpcClient) BscGetBlockByHash(blockHash string) (*BlockResponse, error) {
	response := new(BlockResponse)
	if err := r.c.Call(response, "eth_getBlockByHash", blockHash, true); err != nil {
		return nil, err
	}
	return response, nil
}
