package jwt_tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// TokenBuilder 用于构建和管理 JWT 令牌
type TokenBuilder struct {
	Meta        map[string]any // 元数据
	Config      *JwtConfig     // 配置
	ValidMethod ValidMethod    // 验证方法
	TokenStr    string         // 令牌字符串
}

// JwtConfig 包含 JWT 相关的配置
type JwtConfig struct {
	SecretKey     string            // 密钥
	SigningMethod jwt.SigningMethod // 签名方法
	ExpireTime    time.Duration     // 过期时间
}

// ValidMethod 是一个函数类型，用于验证元数据
type ValidMethod func(meta map[string]any) error

// NewTokenBuilder 创建一个新的 TokenBuilder 实例
func NewTokenBuilder(config JwtConfig) *TokenBuilder {
	return &TokenBuilder{
		Config: &config,
	}
}

// SetMeta 设置令牌的元数据
func (t *TokenBuilder) SetMeta(meta map[string]any) *TokenBuilder {
	t.Meta = meta
	return t
}

// GetMeta 获取令牌的元数据
func (t *TokenBuilder) GetMeta() map[string]any {
	return t.Meta
}

// SetToken 设置令牌字符串
func (t *TokenBuilder) SetToken(token string) *TokenBuilder {
	t.TokenStr = token
	return t
}

// GetToken 获取令牌字符串
func (t *TokenBuilder) GetToken() string {
	return t.TokenStr
}

// RegisterValidateFunc 注册一个验证函数
func (t *TokenBuilder) RegisterValidateFunc(f ValidMethod) *TokenBuilder {
	t.ValidMethod = f
	return t
}

// GenerateToken 生成一个 JWT 令牌
func (t *TokenBuilder) GenerateToken() (string, error) {
	if len(t.Meta) == 0 {
		return "", fmt.Errorf("meta is empty")
	}

	// 设置加密算法，如果未指定则使用默认的 HS256
	signingMethod := t.Config.SigningMethod
	if signingMethod == nil {
		signingMethod = jwt.SigningMethodHS256
	}

	token := jwt.New(signingMethod)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(t.Config.ExpireTime).Unix()

	// 将元数据添加到声明中
	for k, v := range t.Meta {
		claims[k] = v
	}

	// 生成签名字符串
	tokenString, err := token.SignedString([]byte(t.Config.SecretKey))
	if err != nil {
		return "", fmt.Errorf("token generate failed: %w", err)
	}

	t.TokenStr = tokenString
	return tokenString, nil
}

// VerifyToken 验证 JWT 令牌并更新元数据
func (t *TokenBuilder) VerifyToken() error {
	if t.TokenStr == "" {
		return fmt.Errorf("token is empty")
	}

	token, err := jwt.Parse(t.TokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Config.SecretKey), nil
	})
	if err != nil {
		return fmt.Errorf("token parsed error: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	t.Meta = claims

	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("token is invalid")
	}

	expireTime := time.Unix(int64(exp), 0)
	if expireTime.Before(time.Now()) {
		return fmt.Errorf("token is expired")
	}

	if t.ValidMethod != nil {
		if err := t.ValidMethod(claims); err != nil {
			return fmt.Errorf("token is invalid: %w", err)
		}
	}

	return nil
}
