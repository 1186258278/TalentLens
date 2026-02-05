<template>
  <div class="app-container">
    <!-- macOS 风格标题栏 -->
    <TitleBar :title="$t('app.title')">
      <div class="titlebar-actions">
        <span class="job-badge">
          <el-icon><Briefcase /></el-icon>
          {{ jobTitle }}
        </span>
        <LanguageSwitcher />
        <button class="toolbar-btn" @click="router.push('/settings')">
          <el-icon><Setting /></el-icon>
        </button>
      </div>
    </TitleBar>

    <!-- 全局分析进度条 -->
    <Transition name="progress-slide">
      <div v-if="resumeStore.isAnalyzing && resumeStore.batchProgress.total > 0" class="global-progress">
        <div class="progress-info">
          <span class="progress-text">
            {{ $t('analysis.batchProgress', { current: resumeStore.batchProgress.current, total: resumeStore.batchProgress.total }) }}
          </span>
          <span class="progress-pct">{{ Math.round((resumeStore.batchProgress.current / resumeStore.batchProgress.total) * 100) }}%</span>
        </div>
        <div class="progress-track">
          <div class="progress-fill" :style="{ width: (resumeStore.batchProgress.current / resumeStore.batchProgress.total * 100) + '%' }"></div>
        </div>
      </div>
    </Transition>

    <main class="app-main">
      <!-- 左侧面板 -->
      <aside class="left-panel">
        <DropZone @files-added="handleFilesAdded" />

        <div class="resume-section">
          <div class="section-bar">
            <span class="section-label">
              <el-icon><Document /></el-icon>
              {{ $t('home.resumeList') }} ({{ resumeStore.resumes.length }})
            </span>
            <div class="section-btns">
              <button class="action-btn primary" @click="handleStartAnalysis" :disabled="resumeStore.isAnalyzing">
                <el-icon v-if="!resumeStore.isAnalyzing"><VideoPlay /></el-icon>
                <el-icon v-else class="spin"><Loading /></el-icon>
                {{ $t('home.startAnalysis') }}
              </button>
              <button class="action-btn" @click="handleClearAll">
                <el-icon><Delete /></el-icon>
              </button>
            </div>
          </div>

          <div class="resume-list">
            <ResumeCard
              v-for="resume in sortedResumes"
              :key="resume.id"
              :resume="resume"
              :selected="resumeStore.selectedId === resume.id"
              @click="resumeStore.selectResume(resume.id)"
              @re-analyze="resumeStore.reAnalyze"
              @delete="resumeStore.deleteResume"
            />
            <el-empty v-if="resumeStore.resumes.length === 0" :description="$t('home.dragHint')" />
          </div>
        </div>
      </aside>

      <!-- 右侧预览面板 -->
      <section class="right-panel">
        <div class="preview-section">
          <div class="section-bar">
            <span class="section-label">
              <el-icon><View /></el-icon>
              {{ $t('home.preview') }}
            </span>
          </div>

          <div class="preview-content">
            <div v-if="resumeStore.selectedResume" class="resume-detail">
              <!-- 基本信息 -->
              <div class="detail-header">
                <el-icon class="file-icon"><Document /></el-icon>
                <div class="file-info">
                  <h3>{{ resumeStore.selectedResume.fileName }}</h3>
                  <span class="file-meta">
                    {{ formatFileSize(resumeStore.selectedResume.fileSize) }} · {{ getFileTypeLabel(resumeStore.selectedResume.fileType) }}
                  </span>
                </div>
              </div>

              <!-- 分析结果 -->
              <div v-if="resumeStore.selectedResume.status === 'done' && resumeStore.selectedResume.analysis" class="analysis-result">
                <!-- 候选人信息卡片 -->
                <div v-if="resumeStore.selectedResume.analysis.candidateName" class="candidate-card">
                  <div class="candidate-name">{{ resumeStore.selectedResume.analysis.candidateName }}</div>
                  <div class="candidate-meta">
                    <span v-if="resumeStore.selectedResume.analysis.currentRole">{{ resumeStore.selectedResume.analysis.currentRole }}</span>
                    <span v-if="resumeStore.selectedResume.analysis.workYears">{{ resumeStore.selectedResume.analysis.workYears }}经验</span>
                    <span v-if="resumeStore.selectedResume.analysis.education">{{ resumeStore.selectedResume.analysis.education }}</span>
                  </div>
                </div>

                <!-- 评分面板 -->
                <div class="score-section">
                  <div class="main-score" :class="getScoreClass(resumeStore.selectedResume.score)">
                    <span class="score-value" :key="resumeStore.selectedResume.id">{{ animatedDetailScore }}</span>
                    <span class="score-label">{{ $t('analysis.totalScore') }}</span>
                  </div>
                  <div class="score-breakdown">
                    <div class="score-item">
                      <span class="label">{{ $t('analysis.skillMatch') }}</span>
                      <el-progress :percentage="resumeStore.selectedResume.analysis.skillMatch" :color="getProgressColor(resumeStore.selectedResume.analysis.skillMatch)" />
                    </div>
                    <div class="score-item">
                      <span class="label">{{ $t('analysis.experienceMatch') }}</span>
                      <el-progress :percentage="resumeStore.selectedResume.analysis.experienceMatch" :color="getProgressColor(resumeStore.selectedResume.analysis.experienceMatch)" />
                    </div>
                    <div class="score-item">
                      <span class="label">{{ $t('analysis.educationMatch') }}</span>
                      <el-progress :percentage="resumeStore.selectedResume.analysis.educationMatch" :color="getProgressColor(resumeStore.selectedResume.analysis.educationMatch)" />
                    </div>
                  </div>
                </div>

                <!-- 推荐结论 -->
                <div class="recommendation-section">
                  <div class="recommendation-badge" :class="resumeStore.selectedResume.analysis.recommendation">
                    {{ $t(`analysis.recommendations.${resumeStore.selectedResume.analysis.recommendation}`) }}
                  </div>
                </div>

                <!-- 评分详解 -->
                <div v-if="resumeStore.selectedResume.analysis.skillDetail" class="detail-cards">
                  <div class="detail-card">
                    <h4>{{ $t('analysis.skillMatch') }}详解</h4>
                    <p>{{ resumeStore.selectedResume.analysis.skillDetail }}</p>
                  </div>
                  <div class="detail-card">
                    <h4>{{ $t('analysis.experienceMatch') }}详解</h4>
                    <p>{{ resumeStore.selectedResume.analysis.experienceDetail }}</p>
                  </div>
                  <div class="detail-card">
                    <h4>{{ $t('analysis.educationMatch') }}详解</h4>
                    <p>{{ resumeStore.selectedResume.analysis.educationDetail }}</p>
                  </div>
                </div>

                <!-- 优势与不足 -->
                <div class="strength-weakness">
                  <div class="sw-block strengths">
                    <h4><el-icon><CircleCheck /></el-icon> {{ $t('analysis.strengths') }}</h4>
                    <ul>
                      <li v-for="(item, idx) in resumeStore.selectedResume.analysis.strengths" :key="idx">{{ item }}</li>
                    </ul>
                  </div>
                  <div class="sw-block weaknesses">
                    <h4><el-icon><Warning /></el-icon> {{ $t('analysis.weaknesses') }}</h4>
                    <ul>
                      <li v-for="(item, idx) in resumeStore.selectedResume.analysis.weaknesses" :key="idx">{{ item }}</li>
                    </ul>
                  </div>
                </div>

                <!-- 风险提示 -->
                <div v-if="resumeStore.selectedResume.analysis.risks && resumeStore.selectedResume.analysis.risks.length > 0" class="risk-section">
                  <h4><el-icon><Warning /></el-icon> 风险提示</h4>
                  <ul>
                    <li v-for="(item, idx) in resumeStore.selectedResume.analysis.risks" :key="idx">{{ item }}</li>
                  </ul>
                </div>

                <!-- 面试建议 -->
                <div v-if="resumeStore.selectedResume.analysis.interviewSuggestions && resumeStore.selectedResume.analysis.interviewSuggestions.length > 0" class="interview-section">
                  <h4><el-icon><ChatLineSquare /></el-icon> 面试建议</h4>
                  <ul>
                    <li v-for="(item, idx) in resumeStore.selectedResume.analysis.interviewSuggestions" :key="idx">{{ item }}</li>
                  </ul>
                </div>

                <!-- AI 总结 -->
                <div class="summary-section">
                  <h4><el-icon><ChatLineSquare /></el-icon> {{ $t('analysis.summary') }}</h4>
                  <p>{{ resumeStore.selectedResume.analysis.summary }}</p>
                </div>
              </div>

              <!-- 状态提示 -->
              <div v-else-if="resumeStore.selectedResume.status === 'pending'" class="status-hint pending">
                <el-icon><Clock /></el-icon>
                <span>{{ $t('analysis.pendingHint') }}</span>
              </div>
              <div v-else-if="resumeStore.selectedResume.status === 'analyzing'" class="status-hint analyzing">
                <el-icon class="spin"><Loading /></el-icon>
                <span>{{ $t('analysis.analyzingHint') }}</span>
              </div>
              <div v-else-if="resumeStore.selectedResume.status === 'error'" class="status-hint error">
                <el-icon><CircleClose /></el-icon>
                <span>{{ $t('analysis.errorHint') }}</span>
              </div>
            </div>

            <el-empty v-else :description="$t('home.selectToPreview')" />
          </div>
        </div>
      </section>
    </main>

    <AIConfigGuide v-model="showConfigGuide" />

    <!-- 开发者调试面板 (F12 切换) -->
    <DevPanel />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
  Setting, Briefcase, Document, VideoPlay, Delete, View,
  CircleCheck, Warning, ChatLineSquare, Clock, Loading, CircleClose
} from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import TitleBar from '../components/TitleBar.vue'
import DropZone from '../components/DropZone.vue'
import ResumeCard from '../components/ResumeCard.vue'
import LanguageSwitcher from '../components/LanguageSwitcher.vue'
import AIConfigGuide from '../components/AIConfigGuide.vue'
import DevPanel from '../components/DevPanel.vue'
import { useResumeStore, type FileInfo } from '../composables/useResumeStore'

