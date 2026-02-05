import { defineStore } from 'pinia'
import { ref, computed, onMounted } from 'vue'

// 尝试导入 Wails 运行时（可能不存在于纯前端开发模式）
let WailsRuntime: any = null
let WailsApp: any = null
let isWailsEnv = false

// 异步加载 Wails 绑定
async function loadWailsBindings() {
  try {
    WailsRuntime = await import('../../wailsjs/runtime/runtime')
    WailsApp = await import('../../wailsjs/go/main/App')
    isWailsEnv = true
    console.log('✅ Wails 环境已检测到')
  } catch (e) {
    console.log('⚠️ 非 Wails 环境，使用 Mock 模式')
    isWailsEnv = false
  }
}

// 文件信息接口（从 DropZone 传入）
export interface FileInfo {
  name: string
  path: string
  size: number
  type: string
  lastModified: number
}

// 简历类型定义
export interface Resume {
  id: string
  fileName: string
  filePath: string
  fileType: string
  fileSize: number
  content?: string
  status: 'pending' | 'analyzing' | 'done' | 'error'
  score?: number
  analysis?: {
    overallScore: number
    experienceMatch: number
    skillMatch: number
    educationMatch: number
    skillDetail: string
    experienceDetail: string
    educationDetail: string
    candidateName: string
    workYears: string
    education: string
    currentRole: string
    summary: string
    strengths: string[]
    weaknesses: string[]
    risks: string[]
    recommendation: string
    interviewSuggestions: string[]
  }
  createdAt: string
  analyzedAt?: string
}

// 开发者日志条目
export interface DevLogEntry {
  time: string
  level: 'info' | 'warn' | 'error'
  message: string
}

// 开发者日志缓冲区（全局共享）
const MAX_LOG_ENTRIES = 200
const _devLogs: DevLogEntry[] = []
let _devLogListeners: Array<() => void> = []

export function getDevLogs(): DevLogEntry[] {
  return _devLogs
}

export function onDevLogChange(fn: () => void) {
  _devLogListeners.push(fn)
  return () => { _devLogListeners = _devLogListeners.filter(f => f !== fn) }
}

function devLog(level: DevLogEntry['level'], message: string) {
  const entry: DevLogEntry = {
    time: new Date().toLocaleTimeString('zh-CN', { hour12: false, hour: '2-digit', minute: '2-digit', second: '2-digit', fractionalSecondDigits: 3 } as any),
    level,
    message
  }
  _devLogs.push(entry)
  if (_devLogs.length > MAX_LOG_ENTRIES) _devLogs.shift()
  _devLogListeners.forEach(fn => fn())
  
  // 同时输出到控制台
  const prefix = level === 'error' ? '❌' : level === 'warn' ? '⚠️' : '📋'
  console.log(`${prefix} [${entry.time}] ${message}`)
}

// 批量分析进度
export interface BatchProgress {
  current: number
  total: number
  currentResumeId: string | null
}

