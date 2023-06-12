package abstract

type BlockHex struct {
	Number              string           `json:"number"`              // 块号
	Hash                string           `json:"hash"`                // 块哈希
	ParentHash          string           `json:"parentHash"`          // 父块哈希
	Timestamp           string           `json:"timestamp"`           // 时间戳
	Difficulty          string           `json:"difficulty"`          // 难度
	GasLimit            string           `json:"gasLimit"`            // 燃料限制
	GasUsed             string           `json:"gasUsed"`             // 已使用的燃料量
	Miner               string           `json:"miner"`               // 挖矿矿工地址
	Nonce               string           `json:"nonce"`               // 随机数（仅适用于 PoW 挖矿）
	TransactionHashRoot string           `json:"transactionHashRoot"` // 交易哈希的根节点
	ReceiptRootHash     string           `json:"receiptRootHash"`     // 收据哈希的根节点
	StateRootHash       string           `json:"stateRootHash"`       // 系统状态的根节点
	ExtraData           string           `json:"extraData"`           // 额外数据
	Transactions        []TransactionHex `json:"transactions"`        // 交易列表
}

type TransactionHex struct {
	Hash             string `json:"hash"`             // 交易的哈希值
	Nonce            string `json:"nonce"`            // 发送方地址的交易号
	BlockHash        string `json:"blockHash"`        // 交易被包含的区块的哈希值
	BlockNumber      string `json:"blockNumber"`      // 交易被包含的区块号
	TransactionIndex string `json:"transactionIndex"` // 交易在区块中的索引位置
	From             string `json:"from"`             // 发送方地址
	To               string `json:"to"`               // 接收方地址
	Value            string `json:"value"`            // 转账的以太币数量
	Gas              string `json:"gas"`              // 交易提供的gas上限
	GasPrice         string `json:"gasPrice"`         // 交易愿意支付的gas价格
	Input            string `json:"input"`            // 调用合约函数的输入数据
}
