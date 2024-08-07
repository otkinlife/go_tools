

# JWT Tools

`jwt_tools` 是一个用于生成和验证 JWT 令牌的 Go 包。它提供了简单易用的接口来创建和管理 JWT 令牌。

## 安装

使用 `go get` 命令安装此包：

```bash
go get github.com/yourusername/jwt_tools
```

## 导入

在你的代码中导入此包：

```go
import "github.com/yourusername/jwt_tools"
```

## 使用

### 创建 TokenBuilder

首先，创建一个 `JwtConfig` 实例并初始化 `TokenBuilder`：

```go
config := jwt_tools.JwtConfig{
    SecretKey:     "your-secret-key",
    SigningMethod: jwt.SigningMethodHS256,
    ExpireTime:    time.Hour * 24, // 令牌过期时间为24小时
}

tokenBuilder := jwt_tools.NewTokenBuilder(config)
```

### 设置元数据

设置令牌的元数据：

```go
meta := map[string]any{
    "user_id": 12345,
    "role":    "admin",
}

tokenBuilder.SetMeta(meta)
```

### 生成令牌

生成 JWT 令牌：

```go
token, err := tokenBuilder.GenerateToken()
if err != nil {
    fmt.Println("生成令牌失败:", err)
    return
}

fmt.Println("生成的令牌:", token)
```

### 验证令牌

验证 JWT 令牌：

```go
tokenBuilder.SetToken(token)
err = tokenBuilder.VerifyToken()
if err != nil {
    fmt.Println("验证令牌失败:", err)
    return
}

fmt.Println("令牌验证成功, 元数据:", tokenBuilder.GetMeta())
```

### 注册验证函数

你可以注册一个自定义的验证函数来进一步验证元数据：

```go
validateFunc := func(meta map[string]any) error {
    if meta["role"] != "admin" {
        return fmt.Errorf("用户角色无效")
    }
    return nil
}

tokenBuilder.RegisterValidateFunc(validateFunc)
```

## 完整示例

以下是一个完整的示例，展示如何生成和验证 JWT 令牌：

```go
package main

import (
    "fmt"
    "github.com/yourusername/jwt_tools"
    "github.com/golang-jwt/jwt/v4"
    "time"
)

func main() {
    config := jwt_tools.JwtConfig{
        SecretKey:     "your-secret-key",
        SigningMethod: jwt.SigningMethodHS256,
        ExpireTime:    time.Hour * 24,
    }

    tokenBuilder := jwt_tools.NewTokenBuilder(config)

    meta := map[string]any{
        "user_id": 12345,
        "role":    "admin",
    }

    tokenBuilder.SetMeta(meta)

    token, err := tokenBuilder.GenerateToken()
    if err != nil {
        fmt.Println("生成令牌失败:", err)
        return
    }

    fmt.Println("生成的令牌:", token)

    tokenBuilder.SetToken(token)
    err = tokenBuilder.VerifyToken()
    if err != nil {
        fmt.Println("验证令牌失败:", err)
        return
    }

    fmt.Println("令牌验证成功, 元数据:", tokenBuilder.GetMeta())
}
```