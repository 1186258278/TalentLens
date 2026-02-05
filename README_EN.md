<p align="center">
  <h1 align="center">TalentLens</h1>
  <p align="center">AI-Powered Resume Screening Tool for HR</p>
</p>

<p align="center">
  English | <a href="README.md">中文</a>
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

## What is TalentLens?

**TalentLens** is a lightweight, elegant desktop application that helps HR professionals screen resumes efficiently using AI. Simply drag-and-drop resumes, and TalentLens will analyze, score, and rank candidates based on your job requirements.

Available on **Windows** and **macOS**.

---

## Features

| Feature | Description |
|---------|-------------|
| **Drag & Drop** | Upload resumes in PDF / Word / Image formats |
| **AI Analysis** | Deep analysis via OpenAI-compatible APIs |
| **Multi-dimensional Scoring** | Skills match, experience match, education match |
| **Recommendations** | AI provides detailed recommendation with reasoning |
| **Job Templates** | Built-in templates for common job positions |
| **Provider Presets** | Pre-configured for DeepSeek / OpenAI / Zhipu / Moonshot / SiliconFlow |
| **i18n** | Simplified Chinese / Traditional Chinese / English |
| **macOS-style UI** | Frosted glass titlebar, traffic light buttons, polished interface |
| **Real-time Progress** | Live visualization of analysis progress |
| **Dev Console** | Built-in developer debug panel (F12) |

---

## Quick Start

### Download

Download the latest release from the [Releases](https://github.com/1186258278/TalentLens/releases) page:

- **Windows**: `TalentLens-windows-amd64.exe`
- **macOS (Intel)**: `TalentLens-macos-amd64.zip`
- **macOS (Apple Silicon)**: `TalentLens-macos-arm64.zip`

Just download and run. No installation required.

### First-time Setup

1. Open the app and click the **gear icon** in the top-right to open Settings
2. Select an AI provider (DeepSeek recommended - free quota for new users)
3. Follow the guide to obtain and enter your API Key
4. Configure job requirements (use built-in templates for quick setup)
5. Return to the home page, drag in resumes and start analyzing

---

## Supported AI Providers

| Provider | Recommended Model | Notes |
|----------|------------------|-------|
| DeepSeek | DeepSeek-V3.2 | Recommended, excellent cost-performance ratio |
| OpenAI | GPT-5 Mini | Latest GPT-5 family |
| Zhipu AI | GLM-4.7-Flash | Chinese LLM, free tier available |
| Moonshot (Kimi) | Kimi K2.5 | Latest multimodal agent model |
| SiliconFlow | DeepSeek-V3.2 | Model aggregator, 100+ models |
| Custom | - | Any OpenAI-compatible API |

---

## Build from Source

### Prerequisites

- [Go](https://go.dev/) >= 1.21
- [Node.js](https://nodejs.org/) >= 18
- [Wails CLI](https://wails.io/) v2

### Build

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

# Cross-platform builds
wails build -platform darwin/arm64   # macOS Apple Silicon
wails build -platform darwin/amd64   # macOS Intel
wails build -platform windows/amd64  # Windows
```

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.23, Wails v2 |
| Frontend | Vue 3, TypeScript, Vite |
| UI Library | Element Plus |
| State Management | Pinia |
| i18n | vue-i18n |
| Styling | SCSS, macOS Design System |

---

## Project Structure

```
TalentLens/
├── app.go                 # Go backend (AI client, file processing, API)
├── wails.json             # Wails configuration
├── frontend/
│   ├── src/
│   │   ├── views/         # Page views (HomeView, SettingsView)
│   │   ├── components/    # Components (TitleBar, ResumeCard, DropZone...)
│   │   ├── composables/   # Pinia stores
│   │   ├── locales/       # i18n translations
│   │   ├── styles/        # macOS theme variables
│   │   └── data/          # Job presets, provider configs
│   └── index.html
├── build/                 # Build assets & installers
├── docs/                  # Documentation
├── .github/workflows/     # CI/CD auto-build
├── CHANGELOG.md           # Version changelog
└── CONTRIBUTING.md        # Contribution guide
```

---

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines and submission process.

---

## License

This project is licensed under the **GNU General Public License v3.0** - see the [LICENSE](LICENSE) file for details.

---

## About

**TalentLens** is developed by **Wuhan QingChen TianXia Network Technology Co., Ltd.** (武汉晴辰天下网络科技有限公司)

- Website: [https://qingchencloud.com](https://qingchencloud.com)
- Products: [https://qt.cool](https://qt.cool)
- Brand: QingChen Cloud / 晴辰云

---

<p align="center">Made with ❤️ by QingChen Cloud</p>
