package gobsc

import "log"

// call
func (r *rpcClient) call(param interface{}) (string, error) {
	var response string
	err := r.c.Call(&response, "eth_call", param, "latest")
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}
