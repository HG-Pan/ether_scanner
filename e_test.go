package main

import (
	"sync/atomic"
	"testing"
)

func TestAtomicAddInt64(t *testing.T) {
	var currentBlockNum int64 = 0
	blockRange := int64(10)

	// 调用被测试的代码
	blockNum := atomic.AddInt64(&currentBlockNum, blockRange)

	// 验证结果是否符合预期
	expectedBlockNum := blockRange
	if blockNum != expectedBlockNum {
		t.Errorf("Expected blockNum to be %d, but got %d", expectedBlockNum, blockNum)
	}
}
