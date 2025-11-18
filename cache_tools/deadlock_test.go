package cache_tools

import (
	"testing"
	"time"
)

// TestDeadlockWithoutInit 测试未调用Init时的死锁问题
func TestDeadlockWithoutInit(t *testing.T) {
	// 直接使用 NewCache() 创建缓存，没有调用 Init()
	cache := NewCache()

	// 创建一个超时通道来检测死锁
	done := make(chan bool, 1)

	go func() {
		// 这里会死锁，因为 LimitCh 是 nil
		cache.SetString("test", "value")
		done <- true
	}()

	// 等待2秒，如果没有完成说明发生了死锁
	select {
	case <-done:
		t.Log("操作成功完成")
	case <-time.After(2 * time.Second):
		t.Fatal("检测到死锁：向 nil channel 发送数据导致永久阻塞")
	}
}

// TestDeadlockWithExpiration 测试带过期时间的方法
func TestDeadlockWithExpiration(t *testing.T) {
	cache := NewCache()

	done := make(chan bool, 1)

	go func() {
		// 这里也会死锁
		cache.SetStringWithExpiration("test", "value", 1*time.Second)
		done <- true
	}()

	select {
	case <-done:
		t.Log("操作成功完成")
	case <-time.After(2 * time.Second):
		t.Fatal("检测到死锁：SetStringWithExpiration 向 nil channel 发送数据导致永久阻塞")
	}
}

// TestNormalWithInit 测试调用Init后的正常情况
func TestNormalWithInit(t *testing.T) {
	// 调用 Init 初始化
	err := Init(Config{
		MaxSize:  1024,
		PlanTime: "23:59:00",
	})
	if err != nil {
		t.Fatalf("Init 失败: %v", err)
	}

	done := make(chan bool, 1)

	go func() {
		GlobalCache.SetString("test", "value")
		done <- true
	}()

	select {
	case <-done:
		t.Log("Init 后操作成功完成")
	case <-time.After(2 * time.Second):
		t.Fatal("即使调用了 Init 也发生了死锁")
	}
}
