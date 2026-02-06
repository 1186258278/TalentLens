<template>
  <div class="settings-page">
    <!-- macOS 风格标题栏 -->
    <TitleBar :title="$t('settings.title')" :show-back="true">
      <LanguageSwitcher />
    </TitleBar>

    <!-- 主体：侧边栏 + 内容区 -->
    <div class="settings-body">
      <!-- 侧边栏导航 -->
      <aside class="sidebar">
        <nav class="nav-list">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="['nav-item', { active: activeTab === tab.id }]"
            @click="activeTab = tab.id"
          >
            <el-icon class="nav-icon"><component :is="tab.icon" /></el-icon>
            <span class="nav-label">{{ $t(tab.label) }}</span>
          </button>
        </nav>

        <div class="sidebar-bottom">
          <button class="save-btn" @click="saveSettings">
            <svg viewBox="0 0 16 16" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round">
              <path d="M13.3 14H2.7c-.7 0-1.4-.6-1.4-1.3V3.3C1.3 2.6 2 2 2.7 2h7.7l3.3 3.3v7.4c0 .7-.6 1.3-1.4 1.3z"/>
              <path d="M11.3 14V9.3H4.7V14"/>
              <path d="M4.7 2v4h5.3"/>
            </svg>
            {{ $t('common.save') }}
          </button>
        </div>
      </aside>

      <!-- 内容区 -->
      <main class="content-area">
        <!-- AI 模型配置 -->
        <section v-show="activeTab === 'ai'" class="content-section">
          <div class="section-head">
            <h2>{{ $t('settings.aiConfig') }}</h2>
            <p>{{ $t('settings.aiConfigDesc') }}</p>
          </div>

          <div class="form-card">
            <div class="form-item">
              <label>{{ $t('ai.provider') }}</label>
              <el-select v-model="aiForm.provider" class="full-width" @change="handleProviderChange">
                <el-option
                  v-for="provider in providers"
                  :key="provider.id"
                  :label="provider.name"
                  :value="provider.id"
                >
                  <div class="option-row">
                    <span>{{ provider.name }}</span>
                    <el-tag v-if="provider.recommended" size="small" type="success">{{ $t('ai.recommended') }}</el-tag>
                  </div>
                </el-option>
              </el-select>
            </div>
          </div>

          <!-- 配置引导 -->
          <ProviderGuide :provider="currentProvider" />

          <div class="form-card">
            <div v-if="aiForm.provider === 'custom'" class="form-item">
              <label>{{ $t('ai.baseUrl') }}</label>
              <el-input v-model="aiForm.baseURL" :placeholder="$t('ai.baseUrlPlaceholder')" class="full-width" />
            </div>

            <div class="form-item">
              <label>{{ $t('ai.apiKey') }}</label>
              <el-input v-model="aiForm.apiKey" type="password" show-password :placeholder="$t('ai.apiKeyPlaceholder')" class="full-width" />
            </div>

            <div class="form-item">
              <label>{{ $t('ai.model') }}</label>
              <el-select
                v-if="currentProvider && currentProvider.models.length > 0"
                v-model="aiForm.model"
                class="full-width"
              >
                <el-option
                  v-for="model in currentProvider.models"
                  :key="model.id"
                  :label="model.name"
                  :value="model.id"
                >
                  <div class="option-row">
                    <span>{{ model.name }}</span>
                    <el-tag v-if="model.recommended" size="small" type="success">{{ $t('ai.recommended') }}</el-tag>
                  </div>
                </el-option>
              </el-select>
              <el-input v-else v-model="aiForm.model" placeholder="gpt-4o-mini" class="full-width" />
            </div>

            <div class="form-item actions">
              <el-button @click="testConnection" :loading="testing" round>
                <el-icon><Connection /></el-icon>
                {{ $t('ai.testConnection') }}
              </el-button>
            </div>
          </div>
        </section>

        <!-- 岗位需求 -->
        <section v-show="activeTab === 'job'" class="content-section">
          <div class="section-head">
            <h2>{{ $t('settings.jobRequirements') }}</h2>
            <p>{{ $t('settings.jobRequirementsDesc') }}</p>
          </div>

          <JobPresetPicker :selected-id="selectedPresetId" @select="handlePresetSelect" />

          <div class="info-badge">
            <span class="info-badge-label">{{ $t('job.currentJob') }}:</span>
            <span class="info-badge-value">{{ jobForm.title || $t('job.customJob') }}</span>
          </div>

          <div class="form-card">
            <div class="form-item">
              <label>{{ $t('job.title') }}</label>
              <el-input v-model="jobForm.title" :placeholder="$t('job.titlePlaceholder')" class="full-width" />
            </div>

            <div class="form-item">
              <label>{{ $t('job.requiredSkills') }}</label>
              <el-select v-model="jobForm.requiredSkills" multiple filterable allow-create default-first-option :placeholder="$t('common.add')" class="full-width">
                <el-option v-for="skill in commonSkills" :key="skill" :label="skill" :value="skill" />
              </el-select>
            </div>

            <div class="form-row">
              <div class="form-item flex-1">
                <label>{{ $t('job.experienceYears') }}</label>
                <div class="inline-hint">
                  <el-input-number v-model="jobForm.experienceYears" :min="0" :max="30" />
                  <span class="hint-text">{{ $t('common.year') }}</span>
                </div>
              </div>
              <div class="form-item flex-1">
                <label>{{ $t('job.educationLevel') }}</label>
                <el-select v-model="jobForm.educationLevel" class="full-width">
                  <el-option :label="$t('job.education.any')" value="" />
                  <el-option :label="$t('job.education.college')" value="大专" />
                  <el-option :label="$t('job.education.bachelor')" value="本科" />
                  <el-option :label="$t('job.education.master')" value="硕士" />
                  <el-option :label="$t('job.education.phd')" value="博士" />
                </el-select>
              </div>
            </div>
          </div>
        </section>

        <!-- 分析设置 -->
        <section v-show="activeTab === 'analysis'" class="content-section">
          <div class="section-head">
            <h2>{{ $t('settings.analysisSettings') }}</h2>
            <p>{{ $t('settings.analysisSettingsDesc') }}</p>
          </div>

          <div class="form-card">
            <div class="form-item">
              <label>{{ $t('analysis.concurrent') }}</label>
              <div class="inline-hint">
                <el-input-number v-model="analysisForm.maxConcurrent" :min="1" :max="10" />
                <span class="hint-text">{{ $t('analysis.concurrentHint') }}</span>
              </div>
            </div>

            <div class="form-item">
              <label>{{ $t('analysis.autoStart') }}</label>
              <div class="switch-row">
                <el-switch v-model="analysisForm.autoStart" />
                <span class="switch-hint">{{ $t('analysis.autoStartDesc') }}</span>
              </div>
            </div>

            <div class="form-item actions">
              <el-button @click="resetSettings" type="danger" plain round>
                <el-icon><Refresh /></el-icon>
                {{ $t('settings.resetAllSettings') }}
              </el-button>
            </div>

            <div class="form-item" style="margin-top: 24px;">
              <label>{{ $t('settings.dataPath') }}</label>
              <div class="data-path-row">
                <span class="path-text">{{ dataDir || '~/Documents/TalentLens' }}</span>
                <el-button size="small" round @click="openDataDir">{{ $t('settings.openDataDir') }}</el-button>
              </div>
            </div>
          </div>
        </section>

        <!-- 关于 -->
        <section v-show="activeTab === 'about'" class="content-section about-section">
          <div class="about-center">
            <!-- Logo 区域 -->
            <div class="about-hero">
              <div class="about-icon">
                <svg viewBox="0 0 48 48" fill="none" width="48" height="48">
                  <rect width="48" height="48" rx="12" fill="url(#aboutGrad)"/>
                  <path d="M14 16h20M14 24h14M14 32h18" stroke="#fff" stroke-width="2.5" stroke-linecap="round"/>
                  <circle cx="36" cy="32" r="6" fill="#fff" fill-opacity="0.3"/>
                  <path d="M34 32l2 2 3-3" stroke="#fff" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
                  <defs><linearGradient id="aboutGrad" x1="0" y1="0" x2="48" y2="48"><stop stop-color="#007AFF"/><stop offset="1" stop-color="#5856D6"/></linearGradient></defs>
                </svg>
              </div>
              <h1 class="about-title">TalentLens</h1>
              <span class="about-version">v{{ appVersion || '1.1.0' }}</span>
              <p class="about-subtitle">AI-Powered Resume Screening Tool</p>
            </div>

            <!-- 更新检测 -->
            <div class="about-update-card">
              <div class="update-row">
                <div class="update-info">
                  <span class="update-label">{{ $t('about.version') }}</span>
                  <span class="update-value">v{{ appVersion || '1.1.0' }}</span>
                </div>
                <el-button @click="checkForUpdate" :loading="updateChecking" size="small" round>
                  {{ updateChecking ? $t('about.checking') : $t('about.checkUpdate') }}
                </el-button>
              </div>
              <div v-if="updateResult" class="update-result-row">
                <div v-if="updateResult.hasUpdate" class="update-available">
                  <span class="new-badge">NEW</span>
                  <span class="new-text">{{ $t('about.newVersion', { version: updateResult.latestVersion }) }}</span>
                  <el-button type="primary" size="small" round @click="openLink(updateResult.releaseURL)">
                    {{ $t('about.download') }}
                  </el-button>
                </div>
                <span v-else-if="!updateResult.error" class="up-to-date">{{ $t('about.noUpdate') }}</span>
              </div>
            </div>

            <!-- 链接 -->
            <div class="about-links-card">
              <h3 class="links-title">{{ $t('about.links') }}</h3>
              <div class="links-grid">
                <button class="link-card" @click="openLink('https://github.com/1186258278/TalentLens')">
                  <svg class="link-svg" viewBox="0 0 16 16" fill="currentColor" width="20" height="20"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/></svg>
                  <span>{{ $t('about.github') }}</span>
                </button>
                <button class="link-card" @click="openLink('https://talentlens.qt.cool')">
                  <span class="link-emoji">🌐</span>
                  <span>{{ $t('about.website') }}</span>
                </button>
                <button class="link-card" @click="openLink('https://qingchencloud.com')">
                  <span class="link-emoji">🏢</span>
                  <span>{{ $t('about.company') }}</span>
                </button>
              </div>
            </div>

            <!-- 版权 -->
            <p class="about-copyright">Copyright &copy; 2025 {{ $t('about.copyright') }}</p>
          </div>
        </section>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, markRaw } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { Connection, Briefcase, Setting, Refresh, InfoFilled, Link } from '@element-plus/icons-vue'

