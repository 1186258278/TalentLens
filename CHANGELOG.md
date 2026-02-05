# Changelog / 版本变更日志

所有重要变更都会记录在此文件中。  
格式参考 [Keep a Changelog](https://keepachangelog.com/zh-CN/)。

---

## [1.0.0] - 2025-02-06

### 新增 (Added)

- 核心功能：拖拽上传简历（PDF / Word / 图片格式）
- AI 智能分析：基于 OpenAI 兼容接口的多维度简历评估
- 综合评分系统：技能匹配度、经验匹配度、学历匹配度
- AI 推荐结论：强烈推荐 / 推荐 / 可考虑 / 不推荐
- AI 详细分析：优势亮点、待改进项、综合总结
- 岗位模板：内置多种常见岗位模板，支持自定义
- 多 AI 服务商：预设 DeepSeek / OpenAI / 智谱 / Moonshot / SiliconFlow / 自定义
- 服务商配置引导：每个服务商都有详细的 API Key 获取教程
- 国际化：支持简体中文、繁体中文、English
- macOS 风格 UI：毛玻璃标题栏、交通灯窗口按钮
- 侧边栏 Tab 设置页面
- 实时分析进度条（全局 + 单个简历）
- 分析进度阶段文字（准备分析 / AI 分析中 / 解析结果）
- AI 配置检查：未配置时弹出引导对话框
- 开发者调试面板（F12 快捷键）
- 跨平台支持：Windows + macOS
- GPL v3 开源协议

### 技术栈

- 后端：Go 1.23 + Wails v2
- 前端：Vue 3 + TypeScript + Vite + Element Plus
- 状态管理：Pinia
- 国际化：vue-i18n
- 样式：SCSS + macOS Design System
