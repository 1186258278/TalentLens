# 贡献指南 / Contributing Guide

感谢你对 TalentLens 的关注！欢迎提交 Issue 和 Pull Request。

---

## 开发环境搭建

### 前置要求

| 工具 | 版本 | 说明 |
|------|------|------|
| [Go](https://go.dev/) | >= 1.21 | 后端运行时 |
| [Node.js](https://nodejs.org/) | >= 18 | 前端构建 |
| [Wails CLI](https://wails.io/) | v2 | 桌面应用框架 |

### 安装 Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 克隆并运行

```bash
git clone https://github.com/1186258278/TalentLens.git
cd TalentLens

# 安装前端依赖
cd frontend && npm install && cd ..

# 启动开发模式 (热重载)
wails dev
```

---

## 项目结构

```
TalentLens/
├── app.go                     # Go 后端主文件
├── wails.json                 # Wails 配置
├── frontend/
│   ├── src/
│   │   ├── views/             # 页面组件
│   │   ├── components/        # 通用组件
│   │   ├── composables/       # Pinia Store
│   │   ├── locales/           # i18n 翻译
│   │   ├── styles/            # 全局样式
│   │   └── data/              # 静态数据
│   └── index.html
├── build/                     # 构建资源
├── docs/                      # 文档
└── .github/workflows/         # CI/CD
```

---

## 代码规范

### Go 后端

- 遵循 [Effective Go](https://go.dev/doc/effective_go) 规范
- 使用 `gofmt` 格式化代码
- 函数注释使用中文
- 错误处理不要忽略，至少记录日志

### Vue 前端

- 使用 Composition API (`<script setup>`)
- 组件文件使用 PascalCase 命名
- 单个 Vue 文件不超过 500 行
- 样式使用 `<style scoped lang="scss">`
- 使用 macOS 主题变量 (`@import '../styles/macos-theme.scss'`)

### 国际化

- 所有用户可见的文字必须使用 `$t()` 或 `t()` 函数
- 翻译文件位于 `frontend/src/locales/`
- 新增文字需同时更新三个语言文件：`zh-CN.ts`, `zh-TW.ts`, `en-US.ts`

---

## 提交规范

使用 [Conventional Commits](https://www.conventionalcommits.org/) 格式：

```
<类型>(<范围>): <描述>

<正文>
```

### 类型

| 类型 | 说明 |
|------|------|
| `feat` | 新功能 |
| `fix` | 修复 Bug |
| `docs` | 文档变更 |
| `style` | 代码格式（不影响功能） |
| `refactor` | 重构（不是新功能或修复） |
| `perf` | 性能优化 |
| `test` | 测试相关 |
| `chore` | 构建工具或辅助工具变更 |
| `ci` | CI/CD 配置变更 |

### 示例

```
feat(analysis): add batch analysis progress display
fix(settings): fix API key not persisting after save
docs: update README with macOS build instructions
ci: add GitHub Actions multi-platform build workflow
```

---

## 分支策略

| 分支 | 用途 |
|------|------|
| `main` | 稳定版本，保护分支 |
| `develop` | 开发分支，日常开发合并到此 |
| `feature/*` | 功能分支，从 develop 创建 |
| `fix/*` | 修复分支 |
| `release/*` | 发布准备分支 |

### 工作流

1. 从 `develop` 创建功能分支：`git checkout -b feature/my-feature develop`
2. 开发完成后提交 Pull Request 到 `develop`
3. 代码审查通过后合并
4. 发布时从 `develop` 创建 `release/vX.Y.Z` 分支
5. 测试通过后合并到 `main` 并打标签

---

## 发布流程

```bash
# 1. 确认所有功能已合并到 main
# 2. 更新 CHANGELOG.md
# 3. 更新版本号 (wails.json 中的 productVersion)
# 4. 提交并打标签
git tag v1.0.0
git push origin v1.0.0

# 5. GitHub Actions 自动构建并发布 Release
#    - Windows amd64/arm64
#    - macOS amd64/arm64
```

---

## 报告问题

提交 Issue 时请包含以下信息：

1. **操作系统**: Windows / macOS，版本号
2. **应用版本**: TalentLens 版本号
3. **复现步骤**: 如何触发该问题
4. **期望行为**: 你认为应该发生什么
5. **实际行为**: 实际发生了什么
6. **截图/日志**: 如有，请附上 (F12 打开开发者面板查看日志)

---

## 许可证

提交代码即表示你同意将代码贡献在 [GPL v3](LICENSE) 许可证下。
