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
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
	"github.com/xuri/excelize/v2"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 版本信息
const AppVersion = "1.1.0"
const GitHubRepo = "1186258278/TalentLens"

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

// Project 招聘项目
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	JobConfig JobConfig `json:"job_config"`
	ResumeIDs []string  `json:"resume_ids"`
	Status    string    `json:"status"` // draft/analyzing/completed
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Resume 简历结构
type Resume struct {
	ID        string          `json:"id"`
	ProjectID string          `json:"project_id"`
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
	OverallScore    float64 `json:"overall_score"`
	SkillMatch      float64 `json:"skill_match"`
	ExperienceMatch float64 `json:"experience_match"`
	EducationMatch  float64 `json:"education_match"`
	Recommendation  string `json:"recommendation"`

	// 详细分析维度
	SkillDetail      string `json:"skill_detail"`
	ExperienceDetail string `json:"experience_detail"`
	EducationDetail  string `json:"education_detail"`

	// 候选人信息提取
	CandidateName string `json:"candidate_name"`
	WorkYears     string `json:"work_years"`
	Education     string `json:"education"`
	CurrentRole   string `json:"current_role"`

	// 详细评价
	Strengths  []string `json:"strengths"`
	Weaknesses []string `json:"weaknesses"`
	Risks      []string `json:"risks"`
	Summary    string   `json:"summary"`

	// 面试建议
	InterviewSuggestions []string `json:"interview_suggestions"`

	AnalyzedAt string `json:"analyzed_at"`
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
	ctx             context.Context
	config          Config
	activeProjectID string // 当前活跃的项目ID（前端设置）
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

	// 监听原生文件拖拽（Wails 提供真实文件路径）
	runtime.OnFileDrop(ctx, func(x, y int, paths []string) {
		log.Printf("[OnFileDrop] 收到 %d 个文件, 项目=%s", len(paths), a.activeProjectID)
		supportedExts := map[string]bool{
			".pdf": true, ".docx": true, ".doc": true,
			".jpg": true, ".jpeg": true, ".png": true, ".bmp": true, ".gif": true, ".webp": true,
		}
		added := 0
		for _, fp := range paths {
			ext := strings.ToLower(filepath.Ext(fp))
			if !supportedExts[ext] {
				continue
			}
			info, err := os.Stat(fp)
			if err != nil {
				continue
			}
			id := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(fp))

			if a.activeProjectID != "" {
				a.RegisterResumeToProject(a.activeProjectID, id, filepath.Base(fp), fp, ext, info.Size())
			} else {
				a.RegisterResume(id, filepath.Base(fp), fp, ext, info.Size())
			}

			// 通知前端（包含提取的内容）
			resumePath := filepath.Join(a.getDataDir(), "resumes", id+".json")
			data, _ := os.ReadFile(resumePath)
			var resume Resume
			json.Unmarshal(data, &resume)
			runtime.EventsEmit(a.ctx, "resume:dropped", &resume)
			added++
		}
		log.Printf("[OnFileDrop] 成功添加 %d 个文件", added)
	})

	log.Println("TalentLens 已启动")
}

// SetActiveProject 前端告知后端当前活跃项目
func (a *App) SetActiveProject(projectID string) {
	a.activeProjectID = projectID
	log.Printf("[SetActiveProject] 当前项目: %s", projectID)
}