const { t } = useI18n()
const router = useRouter()
const resumeStore = useResumeStore()

const jobTitle = ref('高级Go开发工程师')
const showConfigGuide = ref(false)
const animatedDetailScore = ref(0)
let scoreAnimationFrame: number | null = null

watch(
  () => resumeStore.selectedResume?.score,
  (newScore) => {
    if (newScore && resumeStore.selectedResume?.status === 'done') {
      animateDetailScore(newScore)
    } else {
      animatedDetailScore.value = 0
    }
  },
  { immediate: true }
)

function animateDetailScore(targetScore: number) {
  const startScore = 0
  const startTime = performance.now()
  const duration = 1000
  function update(currentTime: number) {
    const elapsed = currentTime - startTime
    const progress = Math.min(elapsed / duration, 1)
    const easeProgress = progress === 1 ? 1 : 1 - Math.pow(2, -10 * progress)
    animatedDetailScore.value = Math.round(startScore + (targetScore - startScore) * easeProgress)
    if (progress < 1) scoreAnimationFrame = requestAnimationFrame(update)
  }
  if (scoreAnimationFrame) cancelAnimationFrame(scoreAnimationFrame)
  scoreAnimationFrame = requestAnimationFrame(update)
}

function checkAIConfig(): boolean {
  const saved = localStorage.getItem('goresume_settings')
  if (!saved) return false
  try {
    const settings = JSON.parse(saved)
    return settings.ai?.apiKey && settings.ai.apiKey.length >= 10
  } catch { return false }
}

