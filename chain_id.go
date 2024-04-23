package gobsc

// BscChainID get chain id
func (r *rpcClient) BscChainID() (string, error) {
	var response string
	err := r.c.Call(&response, "eth_chainId")
	if err != nil {
		return "", err
	}
	chanId := parseBscResponseString(response)
	return chanId, nil
}
