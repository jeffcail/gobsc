package gobsc

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

var (
	jsonRpcUrl = "https://go.getblock.io/e1f042a68e834f36855546391dfeac82"
)

func init() {
	NewClient(jsonRpcUrl)
}

func TestEthClient(t *testing.T) {

	tests := map[string]struct {
		test func(t *testing.T)
	}{

		"BacGenerateAccount": {
			func(t *testing.T) { testBacGenerateAccount(t) },
		},

		"TestNewClient": {
			func(t *testing.T) { testNewClient(t) },
		},

		"BscChainID": {
			func(t *testing.T) { testBscChainID(t) },
		},

		"BacEstimateGas": {
			func(t *testing.T) { testBacEstimateGas(t) },
		},

		"BscBlockLatestNumber": {
			func(t *testing.T) { testBscBlockLatestNumber(t) },
		},

		"BscGetBlockByNumber": {
			func(t *testing.T) { testBscGetBlockByNumber(t) },
		},

		"GetBlockByHash": {
			func(t *testing.T) { testGetBlockByHash(t) },
		},
	}

	t.Parallel()
	for name, tt := range tests {
		t.Run(name, tt.test)
	}
}

// address: 0xB74B69B77ADc4a54D9f42d4F899C7EEC374e3AE5  privateKey: 0x849f57fc2d79cec5f56288733ba84e96646571b68356d565f57d8665965c61f6
// address: 0xbB21069dE3c65e877Baf976cEba5F634321549BD  privateKey: 0xfcd038318a714ed11813d801b4ed1d00342e3306ee515b9e4f8cea64bf64ee3e
func testBacGenerateAccount(t *testing.T) {
	address, privateKey, err := BacGenerateAccount()
	if err != nil {
		t.Fatalf("error: %s\n", err)
	}
	t.Logf("address: %s\n", address)
	t.Logf("privateKey: %s\n", privateKey)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, "generate bsc address", address)
	//	assert.Equal(t, "generate bsc private key", privateKey)
	//}
}

func testNewClient(t *testing.T) {

	t.Logf("client created 【%v】", Cli)
}

func testBscChainID(t *testing.T) {
	chainID, err := Cli.BscChainID()
	if err != nil {
		t.Fatalf("error: %s\n", err)
	}
	t.Logf("chainID: %x\n", chainID)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, "get chain id", chainID)
	//}
}

func testBacEstimateGas(t *testing.T) {

	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x08265dA01E1A65d62b903c7B34c08cB389bF3D99"),
		To:   &common.Address{},
		//Gas:   2100,
		Value: big.NewInt(1),
	}
	estimateGas, err := Cli.BacEstimateGas(callMsg)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	t.Logf("estimateGas: %v\n", estimateGas)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, "estimated gas", estimateGas)
	//}
}

func testBscBlockLatestNumber(t *testing.T) {
	number, err := Cli.BscBlockLatestNumber()
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}
	t.Logf("block number: %s\n", number)
	//if assert.NoError(t, err) {
	//	assert.Equal(t, "get latest block number", number)
	//}
}

func testBscGetBlockByNumber(t *testing.T) {
	var blockNumber string
	blockNumber = "39708840"
	response, err := Cli.BscGetBlockByNumber(blockNumber)

	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	for _, transaction := range response.Transactions {
		fmt.Println(transaction.Hash)
	}

	t.Logf("block hash: %s\n", response.Hash)
	t.Logf("et block by block number: %v\n", response)

	//if assert.NoError(t, err) {
	//	assert.Equal(t, "block hash", response.Hash) // 0xe6a747a590b0e7d36dd816b1cd44b0849acca7ea6f2c5c2e32a52bff01c6a7a1
	//	assert.Equal(t, "get block by block number", response)
	//}
}

func testGetBlockByHash(t *testing.T) {
	var blockHash string
	blockHash = "0xe6a747a590b0e7d36dd816b1cd44b0849acca7ea6f2c5c2e32a52bff01c6a7a1"
	response, err := Cli.BscGetBlockByHash(blockHash)

	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	for _, transaction := range response.Transactions {
		fmt.Println(transaction.Hash)
	}

	t.Logf("block hash: %s\n", response.Hash)
	t.Logf("et block by block number: %v\n", response)

	//if assert.NoError(t, err) {
	//	assert.Equal(t, "get block by block hash", response)
	//}
}