const sortedResumes = computed(() => {
  return [...resumeStore.resumes].sort((a, b) => {
    if (a.status === 'done' && b.status !== 'done') return -1
    if (a.status !== 'done' && b.status === 'done') return 1
    if (a.status === 'done' && b.status === 'done') return (b.score || 0) - (a.score || 0)
    if (a.status === 'analyzing' && b.status === 'pending') return -1
    if (a.status === 'pending' && b.status === 'analyzing') return 1
    return 0
  })
})

function loadJobTitle() {
  const saved = localStorage.getItem('goresume_settings')
  if (saved) {
    try {
      const settings = JSON.parse(saved)
      if (settings.job?.title) jobTitle.value = settings.job.title
    } catch {}
  }
}

function handleFilesAdded(files: FileInfo[]) {
  resumeStore.addResumes(files)
}

function handleStartAnalysis() {
  const pendingCount = resumeStore.resumes.filter(r => r.status === 'pending').length
  if (pendingCount === 0) { ElMessage.warning(t('guide.noResumesToAnalyze')); return }
  if (!checkAIConfig()) { showConfigGuide.value = true; return }
  resumeStore.startAnalysis()
}

function handleClearAll() { resumeStore.clearAll() }

function formatFileSize(bytes: number): string {
  if (bytes < 1024) return bytes + ' B'
  if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB'
  return (bytes / (1024 * 1024)).toFixed(1) + ' MB'
}

