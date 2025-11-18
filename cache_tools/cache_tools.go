package cache_tools

import (
	"fmt"
	"time"
)

// CacheManager 缓存管理器统一接口
type CacheManager struct {
	cache   *Cache
	watcher *Watcher
}

// NewCacheManager 创建缓存管理器
func NewCacheManager() *CacheManager {
	return &CacheManager{
		cache: NewCache(),
	}
}

// InitCacheManager 初始化缓存管理器
// maxSize: 最大缓存大小(字节)
// clearTime: 定时清理时间，格式: "15:04:05"
func (cm *CacheManager) Init(maxSize int64, clearTime string) error {
	cm.watcher = NewWatcher(cm.cache)
	cm.watcher.SetMaxSize(maxSize)

	if clearTime != "" {
		t, err := time.Parse("15:04:05", clearTime)
		if err != nil {
			return err
		}
		cm.watcher.SetClearPlanTime(t.Hour(), t.Minute(), t.Second())
	}

	// 初始化全局变量
	GlobalCache = cm.cache
	GlobalWatcher = cm.watcher
	LimitCh = make(chan int, 1)

	// 启动监控协程
	go cm.watcher.WatchClear()
	go cm.watcher.WatchLimit()
	go cm.watcher.WatchExpiration()

	return nil
}

// SetString 设置字符串缓存
func (cm *CacheManager) SetString(key, value string) {
	cm.cache.SetString(key, value)
}

// SetStringWithTTL 设置带过期时间的字符串缓存
func (cm *CacheManager) SetStringWithTTL(key, value string, ttl time.Duration) {
	cm.cache.SetStringWithExpiration(key, value, ttl)
}

// GetString 获取字符串缓存
func (cm *CacheManager) GetString(key string) (string, error) {
	return cm.cache.GetString(key)
}

// SetJSON 设置JSON对象缓存
func (cm *CacheManager) SetJSON(key string, data interface{}) error {
	return cm.cache.SetDataWithJson(key, data)
}

// SetJSONWithTTL 设置带过期时间的JSON对象缓存
func (cm *CacheManager) SetJSONWithTTL(key string, data interface{}, ttl time.Duration) error {
	return cm.cache.SetDataWithJsonWithExpiration(key, data, ttl)
}

// GetJSON 获取JSON对象缓存
func (cm *CacheManager) GetJSON(key string, result interface{}) error {
	return cm.cache.LoadDataFromJson(key, result)
}

// Delete 删除缓存
func (cm *CacheManager) Delete(key string) error {
	return cm.cache.Delete(key)
}

// Clear 清空所有缓存
func (cm *CacheManager) Clear() {
	cm.cache.Clear()
}

// Stats 获取缓存统计信息
func (cm *CacheManager) Stats() CacheStats {
	return CacheStats{
		Size:     cm.cache.size,
		KeyCount: len(cm.cache.keyList),
	}
}

// CacheStats 缓存统计信息
type CacheStats struct {
	Size     int64 // 缓存总大小(字节)
	KeyCount int   // 缓存key数量
}

// Global functions for backward compatibility
var defaultCacheManager *CacheManager

// InitDefault 初始化默认缓存管理器
func InitDefault(maxSize int64, clearTime string) error {
	defaultCacheManager = NewCacheManager()
	return defaultCacheManager.Init(maxSize, clearTime)
}

// Set 设置缓存(使用默认管理器)
func Set(key string, value interface{}) error {
	if defaultCacheManager == nil {
		return fmt.Errorf("cache manager not initialized, call InitDefault first")
	}

	switch v := value.(type) {
	case string:
		defaultCacheManager.SetString(key, v)
		return nil
	default:
		return defaultCacheManager.SetJSON(key, v)
	}
}

// SetWithTTL 设置带过期时间的缓存(使用默认管理器)
func SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	if defaultCacheManager == nil {
		return fmt.Errorf("cache manager not initialized, call InitDefault first")
	}

	switch v := value.(type) {
	case string:
		defaultCacheManager.SetStringWithTTL(key, v, ttl)
		return nil
	default:
		return defaultCacheManager.SetJSONWithTTL(key, v, ttl)
	}
}

// Get 获取缓存(使用默认管理器)
func Get(key string, result interface{}) error {
	if defaultCacheManager == nil {
		return fmt.Errorf("cache manager not initialized, call InitDefault first")
	}

	switch result.(type) {
	case *string:
		str, err := defaultCacheManager.GetString(key)
		if err != nil {
			return err
		}
		*(result.(*string)) = str
		return nil
	default:
		return defaultCacheManager.GetJSON(key, result)
	}
}

// Delete 删除缓存(使用默认管理器)
func Delete(key string) error {
	if defaultCacheManager == nil {
		return fmt.Errorf("cache manager not initialized, call InitDefault first")
	}
	return defaultCacheManager.Delete(key)
}

// Clear 清空缓存(使用默认管理器)
func Clear() error {
	if defaultCacheManager == nil {
		return fmt.Errorf("cache manager not initialized, call InitDefault first")
	}
	defaultCacheManager.Clear()
	return nil
}

// GetStats 获取缓存统计信息(使用默认管理器)
func GetStats() (CacheStats, error) {
	if defaultCacheManager == nil {
		return CacheStats{}, fmt.Errorf("cache manager not initialized, call InitDefault first")
	}
	return defaultCacheManager.Stats(), nil
}