import TitleBar from '../components/TitleBar.vue'
import LanguageSwitcher from '../components/LanguageSwitcher.vue'
import ProviderGuide from '../components/ProviderGuide.vue'
import JobPresetPicker from '../components/JobPresetPicker.vue'
import { providers, getProviderById, getRecommendedProvider, type Provider } from '../data/providers'
import { type JobPreset } from '../data/jobPresets'

const router = useRouter()
const { t } = useI18n()

const activeTab = ref('ai')
const testing = ref(false)
const selectedPresetId = ref<string>('')
const dataDir = ref('')

const tabs = [
  { id: 'ai', label: 'settings.aiConfig', icon: markRaw(Connection) },
  { id: 'job', label: 'settings.jobRequirements', icon: markRaw(Briefcase) },
  { id: 'analysis', label: 'settings.analysisSettings', icon: markRaw(Setting) },
  { id: 'about', label: 'about.title', icon: markRaw(InfoFilled) }
]

// 版本检测状态
const appVersion = ref('')
const updateChecking = ref(false)
const updateResult = ref<{ hasUpdate: boolean; latestVersion: string; releaseURL: string; error: string } | null>(null)

const aiForm = reactive({
  provider: 'deepseek',
  baseURL: 'https://api.deepseek.com',
  apiKey: '',
  model: 'deepseek-chat'
})

