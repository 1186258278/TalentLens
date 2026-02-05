# 快速开始指南

本文档帮助开发者快速搭建 GoResumeReview 开发环境。

## 环境要求

| 工具 | 最低版本 | 推荐版本 | 说明 |
|------|----------|----------|------|
| Go | 1.21 | 1.22 | 后端开发 |
| Node.js | 18 | 20 LTS | 前端构建 |
| Wails CLI | 2.7 | 最新 | 桌面应用框架 |
| Git | 2.0 | 最新 | 版本控制 |
| Windows | 10/11 | 11 | 目标平台 |

## 安装步骤

### 1. 安装 Go

```bash
# 下载地址：https://go.dev/dl/

# 验证安装
go version

# 输出应该类似：go version go1.22.0 windows/amd64
```

### 2. 安装 Node.js

```bash
# 下载地址：https://nodejs.org/

# 验证安装
node --version
npm --version

# 输出应该类似：
# v20.10.0
# 10.7.0
```

### 3. 安装 Wails CLI

```bash
# 使用 Go 安装
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 验证安装
wails version

# 输出应该类似：v2.7.0
```

### 4. 安装 Tesseract OCR（可选）

如果需要支持图片简历解析，需要安装 Tesseract：

```bash
# 使用 Chocolatey
choco install tesseract

# 或者下载安装包：https://github.com/tesseract-ocr/tesseract/wiki/Downloads

# 安装中文语言包（可选）
# 下载 chi_sim.traineddata 到 Tesseract 的 tessdata 目录
```

## 项目初始化

### 1. 克隆项目

```bash
# 克隆项目（替换为实际仓库地址）
git clone https://github.com/yourusername/GoResumeReview.git
cd GoResumeReview
```

### 2. 安装前端依赖

```bash
cd frontend
npm install

# 如果遇到网络问题，使用淘宝源
npm install --registry=https://registry.npmmirror.com
```

### 3. 安装后端依赖

```bash
# 返回项目根目录
cd ..

# 安装 Go 依赖
go mod download

# 或者使用国内镜像
GOPROXY=https://goproxy.cn,direct go mod download
```

### 4. 验证环境

```bash
# 运行 Wails 医生命令，检查环境配置
wails doctor

# 应该看到类似输出：
# [✓] go: go1.22.0 [C:\Program Files\Go\]
# [✓] node: v20.10.0 [C:\Program Files\nodejs\]
# [✓] wails: v2.7.0
```

## 开发运行

### 1. 开发模式启动

```bash
wails dev
```

首次运行会：
- 编译前端资源
- 编译后端
- 启动开发服务器
- 打开应用窗口

### 2. 常用开发命令

```bash
# 仅构建前端
cd frontend
npm run build

# 仅构建后端
go build -o bin/app.exe

# 运行前端开发服务器
cd frontend
npm run dev

# 运行后端（配合前端开发服务器）
go run main.go
```

### 3. 调试技巧

#### 后端调试

```bash
# 使用 Delve 调试器
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试
dlv debug .
```

#### 前端调试

在 Chrome DevTools 中：
- 打开应用窗口
- 按 `F12` 打开开发者工具
- 在 Console 查看日志
- 在 Network 查看 API 请求

## 构建发布

### 1. 本地测试构建

```bash
# 构建 Windows 版本
wails build

# 生成的二进制文件位于：build/bin/GoResumeReview.exe
```

### 2. 制作安装包

```bash
# 先构建
wails build

# 生成的安装包位于：build/bin/
# 需要使用 NSIS 或 Inno Setup 制作安装程序
```

## 常见问题

### Q1: wails dev 启动失败

**问题**：`Unable to compile application`

**解决**：
```bash
# 1. 检查 Go 环境
go version

# 2. 清理 Go 缓存
go clean -cache

# 3. 重新下载依赖
go mod tidy
go mod download
```

### Q2: 前端依赖安装失败

**问题**：npm install 超时或失败

**解决**：
```bash
# 使用淘宝镜像
npm install --registry=https://registry.npmmirror.com

# 或者使用 npx 直接运行
npx wails dev
```

### Q3: Tesseract 找不到

**问题**：OCR 功能不可用

**解决**：
```bash
# 1. 确认 Tesseract 已安装
tesseract --version

# 2. 配置 Tesseract 路径（可选）
# 在系统环境变量中添加 Tesseract 路径
```

### Q4: 端口被占用

**问题**：开发服务器端口冲突

**解决**：
```bash
# 查看占用端口的进程
netstat -ano | findstr :34115

# 修改 Wails 配置中的端口
# 在 wails.json 中修改 "info": { "端口": "新端口号" }
```

## 开发工具推荐

### IDE

- **VS Code** - 免费、轻量、插件丰富
- **GoLand** - 专业 Go IDE（需付费）

### 推荐 VS Code 插件

| 插件名称 | 功能 |
|----------|------|
| Go | Go 语言支持 |
| Vue - Official | Vue 3 语言支持 |
| ESLint | 代码检查 |
| Prettier | 代码格式化 |
| Error Lens | 错误提示增强 |
| GitLens | Git 可视化 |

### Windows 开发工具

| 工具 | 用途 |
|------|------|
| Process Explorer | 进程分析 |
| API Monitor | API 调用监控 |
| Windows Terminal | 终端替代 |

## 下一步

- 阅读[架构设计文档](design/architecture.md)
- 了解[文档解析器设计](design/document-parser.md)
- 查看 [AI 分析引擎设计](design/ai-engine.md)
- 开始[开发任务](DEVELOPMENT.md)