function getFileTypeLabel(fileType: string): string {
  const ext = fileType.toLowerCase()
  if (ext === '.pdf') return 'PDF'
  if (ext === '.docx' || ext === '.doc') return 'Word'
  if (['.jpg', '.jpeg', '.png', '.bmp', '.gif', '.webp'].includes(ext)) return t('file.image')
  return t('file.unknown')
}

function getScoreClass(score: number | undefined): string {
  if (!score) return 'low'
  if (score >= 80) return 'high'
  if (score >= 60) return 'medium'
  return 'low'
}

function getProgressColor(value: number): string {
  if (value >= 80) return '#34C759'
  if (value >= 60) return '#FF9500'
  return '#FF3B30'
}

onMounted(async () => {
  loadJobTitle()
  await resumeStore.initWailsEvents()
})
</script>

<style scoped lang="scss">
@import '../styles/macos-theme.scss';

// 主容器
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: $bg-secondary;
  font-family: $font-family;
}

// 标题栏内的操作区
.titlebar-actions {
  display: flex;
  align-items: center;
  gap: 8px;

  .job-badge {
    display: flex;
    align-items: center;
    gap: 5px;
    font-size: 12px;
    color: $text-secondary;
    padding: 3px 10px;
    background: $bg-hover;
    border-radius: 4px;
    border: 1px solid $separator;

    .el-icon { font-size: 12px; color: $gray-1; }
  }

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
}

// 全局分析进度条
.global-progress {
  padding: 8px 18px;
  background: $bg-primary;
  border-bottom: 1px solid $separator;

  .progress-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 6px;
  }

  .progress-text {
    font-size: 12px;
    font-weight: 500;
    color: $system-blue;
  }

  .progress-pct {
    font-size: 11px;
    font-weight: 600;
    color: $text-secondary;
  }

  .progress-track {
    height: 3px;
    background: rgba(0, 122, 255, 0.1);
    border-radius: 2px;
    overflow: hidden;
  }

  .progress-fill {
    height: 100%;
    background: $system-blue;
    border-radius: 2px;
    transition: width 0.5s ease;
  }
}

.progress-slide-enter-active, .progress-slide-leave-active {
  transition: all 0.3s ease;
}
.progress-slide-enter-from, .progress-slide-leave-to {
  opacity: 0;
  max-height: 0;
  padding-top: 0;
  padding-bottom: 0;
}

// 主体
.app-main {
  display: flex;
  flex: 1;
  overflow: hidden;
}

// 左侧面板
.left-panel {
  width: 380px;
  display: flex;
  flex-direction: column;
  border-right: 1px solid $separator;
  background: $bg-primary;
  flex-shrink: 0;

  .resume-section {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
  }
}

// 右侧面板
.right-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: $bg-primary;
}

// 区块头部
.section-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 16px;
  border-bottom: 1px solid $separator;
  background: rgba(246, 246, 246, 0.5);

  .section-label {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 13px;
    font-weight: 600;
    color: $text-primary;

    .el-icon { color: $gray-1; font-size: 15px; }
  }

  .section-btns {
    display: flex;
    gap: 6px;
  }
}