// SelectResumeFiles 打开原生文件选择对话框，返回添加数量
func (a *App) SelectResumeFiles(projectID string) int {
	files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择简历文件",
		Filters: []runtime.FileFilter{
			{DisplayName: "简历文件 (PDF/Word/图片)", Pattern: "*.pdf;*.docx;*.doc;*.jpg;*.jpeg;*.png;*.bmp;*.gif;*.webp"},
			{DisplayName: "PDF 文件", Pattern: "*.pdf"},
			{DisplayName: "Word 文件", Pattern: "*.docx;*.doc"},
			{DisplayName: "图片文件", Pattern: "*.jpg;*.jpeg;*.png;*.bmp;*.gif;*.webp"},
			{DisplayName: "所有文件", Pattern: "*.*"},
		},
	})
	if err != nil || len(files) == 0 {
		return 0
	}

	count := 0
	for _, fp := range files {
		info, err := os.Stat(fp)
		if err != nil {
			continue
		}
		ext := strings.ToLower(filepath.Ext(fp))
		id := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(fp))

		if projectID != "" {
			a.RegisterResumeToProject(projectID, id, filepath.Base(fp), fp, ext, info.Size())
		} else {
			a.RegisterResume(id, filepath.Base(fp), fp, ext, info.Size())
		}

		// 通知前端
		resumePath := filepath.Join(a.getDataDir(), "resumes", id+".json")
		data, _ := os.ReadFile(resumePath)
		var resume Resume
		json.Unmarshal(data, &resume)
		runtime.EventsEmit(a.ctx, "resume:dropped", &resume)
		count++
	}

	log.Printf("[SelectResumeFiles] 选择了 %d 个文件", count)
	return count
}

// getDataDir 获取固定数据存储目录
// Windows: %USERPROFILE%/Documents/TalentLens
// macOS:   ~/Documents/TalentLens
// Linux:   ~/Documents/TalentLens
func (a *App) getDataDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = os.Getenv("HOME")
		if homeDir == "" {
			homeDir = os.Getenv("USERPROFILE")
		}
	}
	dir := filepath.Join(homeDir, "Documents", "TalentLens")
	os.MkdirAll(dir, 0755)
	return dir
}

// GetDataDir 暴露数据目录路径给前端
func (a *App) GetDataDir() string {
	return a.getDataDir()
}

// OpenDataDir 用系统文件管理器打开数据目录
func (a *App) OpenDataDir() {
	dir := a.getDataDir()
	runtime.BrowserOpenURL(a.ctx, dir)
}

func (a *App) getConfigPath() string {
	return filepath.Join(a.getDataDir(), "config.json")
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
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".txt", ".md":
		data, err := os.ReadFile(filePath)
		if err != nil {
			return ""
		}
		return a.truncateContent(string(data), 50000)
	case ".pdf":
		content := a.extractFromPDF(filePath)
		if content != "" {
			return content
		}
		// PDF 库提取失败时回退到原始方式
		log.Println("[extractText] PDF 库提取失败，回退到原始方式")
		data, err := os.ReadFile(filePath)
		if err != nil {
			return ""
		}
		return a.extractFromBytes(data)
	default:
		data, err := os.ReadFile(filePath)
		if err != nil {
			return ""
		}
		return a.extractFromBytes(data)
	}
}

// extractFromPDF 使用 ledongthuc/pdf 库提取 PDF 文本（支持中文）
func (a *App) extractFromPDF(filePath string) string {
	f, r, err := pdf.Open(filePath)
	if err != nil {
		log.Printf("[extractFromPDF] 打开 PDF 失败: %v", err)
		return ""
	}
	defer f.Close()

	var buf bytes.Buffer
	reader, err := r.GetPlainText()
	if err != nil {
		log.Printf("[extractFromPDF] 提取文本失败: %v", err)
		return ""
	}
	buf.ReadFrom(reader)

	content := strings.TrimSpace(buf.String())
	if content == "" {
		log.Printf("[extractFromPDF] PDF 提取结果为空: %s", filePath)
		return ""
	}

	// 清理多余空白行
	lines := strings.Split(content, "\n")
	var cleaned []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}
	result := strings.Join(cleaned, "\n")

	log.Printf("[extractFromPDF] 提取成功: %s, 长度=%d 字符", filepath.Base(filePath), len(result))
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
	dir := filepath.Join(a.getDataDir(), "resumes")
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
	dir := filepath.Join(a.getDataDir(), "resumes")
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
	path := filepath.Join(a.getDataDir(), "resumes", id+".json")
	return os.Remove(path)
}

