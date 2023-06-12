package main

func main() {
	ethClient := getEthClient()
	//getLatestBlock(ethClient)
	getBlockByNumber(ethClient)
}
