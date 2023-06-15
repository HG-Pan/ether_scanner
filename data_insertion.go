package main

import (
	"context"
	"encoding/json"
	"ether_scaenner/abstract"
	"ether_scaenner/config"
	"ether_scaenner/infura"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"sync/atomic"
	"time"
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

func getBlockByNumber(blockNum int64) {

	block, err := getEthClient().GetBlockByNumber(blockNum)
	if err != nil {
		// handle error
	}
	jsonData, err := json.Marshal(block)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func getBlockByNum(blockNum int64) (block *abstract.Block) {
	block, err := getEthClient().GetBlockByNumber(blockNum)
	if err != nil {
		log.Fatal(err)
	}
	return block
}

func getLatestNum() (latestBlock int64) {
	latestBlock, err := getEthClient().GetLatestNum()
	if err != nil {
		// handle error
	}
	return latestBlock
}

func insetBlockInfoCurrent(currentNum int64, database string, table string) {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.TODO())

	// 选择数据库和集合
	collection := client.Database(database).Collection(table)

	// 创建一个错误通道来传递插入错误
	errChan := make(chan error)

	go func() {

		for {
			latestNum := getLatestNum()
			// 如果当前块号大于最新块号，则等待一秒再尝试
			if currentNum > latestNum {
				time.Sleep(1 * time.Second)
				continue
			}

			// 在当前块和最新块之间插入所有的块
			for currentNum <= latestNum {
				// 获取并插入当前块的数据
				block := getBlockByNum(currentNum)
				_, err = collection.InsertOne(context.TODO(), block)
				if err != nil {
					// 将错误发送到错误通道并返回
					errChan <- err
					currentNum++
				}
				// 准备插入下一个块的数据
				currentNum++
				// 插入完成后，等待一秒再继续
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 检查错误通道
	for err := range errChan {
		// 处理错误
		log.Println(err)
	}
}

func initUniqueIndex(database string, table string, uniIndex string) {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.TODO())

	// 选择数据库和集合
	collection := client.Database(database).Collection(table)

	// 创建唯一索引
	indexModel := mongo.IndexModel{
		Keys:    bson.M{uniIndex: 1},
		Options: options.Index().SetUnique(true),
	}

	// 检查索引是否存在
	cursor, err := collection.Indexes().List(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var index bson.M
		if err := cursor.Decode(&index); err != nil {
			log.Fatal(err)
		}

		if _, ok := index["key"]; ok && index["key"].(bson.M)["number"] != nil {
			// 索引已存在
			log.Println("唯一索引已存在")
			return
		}
	}

	// 创建唯一索引
	_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("唯一索引已成功创建")
}

func insetBlockInfoHistory(latestNum int64, database string, table string) {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.TODO())

	// 选择数据库和集合
	collection := client.Database(database).Collection(table)

	var currentNum = atomic.LoadInt64(getTableMaxNum("ethereum", "blocks")) // 创建一个 WaitGroup 来等待所有 Goroutine 完成
	var wg sync.WaitGroup
	// 设置 Goroutine 的数量
	goroutines := 10 // 根据需求设置合适的 Goroutine 数量
	// 计算每个 Goroutine 需要处理的区块范围
	blockRange := (latestNum - currentNum + 1) / int64(goroutines)
	remainder := (latestNum - currentNum + 1) % int64(goroutines)
	// 循环创建 Goroutine
	for i := 0; i < goroutines; i++ {
		// 增加 WaitGroup 的计数器
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 计算该 Goroutine 需要处理的区块范围
			startBlock := currentNum + int64(i)*blockRange
			endBlock := startBlock + blockRange - 1
			// 对最后一个 Goroutine 处理剩余的区块范围
			if i == goroutines-1 {
				endBlock += remainder
			}
			// 在 Goroutine 中插入区块数据
			for blockNum := startBlock; blockNum <= endBlock; blockNum++ {
				// 处理该区块的逻辑
				block := getBlockByNum(blockNum)
				_, err := collection.InsertOne(context.TODO(), block)
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("区块 %d 已成功插入到MongoDB\n", block.Number)
				}
			}
		}(i)
	}
	// 等待所有 Goroutine 完成
	wg.Wait()
	log.Println("所有区块已成功插入到MongoDB")
}

func insetBlockInfo() {
	initUniqueIndex("ethereum", "blocks", "number")
	latestNum := getLatestNum()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		insetBlockInfoHistory(latestNum, "ethereum", "blocks")
		wg.Done()
	}()
	go func() {
		insetBlockInfoCurrent(latestNum, "ethereum", "blocks")
		wg.Done()
	}()
	wg.Wait()
}

// getDBBlockByNumber("ethereum","blocks")
func getDBBlockByNumber(blockNumber int64, database string, table string) (*abstract.Block, error) {
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		return nil, err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {

		}
	}(client, context.TODO())

	// 选择数据库和集合
	collection := client.Database(database).Collection(table)

	// 构建查询条件
	filter := bson.M{"number": blockNumber}

	// 执行查询
	var block abstract.Block
	err = collection.FindOne(context.TODO(), filter).Decode(&block)
	if err != nil {
		return nil, err
	}
	log.Println(block)
	return &block, nil
}

func getTableMaxNum(database string, table string) *int64 {
	// 连接到 MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		return nil
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal("failed to connect")
		}
	}(client, context.TODO())

	collection := client.Database(database).Collection(table)
	// 构建条件
	pipeline := bson.A{
		bson.M{"$group": bson.M{"_id": nil, "maxNumber": bson.M{"$max": "$number"}}},
		bson.M{"$project": bson.M{"_id": 0, "maxNumber": 1}},
	}
	// 执行聚合查询
	var result struct {
		MaxNumber int64 `bson:"maxNumber"`
	}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal("failed to close")
		}
	}(cursor, context.TODO())

	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&result); err != nil {
			return nil
		}
	}
	maxNum := result.MaxNumber
	log.Println("The maximum value in the current table: ", maxNum)
	return &maxNum
}
