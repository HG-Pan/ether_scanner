package infura

import (
	"fmt"
	"strconv"
)

func parseHexInt64(hexStr string) (int64, error) {
	value, err := strconv.ParseInt(hexStr[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse hexInt64: %v", err)
	}
	return value, nil
}