func (a *App) ReAnalyzeResume(id string) error {
	path := filepath.Join(a.getDataDir(), "resumes", id+".json")
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
	dir := filepath.Join(a.getDataDir(), "resumes")
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

// GetFreshResumeContent 重新从原始文件提取内容并返回（同时更新缓存）
func (a *App) GetFreshResumeContent(id string) (string, error) {
	path := filepath.Join(a.getDataDir(), "resumes", id+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("简历不存在")
	}

	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return "", fmt.Errorf("解析失败")
	}

	// 重新从原始文件提取
	if resume.FilePath != "" && resume.FilePath != resume.FileName {
		freshContent := a.extractText(resume.FilePath)
		if freshContent != "" && len(freshContent) > 10 {
			resume.Content = freshContent
			a.saveResume(&resume) // 更新磁盘缓存
			log.Printf("[GetFreshResumeContent] 重新提取成功: %s, 长度=%d", resume.FileName, len(freshContent))
			return freshContent, nil
		}
	}

	// 回退到已有内容
	return resume.Content, nil
}

// ============================================
// 招聘项目管理
// ============================================

func (a *App) getProjectsDir() string {
	dir := filepath.Join(a.getDataDir(), "projects")
	os.MkdirAll(dir, 0755)
	return dir
}

func (a *App) saveProject(p *Project) {
	data, _ := json.MarshalIndent(p, "", "  ")
	os.WriteFile(filepath.Join(a.getProjectsDir(), p.ID+".json"), data, 0644)
}

// CreateProject 创建招聘项目
func (a *App) CreateProject(name string, jobCfg *JobConfig) *Project {
	p := &Project{
		ID:        fmt.Sprintf("proj_%d", time.Now().UnixNano()),
		Name:      name,
		JobConfig: *jobCfg,
		ResumeIDs: []string{},
		Status:    "draft",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	a.saveProject(p)
	log.Printf("[CreateProject] 创建项目: %s (%s)", p.Name, p.ID)
	return p
}

// GetProjects 获取所有项目列表
func (a *App) GetProjects() []*Project {
	dir := a.getProjectsDir()
	entries, _ := os.ReadDir(dir)
	var projects []*Project
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}
		var p Project
		if json.Unmarshal(data, &p) == nil {
			projects = append(projects, &p)
		}
	}
	// 按更新时间倒序
	for i := 0; i < len(projects); i++ {
		for j := i + 1; j < len(projects); j++ {
			if projects[j].UpdatedAt.After(projects[i].UpdatedAt) {
				projects[i], projects[j] = projects[j], projects[i]
			}
		}
	}
	return projects
}

// GetProject 获取单个项目
func (a *App) GetProject(id string) *Project {
	data, err := os.ReadFile(filepath.Join(a.getProjectsDir(), id+".json"))
	if err != nil {
		return nil
	}
	var p Project
	json.Unmarshal(data, &p)
	return &p
}

// UpdateProject 更新项目
func (a *App) UpdateProject(p *Project) error {
	p.UpdatedAt = time.Now()
	a.saveProject(p)
	return nil
}

// DeleteProject 删除项目及其关联简历
func (a *App) DeleteProject(id string) error {
	p := a.GetProject(id)
	if p != nil {
		// 删除关联简历
		for _, rid := range p.ResumeIDs {
			a.DeleteResume(rid)
		}
	}
	return os.Remove(filepath.Join(a.getProjectsDir(), id+".json"))
}

// GetProjectResumes 获取项目下的所有简历
func (a *App) GetProjectResumes(projectID string) []*Resume {
	p := a.GetProject(projectID)
	if p == nil {
		return nil
	}
	var resumes []*Resume
	for _, rid := range p.ResumeIDs {
		path := filepath.Join(a.getDataDir(), "resumes", rid+".json")
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var r Resume
		if json.Unmarshal(data, &r) == nil {
			resumes = append(resumes, &r)
		}
	}
	return resumes
}