const jobForm = reactive({
  title: '高级Go开发工程师',
  requiredSkills: ['Go', 'MySQL', 'Redis'] as string[],
  experienceYears: 5,
  educationLevel: '本科'
})

const analysisForm = reactive({
  maxConcurrent: 3,
  autoStart: false
})

const commonSkills = [
  'Go', 'Python', 'Java', 'JavaScript', 'TypeScript',
  'React', 'Vue', 'Node.js', 'MySQL', 'Redis',
  'Docker', 'Kubernetes', 'Linux', 'Git',
  '需求分析', 'PRD撰写', '数据分析', '用户研究'
]

const currentProvider = computed<Provider | null>(() => {
  return getProviderById(aiForm.provider) || null
})

function handleProviderChange(providerId: string) {
  const provider = getProviderById(providerId)
  if (provider) {
    aiForm.baseURL = provider.baseURL
    aiForm.model = provider.defaultModel
  }
}

function handlePresetSelect(preset: JobPreset | null) {
  if (preset) {
    selectedPresetId.value = preset.id
    jobForm.title = preset.name
    jobForm.requiredSkills = [...preset.requiredSkills]
    jobForm.experienceYears = preset.experienceYears
    jobForm.educationLevel = preset.educationLevel
  } else {
    selectedPresetId.value = 'custom'
  }
}