// macOS 风格操作按钮
.action-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 5px 12px;
  border: 1px solid $gray-4;
  border-radius: $radius-sm;
  background: $bg-primary;
  color: $text-primary;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all $transition-fast;
  font-family: $font-family;

  .el-icon { font-size: 13px; }

  &:hover { background: $bg-hover; border-color: $gray-3; }
  &:active { transform: scale(0.97); }

  &.primary {
    background: $system-blue;
    border-color: $system-blue;
    color: white;

    &:hover { background: $system-blue-hover; }
    &:disabled {
      opacity: 0.6;
      cursor: not-allowed;
    }
  }
}

.spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

// 简历列表
.resume-list {
  flex: 1;
  overflow-y: auto;
  padding: 10px;
  background: $bg-secondary;

  &::-webkit-scrollbar { width: 6px; }
  &::-webkit-scrollbar-track { background: transparent; }
  &::-webkit-scrollbar-thumb { background: rgba(0,0,0,0.12); border-radius: 3px; }

  :deep(.el-empty) {
    padding: 40px 20px;
    .el-empty__description { color: $text-tertiary; font-size: 13px; }
  }
}

// 预览区
.preview-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.preview-content {
  flex: 1;
  overflow: auto;
  padding: 20px;
  background: $bg-secondary;

  :deep(.el-empty) {
    padding: 60px 20px;
    .el-empty__description { color: $text-tertiary; }
  }
}

