package main

import (
	"bytes"
	"context"
	"encoding/json"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// Config 配置结构
type Config struct {
	AI   AIConfig   `json:"ai"`
	Job  JobConfig  `json:"job"`
}

// AIConfig AI配置
type AIConfig struct {
	Provider   string `json:"provider"`
	BaseURL    string `json:"base_url"`
	APIKey     string `json:"api_key"`
	Model      string `json:"model"`
	MaxRetries int    `json:"max_retries"`
	Timeout    int    `json:"timeout"`
}

// JobConfig 岗位配置
type JobConfig struct {
	Title           string   `json:"title"`
	Requirements    []string `json:"requirements"`
	RequiredSkills  []string `json:"required_skills"`
	ExperienceYears int      `json:"experience_years"`
	EducationLevel  string   `json:"education_level"`
}

// Resume 简历结构
type Resume struct {
	ID        string          `json:"id"`
	FileName  string          `json:"file_name"`
	FilePath  string          `json:"file_path"`
	FileType  string          `json:"file_type"`
	FileSize  int64           `json:"file_size"`
	Content   string          `json:"content"`
	Status    string          `json:"status"`
	Score     int             `json:"score"`
	Analysis  *AnalysisResult `json:"analysis,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
}

// AnalysisResult AI分析结果
type AnalysisResult struct {
	OverallScore     int      `json:"overall_score"`
	SkillMatch       int      `json:"skill_match"`
	ExperienceMatch  int      `json:"experience_match"`
	EducationMatch   int      `json:"education_match"`
	Recommendation   string   `json:"recommendation"`
	Strengths        []string `json:"strengths"`
	Weaknesses       []string `json:"weaknesses"`
	Summary          string   `json:"summary"`
	AnalyzedAt       string   `json:"analyzed_at"`
}

// OpenAI API 请求/响应结构
type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature,omitempty"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type App struct {
	ctx    context.Context
	config Config
}

func NewApp() *App {
	return &App{}
}

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "TalentLens",
		Width:            1200,
		Height:           800,
		MinWidth:         900,
		MinHeight:        600,
		Frameless:        true,
		DisableResize:    false,
		StartHidden:      false,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
}

// WindowMinimize 最小化窗口
func (a *App) WindowMinimize() {
	runtime.WindowMinimise(a.ctx)
}

// WindowMaximize 最大化/还原窗口
func (a *App) WindowMaximize() {
	runtime.WindowToggleMaximise(a.ctx)
}

// WindowClose 关闭窗口
func (a *App) WindowClose() {
	runtime.Quit(a.ctx)
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.loadConfig()

	// 监听文件拖拽
	runtime.OnFileDrop(ctx, func(x, y int, paths []string) {
		log.Printf("收到文件: %d个", len(paths))
		for _, path := range paths {
			a.processFile(path)
		}
	})

	log.Println("GoResumeReview 已启动")
}

func (a *App) getConfigPath() string {
	dir := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview")
	os.MkdirAll(dir, 0755)
	return filepath.Join(dir, "config.json")
}

func (a *App) loadConfig() {
	data, err := os.ReadFile(a.getConfigPath())
	if err != nil {
		a.config = Config{
			AI: AIConfig{
				Provider:   "openai",
				BaseURL:    "https://api.openai.com/v1",
				Model:      "gpt-4o",
				MaxRetries: 3,
				Timeout:    60,
			},
			Job: JobConfig{
				Title:          "高级Go开发工程师",
				RequiredSkills: []string{"Go", "MySQL", "Redis"},
			},
		}
		return
	}
	json.Unmarshal(data, &a.config)
}

func (a *App) processFile(filePath string) {
	info, _ := os.Stat(filePath)
	ext := strings.ToLower(filepath.Ext(filePath))

	// 简单解析文本
	content := a.extractText(filePath)

	resume := &Resume{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		FileName:  filepath.Base(filePath),
		FilePath:  filePath,
		FileType:  ext,
		FileSize:  info.Size(),
		Content:   content,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// 保存
	a.saveResume(resume)

	// 发送到前端
	runtime.EventsEmit(a.ctx, "resume:added", resume)
}

func (a *App) extractText(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".txt", ".md":
		return string(data)
	case ".pdf":
		return a.extractFromPDF(data)
	default:
		return a.extractFromBytes(data)
	}
}

func (a *App) extractFromPDF(data []byte) string {
	content := string(data)
	var lines []string

	for len(content) > 0 {
		start := strings.Index(content, "(")
		end := strings.Index(content, ")")
		if start == -1 || end == -1 || end <= start {
			break
		}
		text := content[start+1 : end]
		clean := ""
		for _, ch := range text {
			if ch >= 32 && ch <= 126 {
				clean += string(ch)
			}
		}
		if len(clean) > 2 {
			lines = append(lines, clean)
		}
		content = content[end+1:]
	}

	result := strings.Join(lines, "\n")
	if len(result) > 50000 {
		result = result[:50000]
	}
	return result
}

func (a *App) extractFromBytes(data []byte) string {
	content := string(data)
	lines := strings.Split(content, "\n")
	var clean []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			clean = append(clean, line)
		}
	}
	result := strings.Join(clean, "\n")
	if len(result) > 50000 {
		result = result[:50000]
	}
	return result
}

func (a *App) saveResume(r *Resume) {
	dir := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes")
	os.MkdirAll(dir, 0755)
	data, _ := json.MarshalIndent(r, "", "  ")
	os.WriteFile(filepath.Join(dir, r.ID+".json"), data, 0644)
}

// RegisterResume 前端拖入简历后，通知后端注册并保存到磁盘
// 前端通过 HTML5 拖拽添加文件时，后端无感知，需要前端主动调用此方法
func (a *App) RegisterResume(id string, fileName string, filePath string, fileType string, fileSize int64) (bool, string) {
	log.Printf("[RegisterResume] id=%s, file=%s, path=%s", id, fileName, filePath)

	// 提取文件内容
	content := ""
	if filePath != "" && filePath != fileName {
		// 有真实路径，尝试读取文件内容
		content = a.extractText(filePath)
		log.Printf("[RegisterResume] 提取内容长度: %d", len(content))
	}

	if content == "" {
		content = fmt.Sprintf("[简历文件: %s, 类型: %s, 大小: %d bytes]", fileName, fileType, fileSize)
		log.Printf("[RegisterResume] 使用占位内容")
	}

	resume := &Resume{
		ID:        id,
		FileName:  fileName,
		FilePath:  filePath,
		FileType:  fileType,
		FileSize:  fileSize,
		Content:   content,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	a.saveResume(resume)
	log.Printf("[RegisterResume] 简历已保存: %s", id)
	return true, "简历已注册: " + fileName
}

func (a *App) GetConfig() *Config {
	return &a.config
}

func (a *App) SaveConfig(cfg *Config) error {
	a.config = *cfg
	data, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(a.getConfigPath(), data, 0644)
}

func (a *App) GetResumes() []*Resume {
	dir := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes")
	os.MkdirAll(dir, 0755)

	var resumes []*Resume
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		data, _ := os.ReadFile(filepath.Join(dir, entry.Name()))
		var r Resume
		json.Unmarshal(data, &r)
		resumes = append(resumes, &r)
	}
	return resumes
}

func (a *App) DeleteResume(id string) error {
	path := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes", id+".json")
	return os.Remove(path)
}

func (a *App) ReAnalyzeResume(id string) error {
	path := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes", id+".json")
	data, _ := os.ReadFile(path)
	var r Resume
	json.Unmarshal(data, &r)
	r.Status = "pending"
	data, _ = json.MarshalIndent(r, "", "  ")
	os.WriteFile(path, data, 0644)
	runtime.EventsEmit(a.ctx, "resume:updated", &r)
	return nil
}

func (a *App) ClearResumes() error {
	dir := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes")
	os.RemoveAll(dir)
	os.MkdirAll(dir, 0755)
	return nil
}

func (a *App) GetResumeText(id string) (string, error) {
	resumes := a.GetResumes()
	for _, r := range resumes {
		if r.ID == id {
			return r.Content, nil
		}
	}
	return "", nil
}

// TestAIConnection 测试AI连接
func (a *App) TestAIConnection(cfg *AIConfig) (bool, string) {
	if cfg.APIKey == "" {
		return false, "API Key 不能为空"
	}
	if cfg.BaseURL == "" {
		return false, "Base URL 不能为空"
	}

	// 构造测试请求
	reqBody := ChatRequest{
		Model: cfg.Model,
		Messages: []ChatMessage{
			{Role: "user", Content: "Hello, this is a connection test. Please respond with 'OK'."},
		},
		MaxTokens: 10,
	}

	body, _ := json.Marshal(reqBody)
	url := strings.TrimSuffix(cfg.BaseURL, "/") + "/chat/completions"

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return false, fmt.Sprintf("创建请求失败: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.APIKey)

	client := &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Sprintf("连接失败: %v", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == 401 {
		return false, "API Key 无效或已过期"
	}
	if resp.StatusCode == 403 {
		return false, "API Key 权限不足"
	}
	if resp.StatusCode == 429 {
		return false, "请求过于频繁，请稍后再试"
	}
	if resp.StatusCode >= 500 {
		return false, "服务器错误，请稍后再试"
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		return false, fmt.Sprintf("解析响应失败: %v", err)
	}

	if chatResp.Error != nil {
		return false, fmt.Sprintf("AI 错误: %s", chatResp.Error.Message)
	}

	if len(chatResp.Choices) == 0 {
		return false, "AI 没有返回任何响应"
	}

	return true, "连接成功！AI 服务正常"
}

// AnalyzeResume 分析单个简历
func (a *App) AnalyzeResume(resumeID string, cfg *AIConfig, jobCfg *JobConfig) (*AnalysisResult, error) {
	// 读取简历
	path := filepath.Join(os.Getenv("APPDATA"), "GoResumeReview", "resumes", resumeID+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("简历不存在: %v", err)
	}

	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return nil, fmt.Errorf("解析简历失败: %v", err)
	}

	// 更新状态为分析中
	resume.Status = "analyzing"
	a.saveResume(&resume)
	runtime.EventsEmit(a.ctx, "analysis:progress", map[string]interface{}{
		"id":       resumeID,
		"status":   "analyzing",
		"progress": 10,
	})

	// 构建 Prompt - 进度 30%
	prompt := a.buildAnalysisPrompt(&resume, jobCfg)
	runtime.EventsEmit(a.ctx, "analysis:progress", map[string]interface{}{
		"id":       resumeID,
		"status":   "analyzing",
		"progress": 30,
	})

	// 调用 AI - 进度 50%
	runtime.EventsEmit(a.ctx, "analysis:progress", map[string]interface{}{
		"id":       resumeID,
		"status":   "analyzing",
		"progress": 50,
	})
	result, err := a.callAI(cfg, prompt)
	if err != nil {
		resume.Status = "error"
		a.saveResume(&resume)
		runtime.EventsEmit(a.ctx, "analysis:error", map[string]interface{}{
			"id":    resumeID,
			"error": err.Error(),
		})
		return nil, err
	}

	// 解析 AI 返回结果 - 进度 80%
	runtime.EventsEmit(a.ctx, "analysis:progress", map[string]interface{}{
		"id":       resumeID,
		"status":   "analyzing",
		"progress": 80,
	})
	analysis, err := a.parseAnalysisResult(result)
	if err != nil {
		resume.Status = "error"
		a.saveResume(&resume)
		runtime.EventsEmit(a.ctx, "analysis:error", map[string]interface{}{
			"id":    resumeID,
			"error": "解析AI返回失败: " + err.Error(),
		})
		return nil, err
	}

	// 更新简历状态 - 进度 100%
	runtime.EventsEmit(a.ctx, "analysis:progress", map[string]interface{}{
		"id":       resumeID,
		"status":   "analyzing",
		"progress": 100,
	})
	resume.Status = "done"
	resume.Score = analysis.OverallScore
	resume.Analysis = analysis
	a.saveResume(&resume)

	// 发送完成事件
	runtime.EventsEmit(a.ctx, "analysis:completed", map[string]interface{}{
		"id":       resumeID,
		"score":    analysis.OverallScore,
		"analysis": analysis,
	})

	return analysis, nil
}

// StartBatchAnalysis 批量分析简历
func (a *App) StartBatchAnalysis(resumeIDs []string, cfg *AIConfig, jobCfg *JobConfig) {
	go func() {
		total := len(resumeIDs)
		for i, id := range resumeIDs {
			// 发送进度
			runtime.EventsEmit(a.ctx, "batch:progress", map[string]interface{}{
				"current":  i + 1,
				"total":    total,
				"resumeId": id,
			})

			_, err := a.AnalyzeResume(id, cfg, jobCfg)
			if err != nil {
				log.Printf("分析简历 %s 失败: %v", id, err)
			}

			// 防止请求过快
			time.Sleep(500 * time.Millisecond)
		}

		runtime.EventsEmit(a.ctx, "batch:completed", map[string]interface{}{
			"total": total,
		})
	}()
}

// buildAnalysisPrompt 构建分析提示词
func (a *App) buildAnalysisPrompt(resume *Resume, jobCfg *JobConfig) string {
	skills := strings.Join(jobCfg.RequiredSkills, "、")
	requirements := strings.Join(jobCfg.Requirements, "\n- ")

	prompt := fmt.Sprintf("你是一位专业的HR招聘专家，请根据以下岗位要求分析这份简历。\n\n"+
		"## 岗位信息\n"+
		"- 岗位名称: %s\n"+
		"- 工作年限要求: %d 年以上\n"+
		"- 学历要求: %s\n"+
		"- 必备技能: %s\n"+
		"- 其他要求:\n- %s\n\n"+
		"## 简历内容\n"+
		"文件名: %s\n"+
		"---\n%s\n---\n\n"+
		"## 分析要求\n"+
		"请严格按照以下JSON格式返回分析结果，不要包含其他内容：\n\n"+
		"{\n"+
		"  \"overall_score\": 85,\n"+
		"  \"skill_match\": 90,\n"+
		"  \"experience_match\": 80,\n"+
		"  \"education_match\": 85,\n"+
		"  \"recommendation\": \"strong_recommend\",\n"+
		"  \"strengths\": [\"优势1\", \"优势2\", \"优势3\"],\n"+
		"  \"weaknesses\": [\"不足1\", \"不足2\"],\n"+
		"  \"summary\": \"一句话总结候选人与岗位的匹配情况\"\n"+
		"}\n\n"+
		"评分说明：\n"+
		"- overall_score: 综合评分 (0-100)\n"+
		"- skill_match: 技能匹配度 (0-100)\n"+
		"- experience_match: 经验匹配度 (0-100)\n"+
		"- education_match: 学历匹配度 (0-100)\n"+
		"- recommendation: 推荐等级，只能是以下四种之一：\n"+
		"  - \"strong_recommend\" (强烈推荐，85分以上)\n"+
		"  - \"recommend\" (推荐，70-84分)\n"+
		"  - \"consider\" (可以考虑，60-69分)\n"+
		"  - \"not_recommend\" (不推荐，60分以下)\n\n"+
		"请确保返回的是合法的JSON格式。",
		jobCfg.Title,
		jobCfg.ExperienceYears,
		jobCfg.EducationLevel,
		skills,
		requirements,
		resume.FileName,
		a.truncateContent(resume.Content, 8000),
	)

	return prompt
}

// truncateContent 截断过长内容
func (a *App) truncateContent(content string, maxLen int) string {
	if len(content) <= maxLen {
		return content
	}
	return content[:maxLen] + "\n...(内容已截断)"
}

// callAI 调用AI接口
func (a *App) callAI(cfg *AIConfig, prompt string) (string, error) {
	reqBody := ChatRequest{
		Model: cfg.Model,
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: "你是一位专业的HR招聘专家，擅长分析简历与岗位的匹配度。请严格按照用户要求的JSON格式返回结果。",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.3,
		MaxTokens:   2000,
	}

	body, _ := json.Marshal(reqBody)
	url := strings.TrimSuffix(cfg.BaseURL, "/") + "/chat/completions"

	var lastErr error
	for attempt := 0; attempt < cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		if err != nil {
			lastErr = err
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+cfg.APIKey)

		client := &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 429 {
			lastErr = fmt.Errorf("请求过于频繁")
			continue
		}

		if resp.StatusCode >= 400 {
			lastErr = fmt.Errorf("API错误: %d - %s", resp.StatusCode, string(respBody))
			continue
		}

		var chatResp ChatResponse
		if err := json.Unmarshal(respBody, &chatResp); err != nil {
			lastErr = fmt.Errorf("解析响应失败: %v", err)
			continue
		}

		if chatResp.Error != nil {
			lastErr = fmt.Errorf("AI错误: %s", chatResp.Error.Message)
			continue
		}

		if len(chatResp.Choices) == 0 {
			lastErr = fmt.Errorf("AI未返回结果")
			continue
		}

		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("重试%d次后失败: %v", cfg.MaxRetries, lastErr)
}

// parseAnalysisResult 解析AI返回的分析结果
func (a *App) parseAnalysisResult(content string) (*AnalysisResult, error) {
	// 尝试提取JSON
	jsonStr := content

	// 如果包含```json代码块，提取其中内容
	if idx := strings.Index(content, "```json"); idx != -1 {
		start := idx + 7
		end := strings.Index(content[start:], "```")
		if end != -1 {
			jsonStr = content[start : start+end]
		}
	} else if idx := strings.Index(content, "```"); idx != -1 {
		start := idx + 3
		end := strings.Index(content[start:], "```")
		if end != -1 {
			jsonStr = content[start : start+end]
		}
	}

	// 尝试找到JSON对象
	jsonStr = strings.TrimSpace(jsonStr)
	if !strings.HasPrefix(jsonStr, "{") {
		// 尝试用正则提取
		re := regexp.MustCompile(`\{[\s\S]*\}`)
		matches := re.FindString(content)
		if matches != "" {
			jsonStr = matches
		}
	}

	var result AnalysisResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("JSON解析失败: %v, 内容: %s", err, jsonStr[:min(200, len(jsonStr))])
	}

	// 验证并修正数据
	result.OverallScore = clamp(result.OverallScore, 0, 100)
	result.SkillMatch = clamp(result.SkillMatch, 0, 100)
	result.ExperienceMatch = clamp(result.ExperienceMatch, 0, 100)
	result.EducationMatch = clamp(result.EducationMatch, 0, 100)

	// 验证推荐等级
	validRecs := map[string]bool{
		"strong_recommend": true,
		"recommend":        true,
		"consider":         true,
		"not_recommend":    true,
	}
	if !validRecs[result.Recommendation] {
		// 根据分数自动设置
		switch {
		case result.OverallScore >= 85:
			result.Recommendation = "strong_recommend"
		case result.OverallScore >= 70:
			result.Recommendation = "recommend"
		case result.OverallScore >= 60:
			result.Recommendation = "consider"
		default:
			result.Recommendation = "not_recommend"
		}
	}

	result.AnalyzedAt = time.Now().Format(time.RFC3339)

	return &result, nil
}

func clamp(value, minVal, maxVal int) int {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
