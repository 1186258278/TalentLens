<template>
  <div class="project-page">
    <TitleBar title="TalentLens">
      <div class="titlebar-actions">
        <LanguageSwitcher />
        <button class="toolbar-btn" @click="router.push('/settings')">
          <el-icon><Setting /></el-icon>
        </button>
      </div>
    </TitleBar>

    <main class="project-main">
      <!-- 顶部操作栏 -->
      <div class="top-bar">
        <h2 class="page-title">{{ $t('project.title') }}</h2>
        <button class="create-btn" @click="showCreateDialog = true">
          <svg viewBox="0 0 16 16" fill="currentColor"><path d="M8 2a.5.5 0 0 1 .5.5v5h5a.5.5 0 0 1 0 1h-5v5a.5.5 0 0 1-1 0v-5h-5a.5.5 0 0 1 0-1h5v-5A.5.5 0 0 1 8 2z"/></svg>
          {{ $t('project.create') }}
        </button>
      </div>

      <!-- 项目列表 -->
      <div class="project-grid">
        <div
          v-for="project in projectStore.projects"
          :key="project.id"
          class="project-card"
          @click="openProject(project.id)"
        >
          <div class="card-top">
            <h3 class="card-title">{{ project.name }}</h3>
            <button class="card-menu" @click.stop="handleDelete(project.id)" title="删除">
              <svg viewBox="0 0 16 16" fill="currentColor"><path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"/><path fill-rule="evenodd" d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4L4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118z"/></svg>
            </button>
          </div>
          <div class="card-meta">
            <span class="meta-item">{{ project.job_config?.title || '未设置岗位' }}</span>
          </div>
          <div class="card-stats">
            <span class="stat">{{ project.resume_ids?.length || 0 }} 份简历</span>
            <span class="stat-badge" :class="project.status">{{ statusText(project.status) }}</span>
          </div>
          <div class="card-time">{{ formatDate(project.created_at) }}</div>
        </div>

        <!-- 空状态 -->
        <div v-if="projectStore.projects.length === 0 && !projectStore.loading" class="empty-state">
          <div class="empty-icon">
            <svg viewBox="0 0 48 48" fill="none" stroke="currentColor" stroke-width="1.5"><rect x="6" y="10" width="36" height="28" rx="4"/><path d="M6 18h36"/><circle cx="14" cy="14" r="1.5" fill="currentColor"/><circle cx="20" cy="14" r="1.5" fill="currentColor"/></svg>
          </div>
          <p>{{ $t('project.empty') }}</p>
          <button class="create-btn" @click="showCreateDialog = true">
            {{ $t('project.createFirst') }}
          </button>
        </div>
      </div>
    </main>

    <!-- 新建项目对话框 -->
    <el-dialog v-model="showCreateDialog" :title="$t('project.createDialogTitle')" width="520px" :close-on-click-modal="false">
      <div class="create-form">
        <div class="form-item">
          <label>{{ $t('project.projectName') }}</label>
          <el-input v-model="newProject.name" :placeholder="$t('project.projectNamePlaceholder')" />
        </div>

        <div class="form-item">
          <label>{{ $t('job.title') }}</label>
          <el-input v-model="newProject.jobTitle" :placeholder="$t('job.titlePlaceholder')" />
        </div>

        <div class="form-item">
          <label>{{ $t('job.requiredSkills') }}</label>
          <el-select v-model="newProject.skills" multiple filterable allow-create default-first-option :placeholder="$t('common.add')" style="width: 100%">
            <el-option v-for="s in commonSkills" :key="s" :label="s" :value="s" />
          </el-select>
        </div>

        <div class="form-row">
          <div class="form-item flex-1">
            <label>{{ $t('job.experienceYears') }}</label>
            <el-input-number v-model="newProject.expYears" :min="0" :max="30" />
          </div>
          <div class="form-item flex-1">
            <label>{{ $t('job.educationLevel') }}</label>
            <el-select v-model="newProject.eduLevel" style="width: 100%">
              <el-option :label="$t('job.education.any')" value="" />
              <el-option :label="$t('job.education.college')" value="大专" />
              <el-option :label="$t('job.education.bachelor')" value="本科" />
              <el-option :label="$t('job.education.master')" value="硕士" />
              <el-option :label="$t('job.education.phd')" value="博士" />
            </el-select>
          </div>
        </div>
      </div>

      <template #footer>
        <el-button @click="showCreateDialog = false">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleCreate" :disabled="!newProject.name.trim()">{{ $t('common.confirm') }}</el-button>
      </template>
    </el-dialog>

    <DevPanel />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Setting } from '@element-plus/icons-vue'
import TitleBar from '../components/TitleBar.vue'
import LanguageSwitcher from '../components/LanguageSwitcher.vue'
import DevPanel from '../components/DevPanel.vue'
import { useProjectStore } from '../composables/useProjectStore'

const router = useRouter()
const { t } = useI18n()
const projectStore = useProjectStore()

const showCreateDialog = ref(false)
const newProject = reactive({
  name: '',
  jobTitle: '高级Go开发工程师',
  skills: ['Go', 'MySQL', 'Redis'] as string[],
  expYears: 3,
  eduLevel: '本科'
})

const commonSkills = [
  'Go', 'Python', 'Java', 'JavaScript', 'TypeScript',
  'React', 'Vue', 'Node.js', 'MySQL', 'Redis',
  'Docker', 'Kubernetes', 'Linux', 'Git'
]