async function testConnection() {
  if (!aiForm.apiKey) {
    ElMessage.warning(t('ai.apiKeyPlaceholder'))
    return
  }
  testing.value = true
  try {
    let WailsApp: any = null
    try { WailsApp = await import('../../wailsjs/go/main/App') } catch {}
    if (WailsApp) {
      const config = { provider: aiForm.provider, base_url: aiForm.baseURL, api_key: aiForm.apiKey, model: aiForm.model, max_retries: 3, timeout: 30 }
      const result = await WailsApp.TestAIConnection(config)
      let success: boolean, message: string
      if (Array.isArray(result)) { [success, message] = result } else { success = !!result; message = typeof result === 'string' ? result : '' }
      if (success) { ElMessage.success(message || t('ai.connectionSuccess')) } else { ElMessage.error(message || t('ai.connectionFailed')) }
    } else {
      await new Promise(resolve => setTimeout(resolve, 1500))
      ElMessage.success(t('ai.connectionSuccess') + ' (Mock)')
    }
  } catch (error: any) {
    ElMessage.error(error.message || t('ai.connectionFailed'))
  } finally {
    testing.value = false
  }
}

function saveSettings() {
  const settings = {
    ai: { provider: aiForm.provider, baseURL: aiForm.baseURL, apiKey: aiForm.apiKey, model: aiForm.model },
    job: { title: jobForm.title, requiredSkills: [...jobForm.requiredSkills], experienceYears: jobForm.experienceYears, educationLevel: jobForm.educationLevel },
    analysis: { maxConcurrent: analysisForm.maxConcurrent, autoStart: analysisForm.autoStart }
  }
  try {
    localStorage.setItem('goresume_settings', JSON.stringify(settings))
    ElMessage.success(t('settings.saved'))
  } catch (e) {
    ElMessage.error('保存失败')
  }
}

function resetSettings() {
  const recommended = getRecommendedProvider()
  aiForm.provider = recommended.id
  aiForm.baseURL = recommended.baseURL
  aiForm.apiKey = ''
  aiForm.model = recommended.defaultModel
  jobForm.title = '高级Go开发工程师'
  jobForm.requiredSkills = ['Go', 'MySQL', 'Redis']
  jobForm.experienceYears = 5
  jobForm.educationLevel = '本科'
  analysisForm.maxConcurrent = 3
  analysisForm.autoStart = false
  selectedPresetId.value = ''
  localStorage.removeItem('goresume_settings')
  ElMessage.success(t('settings.resetDone'))
}

