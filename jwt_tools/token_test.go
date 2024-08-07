package jwt_tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestTokenBuilder_GenerateToken(t *testing.T) {
	tb := NewTokenBuilder(JwtConfig{
		SecretKey:     "test",
		SigningMethod: jwt.SigningMethodHS256,
		ExpireTime:    3600 * time.Second,
	})
	token, err := tb.SetMeta(map[string]any{
		"uid":  1,
		"name": "test",
	}).GenerateToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	tb.RegisterValidateFunc(func(meta map[string]any) error {
		if meta["uid"] == nil {
			return fmt.Errorf("uid is nil")
		}
		t.Log("uid is", meta["uid"])
		t.Log("name is", meta["name"])
		return nil
	})

	err = tb.VerifyToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("token is valid")
	fmt.Println(tb.GetMeta())
	return
}
