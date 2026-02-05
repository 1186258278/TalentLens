<template>
  <div
    class="drop-zone"
    :class="{ 'is-dragging': isDragging }"
    @dragenter="handleDragEnter"
    @dragleave="handleDragLeave"
    @dragover.prevent="handleDragOver"
    @drop.prevent="handleDrop"
    @click="handleClick"
  >
    <input
      ref="fileInput"
      type="file"
      multiple
      accept=".pdf,.docx,.doc,.jpg,.jpeg,.png,.bmp,.gif,.webp"
      @change="handleFileSelect"
      hidden
    />

    <div class="drop-content">
      <el-icon class="drop-icon"><UploadFilled /></el-icon>
      <p class="drop-text">{{ dropText }}</p>
      <p class="drop-hint">{{ dropHint }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { UploadFilled } from '@element-plus/icons-vue'

const { t } = useI18n()

// 定义文件信息接口
export interface FileInfo {
  name: string
  path: string
  size: number
  type: string
  lastModified: number
}

const emit = defineEmits<{
  (e: 'files-added', files: FileInfo[]): void
}>()

const fileInput = ref<HTMLInputElement | null>(null)
const isDragging = ref(false)
const dragCounter = ref(0)

// 拖拽文本
const dropText = computed(() => {
  return isDragging.value ? t('home.dropFiles') : t('home.dropHere')
})

const dropHint = computed(() => t('home.supportedFormats'))

// 支持的文件扩展名
const supportedExtensions = ['.pdf', '.docx', '.doc', '.jpg', '.jpeg', '.png', '.bmp', '.gif', '.webp']

// 检查文件是否支持
function isFileSupported(file: File): boolean {
  const ext = '.' + file.name.split('.').pop()?.toLowerCase()
  return supportedExtensions.includes(ext)
}

// 从 File 对象提取文件信息
function extractFileInfo(file: File): FileInfo {
  // Wails/Electron 环境下 File 对象可能有 path 属性
  const filePath = (file as any).path || file.name
  return {
    name: file.name,
    path: filePath,
    size: file.size,
    type: file.type,
    lastModified: file.lastModified
  }
}

// 处理拖拽进入
function handleDragEnter(e: DragEvent) {
  e.preventDefault()
  dragCounter.value++
  if (dragCounter.value > 0) {
    isDragging.value = true
  }
}

// 处理拖拽离开
function handleDragLeave(e: DragEvent) {
  e.preventDefault()
  dragCounter.value--
  if (dragCounter.value === 0) {
    isDragging.value = false
  }
}

// 处理拖拽经过
function handleDragOver(e: DragEvent) {
  e.preventDefault()
}

// 处理释放文件
function handleDrop(e: DragEvent) {
  e.preventDefault()
  dragCounter.value = 0
  isDragging.value = false

  const files = e.dataTransfer?.files
  if (files && files.length > 0) {
    const validFiles = Array.from(files)
      .filter(isFileSupported)
      .map(extractFileInfo)
    
    console.log('拖放的文件:', validFiles)
    
    if (validFiles.length > 0) {
      emit('files-added', validFiles)
    }
  }
}

// 处理点击选择文件
function handleClick() {
  fileInput.value?.click()
}

// 处理文件选择
function handleFileSelect(e: Event) {
  const input = e.target as HTMLInputElement
  if (input.files && input.files.length > 0) {
    const validFiles = Array.from(input.files)
      .filter(isFileSupported)
      .map(extractFileInfo)
    
    console.log('选择的文件:', validFiles)
    
    if (validFiles.length > 0) {
      emit('files-added', validFiles)
    }
    // 清空input，允许重复选择同一文件
    input.value = ''
  }
}
</script>

<style scoped lang="scss">
// ==========================================
// 设计系统变量
// ==========================================
$primary: #1e3a5f;
$primary-light: #2d5a87;
$primary-lighter: #e8f0f7;
$accent: #0ea5e9;

$bg-section: #fafbfc;
$border: #e2e7ed;
$text-secondary: #5a6678;
$text-muted: #8b95a5;

$radius-md: 10px;

// ==========================================
// 拖拽上传区域
// ==========================================
.drop-zone {
  margin: 16px;
  border: 2px dashed $border;
  border-radius: $radius-md;
  padding: 36px 24px;
  text-align: center;
  cursor: pointer;
  transition: all 0.25s ease;
  background-color: $bg-section;
  position: relative;
  overflow: hidden;

  // 装饰性背景图案
  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-image: radial-gradient(circle at 2px 2px, rgba(30, 58, 95, 0.03) 1px, transparent 0);
    background-size: 20px 20px;
    pointer-events: none;
    opacity: 0;
    transition: opacity 0.25s ease;
  }

  &:hover {
    border-color: $primary-light;
    background-color: $primary-lighter;
    box-shadow: 0 4px 12px rgba(30, 58, 95, 0.08);

    &::before {
      opacity: 1;
    }

    .drop-icon {
      color: $primary;
      transform: translateY(-2px);
    }

    .drop-text {
      color: $primary;
    }
  }

  &.is-dragging {
    border-color: $primary;
    border-style: solid;
    background-color: $primary-lighter;
    transform: scale(1.01);
    box-shadow: 0 8px 24px rgba(30, 58, 95, 0.12);

    &::before {
      opacity: 1;
    }

    .drop-icon {
      color: $primary;
      transform: translateY(-4px) scale(1.05);
    }

    .drop-text {
      color: $primary;
      font-weight: 600;
    }
  }

  .drop-content {
    pointer-events: none;
    position: relative;
    z-index: 1;
  }

  .drop-icon {
    font-size: 52px;
    color: $text-muted;
    margin-bottom: 14px;
    transition: all 0.25s ease;
  }

  .drop-text {
    font-size: 15px;
    color: $text-secondary;
    margin: 0 0 8px 0;
    font-weight: 500;
    transition: all 0.25s ease;
  }

  .drop-hint {
    font-size: 12px;
    color: $text-muted;
    margin: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;

    // 文件类型标签样式
    &::before {
      content: 'PDF';
      padding: 2px 6px;
      background-color: rgba(30, 58, 95, 0.08);
      border-radius: 3px;
      font-size: 10px;
      font-weight: 600;
      color: $primary-light;
      letter-spacing: 0.5px;
    }

    &::after {
      content: 'Word';
      padding: 2px 6px;
      background-color: rgba(30, 58, 95, 0.08);
      border-radius: 3px;
      font-size: 10px;
      font-weight: 600;
      color: $primary-light;
      letter-spacing: 0.5px;
    }
  }
}
</style>