async function openDataDir() {
  try {
    const WailsApp = await import('../../wailsjs/go/main/App')
    await WailsApp.OpenDataDir()
  } catch {}
}

async function checkForUpdate() {
  updateChecking.value = true
  updateResult.value = null
  try {
    const WailsApp = await import('../../wailsjs/go/main/App')
    const result = await WailsApp.CheckForUpdate()
    updateResult.value = result as any
    if (result.error) {
      ElMessage.warning(t('about.updateError') + ': ' + result.error)
    } else if (result.hasUpdate) {
      ElMessage.success(t('about.newVersion', { version: result.latestVersion }))
    } else {
      ElMessage.info(t('about.noUpdate'))
    }
  } catch {
    ElMessage.error(t('about.updateError'))
  } finally {
    updateChecking.value = false
  }
}

async function openLink(url: string) {
  try {
    const WailsApp = await import('../../wailsjs/go/main/App')
    await WailsApp.OpenURL(url)
  } catch {
    window.open(url, '_blank')
  }
}

onMounted(async () => {
  // 加载数据目录路径和版本号
  try {
    const WailsApp = await import('../../wailsjs/go/main/App')
    dataDir.value = await WailsApp.GetDataDir()
    appVersion.value = await WailsApp.GetAppVersion()
  } catch {
    appVersion.value = '1.1.0'
  }

  const saved = localStorage.getItem('goresume_settings')
  if (saved) {
    try {
      const s = JSON.parse(saved)
      if (s.ai) { aiForm.provider = s.ai.provider || aiForm.provider; aiForm.baseURL = s.ai.baseURL || aiForm.baseURL; aiForm.apiKey = s.ai.apiKey || ''; aiForm.model = s.ai.model || aiForm.model }
      if (s.job) { jobForm.title = s.job.title || jobForm.title; jobForm.requiredSkills = s.job.requiredSkills || jobForm.requiredSkills; jobForm.experienceYears = s.job.experienceYears ?? jobForm.experienceYears; jobForm.educationLevel = s.job.educationLevel || jobForm.educationLevel }
      if (s.analysis) { analysisForm.maxConcurrent = s.analysis.maxConcurrent ?? analysisForm.maxConcurrent; analysisForm.autoStart = s.analysis.autoStart ?? false }
    } catch {}
  }
})
</script>

<style scoped lang="scss">
@import '../styles/macos-theme.scss';

// 页面容器
.settings-page {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: $bg-secondary;
  overflow: hidden;
  font-family: $font-family;
}

// 主体
.settings-body {
  display: flex;
  flex: 1;
  overflow: hidden;
}

// 侧边栏
.sidebar {
  width: 200px;
  background: $bg-sidebar;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-right: 1px solid $separator;
  display: flex;
  flex-direction: column;
  flex-shrink: 0;
}

.nav-list {
  flex: 1;
  padding: 12px 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  width: 100%;
  height: $nav-item-height;
  padding: 0 12px;
  border: none;
  border-radius: $radius-sm;
  background: transparent;
  color: $text-primary;
  font-size: $nav-text-size;
  font-weight: 400;
  cursor: pointer;
  transition: background $transition-fast;
  text-align: left;
  font-family: $font-family;

  &:hover {
    background: $bg-hover;
  }

  &.active {
    background: $bg-active;
    font-weight: 500;

    .nav-icon {
      color: $system-blue;
    }
  }

  .nav-icon {
    font-size: $nav-icon-size;
    color: $gray-1;
    transition: color $transition-fast;
  }
}

// 侧边栏底部
.sidebar-bottom {
  padding: 12px 8px;
  border-top: 1px solid $separator;
}