// 使用 Pinia 创建单例 Store
export const useResumeStore = defineStore('resume', () => {
  // 状态
  const resumes = ref<Resume[]>([])
  const isAnalyzing = ref(false)
  const selectedId = ref<string | null>(null)
  
  // 分析进度状态
  const batchProgress = ref<BatchProgress>({ current: 0, total: 0, currentResumeId: null })

  // 计算属性
  const selectedResume = computed(() => {
    if (!selectedId.value) return null
    return resumes.value.find(r => r.id === selectedId.value) || null
  })

  const pendingCount = computed(() => {
    return resumes.value.filter(r => r.status === 'pending').length
  })

  const doneCount = computed(() => {
    return resumes.value.filter(r => r.status === 'done').length
  })

  // 添加简历（接受 FileInfo 数组）
  async function addResumes(files: FileInfo[]) {
    for (const file of files) {
      // 检查是否已存在相同文件名的简历
      if (resumes.value.some(r => r.fileName === file.name && r.fileSize === file.size)) {
        devLog('warn', `跳过重复文件: ${file.name}`)
        continue
      }

      const ext = '.' + file.name.split('.').pop()?.toLowerCase()
      const id = Date.now().toString() + Math.random().toString(36).substring(2, 11)

      const newResume: Resume = {
        id,
        fileName: file.name,
        filePath: file.path,
        fileType: ext,
        fileSize: file.size,
        status: 'pending',
        createdAt: new Date().toISOString()
      }

      resumes.value.unshift(newResume)
      devLog('info', `添加简历: ${file.name} (id: ${id})`)

      // 在 Wails 环境下同步到后端磁盘
      if (isWailsEnv && WailsApp) {
        try {
          const result = await WailsApp.RegisterResume(id, file.name, file.path, ext, file.size)
          if (Array.isArray(result)) {
            devLog('info', `后端注册: ${result[1]}`)
          } else {
            devLog('info', `后端注册完成`)
          }
        } catch (err: any) {
          devLog('error', `后端注册失败: ${err.message || err}`)
        }
      }
    }
  }

  // 删除简历
  function deleteResume(id: string) {
    const index = resumes.value.findIndex(r => r.id === id)
    if (index !== -1) {
      resumes.value.splice(index, 1)
      // 如果删除的是当前选中的，清除选中状态
      if (selectedId.value === id) {
        selectedId.value = null
      }
    }
  }

  // 选择简历
  function selectResume(id: string | null) {
    selectedId.value = id
  }

  // 重新分析
  function reAnalyze(id: string) {
    const resume = resumes.value.find(r => r.id === id)
    if (resume) {
      resume.status = 'pending'
      resume.score = undefined
      resume.analysis = undefined

      // 模拟分析过程
      setTimeout(() => {
        resume.status = 'analyzing'
        setTimeout(() => {
          resume.status = 'done'
          resume.score = Math.floor(Math.random() * 30) + 70
          resume.analysis = {
            overallScore: resume.score,
            experienceMatch: Math.floor(Math.random() * 20) + 80,
            skillMatch: Math.floor(Math.random() * 25) + 75,
            educationMatch: Math.floor(Math.random() * 15) + 85,
            summary: 'AI分析完成',
            strengths: ['技能匹配', '经验丰富'],
            weaknesses: ['可提升空间'],
            recommendation: 'recommend'
          }
          resume.analyzedAt = new Date().toISOString()
        }, 2000)
      }, 500)
    }
  }

  // 模拟的分析结果模板
  const mockStrengths = [
    '技术栈与岗位要求高度匹配',
    '有丰富的项目经验',
    '学历背景符合要求',
    '具备良好的沟通能力',
    '有团队管理经验',
    '掌握核心技能',
    '工作经验丰富',
    '有相关行业背景'
  ]

  const mockWeaknesses = [
    '部分技能需要加强',
    '项目经验略显不足',
    '缺少某些加分技能',
    '工作年限稍短',
    '缺乏大型项目经验',
    '需要补充相关认证'
  ]

  // 获取 AI 配置
  function getAIConfig() {
    const saved = localStorage.getItem('goresume_settings')
    if (!saved) return null
    try {
      const settings = JSON.parse(saved)
      return {
        provider: settings.ai?.provider || 'deepseek',
        base_url: settings.ai?.baseURL || 'https://api.deepseek.com/v1',
        api_key: settings.ai?.apiKey || '',
        model: settings.ai?.model || 'deepseek-chat',
        max_retries: 3,
        timeout: 60
      }
    } catch {
      return null
    }
  }

  // 获取岗位配置
  function getJobConfig() {
    const saved = localStorage.getItem('goresume_settings')
    if (!saved) {
      return {
        title: '高级开发工程师',
        requirements: ['熟悉相关技术栈', '有团队协作经验'],
        required_skills: ['开发', '设计', '测试'],
        experience_years: 3,
        education_level: '本科'
      }
    }
    try {
      const settings = JSON.parse(saved)
      return {
        title: settings.job?.title || '高级开发工程师',
        requirements: settings.job?.requirements || ['熟悉相关技术栈'],
        required_skills: settings.job?.requiredSkills || ['开发'],
        experience_years: settings.job?.experienceYears || 3,
        education_level: settings.job?.educationLevel || '本科'
      }
    } catch {
      return {
        title: '高级开发工程师',
        requirements: [],
        required_skills: [],
        experience_years: 3,
        education_level: '本科'
      }
    }
  }

  // 开始分析所有待分析的简历
  async function startAnalysis() {
    if (isAnalyzing.value) return

    isAnalyzing.value = true

    const pendingResumes = resumes.value.filter(r => r.status === 'pending')

    // 如果在 Wails 环境下，使用后端 API
    if (isWailsEnv && WailsApp) {
      const aiConfig = getAIConfig()
      const jobConfig = getJobConfig()

      if (!aiConfig || !aiConfig.api_key) {
        devLog('error', 'AI 配置无效或缺少 API Key')
        isAnalyzing.value = false
        return
      }

      // 使用 Wails 后端批量分析
      const resumeIds = pendingResumes.map(r => r.id)
      devLog('info', `启动后端批量分析: ${resumeIds.length} 份简历, ids=${resumeIds.join(',')}`)
      try {
        await WailsApp.StartBatchAnalysis(resumeIds, aiConfig, jobConfig)
        devLog('info', '后端 StartBatchAnalysis 调用成功')
      } catch (err: any) {
        devLog('error', `启动批量分析失败: ${err.message || err}`)
        isAnalyzing.value = false
      }
      // 注意：isAnalyzing 会在接收到 batch:completed 事件时设置为 false
      return
    }

    // 非 Wails 环境，使用 Mock 模式
    const total = pendingResumes.length
    for (let i = 0; i < pendingResumes.length; i++) {
      const resume = pendingResumes[i]
      
      // 更新批量进度
      batchProgress.value = { current: i + 1, total, currentResumeId: resume.id }
      
      resume.status = 'analyzing'
      ;(resume as any).progress = 10

      // 模拟进度阶段
      await new Promise(resolve => setTimeout(resolve, 500))
      ;(resume as any).progress = 30
      
      await new Promise(resolve => setTimeout(resolve, 800))
      ;(resume as any).progress = 70
      
      await new Promise(resolve => setTimeout(resolve, 500))
      ;(resume as any).progress = 100

      resume.status = 'done'
      resume.score = Math.floor(Math.random() * 35) + 60
      
      const expMatch = Math.floor(Math.random() * 25) + 70
      const skillMatch = Math.floor(Math.random() * 30) + 65
      const eduMatch = Math.floor(Math.random() * 20) + 75

      // 随机选择优势和不足
      const shuffledStrengths = [...mockStrengths].sort(() => Math.random() - 0.5)
      const shuffledWeaknesses = [...mockWeaknesses].sort(() => Math.random() - 0.5)

      resume.analysis = {
        overallScore: resume.score,
        experienceMatch: expMatch,
        skillMatch: skillMatch,
        educationMatch: eduMatch,
        skillDetail: '技能匹配分析详情 (Mock 模式)',
        experienceDetail: '经验匹配分析详情 (Mock 模式)',
        educationDetail: '学历匹配分析详情 (Mock 模式)',
        candidateName: resume.fileName.replace(/\.[^/.]+$/, ''),
        workYears: '3年',
        education: '本科',
        currentRole: '开发工程师',
        summary: `该候选人（${resume.fileName.replace(/\.[^/.]+$/, '')}）整体表现${resume.score >= 80 ? '优秀' : resume.score >= 70 ? '良好' : '一般'}。技能匹配度${skillMatch}%，工作经验匹配度${expMatch}%。${resume.score >= 75 ? '建议安排面试进一步了解。' : '可根据实际情况考虑是否进入下一轮。'}`,
        strengths: shuffledStrengths.slice(0, 2 + Math.floor(Math.random() * 2)),
        weaknesses: shuffledWeaknesses.slice(0, 1 + Math.floor(Math.random() * 2)),
        risks: ['Mock 模式暂无风险分析'],
        recommendation: resume.score >= 85 ? 'strong_recommend' : 
                       resume.score >= 70 ? 'recommend' : 
                       resume.score >= 55 ? 'consider' : 'not_recommend',
        interviewSuggestions: ['Mock 模式暂无面试建议']
      }
      resume.analyzedAt = new Date().toISOString()
    }

    isAnalyzing.value = false
    batchProgress.value = { current: 0, total: 0, currentResumeId: null }
  }

  // 清空所有简历
  function clearAll() {
    resumes.value = []
    selectedId.value = null
  }

  // 初始化 Wails 事件监听
  async function initWailsEvents() {
    await loadWailsBindings()
    
    if (!isWailsEnv || !WailsRuntime) {
      console.log('⚠️ 跳过 Wails 事件监听初始化')
      return
    }

    // 监听简历添加事件
    WailsRuntime.EventsOn('resume:added', (data: any) => {
      devLog('info', `收到后端简历添加事件: ${data.file_name}`)
      const newResume: Resume = {
        id: data.id,
        fileName: data.file_name,
        filePath: data.file_path,
        fileType: data.file_type,
        fileSize: data.file_size,
        content: data.content,
        status: data.status as Resume['status'],
        score: data.score,
        createdAt: data.created_at
      }
      // 检查是否已存在
      if (!resumes.value.some(r => r.id === newResume.id)) {
        resumes.value.unshift(newResume)
      }
    })

    // 监听分析进度事件（含进度百分比）
    WailsRuntime.EventsOn('analysis:progress', (data: any) => {
      devLog('info', `分析进度: id=${data.id}, progress=${data.progress}%`)
      const resume = resumes.value.find(r => r.id === data.id)
      if (resume) {
        resume.status = data.status || 'analyzing'
        // 存储进度到 resume 对象上（动态属性）
        ;(resume as any).progress = data.progress || 0
      }
    })

    // 监听分析完成事件
    WailsRuntime.EventsOn('analysis:completed', (data: any) => {
      devLog('info', `分析完成: id=${data.id}, score=${data.score}`)
      const resume = resumes.value.find(r => r.id === data.id)
      if (resume) {
        resume.status = 'done'
        resume.score = data.score
        const a = data.analysis
        resume.analysis = {
          overallScore: a.overall_score,
          experienceMatch: a.experience_match,
          skillMatch: a.skill_match,
          educationMatch: a.education_match,
          skillDetail: a.skill_detail || '',
          experienceDetail: a.experience_detail || '',
          educationDetail: a.education_detail || '',
          candidateName: a.candidate_name || '',
          workYears: a.work_years || '',
          education: a.education || '',
          currentRole: a.current_role || '',
          summary: a.summary,
          strengths: a.strengths || [],
          weaknesses: a.weaknesses || [],
          risks: a.risks || [],
          recommendation: a.recommendation,
          interviewSuggestions: a.interview_suggestions || []
        }
        resume.analyzedAt = data.analysis.analyzed_at
      }
    })

    // 监听分析错误事件
    WailsRuntime.EventsOn('analysis:error', (data: any) => {
      devLog('error', `分析失败: id=${data.id}, error=${data.error}`)
      const resume = resumes.value.find(r => r.id === data.id)
      if (resume) {
        resume.status = 'error'
      }
    })

    // 监听批量分析进度（更新全局进度状态）
    WailsRuntime.EventsOn('batch:progress', (data: any) => {
      devLog('info', `批量进度: ${data.current}/${data.total}`)
      batchProgress.value = {
        current: data.current,
        total: data.total,
        currentResumeId: data.resumeId || null
      }
    })

    // 监听批量分析完成
    WailsRuntime.EventsOn('batch:completed', (data: any) => {
      devLog('info', `批量分析完成, 共 ${data.total} 份`)
      isAnalyzing.value = false
      batchProgress.value = { current: 0, total: 0, currentResumeId: null }
    })

    console.log('✅ Wails 事件监听已初始化')
  }

  // 检查是否为 Wails 环境
  function isWailsEnvironment() {
    return isWailsEnv
  }

  return {
    // 状态
    resumes,
    isAnalyzing,
    selectedId,
    batchProgress,
    // 计算属性
    selectedResume,
    pendingCount,
    doneCount,
    // 方法
    addResumes,
    deleteResume,
    selectResume,
    reAnalyze,
    startAnalysis,
    clearAll,
    initWailsEvents,
    isWailsEnvironment
  }
})
