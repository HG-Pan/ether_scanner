package infura

import (
	"ether_scaenner/abstract"
	"ether_scaenner/config"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"strconv"
)

type JsonRpcClient struct {
	client *rpc.Client
}

func NewJsonRpcClient() (*JsonRpcClient, error) {

	client, err := rpc.Dial(config.TargetUrl)
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}
	return &JsonRpcClient{client: client}, nil
}

func (ijrc *JsonRpcClient) GetBlockByNumber(blockNumber int64) (*abstract.Block, error) {
	var blockHex abstract.BlockHex
	hexBlockNumber := fmt.Sprintf("0x%x", blockNumber) // 将 blockNumber 转换为带有 0x 前缀的十六进制字符串
	err := ijrc.client.Call(&blockHex, "eth_getBlockByNumber", hexBlockNumber, true)
	if err != nil {
		log.Printf("eth_getBlockByNumber: %v", err)
		return nil, err
	}

	block := abstract.Block{
		Number:     blockNumber,
		Hash:       blockHex.Hash,
		ParentHash: blockHex.ParentHash,
		Difficulty: blockHex.Difficulty,
		Miner:      blockHex.Miner,
	}

	gasLimit, err := parseHexInt64(blockHex.GasLimit)
	if err != nil {
		log.Printf("Failed to parse gasLimit: %v", err)
		return nil, err
	}
	block.GasLimit = gasLimit

	gasUsed, err := parseHexInt64(blockHex.GasUsed)
	if err != nil {
		log.Printf("Failed to parse gasUsed: %v", err)
		return nil, err
	}
	block.GasUsed = gasUsed

	nonce, err := parseHexInt64(blockHex.Nonce)
	if err != nil {
		log.Printf("Failed to parse nonce: %v", err)
		return nil, err
	}
	block.Nonce = nonce

	timestamp, err := parseHexInt64(blockHex.Timestamp)
	if err != nil {
		log.Printf("Failed to parse timestamp: %v", err)
		return nil, err
	}
	block.Timestamp = timestamp
	transactions, err := convertTransactionHexToTransaction(blockHex.Transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to convert transactions: %v", err)
	}
	block.Transactions = transactions
	return &block, nil
}

func convertTransactionHexToTransaction(transactions []abstract.TransactionHex) ([]abstract.Transaction, error) {
	converted := make([]abstract.Transaction, len(transactions))
	for i, txHex := range transactions {
		value, err := strconv.ParseInt(txHex.Value, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction value: %v", err)
		}
		nonce, err := strconv.ParseInt(txHex.Nonce, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction value: %v", err)
		}

		gas, err := strconv.ParseInt(txHex.Gas, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction gas: %v", err)
		}

		gasPrice, err := strconv.ParseInt(txHex.GasPrice, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction gas price: %v", err)
		}

		transactionIndex, err := strconv.ParseInt(txHex.TransactionIndex, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction block number: %v", err)
		}
		blockNumber, err := strconv.ParseInt(txHex.BlockNumber, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction block number: %v", err)
		}

		converted[i] = abstract.Transaction{
			Hash:             txHex.Hash,
			Nonce:            nonce,
			BlockHash:        txHex.BlockHash,
			BlockNumber:      blockNumber,
			TransactionIndex: transactionIndex,
			From:             txHex.From,
			To:               txHex.To,
			Value:            value,
			Gas:              gas,
			GasPrice:         gasPrice,
			Input:            txHex.Input,
		}
	}

	return converted, nil
}

func (ijrc *JsonRpcClient) GetLatestNum() (int64, error) {
	var blockNumberHex string
	err := ijrc.client.Call(&blockNumberHex, "eth_blockNumber")
	if err != nil {
		log.Fatalf("eth_blockNumber: %v", err)
		return 0, err
	}

	// 去掉 "0x" 前缀
	blockNumberHex = blockNumberHex[2:]

	// 将十六进制字符串转换为 int64
	blockNumber, err := strconv.ParseInt(blockNumberHex, 16, 64)
	if err != nil {
		return 0, err
	}
	return blockNumber, nil
}
