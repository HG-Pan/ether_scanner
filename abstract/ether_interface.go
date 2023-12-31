package abstract

type Client interface {
	GetBlockByNumber(blockNumber int64) (*Block, error)
	GetLatestNum() (int64, error)
}

type EthClient struct {
	Client Client
}

func (ec *EthClient) GetBlockByNumber(blockNumber int64) (*Block, error) {
	return ec.Client.GetBlockByNumber(blockNumber)
}

func (ec *EthClient) GetLatestNum() (int64, error) {
	return ec.Client.GetLatestNum()
}
