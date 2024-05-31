当然可以！以下是一个详细的README模板，你可以根据需要进行修改和扩展：

---

# goCliGPT

goCliGPT 是一个基于Go语言开发的命令行聊天助手，目前暂时支持与Qwen和ChatGPT两种API的交互，并可以保存和加载对话历史记录。

## 功能特性

- 支持与在线的LLM API进行交互对话
- 保存和加载对话历史记录
- 根据时间戳将对话历史记录保存为不同的JSON文件
- 交互模式下支持持续对话

## 安装

1. 克隆项目仓库：

    ```sh
    git clone https://github.com/yourusername/goCliGPT.git
    cd goCliGPT
    ```

2. 初始化Go模块：

    ```sh
    go mod init goCliGPT
    ```

## 配置

在项目根目录下存在一个 `config.json` 文件，你可以按照以下格式填写配置：

```json
{
    "ChatGPT": {
        "api_url": "https://api.openai.com/v1/completions",
        "api_key": "your_chatgpt_api_key_here",
        "model_name": "text-davinci-003"
    },

    "Qwen": {
        "api_url": "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
        "api_key": "your_qwen_api_key_here",
        "model_name": "qwen-turbo"
    }
}
```

请确保使用有效的API密钥和API URL。

## 使用方法

### 运行项目

在项目根目录下可以直接运行以下命令启动程序：

```sh
go run main.go <api_name>
```

例如，使用Qwen API：

```sh
go run main.go qwen
```

或使用ChatGPT API：

```sh
go run main.go chatgpt
```



### 编译项目

也可以通过编译项目的方式运行
在项目根目录下运行以下命令生成可执行文件 `goCliGPT`：

```sh
go build -o goCliGPT main.go
```

然后在命令行中运行以下命令启动程序：

```sh
goCliGPT <api_name>
```

例如，使用Qwen API：

```sh
goCliGPT qwen
```

或使用ChatGPT API：

```sh
goCliGPT chatgpt
```


### 交互模式

启动程序后，系统会提示你选择加载一个已有的对话历史记录，或新建一个对话：

```sh
Available conversations:
1: history_20240531_110237.json
2: history_20240531_110352.json
Enter the number of the conversation to load or 'new' to start a new conversation:
```

输入 `new` 新建一个对话，或输入对应的数字加载已有对话。

在交互模式下，你可以输入你的问题，程序会返回响应。输入 `new` 可以新建一个对话，输入 `exit` 可以退出并保存当前对话。

## 目录结构

```
goCliGPT/
├── api/
│   ├── chatgpt.go
│   └── qwen.go
├── config/
│   └── config.go
├── history/
│   └── history.go
├── main.go
├── config.json
```

## 贡献

欢迎提交issue和pull request来帮助改进这个项目。

## 许可

此项目遵循MIT许可。

---

你可以根据实际情况对这份README进行调整和完善。如果有任何特定的细节需要补充或修改，请告诉我。
