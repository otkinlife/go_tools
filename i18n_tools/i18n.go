package i18n_tools

import (
	"encoding/json"
	"github.com/spf13/cast"
	"strings"
)

type I18NBuilder struct {
	langConfig map[string]map[string]string // map[key]map[lang]value
	keyConfig  map[string]map[string]string // map[lang]map[key]value
}

// NewI18NBuilder 创建一个 I18NBuilder
// langConfig: 配置map[key]map[lang]value，key 为 key，lang 为语言，value 为对应的值
func NewI18NBuilder(config map[string]map[string]string) *I18NBuilder {
	keyConfig := make(map[string]map[string]string)
	for key, langConfig := range config {
		for lang := range langConfig {
			if _, ok := keyConfig[lang]; !ok {
				keyConfig[lang] = make(map[string]string)
			}
			keyConfig[lang][key] = key
		}
	}
	return &I18NBuilder{
		langConfig: config,
		keyConfig:  keyConfig,
	}
}

// NewI18NBuilderFromJson 从 json 字符串创建一个 I18NBuilder
func NewI18NBuilderFromJson(jsonStr string) (*I18NBuilder, error) {
	config := make(map[string]map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		return nil, err
	}
	return NewI18NBuilder(config), nil
}

// GetConfig 获取配置
func (b *I18NBuilder) GetConfig() map[string]map[string]string {
	return b.langConfig
}

// Get 获取 key 对应的值
func (b *I18NBuilder) Get(key, lang string) string {
	if langConfig, ok := b.langConfig[key]; ok {
		if value, ok := langConfig[lang]; ok {
			return value
		}
	}
	return key
}

// GetWithParams 获取 key 对应的值，并替换其中的参数
func (b *I18NBuilder) GetWithParams(key, lang string, params map[string]any) string {
	value := b.Get(key, lang)
	for k, v := range params {
		value = strings.ReplaceAll(value, k, cast.ToString(v))
	}
	return value
}

// GetConfigWithKey 获取 key 对应的配置
func (b *I18NBuilder) GetConfigWithKey(key string) map[string]string {
	if langConfig, ok := b.langConfig[key]; ok {
		return langConfig
	}
	return nil
}

// GetKeyConfig 获取 lang 对应的配置
func (b *I18NBuilder) GetKeyConfig(lang string) map[string]string {
	if keyConfig, ok := b.keyConfig[lang]; ok {
		return keyConfig
	}
	return nil
}
