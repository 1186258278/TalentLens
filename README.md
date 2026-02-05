<p align="center">
  <h1 align="center">TalentLens</h1>
  <p align="center">AI-Powered Resume Screening Tool for HR</p>
  <p align="center">AI 驱动的智能简历筛选工具，专为 HR 设计</p>
</p>

<p align="center">
  <a href="#features">Features</a> ·
  <a href="#quick-start">Quick Start</a> ·
  <a href="#development">Development</a> ·
  <a href="#license">License</a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.23-00ADD8?logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Vue-3-4FC08D?logo=vue.js" alt="Vue 3">
  <img src="https://img.shields.io/badge/Wails-v2-red" alt="Wails">
  <img src="https://img.shields.io/badge/License-GPL--3.0-blue" alt="License">
  <img src="https://img.shields.io/badge/Platform-Windows-0078D6?logo=windows" alt="Windows">
</p>

---

## What is TalentLens?

**TalentLens** is a lightweight, elegant desktop application that helps HR professionals screen resumes efficiently using AI. Simply drag-and-drop resumes, and TalentLens will analyze, score, and rank candidates based on your job requirements.

**TalentLens** 是一款轻量级桌面应用，帮助 HR 利用 AI 高效筛选简历。只需拖入简历文件，即可自动分析、打分、排序，根据岗位需求给出专业建议。

---

## Features

| Feature | Description |
|---------|-------------|
| **Drag & Drop** | 拖拽上传 PDF / Word / 图片格式简历 |
| **AI Analysis** | 基于 OpenAI 兼容接口进行智能分析 |
| **Scoring** | 多维度评分：技能匹配、经验匹配、学历匹配 |
| **Recommendations** | AI 给出推荐/不推荐结论及详细理由 |
| **Job Templates** | 内置多种岗位模板，一键配置需求 |
| **Provider Presets** | 预设 DeepSeek / OpenAI / 智谱 / Moonshot 等服务商 |
| **i18n** | 支持简体中文 / 繁体中文 / English |
| **macOS Style UI** | 毛玻璃标题栏、交通灯按钮、精致 UI |
| **Dev Console** | 内置开发者调试面板 (F12) |
| **Real-time Progress** | 分析过程实时可视化进度 |

---

## Quick Start

### Download

从 [Releases](https://github.com/1186258278/TalentLens/releases) 下载最新版本的 `TalentLens.exe`，双击运行即可。

### First-time Setup

1. 打开应用，点击右上角 **齿轮图标** 进入设置
2. 选择 AI 服务商（推荐 DeepSeek，新用户有免费额度）
3. 按照引导获取并填入 API Key
4. 配置岗位需求（可使用内置模板）
5. 返回主页，拖入简历开始分析

---

## Development

### Prerequisites

- [Go](https://go.dev/) >= 1.21
- [Node.js](https://nodejs.org/) >= 18
- [Wails CLI](https://wails.io/) v2

### Build from Source

```bash
# Clone the repository
git clone https://github.com/1186258278/TalentLens.git
cd TalentLens

# Install frontend dependencies
cd frontend && npm install && cd ..

# Development mode (hot reload)
wails dev

# Production build
wails build
```

### Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.23, Wails v2 |
| Frontend | Vue 3, TypeScript, Vite |
| UI Library | Element Plus |
| State | Pinia |
| i18n | vue-i18n |
| Styling | SCSS, macOS Design System |

### Project Structure

```
TalentLens/
├── app.go                 # Go backend (AI client, file processing, API)
├── wails.json             # Wails configuration
├── frontend/
│   ├── src/
│   │   ├── views/         # HomeView, SettingsView
│   │   ├── components/    # TitleBar, ResumeCard, DropZone, DevPanel...
│   │   ├── composables/   # Pinia stores
│   │   ├── locales/       # i18n translations
│   │   ├── styles/        # macOS theme variables
│   │   └── data/          # Job presets, provider configs
│   └── index.html
├── build/                 # Build assets & installers
└── docs/                  # Project documentation
```

---

## Supported AI Providers

| Provider | Default Model | Notes |
|----------|--------------|-------|
| DeepSeek | deepseek-chat | Recommended, free quota for new users |
| OpenAI | gpt-4o-mini | Most popular |
| Zhipu AI (智谱) | glm-4-flash | Chinese LLM |
| Moonshot (月之暗面) | moonshot-v1-8k | Chinese LLM |
| SiliconFlow | Qwen/Qwen2.5-7B | Aggregator |
| Custom | - | Any OpenAI-compatible API |

---

## License

This project is licensed under the **GNU General Public License v3.0** - see the [LICENSE](LICENSE) file for details.

---

## About

**TalentLens** is developed by **武汉晴辰天下网络科技有限公司** (Wuhan QingChen TianXia Network Technology Co., Ltd.)

- Website: [https://qingchencloud.com](https://qingchencloud.com)
- Products: [https://qt.cool](https://qt.cool)
- Brand: 晴辰云 / 晴辰

---

<p align="center">Made with ❤️ by QingChen Cloud</p>
