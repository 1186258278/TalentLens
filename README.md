<p align="center">
  <h1 align="center">TalentLens</h1>
  <p align="center">AI 驱动的智能简历筛选工具，专为 HR 设计</p>
</p>

<p align="center">
  <a href="README_EN.md">English</a> | 中文
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23-00ADD8?logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Vue-3-4FC08D?logo=vue.js" alt="Vue 3">
  <img src="https://img.shields.io/badge/Wails-v2-red" alt="Wails">
  <img src="https://img.shields.io/badge/License-GPL--3.0-blue" alt="License">
  <img src="https://img.shields.io/badge/Windows-0078D6?logo=windows&logoColor=white" alt="Windows">
  <img src="https://img.shields.io/badge/macOS-000000?logo=apple&logoColor=white" alt="macOS">
</p>

---

## 简介

**TalentLens** 是一款轻量级桌面应用，帮助 HR 利用 AI 高效筛选简历。只需拖入简历文件，即可自动分析、打分、排序，根据岗位需求给出专业推荐建议。

支持 **Windows** 和 **macOS** 双平台。

---

## 功能特性

| 功能 | 说明 |
|------|------|
| **拖拽上传** | 支持 PDF / Word / 图片格式简历 |
| **AI 智能分析** | 基于 OpenAI 兼容接口，多维度深度分析 |
| **综合评分** | 技能匹配度、经验匹配度、学历匹配度 |
| **推荐结论** | AI 给出推荐/不推荐结论及详细理由 |
| **岗位模板** | 内置多种岗位模板，一键配置需求 |
| **多服务商** | 预设 DeepSeek / OpenAI / 智谱 / Moonshot / SiliconFlow |
| **国际化** | 支持简体中文 / 繁体中文 / English |
| **macOS 风格 UI** | 毛玻璃标题栏、交通灯按钮、精致界面 |
| **实时进度** | 分析过程实时可视化进度显示 |
| **开发者工具** | 内置调试面板 (F12) |

---

## 快速开始

### 下载安装

从 [Releases](https://github.com/1186258278/TalentLens/releases) 页面下载最新版本：

- **Windows**: `TalentLens-windows-amd64.exe`
- **macOS (Intel)**: `TalentLens-macos-amd64.zip`
- **macOS (Apple Silicon)**: `TalentLens-macos-arm64.zip`

下载后双击运行即可，无需安装。

### 首次配置

1. 打开应用，点击右上角 **齿轮图标** 进入设置
2. 选择 AI 服务商（推荐 DeepSeek，新用户有免费额度）
3. 按照引导获取并填入 API Key
4. 配置岗位需求（可使用内置模板）
5. 返回主页，拖入简历开始分析

---

## 支持的 AI 服务商

| 服务商 | 推荐模型 | 说明 |
|--------|---------|------|
| DeepSeek | DeepSeek-V3.2 | 推荐，性价比高，支持对话/深度推理双模式 |
| OpenAI | GPT-5 Mini | 全球最流行，GPT-5 系列最新 |
| 智谱 AI | GLM-4.7-Flash | 国产旗舰，免费调用 |
| Moonshot (Kimi) | Kimi K2.5 | 最新多模态智能体模型 |
| 硅基流动 | DeepSeek-V3.2 | 模型聚合平台，100+ 模型一站调用 |
| 自定义 | - | 任何 OpenAI 兼容接口 |

---

## 从源码构建

### 环境要求

- [Go](https://go.dev/) >= 1.21
- [Node.js](https://nodejs.org/) >= 18
- [Wails CLI](https://wails.io/) v2

### 构建步骤

```bash
# 克隆仓库
git clone https://github.com/1186258278/TalentLens.git
cd TalentLens

# 安装前端依赖
cd frontend && npm install && cd ..

# 开发模式 (热重载)
wails dev

# 生产构建
wails build

# 跨平台构建
wails build -platform darwin/arm64   # macOS Apple Silicon
wails build -platform darwin/amd64   # macOS Intel
wails build -platform windows/amd64  # Windows
```

---

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Go 1.23, Wails v2 |
| 前端 | Vue 3, TypeScript, Vite |
| UI 组件 | Element Plus |
| 状态管理 | Pinia |
| 国际化 | vue-i18n |
| 样式 | SCSS, macOS 设计系统 |

---

## 项目结构

```
TalentLens/
├── app.go                 # Go 后端 (AI 客户端、文件处理、API)
├── wails.json             # Wails 项目配置
├── frontend/
│   ├── src/
│   │   ├── views/         # 页面视图 (HomeView, SettingsView)
│   │   ├── components/    # 组件 (TitleBar, ResumeCard, DropZone...)
│   │   ├── composables/   # Pinia Store
│   │   ├── locales/       # 国际化翻译文件
│   │   ├── styles/        # macOS 主题变量
│   │   └── data/          # 岗位模板、服务商配置
│   └── index.html
├── build/                 # 构建资源和安装脚本
├── docs/                  # 项目文档
├── .github/workflows/     # CI/CD 自动构建
├── CHANGELOG.md           # 版本变更日志
└── CONTRIBUTING.md        # 贡献指南
```

---

## 参与贡献

欢迎参与贡献！请阅读 [CONTRIBUTING.md](CONTRIBUTING.md) 了解开发规范和提交流程。

---

## 许可证

本项目采用 **GNU General Public License v3.0** 开源协议，详见 [LICENSE](LICENSE)。

---

## 关于

**TalentLens** 由 **武汉晴辰天下网络科技有限公司** 开发维护。

- 产品官网：[https://talentlens.qt.cool](https://talentlens.qt.cool)
- 公司官网：[https://qingchencloud.com](https://qingchencloud.com)
- 产品导航：[https://qt.cool](https://qt.cool)
- 品牌：晴辰云 / 晴辰

---

<p align="center">Made with ❤️ by QingChen Cloud</p>
