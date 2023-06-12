package main

import (
	"encoding/json"
	"ethers_scaenner/abstract"
	"ethers_scaenner/infura"
	"fmt"
	"log"
)

func getEthClient() *abstract.EthClient {
	infuraClient, err := infura.NewJsonRpcClient()
	if err != nil {
		// handle error
	}
	ethClient := &abstract.EthClient{
		Client: infuraClient,
	}
	return ethClient
}

func getBlockByNumber(ethClient *abstract.EthClient) {
	var blockNum int64 = 9163864

	block, err := ethClient.GetBlockByNumber(blockNum)
	if err != nil {
		// handle error
	}
	jsonData, err := json.Marshal(block)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func getLatestBlock(ethClient *abstract.EthClient) {
	latestBlock, err := ethClient.GetLatestBlock()
	if err != nil {
		// handle error
	}
	fmt.Println(latestBlock)
}