.save-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  width: 100%;
  padding: 8px 12px;
  border: none;
  border-radius: $radius-sm;
  background: $system-blue;
  color: white;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all $transition-fast;
  font-family: $font-family;

  svg {
    width: 14px;
    height: 14px;
  }

  &:hover {
    background: $system-blue-hover;
  }

  &:active {
    transform: scale(0.98);
  }
}

// 内容区
.content-area {
  flex: 1;
  overflow-y: auto;
  padding: 28px 36px;
  background: $bg-secondary;
}

.content-section {
  max-width: 600px;
  animation: fadeSlide 0.25s ease;
}

@keyframes fadeSlide {
  from { opacity: 0; transform: translateY(6px); }
  to { opacity: 1; transform: translateY(0); }
}

// 区块头部
.section-head {
  margin-bottom: 24px;

  h2 {
    margin: 0 0 6px 0;
    font-size: 20px;
    font-weight: 700;
    color: $text-primary;
    letter-spacing: -0.02em;
  }

  p {
    margin: 0;
    font-size: 13px;
    color: $text-secondary;
    line-height: 1.5;
  }
}

// 表单卡片
.form-card {
  background: $bg-primary;
  border-radius: $radius-lg;
  border: 1px solid $separator;
  padding: 20px;
  margin-bottom: 16px;
}

// 表单项
.form-item {
  margin-bottom: 20px;

  &:last-child {
    margin-bottom: 0;
  }

  &.actions {
    padding-top: 16px;
    border-top: 1px solid $separator;
  }

  label {
    display: block;
    margin-bottom: 6px;
    font-size: 13px;
    font-weight: 500;
    color: $text-primary;
  }
}

.form-row {
  display: flex;
  gap: 16px;
}

.flex-1 {
  flex: 1;
}

.full-width {
  width: 100%;
}

.inline-hint {
  display: flex;
  align-items: center;
  gap: 10px;

  .hint-text {
    font-size: 13px;
    color: $text-secondary;
  }
}

.switch-row {
  display: flex;
  align-items: center;
  gap: 10px;

  .switch-hint {
    font-size: 13px;
    color: $text-secondary;
  }
}

.option-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

// 信息标签
.info-badge {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 14px;
  background: $system-blue-light;
  border-radius: $radius-md;
  margin-bottom: 16px;

  .info-badge-label {
    font-size: 13px;
    color: $text-secondary;
  }

  .info-badge-value {
    font-size: 13px;
    font-weight: 600;
    color: $system-blue;
  }
}

// Element Plus 覆盖
:deep(.el-input__wrapper),
:deep(.el-select__wrapper) {
  box-shadow: 0 0 0 1px $gray-4 inset;
  border-radius: $radius-sm;
  transition: all $transition-fast;
  font-size: 13px;

  &:hover {
    box-shadow: 0 0 0 1px $gray-3 inset;
  }

  &.is-focus {
    box-shadow: 0 0 0 1px $system-blue inset, 0 0 0 3px rgba(0, 122, 255, 0.2);
  }
}

:deep(.el-switch) {
  --el-switch-on-color: #{$system-blue};
}

:deep(.el-select__tags .el-tag) {
  background: $system-blue-light;
  border-color: transparent;
  color: $system-blue;

  .el-tag__close {
    color: $system-blue;
    &:hover { background: $system-blue; color: white; }
  }
}

:deep(.el-button.is-round) {
  font-size: 13px;
  font-weight: 500;
}

:deep(.el-button--danger.is-plain) {
  --el-button-text-color: #{$system-red};
  --el-button-bg-color: rgba(255, 59, 48, 0.06);
  --el-button-border-color: rgba(255, 59, 48, 0.3);
  --el-button-hover-text-color: white;
  --el-button-hover-bg-color: #{$system-red};
  --el-button-hover-border-color: #{$system-red};
}

