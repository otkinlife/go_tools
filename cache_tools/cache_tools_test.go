package cache_tools

import (
	"fmt"
	"testing"
	"time"
)

// 测试基本的字符串缓存功能
func TestBasicStringCache(t *testing.T) {
	// 初始化缓存管理器
	manager := NewCacheManager()
	err := manager.Init(1024*1024, "") // 1MB, 无定时清理
	if err != nil {
		t.Fatalf("Failed to init cache manager: %v", err)
	}

	// 测试设置和获取字符串
	key := "test_key"
	value := "test_value"

	manager.SetString(key, value)

	result, err := manager.GetString(key)
	if err != nil {
		t.Fatalf("Failed to get string: %v", err)
	}

	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}
}

// 测试TTL功能
func TestStringCacheWithTTL(t *testing.T) {
	manager := NewCacheManager()
	err := manager.Init(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init cache manager: %v", err)
	}

	key := "ttl_key"
	value := "ttl_value"

	// 设置1秒过期的缓存
	manager.SetStringWithTTL(key, value, 1*time.Second)

	// 立即获取应该成功
	result, err := manager.GetString(key)
	if err != nil {
		t.Fatalf("Failed to get string: %v", err)
	}

	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}

	// 等待过期
	time.Sleep(2 * time.Second)

	// 再次获取应该为空
	result, err = manager.GetString(key)
	if err != nil {
		t.Fatalf("Failed to get string: %v", err)
	}

	if result != "" {
		t.Errorf("Expected empty string after expiration, got %s", result)
	}
}

// 测试JSON缓存功能
func TestJSONCache(t *testing.T) {
	type TestStruct struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	manager := NewCacheManager()
	err := manager.Init(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init cache manager: %v", err)
	}

	key := "json_key"
	original := TestStruct{
		ID:   1,
		Name: "Alice",
		Age:  25,
	}

	// 设置JSON对象
	err = manager.SetJSON(key, original)
	if err != nil {
		t.Fatalf("Failed to set JSON: %v", err)
	}

	// 获取JSON对象
	var result TestStruct
	err = manager.GetJSON(key, &result)
	if err != nil {
		t.Fatalf("Failed to get JSON: %v", err)
	}

	if result != original {
		t.Errorf("Expected %+v, got %+v", original, result)
	}
}

// 测试删除功能
func TestDelete(t *testing.T) {
	manager := NewCacheManager()
	err := manager.Init(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init cache manager: %v", err)
	}

	key := "delete_key"
	value := "delete_value"

	manager.SetString(key, value)

	// 确认数据存在
	result, err := manager.GetString(key)
	if err != nil {
		t.Fatalf("Failed to get string: %v", err)
	}
	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}

	// 删除数据
	err = manager.Delete(key)
	if err != nil {
		t.Fatalf("Failed to delete: %v", err)
	}

	// 确认数据已删除
	result, err = manager.GetString(key)
	if err != nil {
		t.Fatalf("Failed to get string after delete: %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty string after delete, got %s", result)
	}
}

// 测试清空功能
func TestClear(t *testing.T) {
	manager := NewCacheManager()
	err := manager.Init(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init cache manager: %v", err)
	}

	// 设置多个缓存项
	for i := 0; i < 5; i++ {
		manager.SetString(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
	}

	// 检查统计信息
	stats := manager.Stats()
	if stats.KeyCount == 0 {
		t.Error("Expected non-zero key count before clear")
	}

	// 清空缓存
	manager.Clear()

	// 检查统计信息
	stats = manager.Stats()
	if stats.KeyCount != 0 {
		t.Errorf("Expected zero key count after clear, got %d", stats.KeyCount)
	}
	if stats.Size != 0 {
		t.Errorf("Expected zero size after clear, got %d", stats.Size)
	}
}

// 测试默认缓存管理器
func TestDefaultCacheManager(t *testing.T) {
	// 重置默认管理器
	defaultCacheManager = nil

	// 初始化默认管理器
	err := InitDefault(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init default cache manager: %v", err)
	}

	key := "default_key"
	value := "default_value"

	// 测试字符串操作
	err = Set(key, value)
	if err != nil {
		t.Fatalf("Failed to set with default manager: %v", err)
	}

	var result string
	err = Get(key, &result)
	if err != nil {
		t.Fatalf("Failed to get with default manager: %v", err)
	}

	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}

	// 测试JSON操作
	type TestStruct struct {
		Message string `json:"message"`
	}

	jsonKey := "default_json_key"
	jsonValue := TestStruct{Message: "Hello World"}

	err = Set(jsonKey, jsonValue)
	if err != nil {
		t.Fatalf("Failed to set JSON with default manager: %v", err)
	}

	var jsonResult TestStruct
	err = Get(jsonKey, &jsonResult)
	if err != nil {
		t.Fatalf("Failed to get JSON with default manager: %v", err)
	}

	if jsonResult != jsonValue {
		t.Errorf("Expected %+v, got %+v", jsonValue, jsonResult)
	}

	// 测试统计信息
	stats, err := GetStats()
	if err != nil {
		t.Fatalf("Failed to get stats: %v", err)
	}

	if stats.KeyCount < 2 {
		t.Errorf("Expected at least 2 keys, got %d", stats.KeyCount)
	}
}

// 测试TTL功能（使用默认管理器）
func TestDefaultManagerWithTTL(t *testing.T) {
	// 重置默认管理器
	defaultCacheManager = nil

	err := InitDefault(1024*1024, "")
	if err != nil {
		t.Fatalf("Failed to init default cache manager: %v", err)
	}

	key := "ttl_default_key"
	value := "ttl_default_value"

	// 设置带TTL的缓存
	err = SetWithTTL(key, value, 1*time.Second)
	if err != nil {
		t.Fatalf("Failed to set with TTL: %v", err)
	}

	// 立即获取
	var result string
	err = Get(key, &result)
	if err != nil {
		t.Fatalf("Failed to get: %v", err)
	}

	if result != value {
		t.Errorf("Expected %s, got %s", value, result)
	}

	// 等待过期
	time.Sleep(2 * time.Second)

	// 再次获取应该为空
	result = ""
	err = Get(key, &result)
	if err != nil {
		t.Fatalf("Failed to get after expiration: %v", err)
	}

	if result != "" {
		t.Errorf("Expected empty string after expiration, got %s", result)
	}
}

// 基准测试：字符串设置
func BenchmarkSetString(b *testing.B) {
	manager := NewCacheManager()
	manager.Init(100*1024*1024, "") // 100MB

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.SetString(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
	}
}

// 基准测试：字符串获取
func BenchmarkGetString(b *testing.B) {
	manager := NewCacheManager()
	manager.Init(100*1024*1024, "")

	// 预设数据
	for i := 0; i < 1000; i++ {
		manager.SetString(fmt.Sprintf("key_%d", i), fmt.Sprintf("value_%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.GetString(fmt.Sprintf("key_%d", i%1000))
	}
}

// 基准测试：JSON设置
func BenchmarkSetJSON(b *testing.B) {
	type TestData struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Data []int  `json:"data"`
	}

	manager := NewCacheManager()
	manager.Init(100*1024*1024, "")

	testData := TestData{
		ID:   1,
		Name: "test",
		Data: []int{1, 2, 3, 4, 5},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager.SetJSON(fmt.Sprintf("json_key_%d", i), testData)
	}
}
