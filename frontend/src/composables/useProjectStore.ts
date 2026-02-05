import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { devLog } from './useResumeStore'

// 项目类型定义
export interface Project {
  id: string
  name: string
  job_config: {
    title: string
    requirements: string[]
    required_skills: string[]
    experience_years: number
    education_level: string
  }
  resume_ids: string[]
  status: 'draft' | 'analyzing' | 'completed'
  created_at: string
  updated_at: string
}

// Wails 绑定
let WailsApp: any = null
let isWailsEnv = false

async function loadWails() {
  if (WailsApp) return
  try {
    WailsApp = await import('../../wailsjs/go/main/App')
    isWailsEnv = true
  } catch {
    isWailsEnv = false
  }
}

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>([])
  const currentProjectId = ref<string | null>(null)
  const loading = ref(false)

  const currentProject = computed(() => {
    if (!currentProjectId.value) return null
    return projects.value.find(p => p.id === currentProjectId.value) || null
  })

  // 加载所有项目
  async function loadProjects() {
    await loadWails()
    if (!isWailsEnv || !WailsApp) {
      // Mock 模式
      devLog('info', '项目列表: Mock 模式')
      return
    }

    loading.value = true
    try {
      const result = await WailsApp.GetProjects()
      projects.value = result || []
      devLog('info', `加载了 ${projects.value.length} 个项目`)
    } catch (err: any) {
      devLog('error', `加载项目失败: ${err.message || err}`)
    } finally {
      loading.value = false
    }
  }

  // 创建项目
  async function createProject(name: string, jobConfig: any): Promise<Project | null> {
    await loadWails()
    if (!isWailsEnv || !WailsApp) {
      // Mock
      const mock: Project = {
        id: 'proj_' + Date.now(),
        name,
        job_config: jobConfig,
        resume_ids: [],
        status: 'draft',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      }
      projects.value.unshift(mock)
      devLog('info', `创建 Mock 项目: ${name}`)
      return mock
    }

    try {
      const result = await WailsApp.CreateProject(name, jobConfig)
      await loadProjects()
      devLog('info', `创建项目成功: ${name}`)
      return result
    } catch (err: any) {
      devLog('error', `创建项目失败: ${err.message || err}`)
      return null
    }
  }

  // 删除项目
  async function deleteProject(id: string) {
    await loadWails()
    if (isWailsEnv && WailsApp) {
      try {
        await WailsApp.DeleteProject(id)
        devLog('info', `删除项目: ${id}`)
      } catch (err: any) {
        devLog('error', `删除项目失败: ${err.message || err}`)
      }
    }
    projects.value = projects.value.filter(p => p.id !== id)
    if (currentProjectId.value === id) currentProjectId.value = null
  }

  // 刷新单个项目
  async function refreshProject(id: string) {
    await loadWails()
    if (!isWailsEnv || !WailsApp) return
    try {
      const result = await WailsApp.GetProject(id)
      if (result) {
        const idx = projects.value.findIndex(p => p.id === id)
        if (idx >= 0) projects.value[idx] = result
        else projects.value.unshift(result)
      }
    } catch {}
  }

  // 导出报告
  async function exportReport(projectId: string): Promise<string> {
    await loadWails()
    if (!isWailsEnv || !WailsApp) {
      devLog('warn', 'Mock 模式不支持导出')
      return ''
    }
    try {
      const result = await WailsApp.ExportProjectReport(projectId)
      if (Array.isArray(result)) {
        devLog('info', `导出成功: ${result[0]}`)
        return result[0]
      }
      devLog('info', `导出成功: ${result}`)
      return result
    } catch (err: any) {
      devLog('error', `导出失败: ${err.message || err}`)
      return ''
    }
  }

  // 迁移现有数据
  async function migrateExisting(): Promise<string> {
    await loadWails()
    if (!isWailsEnv || !WailsApp) return ''
    try {
      const result = await WailsApp.MigrateExistingResumes()
      if (result) {
        devLog('info', `数据迁移完成，默认项目ID: ${result}`)
        await loadProjects()
      }
      return result || ''
    } catch { return '' }
  }

  return {
    projects,
    currentProjectId,
    currentProject,
    loading,
    loadProjects,
    createProject,
    deleteProject,
    refreshProject,
    exportReport,
    migrateExisting
  }
})
