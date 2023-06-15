package main

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestGetNextBlockNum(t *testing.T) {
	// 设置初始的 currentBlockNum 和 blockRange
	var currentBlockNum int64 = 100
	blockRange := int64(1)

	// 循环测试多次调用
	for i := 0; i < 100; i++ {
		// 调用原子操作获取下一个区块号
		blockNum := atomic.AddInt64(&currentBlockNum, 1)

		// 验证获取的区块号是否符合预期
		expectedBlockNum := currentBlockNum - blockRange
		fmt.Printf("%d: %d\n:%d\n:%d\n", i, blockNum, currentBlockNum, expectedBlockNum)
		if blockNum != expectedBlockNum {
			t.Errorf("Expected blockNum to be %d, but got %d", expectedBlockNum, blockNum)
		}
	}
}
