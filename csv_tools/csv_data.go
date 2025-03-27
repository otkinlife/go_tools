package csv_tools

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CSVData 结构体用于处理CSV数据
type CSVData struct {
	headers     map[string]int // 列名到列索引的映射
	headersLine []string       // CSV表头
	data        [][]string     // CSV数据内容
	split       string         // 分隔符
	err         error          // 错误信息
}

// LoadFromReader 从io.Reader加载CSV数据
func (c *CSVData) LoadFromReader(r io.Reader) *CSVData {
	reader := csv.NewReader(r)
	if c.split != "" {
		reader.Comma = rune(c.split[0])
	}
	data, err := reader.ReadAll()
	if err != nil {
		c.err = err
		return c
	}

	if len(data) > 0 {
		c.SetHeaders(data[0])
		c.data = data[1:]
	} else {
		c.data = [][]string{}
	}
	return c
}

// LoadFromLines 从字符串二维数组加载CSV数据
func (c *CSVData) LoadFromLines(lines [][]string) *CSVData {
	if len(lines) == 0 {
		c.err = io.EOF
		return c
	}
	c.SetHeaders(lines[0])
	c.SetData(lines[1:])
	return c
}

// GetData 获取CSV数据内容
func (c *CSVData) GetData() [][]string {
	return c.data
}

// SetData 加载CSV数据内容
func (c *CSVData) SetData(data [][]string) {
	c.data = data
}

// AppendData 追加CSV数据
func (c *CSVData) AppendData(data []string) {
	c.data = append(c.data, data)
}

// SetHeaders 加载CSV表头
func (c *CSVData) SetHeaders(headers []string) {
	c.headersLine = headers
	c.headers = make(map[string]int, len(headers))
	for i, header := range headers {
		c.headers[header] = i
	}
}

// GetHeaders 获取CSV表头映射
func (c *CSVData) GetHeaders() map[string]int {
	return c.headers
}

// GetHeaderIndex 获取指定列名的索引
func (c *CSVData) GetHeaderIndex(key string) int {
	if index, ok := c.headers[key]; ok {
		return index
	}
	return -1
}

// GetLineValue 获取指定行和列的值
func (c *CSVData) GetLineValue(key string, line int) string {
	if index, ok := c.headers[key]; ok {
		if line >= 0 && line < len(c.data) && index < len(c.data[line]) {
			return c.data[line][index]
		}
	}
	return ""
}

// SetSplit 设置CSV分隔符
func (c *CSVData) SetSplit(split string) *CSVData {
	if split == "" {
		c.split = ","
	} else {
		c.split = split
	}
	return c
}

// SaveToFile 将CSV数据保存到文件
func (c *CSVData) SaveToFile(filePath string) {
	if c.err != nil {
		return
	}
	// 创建文件，如果没有目录则创建目录
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		err = fmt.Errorf("创建目录失败: %v", err)
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		err = fmt.Errorf("创建文件失败: %v", err)
		return
	}
	defer file.Close()

	// 写入UTF-8 BOM头
	_, err = file.Write([]byte{0xEF, 0xBB, 0xBF})
	if err != nil {
		c.err = fmt.Errorf("写入BOM头失败: %v", err)
		return
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	if err := writer.Write(c.headersLine); err != nil {
		c.err = fmt.Errorf("写入表头失败: %v", err)
	}

	// 写入数据行
	for _, row := range c.data {
		if err := writer.Write(row); err != nil {
			c.err = fmt.Errorf("写入数据行失败: %v", err)
		}
	}
}

// GetError 获取处理过程中的错误
func (c *CSVData) GetError() error {
	return c.err
}

// NewCSVData 创建一个新的CSVData实例
func NewCSVData() *CSVData {
	return &CSVData{
		headers:     make(map[string]int),
		headersLine: []string{},
		data:        [][]string{},
		split:       ",",
	}
}