function statusText(status: string): string {
  const map: Record<string, string> = {
    draft: t('project.statusDraft'),
    analyzing: t('project.statusAnalyzing'),
    completed: t('project.statusCompleted')
  }
  return map[status] || status
}

function formatDate(dateStr: string): string {
  if (!dateStr) return ''
  try {
    return new Date(dateStr).toLocaleDateString('zh-CN')
  } catch { return dateStr }
}

function openProject(id: string) {
  projectStore.currentProjectId = id
  router.push(`/project/${id}`)
}

async function handleCreate() {
  if (!newProject.name.trim()) return
  const jobCfg = {
    title: newProject.jobTitle,
    requirements: [],
    required_skills: newProject.skills,
    experience_years: newProject.expYears,
    education_level: newProject.eduLevel
  }
  const result = await projectStore.createProject(newProject.name.trim(), jobCfg)
  if (result) {
    showCreateDialog.value = false
    ElMessage.success(t('project.createSuccess'))
    // 重置表单
    newProject.name = ''
  }
}

async function handleDelete(id: string) {
  try {
    await ElMessageBox.confirm(t('project.deleteConfirm'), t('common.warning'), { type: 'warning' })
    await projectStore.deleteProject(id)
    ElMessage.success(t('project.deleteSuccess'))
  } catch {}
}

onMounted(async () => {
  await projectStore.loadProjects()
  // 如果没有项目，尝试迁移现有数据
  if (projectStore.projects.length === 0) {
    await projectStore.migrateExisting()
  }
})
</script>

<style scoped lang="scss">
@import '../styles/macos-theme.scss';

.project-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: $bg-secondary;
  font-family: $font-family;
}

.project-main {
  flex: 1;
  overflow-y: auto;
  padding: 28px 40px;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;

  .page-title {
    margin: 0;
    font-size: 22px;
    font-weight: 700;
    color: $text-primary;
    letter-spacing: -0.02em;
  }
}

.create-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  border: none;
  border-radius: $radius-md;
  background: $system-blue;
  color: white;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  font-family: $font-family;
  transition: background $transition-fast;

  svg { width: 14px; height: 14px; }

  &:hover { background: $system-blue-hover; }
  &:active { transform: scale(0.97); }
}

.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 16px;
}

.project-card {
  background: $bg-primary;
  border: 1px solid $separator;
  border-radius: $radius-lg;
  padding: 20px;
  cursor: pointer;
  transition: all $transition-normal;

  &:hover {
    border-color: $system-blue;
    box-shadow: 0 4px 16px rgba(0, 122, 255, 0.1);
    transform: translateY(-2px);
  }

  .card-top {
    display: flex;
    justify-content: space-between;
    align-items: flex-start;
    margin-bottom: 10px;
  }

  .card-title {
    margin: 0;
    font-size: 16px;
    font-weight: 600;
    color: $text-primary;
  }

  .card-menu {
    width: 28px;
    height: 28px;
    border: none;
    border-radius: $radius-sm;
    background: transparent;
    color: $text-tertiary;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all $transition-fast;

    svg { width: 14px; height: 14px; }

    &:hover { background: rgba(255, 59, 48, 0.1); color: $system-red; }
  }

  .card-meta {
    margin-bottom: 12px;

    .meta-item {
      font-size: 13px;
      color: $text-secondary;
    }
  }

  .card-stats {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 10px;

    .stat {
      font-size: 12px;
      color: $text-secondary;
    }

    .stat-badge {
      font-size: 11px;
      font-weight: 500;
      padding: 2px 8px;
      border-radius: 4px;

      &.draft { background: $bg-hover; color: $text-secondary; }
      &.analyzing { background: rgba(0, 122, 255, 0.1); color: $system-blue; }
      &.completed { background: rgba(52, 199, 89, 0.1); color: $system-green; }
    }
  }

  .card-time {
    font-size: 11px;
    color: $text-tertiary;
  }
}

.empty-state {
  grid-column: 1 / -1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 20px;
  text-align: center;

  .empty-icon {
    margin-bottom: 16px;

    svg { width: 48px; height: 48px; color: $text-tertiary; }
  }

  p {
    font-size: 14px;
    color: $text-secondary;
    margin: 0 0 20px 0;
  }
}

.titlebar-actions {
  display: flex;
  align-items: center;
  gap: 8px;

  .toolbar-btn {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    border: none;
    border-radius: $radius-sm;
    background: transparent;
    color: $text-secondary;
    cursor: pointer;
    transition: background $transition-fast;

    .el-icon { font-size: 16px; }
    &:hover { background: $bg-hover; color: $text-primary; }
  }

  :deep(.language-trigger) {
    background: transparent;
    border: 1px solid $gray-4;
    color: $text-primary;
    font-size: 12px;
    padding: 3px 8px;
    border-radius: 4px;
    &:hover { background: $bg-hover; }
  }
}

// 对话框样式
.create-form {
  .form-item {
    margin-bottom: 18px;
    label { display: block; margin-bottom: 6px; font-size: 13px; font-weight: 500; color: $text-primary; }
    &.flex-1 { flex: 1; }
  }
  .form-row { display: flex; gap: 16px; }
}

:deep(.el-dialog) {
  border-radius: $radius-lg;
}
</style>
