package abstract

type Block struct {
	Number              int64         "number"                     // 块号
	Hash                string        `json:"hash"`                // 块哈希
	ParentHash          string        `json:"parentHash"`          // 父块哈希
	Timestamp           int64         "timestamp"                  // 时间戳
	Difficulty          string        `json:"difficulty"`          // 难度
	GasLimit            int64         "gasLimit"                   // 燃料限制
	GasUsed             int64         "gasUsed"                    // 已使用的燃料量
	Miner               string        `json:"miner"`               // 挖矿矿工地址
	Nonce               int64         "nonce"                      // 随机数（仅适用于 PoW 挖矿）
	TransactionHashRoot string        `json:"transactionHashRoot"` // 交易哈希的根节点
	ReceiptRootHash     string        `json:"receiptRootHash"`     // 收据哈希的根节点
	StateRootHash       string        `json:"stateRootHash"`       // 系统状态的根节点
	ExtraData           string        `json:"extraData"`           // 额外数据
	Transactions        []Transaction `json:"transactions"`        // 交易列表
}

type Transaction struct {
	Hash             string `json:"hash"`      // 交易的哈希值
	Nonce            int64  "nonce"            // 发送方地址的交易号
	BlockHash        string `json:"blockHash"` // 交易被包含的区块的哈希值
	BlockNumber      int64  "blockNumber"      // 交易被包含的区块号
	TransactionIndex int64  "transactionIndex" // 交易在区块中的索引位置
	From             string `json:"from"`      // 发送方地址
	To               string `json:"to"`        // 接收方地址
	Value            int64  "value"            // 转账的以太币数量
	Gas              int64  "gas"              // 交易提供的gas上限
	GasPrice         int64  "gasPrice"         // 交易愿意支付的gas价格
	Input            string `json:"input"`     // 调用合约函数的输入数据
}
