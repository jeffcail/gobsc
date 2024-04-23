package gobsc

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// BacGenerateAccount 生成地址和私钥
func BacGenerateAccount() (string, string, error) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	// generate address of use private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	privateKeyStr := hexutil.Encode(crypto.FromECDSA(privateKey))

	return address, privateKeyStr, nil
}
