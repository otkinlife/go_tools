package http_tools

import (
	"os"
	"strings"
	"testing"
)

func TestSendGet(t *testing.T) {
	t.Log("TestSendGet")
	client, _ := NewReqClient("GET", "https://httpbin.org/json")
	err := client.Send()
	if err != nil {
		t.Error(err)
		return
	}
	body, err := client.GetBody()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
}

func TestSetFormAndFile(t *testing.T) {
	t.Log("TestSetFormAndFile")

	// 创建临时测试文件
	tmpFile, err := os.CreateTemp("", "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// 写入测试内容
	_, err = tmpFile.WriteString("test file content")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// 测试先设置表单再设置文件
	client1, err := NewReqClient("POST", "https://httpbin.org/post")
	if err != nil {
		t.Fatal(err)
	}

	client1.SetForm(map[string]string{
		"field1": "value1",
		"field2": "value2",
	})

	err = client1.SetFile("file", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// 验证请求体将包含表单和文件数据
	if len(client1.formFields) != 2 {
		t.Errorf("Expected 2 form fields, got %d", len(client1.formFields))
	}

	if len(client1.files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(client1.files))
	}

	// 测试先设置文件再设置表单
	client2, err := NewReqClient("POST", "https://httpbin.org/post")
	if err != nil {
		t.Fatal(err)
	}

	err = client2.SetFile("file", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	client2.SetForm(map[string]string{
		"field1": "value1",
		"field2": "value2",
	})

	// 验证请求体将包含表单和文件数据
	if len(client2.formFields) != 2 {
		t.Errorf("Expected 2 form fields, got %d", len(client2.formFields))
	}

	if len(client2.files) != 1 {
		t.Errorf("Expected 1 file, got %d", len(client2.files))
	}

	t.Log("Both orders work correctly")
}

func TestCurlWithFiles(t *testing.T) {
	t.Log("TestCurlWithFiles")

	// 创建临时测试文件
	tmpFile, err := os.CreateTemp("", "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// 写入测试内容
	_, err = tmpFile.WriteString("test file content")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// 创建客户端并设置表单和文件
	client, err := NewReqClient("POST", "https://httpbin.org/post")
	if err != nil {
		t.Fatal(err)
	}

	client.SetForm(map[string]string{
		"field1": "value1",
		"field2": "value2",
	})

	err = client.SetFile("file", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// 测试curl生成
	curlCmd, err := client.ConvertToCurlWithFiles()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Generated curl command: %s", curlCmd)

	// 验证curl命令包含正确的元素
	if !strings.Contains(curlCmd, "curl") {
		t.Error("Curl command should contain 'curl'")
	}

	if !strings.Contains(curlCmd, "-X POST") {
		t.Error("Curl command should contain '-X POST'")
	}

	if !strings.Contains(curlCmd, "https://httpbin.org/post") {
		t.Error("Curl command should contain the URL")
	}

	if !strings.Contains(curlCmd, "-F 'field1=value1'") {
		t.Error("Curl command should contain form field")
	}

	if !strings.Contains(curlCmd, "-F 'field2=value2'") {
		t.Error("Curl command should contain form field")
	}

	if !strings.Contains(curlCmd, "-F 'file=@"+tmpFile.Name()+"'") {
		t.Error("Curl command should contain file with @filepath format")
	}

	// 验证不包含文件内容
	if strings.Contains(curlCmd, "test file content") {
		t.Error("Curl command should NOT contain file content")
	}

	t.Log("Curl command generation test passed")
}

func TestPrintCurlWithFiles(t *testing.T) {
	t.Log("TestPrintCurlWithFiles - this should print a curl command")

	// 创建临时测试文件
	tmpFile, err := os.CreateTemp("", "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	// 写入测试内容
	_, err = tmpFile.WriteString("test file content")
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// 创建客户端并启用curl打印
	client, err := NewReqClient("POST", "https://httpbin.org/post")
	if err != nil {
		t.Fatal(err)
	}

	// 启用curl打印
	client.SetIsPrintCurl(true)

	client.SetForm(map[string]string{
		"field1": "value1",
		"field2": "value2",
	})

	err = client.SetFile("file", tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	t.Log("About to send request - curl command should be printed above this:")
	// 这会触发curl命令的打印
	err = client.Send()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Request sent successfully with curl printed in @filepath format")
}