// 自定义滚动条
.content-area {
  &::-webkit-scrollbar { width: 6px; }
  &::-webkit-scrollbar-track { background: transparent; }
  &::-webkit-scrollbar-thumb { background: rgba(0,0,0,0.12); border-radius: 3px; }
  &::-webkit-scrollbar-thumb:hover { background: rgba(0,0,0,0.2); }
}

.data-path-row {
  display: flex;
  align-items: center;
  gap: 12px;

  .path-text {
    flex: 1;
    font-size: 12px;
    color: $text-secondary;
    padding: 8px 12px;
    background: rgba(0, 0, 0, 0.03);
    border-radius: $radius-sm;
    border: 1px solid $separator;
    font-family: 'SF Mono', 'Menlo', monospace;
    word-break: break-all;
  }
}

// 关于页面
.about-section {
  display: flex;
  align-items: flex-start;
  justify-content: center;
}

.about-center {
  max-width: 420px;
  width: 100%;
  margin: 0 auto;
  padding-top: 20px;
}

.about-hero {
  text-align: center;
  margin-bottom: 24px;

  .about-icon {
    margin-bottom: 14px;
    svg { filter: drop-shadow(0 4px 12px rgba(0, 122, 255, 0.25)); }
  }

  .about-title {
    font-size: 26px;
    font-weight: 700;
    color: $text-primary;
    margin: 0 0 8px 0;
    letter-spacing: -0.5px;
  }

  .about-version {
    display: inline-block;
    font-size: 12px;
    font-weight: 600;
    padding: 3px 12px;
    border-radius: 12px;
    background: $system-blue-light;
    color: $system-blue;
    margin-bottom: 8px;
  }

  .about-subtitle {
    font-size: 13px;
    color: $text-tertiary;
    margin: 8px 0 0 0;
  }
}

.about-update-card {
  background: $bg-primary;
  border: 1px solid $separator;
  border-radius: $radius-lg;
  padding: 16px 20px;
  margin-bottom: 16px;

  .update-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .update-info {
    display: flex;
    align-items: center;
    gap: 8px;

    .update-label {
      font-size: 13px;
      color: $text-secondary;
    }
    .update-value {
      font-size: 13px;
      font-weight: 600;
      color: $text-primary;
    }
  }

  .update-result-row {
    margin-top: 12px;
    padding-top: 12px;
    border-top: 1px solid $separator;

    .update-available {
      display: flex;
      align-items: center;
      gap: 10px;

      .new-badge {
        font-size: 10px;
        font-weight: 700;
        padding: 2px 8px;
        border-radius: 4px;
        background: $system-green;
        color: white;
        letter-spacing: 0.5px;
      }
      .new-text {
        flex: 1;
        font-size: 13px;
        font-weight: 500;
        color: $system-green;
      }
    }

    .up-to-date {
      font-size: 13px;
      color: $text-tertiary;
    }
  }
}

.about-links-card {
  background: $bg-primary;
  border: 1px solid $separator;
  border-radius: $radius-lg;
  padding: 16px 20px;
  margin-bottom: 16px;

  .links-title {
    font-size: 13px;
    font-weight: 600;
    color: $text-secondary;
    margin: 0 0 12px 0;
  }

  .links-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 10px;

    .link-card {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 8px;
      padding: 16px 10px;
      border: 1px solid $separator;
      border-radius: $radius-md;
      background: $bg-secondary;
      cursor: pointer;
      font-family: $font-family;
      font-size: 12px;
      font-weight: 500;
      color: $text-primary;
      transition: all $transition-fast;

      .link-svg { color: $text-secondary; }
      .link-emoji { font-size: 20px; }

      &:hover {
        border-color: $system-blue;
        background: $system-blue-light;
        color: $system-blue;
        transform: translateY(-2px);
        box-shadow: $shadow-md;

        .link-svg { color: $system-blue; }
      }
    }
  }
}

.about-copyright {
  text-align: center;
  font-size: 11px;
  color: $text-tertiary;
  margin: 0;
  padding: 8px 0;
}
</style>
