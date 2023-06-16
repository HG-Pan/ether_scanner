package infura

import (
	"fmt"
	"math/big"
	"strings"
)

func parseHexInt64(hexStr string) (int64, error) {
	// 去除可能的前缀 "0x"
	hexStr = strings.TrimPrefix(hexStr, "0x")
	// 使用big.Int来解析十六进制字符串
	n := new(big.Int)
	_, success := n.SetString(hexStr, 16)
	if !success {
		return 0, fmt.Errorf("failed to parse hexInt64: invalid format")
	}
	// 检查是否超出int64范围
	if n.IsInt64() {
		return n.Int64(), nil
	}
	// 超出范围时返回-1
	return -1, nil
}
