# 开发指南

## 开发环境设置

### 前置要求
- Go 1.21 或更高版本
- Git

### 克隆项目
```bash
git clone https://github.com/wshlovercn/anime.git
cd anime
```

### 安装依赖
```bash
go mod download
```

## 项目结构说明

```
anime/
├── cmd/anime/              # 主程序入口
│   └── main.go            # 命令行接口
├── internal/              # 内部包（不对外暴露）
│   ├── generator/         # 动漫生成核心逻辑
│   ├── character/         # 角色管理和一致性控制
│   ├── scene/            # 场景生成
│   └── audio/            # 音频处理
├── pkg/                  # 公共包（可对外暴露）
│   ├── novel/            # 小说解析
│   └── renderer/         # 渲染引擎
└── docs/                 # 文档
```

## 代码规范

### 命名约定
- 包名：小写，简短，有意义
- 接口：名词或名词短语，如 `Generator`、`Parser`
- 函数：驼峰命名，导出函数首字母大写
- 变量：驼峰命名，避免单字母变量（除循环外）

### 注释规范
- 所有导出的类型、函数必须有注释
- 注释以类型/函数名开头
- 使用完整句子，以句号结尾

### 错误处理
- 使用 `error` 类型返回错误
- 不要忽略错误
- 提供有意义的错误信息

## 开发流程

### 1. 创建功能分支
```bash
git checkout -b feature/your-feature-name
```

### 2. 编写代码
遵循上述代码规范

### 3. 测试
```bash
go test ./...
```

### 4. 格式化代码
```bash
go fmt ./...
```

### 5. 提交代码
```bash
git add .
git commit -m "描述你的更改"
git push origin feature/your-feature-name
```

### 6. 创建 Pull Request
在 GitHub 上创建 PR，等待代码审查

## 测试指南

### 单元测试
- 测试文件命名：`xxx_test.go`
- 测试函数命名：`TestXxx`
- 使用 `testing` 包

### 示例
```go
func TestParseNovel(t *testing.T) {
    parser := NewParser()
    novel, err := parser.Parse("test.txt")
    if err != nil {
        t.Fatalf("Parse failed: %v", err)
    }
    if novel.Title == "" {
        t.Error("Expected non-empty title")
    }
}
```

## 常见问题

### 依赖管理
使用 `go mod` 管理依赖：
```bash
go get package-name
go mod tidy
```

### 调试
使用 `dlv` 或 IDE 内置调试器

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

欢迎提交 Issue 和 PR！
