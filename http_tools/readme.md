# Golang HTTP 请求库使用说明

这是一个简单的 Golang HTTP 请求库，可以用于发送 GET 和 POST 请求。

## 安装

首先，你需要将这个库导入到你的项目中。你可以通过 `go get` 命令来获取这个库：

```bash
go get -u github.com/otkinlife/go_tools/http
```

请将 `yourusername` 替换为你的 GitHub 用户名。

## 使用方法

首先，你需要创建一个 `ReqClient` 对象：

```go
req := http.NewReq()
```

然后，你可以设置请求头和请求参数：

```go
req.SetHeaders(map[string]string{
	"Content-Type": "application/json",
})

req.SetQuery(map[string]string{
	"key": "value",
})
```

你也可以设置表单或 JSON 请求体：

```go
req.SetForm(map[string]string{
	"username": "test",
	"password": "test",
})

req.SetJson(map[string]interface{}{
	"username": "test",
	"password": "test",
})
```

然后，你可以使用 `Get` 或 `Post` 方法发送请求：

```go
err := req.Get("http://example.com")
if err != nil {
	// handle error
}

err := req.Post("http://example.com")
if err != nil {
	// handle error
}
```

你还可以使用 `GetRetry` 或 `PostRetry` 方法发送请求并进行重试：

```go
err := req.GetRetry("http://example.com", 3)
if err != nil {
	// handle error
}

err := req.PostRetry("http://example.com", 3)
if err != nil {
	// handle error
}
```
你可以使用 `UploadFile` 方法上传文件：

```go
err := req.UploadFile("fileField", "/path/to/file.txt")
if err != nil {
    // handle error
}

err = req.Post("http://example.com/upload")
if err != nil {
    // handle error
}
```

你可以使用 `GetHttpCode` 方法获取 HTTP 状态码：

```go
code := req.GetHttpCode()
```

你可以使用 `GetBody` 方法获取响应体：

```go
body, err := req.GetBody()
if err != nil {
	// handle error
}
```

你还可以使用 `LoadBody` 方法将响应体加载到一个变量中：

```go
var data map[string]interface{}
err := req.LoadBody(&data)
if err != nil {
	// handle error
}
```