// 简历详情
.resume-detail {
  .detail-header {
    display: flex;
    align-items: center;
    gap: 14px;
    padding: 16px 20px;
    background: $bg-primary;
    border-radius: $radius-lg;
    margin-bottom: 14px;
    border: 1px solid $separator;

    .file-icon { font-size: 36px; color: $system-blue; }

    .file-info {
      h3 { margin: 0 0 3px 0; font-size: 15px; font-weight: 600; color: $text-primary; }
      .file-meta { font-size: 12px; color: $text-secondary; }
    }
  }

  .analysis-result {
    .score-section {
      display: flex;
      gap: 20px;
      padding: 20px;
      background: $bg-primary;
      border-radius: $radius-lg;
      margin-bottom: 14px;
      border: 1px solid $separator;

      .main-score {
        width: 90px;
        height: 90px;
        border-radius: 50%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        color: white;
        flex-shrink: 0;

        &.high { background: linear-gradient(135deg, $system-green, darken($system-green, 10%)); }
        &.medium { background: linear-gradient(135deg, $system-orange, darken($system-orange, 10%)); }
        &.low { background: linear-gradient(135deg, $system-red, darken($system-red, 10%)); }

        .score-value { font-size: 28px; font-weight: 700; line-height: 1; animation: scorePopIn 0.5s ease-out; }
        .score-label { font-size: 10px; opacity: 0.9; margin-top: 3px; }
      }

      .score-breakdown {
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 10px;

        .score-item {
          .label { display: block; font-size: 12px; color: $text-secondary; margin-bottom: 3px; }
          :deep(.el-progress__text) { font-size: 11px; font-weight: 600; }
        }
      }
    }

    .recommendation-section {
      text-align: center;
      margin-bottom: 14px;

      .recommendation-badge {
        display: inline-block;
        padding: 8px 20px;
        border-radius: 16px;
        font-size: 13px;
        font-weight: 600;

        &.strong_recommend { background: linear-gradient(135deg, $system-green, darken($system-green, 10%)); color: white; }
        &.recommend { background: rgba(52, 199, 89, 0.12); color: $system-green; }
        &.consider { background: rgba(255, 149, 0, 0.12); color: $system-orange; }
        &.not_recommend { background: rgba(255, 59, 48, 0.12); color: $system-red; }
      }
    }

    .strength-weakness {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 14px;
      margin-bottom: 14px;

      .sw-block {
        padding: 16px;
        border-radius: $radius-lg;
        border: 1px solid $separator;

        h4 { display: flex; align-items: center; gap: 6px; margin: 0 0 10px 0; font-size: 13px; font-weight: 600; }
        ul { margin: 0; padding-left: 18px; }
        li { font-size: 12px; color: $text-secondary; margin-bottom: 5px; &:last-child { margin-bottom: 0; } }

        &.strengths { background: rgba(52, 199, 89, 0.06); h4 { color: $system-green; } }
        &.weaknesses { background: rgba(255, 149, 0, 0.06); h4 { color: $system-orange; } }
      }
    }

    // 候选人信息卡片
    .candidate-card {
      padding: 14px 18px;
      background: linear-gradient(135deg, rgba(0,122,255,0.06), rgba(0,122,255,0.02));
      border: 1px solid rgba(0,122,255,0.15);
      border-radius: $radius-lg;
      margin-bottom: 14px;

      .candidate-name { font-size: 16px; font-weight: 700; color: $text-primary; margin-bottom: 6px; }
      .candidate-meta {
        display: flex; flex-wrap: wrap; gap: 8px;
        span {
          font-size: 12px; color: $text-secondary;
          padding: 2px 8px; background: rgba(0,0,0,0.04); border-radius: 4px;
        }
      }
    }

    // 评分详解
    .detail-cards {
      display: flex; flex-direction: column; gap: 10px; margin-bottom: 14px;

      .detail-card {
        padding: 14px 16px;
        background: $bg-primary;
        border: 1px solid $separator;
        border-radius: $radius-lg;

        h4 { margin: 0 0 8px 0; font-size: 13px; font-weight: 600; color: $text-primary; }
        p { margin: 0; font-size: 12px; color: $text-secondary; line-height: 1.7; }
      }
    }

    // 风险提示
    .risk-section {
      padding: 14px 16px;
      background: rgba(255, 149, 0, 0.04);
      border: 1px solid rgba(255, 149, 0, 0.15);
      border-radius: $radius-lg;
      margin-bottom: 14px;

      h4 { display: flex; align-items: center; gap: 6px; margin: 0 0 10px 0; font-size: 13px; font-weight: 600; color: $system-orange; }
      ul { margin: 0; padding-left: 18px; }
      li { font-size: 12px; color: $text-secondary; margin-bottom: 5px; line-height: 1.6; }
    }

    // 面试建议
    .interview-section {
      padding: 14px 16px;
      background: rgba(0, 122, 255, 0.04);
      border: 1px solid rgba(0, 122, 255, 0.12);
      border-radius: $radius-lg;
      margin-bottom: 14px;

      h4 { display: flex; align-items: center; gap: 6px; margin: 0 0 10px 0; font-size: 13px; font-weight: 600; color: $system-blue; }
      ul { margin: 0; padding-left: 18px; }
      li { font-size: 12px; color: $text-secondary; margin-bottom: 5px; line-height: 1.6; }
    }

    // AI 总结
    .summary-section {
      padding: 16px;
      background: $bg-primary;
      border-radius: $radius-lg;
      border: 1px solid $separator;

      h4 { display: flex; align-items: center; gap: 6px; margin: 0 0 10px 0; font-size: 13px; font-weight: 600; color: $system-blue; }
      p { margin: 0; font-size: 13px; color: $text-secondary; line-height: 1.7; }
    }
  }

  .status-hint {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 60px 20px;
    text-align: center;

    .el-icon { font-size: 40px; margin-bottom: 12px; }
    span { font-size: 13px; }

    &.pending { .el-icon, span { color: $text-tertiary; } }
    &.analyzing { .el-icon, span { color: $system-blue; } }
    &.error { .el-icon, span { color: $system-red; } }
  }
}

@keyframes scorePopIn {
  0% { transform: scale(0.5); opacity: 0; }
  50% { transform: scale(1.1); }
  100% { transform: scale(1); opacity: 1; }
}

:deep(.el-progress-bar__inner) {
  transition: width 0.5s ease-out !important;
}

// 语言切换器样式
.titlebar-actions {
  :deep(.language-trigger) {
    background: transparent;
    border: 1px solid $gray-4;
    color: $text-primary;
    font-size: 12px;
    padding: 3px 8px;
    border-radius: 4px;

    &:hover { background: $bg-hover; border-color: $gray-3; }
  }
}
</style>