// GetProjectRanking 获取项目排名（按分数降序）
func (a *App) GetProjectRanking(projectID string) []*Resume {
	resumes := a.GetProjectResumes(projectID)
	// 按分数降序排列
	for i := 0; i < len(resumes); i++ {
		for j := i + 1; j < len(resumes); j++ {
			if resumes[j].Score > resumes[i].Score {
				resumes[i], resumes[j] = resumes[j], resumes[i]
			}
		}
	}
	return resumes
}

// GetProjectStats 获取项目统计信息
func (a *App) GetProjectStats(projectID string) map[string]interface{} {
	resumes := a.GetProjectResumes(projectID)
	total := len(resumes)
	analyzed := 0
	recommended := 0
	totalScore := 0
	maxScore := 0

	for _, r := range resumes {
		if r.Status == "done" {
			analyzed++
			totalScore += r.Score
			if r.Score > maxScore {
				maxScore = r.Score
			}
			if r.Analysis != nil && (r.Analysis.Recommendation == "strong_recommend" || r.Analysis.Recommendation == "recommend") {
				recommended++
			}
		}
	}

	avgScore := 0
	if analyzed > 0 {
		avgScore = totalScore / analyzed
	}

	return map[string]interface{}{
		"total":       total,
		"analyzed":    analyzed,
		"avgScore":    avgScore,
		"maxScore":    maxScore,
		"recommended": recommended,
	}
}

// RegisterResumeToProject 注册简历到项目
func (a *App) RegisterResumeToProject(projectID string, id string, fileName string, filePath string, fileType string, fileSize int64) (bool, string) {
	log.Printf("[RegisterResumeToProject] proj=%s, file=%s", projectID, fileName)

	content := ""
	if filePath != "" && filePath != fileName {
		content = a.extractText(filePath)
	}
	if content == "" {
		content = fmt.Sprintf("[简历文件: %s, 类型: %s, 大小: %d bytes]", fileName, fileType, fileSize)
	}

	resume := &Resume{
		ID:        id,
		ProjectID: projectID,
		FileName:  fileName,
		FilePath:  filePath,
		FileType:  fileType,
		FileSize:  fileSize,
		Content:   content,
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	a.saveResume(resume)

	// 更新项目的简历列表
	p := a.GetProject(projectID)
	if p != nil {
		p.ResumeIDs = append(p.ResumeIDs, id)
		a.UpdateProject(p)
	}

	return true, "简历已注册: " + fileName
}

// ImportResumesToProject 批量导入简历文件到项目
func (a *App) ImportResumesToProject(projectID string, filePaths []string) (int, error) {
	count := 0
	supportedExts := map[string]bool{
		".pdf": true, ".docx": true, ".doc": true,
		".jpg": true, ".jpeg": true, ".png": true, ".bmp": true, ".gif": true, ".webp": true,
	}

	for _, fp := range filePaths {
		ext := strings.ToLower(filepath.Ext(fp))
		if !supportedExts[ext] {
			continue
		}
		info, err := os.Stat(fp)
		if err != nil {
			continue
		}
		id := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(fp))
		a.RegisterResumeToProject(projectID, id, filepath.Base(fp), fp, ext, info.Size())
		count++
	}
	return count, nil
}

