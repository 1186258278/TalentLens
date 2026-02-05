# API 接口设计文档

本文档描述 GoResumeReview 前后端通信接口设计。

## 概述

Wails 应用采用 Go-JS 绑定方式前后端通信，无需传统 REST API。

## Go 调用 JavaScript

### 事件机制

```go
// Go 发送事件到前端
runtime.EventsEmit(ctx, "eventName", data)

// Go 发送带进度的事件
runtime.EventsEmit(ctx, "analysis:progress", map[string]interface{}{
    "taskId": "xxx",
    "progress": 50,
    "status": "parsing",
})
```

### 日志

```go
// 调试日志
runtime.LogDebug(ctx, "Debug message")
runtime.LogInfo(ctx, "Info message")
runtime.LogWarning(ctx, "Warning message")
runtime.LogError(ctx, "Error message")
```

## JavaScript 调用 Go

### 绑定方法

Go 端通过 `Bind` 选项暴露方法：

```go
// main.go
type App struct {
    // ...
}

func (a *App) Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func main() {
    err := wails.Run(&options.App{
        // ...
        Bind: []interface{}{
            &App{},
        },
    })
}
```

前端调用：

```typescript
// frontend/src/wailsjs/runtime/runtime.ts
import { Greet } from '../../wailsjs/go/main/App';

// 调用 Go 方法
const result = await Greet('World');
```

## 核心 API 设计

### 1. 简历管理

```go
// 添加简历到分析队列
func (a *App) AddResumes(filePaths []string) (*AddResumesResult, error)

// 获取所有简历
func (a *App) GetResumes() ([]*Resume, error)

// 获取单个简历详情
func (a *App) GetResume(id string) (*Resume, error)

// 删除简历
func (a *App) DeleteResume(id string) error

// 重新分析简历
func (a *App) ReAnalyzeResume(id string) error

// 清空所有简历
func (a *App) ClearResumes() error
```

### 2. 分析控制

```go
// 开始批量分析
func (a *App) StartAnalysis() error

// 暂停分析
func (a *App) PauseAnalysis() error

// 继续分析
func (a *App) ResumeAnalysis() error

// 取消分析
func (a *App) CancelAnalysis() error

// 获取分析进度
func (a *App) GetAnalysisProgress() (*AnalysisProgress, error)
```

### 3. 配置管理

```go
// 获取当前配置
func (a *App) GetConfig() (*Config, error)

// 保存配置
func (a *App) SaveConfig(config *Config) error

// 测试 API 连接
func (a *App) TestAPIConnection() error

// 重置为默认配置
func (a *App) ResetConfig() error
```

### 4. 预览控制

```go
// 获取文件内容（用于预览）
func (a *App) GetFileContent(filePath string) ([]byte, error)

// 简历文本内容
func (a *App) GetResumeText(id string) (string, error)
```

## 数据类型定义

### Resume（简历）

```typescript
interface Resume {
    id: string;
    fileName: string;
    filePath: string;
    fileType: 'pdf' | 'docx' | 'doc' | 'image';
    fileSize: number;
    fileHash: string;
    content?: string;
    status: 'pending' | 'parsing' | 'analyzing' | 'done' | 'error';
    score?: number;
    analysis?: AnalysisResult;
    error?: string;
    createdAt: string;
    analyzedAt?: string;
}
```

### AnalysisResult（分析结果）

```typescript
interface AnalysisResult {
    overallScore: number;
    experienceMatch: number;
    skillMatch: number;
    educationMatch: number;
    summary: string;
    strengths: string[];
    weaknesses: string[];
    recommendation: 'strongly_recommended' | 'recommended' | 'consider' | 'not_recommended';
    detailedAnalysis?: string;
}
```

### AnalysisProgress（分析进度）

```typescript
interface AnalysisProgress {
    total: number;
    completed: number;
    failed: number;
    currentFile?: string;
    status: 'idle' | 'running' | 'paused' | 'completed';
    estimatedTimeRemaining?: number; // 秒
}
```

### Config（配置）

```typescript
interface Config {
    // AI 配置
    ai: {
        provider: 'openai' | 'azure' | 'anthropic' | 'zhipu' | 'deepseek' | 'custom';
        baseUrl: string;
        apiKey: string;
        model: string;
        maxRetries: number;
        timeout: number;
    };
    
    // 岗位需求
    job: {
        title: string;
        requirements: string[];
        requiredSkills: string[];
        preferredSkills: string[];
        experienceYears: number;
        educationLevel: string;
    };
    
    // 分析选项
    analysis: {
        maxConcurrent: number;
        autoStart: boolean;
        showNotifications: boolean;
    };
}
```

## 事件流

### 简历添加事件

```typescript
// 前端监听
window.runtime.EventsOn('resumes:added', (data) => {
    console.log('新增简历数量:', data.count);
    console.log('简历列表:', data.resumes);
});

// 数据格式
interface ResumesAddedEvent {
    count: number;
    resumes: Resume[];
    alreadyAnalyzed: string[]; // 已分析过的简历ID列表
}
```

### 分析进度事件

```typescript
// 前端监听
window.runtime.EventsOn('analysis:progress', (data) => {
    updateProgress(data);
});

// 数据格式
interface AnalysisProgressEvent {
    taskId: string;
    status: 'parsing' | 'analyzing' | 'completed' | 'error';
    progress: number; // 0-100
    currentFile?: string;
    result?: AnalysisResult;
    error?: string;
}
```

### 分析完成事件

```typescript
// 前端监听
window.runtime.EventsOn('analysis:completed', (data) => {
    showNotification('分析完成', `共分析 ${data.completed} 份简历`);
});

// 数据格式
interface AnalysisCompletedEvent {
    total: number;
    completed: number;
    failed: number;
    results: {
        resumeId: string;
        score: number;
        recommendation: string;
    }[];
}
```

## 错误处理

### 错误码定义

```go
const (
    ErrCodeSuccess = 0
    ErrCodeFileNotFound = 1001
    ErrCodeInvalidFormat = 1002
    ErrCodeParseFailed = 1003
    ErrCodeAPIError = 2001
    ErrCodeRateLimited = 2002
    ErrCodeInvalidAPIKey = 2003
    ErrCodeTimeout = 2004
    ErrCodeAnalysisFailed = 3001
    ErrCodeCacheError = 4001
)
```

### 错误响应格式

```typescript
interface APIError {
    code: number;
    message: string;
    details?: any;
}
```

## WebSocket 模拟

由于 Wails 不支持真正的 WebSocket，使用事件模拟实时通信：

```typescript
// 订阅分析进度
const progressUnsubscribe = window.runtime.EventsOn(
    'analysis:progress',
    (data) => {
        // 处理进度更新
    }
);

// 取消订阅
progressUnsubscribe();
```

## 性能优化

### 批量操作

```go
// 推荐：批量添加简历
func (a *App) AddResumesBatch(filePaths []string) (*BatchResult, error) {
    // 内部使用 goroutine 并发处理
    // 返回进度和结果
}
```

### 大文件处理

```go
// 对于大文件，使用流式处理
func (a *App) ParseLargeFile(filePath string, onProgress func(float64)) error {
    // 分块读取
    // 回调进度
}
```
