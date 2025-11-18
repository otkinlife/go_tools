# Cache Tools

一个基于Go语言的高性能缓存工具包，支持LRU淘汰策略、过期时间、JSON序列化等功能。

## 功能特性

- **内存缓存**: 基于 sync.Map 实现的并发安全缓存
- **LRU淘汰**: 智能的最近最少使用淘汰算法
- **TTL支持**: 支持设置缓存过期时间
- **JSON序列化**: 自动处理复杂对象的JSON序列化/反序列化
- **大小限制**: 内置缓存大小监控和自动清理
- **定时清理**: 支持定时清理过期缓存
- **统计信息**: 提供缓存使用统计

## 快速开始

### 基础使用

```go
package main

import (
    "fmt"
    "time"
    "github.com/otkinlife/go_tools/cache_tools"
)

func main() {
    // 初始化默认缓存管理器
    // maxSize: 10MB, clearTime: 每天凌晨3点清理
    err := cache_tools.InitDefault(10*1024*1024, "03:00:00")
    if err != nil {
        panic(err)
    }

    // 设置字符串缓存
    cache_tools.Set("user:123", "John Doe")

    // 获取字符串缓存
    var name string
    err = cache_tools.Get("user:123", &name)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Name:", name) // 输出: Name: John Doe

    // 设置带过期时间的缓存
    cache_tools.SetWithTTL("session:abc", "active", 30*time.Minute)
}
```

### 使用结构体

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    cache_tools.InitDefault(10*1024*1024, "03:00:00")

    user := User{ID: 1, Name: "Alice", Age: 25}

    // 设置结构体缓存(自动JSON序列化)
    err := cache_tools.Set("user:1", user)
    if err != nil {
        panic(err)
    }

    // 获取结构体缓存(自动JSON反序列化)
    var cachedUser User
    err = cache_tools.Get("user:1", &cachedUser)
    if err != nil {
        panic(err)
    }

    fmt.Printf("User: %+v\\n", cachedUser)
}
```

### 高级用法：自定义缓存管理器

```go
func main() {
    // 创建自定义缓存管理器
    manager := cache_tools.NewCacheManager()

    // 初始化: 最大5MB，每晚12点清理
    err := manager.Init(5*1024*1024, "00:00:00")
    if err != nil {
        panic(err)
    }

    // 使用管理器实例
    manager.SetString("key1", "value1")
    manager.SetStringWithTTL("key2", "value2", 1*time.Hour)

    value, err := manager.GetString("key1")
    if err != nil {
        panic(err)
    }
    fmt.Println("Value:", value)

    // 获取缓存统计
    stats := manager.Stats()
    fmt.Printf("Cache size: %d bytes, Keys: %d\\n", stats.Size, stats.KeyCount)
}
```

## API 文档

### 全局函数(使用默认管理器)

#### 初始化
- `InitDefault(maxSize int64, clearTime string) error` - 初始化默认缓存管理器

#### 基本操作
- `Set(key string, value interface{}) error` - 设置缓存
- `SetWithTTL(key string, value interface{}, ttl time.Duration) error` - 设置带过期时间的缓存
- `Get(key string, result interface{}) error` - 获取缓存
- `Delete(key string) error` - 删除缓存
- `Clear() error` - 清空所有缓存
- `GetStats() (CacheStats, error)` - 获取统计信息

### CacheManager 实例方法

#### 创建和初始化
- `NewCacheManager() *CacheManager` - 创建新的缓存管理器
- `Init(maxSize int64, clearTime string) error` - 初始化管理器

#### 字符串操作
- `SetString(key, value string)` - 设置字符串缓存
- `SetStringWithTTL(key, value string, ttl time.Duration)` - 设置带TTL的字符串缓存
- `GetString(key string) (string, error)` - 获取字符串缓存

#### JSON操作
- `SetJSON(key string, data interface{}) error` - 设置JSON对象缓存
- `SetJSONWithTTL(key string, data interface{}, ttl time.Duration) error` - 设置带TTL的JSON缓存
- `GetJSON(key string, result interface{}) error` - 获取JSON对象缓存

#### 管理操作
- `Delete(key string) error` - 删除指定缓存
- `Clear()` - 清空所有缓存
- `Stats() CacheStats` - 获取统计信息

### 数据结构

```go
type CacheStats struct {
    Size     int64 // 缓存总大小(字节)
    KeyCount int   // 缓存key数量
}
```

## 配置说明

### maxSize
缓存最大字节数，超过此限制时会自动触发LRU淘汰策略，清理约1/3的最少使用缓存。

### clearTime
定时清理时间，格式为 "HH:MM:SS"，每天在指定时间清理过期缓存。设置为空字符串则禁用定时清理。

## 注意事项

1. 缓存key会自动进行MD5哈希处理，支持任意字符串作为key
2. JSON序列化/反序列化使用标准库，确保结构体字段可导出
3. 过期时间检查是异步进行的，可能存在短暂延迟
4. LRU淘汰策略基于访问频率和最后使用时间综合计算

## 性能特点

- 基于 sync.Map 实现，并发读写性能优秀
- 内存占用可控，支持大小限制和自动清理
- LRU算法确保热点数据常驻内存
- 异步过期检查，不影响正常读写性能

## 示例项目

查看 `cache_test.go` 文件获取更多使用示例。