// StartProjectAnalysis 对项目中所有待分析的简历进行批量分析
func (a *App) StartProjectAnalysis(projectID string, cfg *AIConfig) {
	if cfg == nil || cfg.APIKey == "" {
		log.Println("[StartProjectAnalysis] AI 未配置，终止")
		runtime.EventsEmit(a.ctx, "analysis:error", map[string]interface{}{
			"id":    "",
			"error": "AI 未配置: 请先在设置中填写 API Key",
		})
		return
	}
	p := a.GetProject(projectID)
	if p == nil {
		return
	}

	go func() {
		p.Status = "analyzing"
		a.saveProject(p)

		resumes := a.GetProjectResumes(projectID)
		var pendingIDs []string
		for _, r := range resumes {
			if r.Status == "pending" || r.Status == "error" {
				pendingIDs = append(pendingIDs, r.ID)
			}
		}

		total := len(pendingIDs)
		for i, id := range pendingIDs {
			runtime.EventsEmit(a.ctx, "batch:progress", map[string]interface{}{
				"current":   i + 1,
				"total":     total,
				"resumeId":  id,
				"projectId": projectID,
			})

			_, err := a.AnalyzeResume(id, cfg, &p.JobConfig)
			if err != nil {
				log.Printf("分析简历 %s 失败: %v", id, err)
			}
			time.Sleep(500 * time.Millisecond)
		}

		p.Status = "completed"
		a.saveProject(p)

		runtime.EventsEmit(a.ctx, "batch:completed", map[string]interface{}{
			"total":     total,
			"projectId": projectID,
		})
	}()
}

// MigrateExistingResumes 将现有简历迁移到默认项目
func (a *App) MigrateExistingResumes() string {
	resumes := a.GetResumes()
	if len(resumes) == 0 {
		return ""
	}

	// 检查是否已有项目
	projects := a.GetProjects()
	if len(projects) > 0 {
		return projects[0].ID
	}

	// 创建默认项目
	jobCfg := &a.config.Job
	p := a.CreateProject("默认项目", jobCfg)

	for _, r := range resumes {
		r.ProjectID = p.ID
		a.saveResume(r)
		p.ResumeIDs = append(p.ResumeIDs, r.ID)
	}
	a.saveProject(p)
	log.Printf("[MigrateExistingResumes] 迁移 %d 份简历到默认项目", len(resumes))
	return p.ID
}

// ExportProjectReport 导出项目分析报告为 Excel
func (a *App) ExportProjectReport(projectID string) (string, error) {
	p := a.GetProject(projectID)
	if p == nil {
		return "", fmt.Errorf("项目不存在")
	}

	resumes := a.GetProjectRanking(projectID)
	if len(resumes) == 0 {
		return "", fmt.Errorf("项目中没有简历")
	}

	f := excelize.NewFile()
	sheet := "候选人排名"
	f.SetSheetName("Sheet1", sheet)

	// 表头
	headers := []string{"排名", "姓名", "文件名", "综合分", "技能匹配", "经验匹配", "学历匹配", "推荐等级", "优势", "不足", "风险", "总结"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// 表头样式
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 11, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"007AFF"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	f.SetRowStyle(sheet, 1, 1, headerStyle)

	// 数据行
	recMap := map[string]string{
		"strong_recommend": "强烈推荐",
		"recommend":        "推荐",
		"consider":         "可考虑",
		"not_recommend":    "不推荐",
	}

	for i, r := range resumes {
		row := i + 2
		name := r.FileName
		strengths := ""
		weaknesses := ""
		risks := ""
		summary := ""
		rec := ""

		if r.Analysis != nil {
			if r.Analysis.CandidateName != "" {
				name = r.Analysis.CandidateName
			}
			strengths = strings.Join(r.Analysis.Strengths, "\n")
			weaknesses = strings.Join(r.Analysis.Weaknesses, "\n")
			risks = strings.Join(r.Analysis.Risks, "\n")
			summary = r.Analysis.Summary
			if v, ok := recMap[r.Analysis.Recommendation]; ok {
				rec = v
			} else {
				rec = r.Analysis.Recommendation
			}
		}

		rowData := []interface{}{
			i + 1, name, r.FileName, r.Score,
			0, 0, 0, rec, strengths, weaknesses, risks, summary,
		}
		if r.Analysis != nil {
			rowData[4] = r.Analysis.SkillMatch
			rowData[5] = r.Analysis.ExperienceMatch
			rowData[6] = r.Analysis.EducationMatch
		}

		for j, val := range rowData {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, val)
		}
	}

	// 设置列宽
	colWidths := map[string]float64{"A": 6, "B": 12, "C": 25, "D": 8, "E": 10, "F": 10, "G": 10, "H": 10, "I": 30, "J": 30, "K": 25, "L": 40}
	for col, w := range colWidths {
		f.SetColWidth(sheet, col, col, w)
	}

	// 保存
	outDir := filepath.Join(a.getDataDir(), "exports")
	os.MkdirAll(outDir, 0755)
	fileName := fmt.Sprintf("%s_排名报告_%s.xlsx", p.Name, time.Now().Format("20060102_150405"))
	outPath := filepath.Join(outDir, fileName)
	if err := f.SaveAs(outPath); err != nil {
		return "", fmt.Errorf("保存失败: %v", err)
	}

	log.Printf("[ExportProjectReport] 导出成功: %s", outPath)
	return outPath, nil
}

