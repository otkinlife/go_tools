# Golang HTTP 请求库使用说明

这是一个简单的 Golang HTTP 请求库，可以用于发送 GET 和 POST 请求。

## 安装

首先，你需要将这个库导入到你的项目中。你可以通过 `go get` 命令来获取这个库：

```bash
go get -u github.com/otkinlife/go_tools/http_tools
```

## 使用方法

```go
    url := "https://xxx"
	cli, err := http_tools.NewReqClient("POST", url)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	//构建请求体
	paramsJson := map[string]any{}
	paramsJson["user_name"] = "test"
	paramsJson["user_id"] = "test"
	if err := cli.SetJson(paramsJson); err != nil {
		return nil, err
	}

	// 构建请求头
	cli.SetHeaders(map[string]string{
		"Authorization": fmt.Sprintf("%s", "token"),
	})

	// 构建URL以及请求参数
	if len(req.Query) > 0 {
		cli.SetQuery(req.Query)
	}

	// 设置超时时间
	cli.SetTimeout(120 * time.Second)
	if err := cli.Send(); err != nil {
		return nil, err
	}

	code := cli.GetHttpCode()
	body, err := cli.GetBody()
	if err != nil {
		return nil, err
	}
```