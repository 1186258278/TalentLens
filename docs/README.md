# GoResumeReview 项目文档

## 项目概述

GoResumeReview 是一款面向 HR 群体的智能简历审核工具，支持 Windows 桌面客户端，采用 Go + Vue3 技术栈开发，实现高并发简历分析和可视化预览功能。

## 文档目录

- [架构设计文档](design/architecture.md)
- [API接口设计](api/README.md)
- [开发指南](guides/getting-started.md)
- [用户手册](guides/user-manual.md)

## 快速开始

```bash
# 克隆项目
git clone https://github.com/yourusername/GoResumeReview.git

# 安装依赖
cd GoResumeReview
wails doctor

# 开发模式运行
wails dev
```

## 技术栈

- **GUI框架**: Wails v2 (Go + Web)
- **前端框架**: Vue3 + TypeScript
- **UI组件库**: Element Plus
- **文档解析**: pdfcpu + docxgo + Tesseract OCR
- **AI集成**: OpenAI 兼容 API
- **数据存储**: SQLite

## 项目状态

当前状态：立项阶段

## 许可证

MIT License