// OpenExportDir 打开导出目录
func (a *App) OpenExportDir() {
	dir := filepath.Join(a.getDataDir(), "exports")
	os.MkdirAll(dir, 0755)
	runtime.BrowserOpenURL(a.ctx, dir)
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
	// 校验 AI 配置
	if cfg == nil || cfg.APIKey == "" {
		return nil, fmt.Errorf("AI 未配置: 请先在设置中填写 API Key")
	}
	if cfg.BaseURL == "" {
		return nil, fmt.Errorf("AI 未配置: 请先在设置中填写 Base URL")
	}

	// 读取简历
	path := filepath.Join(a.getDataDir(), "resumes", resumeID+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("简历不存在: %v", err)
	}

	var resume Resume
	if err := json.Unmarshal(data, &resume); err != nil {
		return nil, fmt.Errorf("解析简历失败: %v", err)
	}

	// 每次分析前重新提取文件内容（避免使用旧解析器缓存的错误内容）
	if resume.FilePath != "" && resume.FilePath != resume.FileName {
		freshContent := a.extractText(resume.FilePath)
		if freshContent != "" && len(freshContent) > 20 {
			log.Printf("[AnalyzeResume] 重新提取内容: %s, 长度=%d", resume.FileName, len(freshContent))
			resume.Content = freshContent
			a.saveResume(&resume) // 更新磁盘缓存
		} else {
			log.Printf("[AnalyzeResume] 重新提取失败或内容过短，使用已有内容")
		}
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
	resume.Score = int(math.Round(analysis.OverallScore))
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
	if cfg == nil || cfg.APIKey == "" {
		log.Println("[StartBatchAnalysis] AI 未配置，终止")
		runtime.EventsEmit(a.ctx, "analysis:error", map[string]interface{}{
			"id":    "",
			"error": "AI 未配置: 请先在设置中填写 API Key",
		})
		return
	}
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

	systemPrompt := "你是一位拥有15年经验的资深人力资源专家和猎头顾问。你的专长是精准评估候选人与岗位的匹配度。\n" +
		"你必须基于简历中的客观事实进行分析，不得凭空臆造简历中没有的信息。\n" +
		"你的评分必须严谨且前后一致，遵循统一的评分标准。\n" +
		"你的分析要全面、专业、有深度，就像撰写一份正式的候选人评估报告。"

	userPrompt := fmt.Sprintf(
		"## 招聘岗位信息\n"+
			"- 岗位名称: %s\n"+
			"- 最低工作年限: %d 年\n"+
			"- 最低学历: %s\n"+
			"- 核心必备技能: %s\n"+
			"- 补充要求:\n- %s\n\n"+
			"## 候选人简历\n"+
			"文件名: %s\n"+
			"```\n%s\n```\n\n"+
			"## 分析任务\n\n"+
			"请对这份简历进行全方位、深度的专业分析。\n\n"+
			"### 评分标准（严格执行）\n\n"+
			"**技能匹配度 (skill_match)**:\n"+
			"- 90-100: 完全掌握所有核心技能，且有相关高级技能加分\n"+
			"- 70-89: 掌握大部分核心技能，个别技能有欠缺\n"+
			"- 50-69: 掌握部分核心技能，有较大技能差距\n"+
			"- 0-49: 核心技能严重不足\n\n"+
			"**经验匹配度 (experience_match)**:\n"+
			"- 90-100: 工作年限超过要求，且有高度相关的项目经验\n"+
			"- 70-89: 工作年限满足要求，有相关经验\n"+
			"- 50-69: 工作年限略不足，或经验相关度不高\n"+
			"- 0-49: 工作年限严重不足，或无相关经验\n\n"+
			"**学历匹配度 (education_match)**:\n"+
			"- 90-100: 学历超过要求，且专业高度对口\n"+
			"- 70-89: 学历满足要求，专业相关\n"+
			"- 50-69: 学历勉强满足，专业有一定偏差\n"+
			"- 0-49: 学历不满足要求\n\n"+
			"**综合评分 (overall_score)** = skill_match * 0.45 + experience_match * 0.35 + education_match * 0.20\n\n"+
			"### 推荐等级（根据综合评分）\n"+
			"- \"strong_recommend\": 综合分 >= 85，各单项均 >= 70\n"+
			"- \"recommend\": 综合分 70-84\n"+
			"- \"consider\": 综合分 55-69\n"+
			"- \"not_recommend\": 综合分 < 55\n\n"+
			"### 输出要求\n\n"+
			"请严格按以下JSON格式输出，不要输出任何其他内容：\n\n"+
			"```json\n"+
			"{\n"+
			"  \"candidate_name\": \"从简历中提取的姓名\",\n"+
			"  \"work_years\": \"从简历中提取的工作年限，如 5年\",\n"+
			"  \"education\": \"从简历中提取的最高学历和学校，如 本科-武汉大学-计算机科学\",\n"+
			"  \"current_role\": \"从简历中提取的当前/最近职位，如 高级Go开发工程师@字节跳动\",\n"+
			"  \"overall_score\": 78,\n"+
			"  \"skill_match\": 82,\n"+
			"  \"experience_match\": 75,\n"+
			"  \"education_match\": 80,\n"+
			"  \"skill_detail\": \"逐项说明每个核心技能的掌握情况，如：Go(精通，有3年生产经验)、MySQL(熟练，简历中有分库分表经验)、Redis(了解，未提及具体使用场景)\",\n"+
			"  \"experience_detail\": \"详细分析工作经历与岗位的匹配程度，包括行业相关度、项目复杂度、职责范围等\",\n"+
			"  \"education_detail\": \"分析学历背景、专业对口程度、是否有相关认证或培训\",\n"+
			"  \"recommendation\": \"recommend\",\n"+
			"  \"strengths\": [\n"+
			"    \"具体的优势1（必须引用简历中的事实依据）\",\n"+
			"    \"具体的优势2\",\n"+
			"    \"具体的优势3\"\n"+
			"  ],\n"+
			"  \"weaknesses\": [\n"+
			"    \"具体的不足1（必须基于岗位要求指出差距）\",\n"+
			"    \"具体的不足2\"\n"+
			"  ],\n"+
			"  \"risks\": [\n"+
			"    \"潜在风险或需关注事项，如频繁跳槽、职业路径不连贯等\"\n"+
			"  ],\n"+
			"  \"interview_suggestions\": [\n"+
			"    \"如果进入面试环节，建议重点考察的问题或方向\"\n"+
			"  ],\n"+
			"  \"summary\": \"2-3句话全面总结该候选人：包括核心亮点、主要短板、综合判断。需要具体且专业，避免空泛表述。\"\n"+
			"}\n"+
			"```\n\n"+
			"注意事项：\n"+
			"1. 所有分析必须基于简历中的客观内容，不得编造简历中不存在的信息\n"+
			"2. strengths 至少3条，weaknesses 至少2条，每条都要具体且有事实依据\n"+
			"3. summary 不能笼统，要结合候选人的具体情况给出有价值的判断\n"+
			"4. 如果简历信息不完整或模糊，请在分析中明确指出\n"+
			"5. 确保返回合法的JSON格式",
		jobCfg.Title,
		jobCfg.ExperienceYears,
		jobCfg.EducationLevel,
		skills,
		requirements,
		resume.FileName,
		a.truncateContent(resume.Content, 10000),
	)

	// 使用 system + user 消息格式
	return systemPrompt + "\n\n---\n\n" + userPrompt
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
	// 拆分 system prompt 和 user prompt
	parts := strings.SplitN(prompt, "\n\n---\n\n", 2)
	systemMsg := parts[0]
	userMsg := prompt
	if len(parts) == 2 {
		userMsg = parts[1]
	}

	reqBody := ChatRequest{
		Model: cfg.Model,
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: systemMsg,
			},
			{
				Role:    "user",
				Content: userMsg,
			},
		},
		Temperature: 0.2,
		MaxTokens:   4000,
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

	// 验证并修正数据：四舍五入并限制在 0-100 范围
	result.OverallScore = clampFloat(math.Round(result.OverallScore), 0, 100)
	result.SkillMatch = clampFloat(math.Round(result.SkillMatch), 0, 100)
	result.ExperienceMatch = clampFloat(math.Round(result.ExperienceMatch), 0, 100)
	result.EducationMatch = clampFloat(math.Round(result.EducationMatch), 0, 100)

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

func clampFloat(value, minVal, maxVal float64) float64 {
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

// ============================================
// 版本检测与链接
// ============================================

// GetAppVersion 返回当前应用版本号
func (a *App) GetAppVersion() string {
	return AppVersion
}

// OpenURL 用系统浏览器打开链接
func (a *App) OpenURL(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}

// CheckForUpdate 检查 GitHub 是否有新版本
func (a *App) CheckForUpdate() map[string]interface{} {
	result := map[string]interface{}{
		"hasUpdate":      false,
		"currentVersion": AppVersion,
		"latestVersion":  AppVersion,
		"releaseURL":     "",
		"error":          "",
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", GitHubRepo)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		result["error"] = fmt.Sprintf("网络错误: %v", err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		result["error"] = fmt.Sprintf("GitHub API 返回 %d", resp.StatusCode)
		return result
	}

	body, _ := io.ReadAll(resp.Body)
	var release struct {
		TagName string `json:"tag_name"`
		HTMLURL string `json:"html_url"`
		Name    string `json:"name"`
	}
	if err := json.Unmarshal(body, &release); err != nil {
		result["error"] = "解析版本信息失败"
		return result
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")
	result["latestVersion"] = latestVersion
	result["releaseURL"] = release.HTMLURL

	if compareVersions(latestVersion, AppVersion) > 0 {
		result["hasUpdate"] = true
	}

	log.Printf("[CheckForUpdate] 当前=%s, 最新=%s, 有更新=%v", AppVersion, latestVersion, result["hasUpdate"])
	return result
}

// compareVersions 比较语义化版本号，返回 1(a>b), 0(a==b), -1(a<b)
func compareVersions(a, b string) int {
	partsA := strings.Split(a, ".")
	partsB := strings.Split(b, ".")

	maxLen := len(partsA)
	if len(partsB) > maxLen {
		maxLen = len(partsB)
	}

	for i := 0; i < maxLen; i++ {
		var numA, numB int
		if i < len(partsA) {
			fmt.Sscanf(partsA[i], "%d", &numA)
		}
		if i < len(partsB) {
			fmt.Sscanf(partsB[i], "%d", &numB)
		}
		if numA > numB {
			return 1
		}
		if numA < numB {
			return -1
		}
	}
	return 0
}
