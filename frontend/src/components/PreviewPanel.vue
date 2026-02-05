<template>
  <div class="pdf-preview">
    <div v-if="loading" class="loading">
      <el-icon class="loading-icon"><Loading /></el-icon>
      <span>加载中...</span>
    </div>
    <div v-else-if="error" class="error">
      <el-icon><Warning /></el-icon>
      <span>{{ error }}</span>
    </div>
    <div v-else class="pdf-container">
      <canvas ref="canvas" class="pdf-canvas"></canvas>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { Loading, Warning } from '@element-plus/icons-vue'

const props = defineProps<{
  filePath: string
}>()

const canvas = ref<HTMLCanvasElement | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)

// PDF.js 相关
let pdfjsLib: any = null
let pdfDoc: any = null
let currentPage = 1
let totalPages = 0

// 加载PDF.js
async function loadPdfJs() {
  try {
    // 动态导入pdf.js
    pdfjsLib = await import('pdfjs-dist')

    // 设置worker
    pdfjsLib.GlobalWorkerOptions.workerSrc = await import('pdfjs-dist/build/pdf.worker.mjs')
  } catch (e) {
    console.error('加载PDF.js失败:', e)
    error.value = 'PDF预览组件加载失败'
  }
}

// 加载PDF文档
async function loadPdf() {
  if (!pdfjsLib || !props.filePath) return

  loading.value = true
  error.value = null

  try {
    // 使用文件URL加载
    const loadingTask = pdfjsLib.getDocument({
      url: props.filePath,
      cMapUrl: 'https://cdn.jsdelivr.net/npm/pdfjs-dist@4.0.0/cmaps/',
      cMapPacked: true
    })

    pdfDoc = await loadingTask.promise
    totalPages = pdfDoc.numPages

    // 渲染第一页
    await renderPage(currentPage)
  } catch (e) {
    console.error('加载PDF失败:', e)
    error.value = '无法加载PDF文件'
  } finally {
    loading.value = false
  }
}

// 渲染页面
async function renderPage(pageNum: number) {
  if (!pdfDoc || !canvas.value) return

  try {
    const page = await pdfDoc.getPage(pageNum)
    const viewport = page.getViewport({ scale: 1.5 })

    const context = canvas.value.getContext('2d')
    if (!context) return

    canvas.value.height = viewport.height
    canvas.value.width = viewport.width

    await page.render({
      canvasContext: context,
      viewport: viewport
    }).promise
  } catch (e) {
    console.error('渲染页面失败:', e)
  }
}

// 组件挂载
onMounted(async () => {
  await loadPdfJs()
  await loadPdf()
})

// 监听文件路径变化
watch(() => props.filePath, async () => {
  await loadPdf()
})

// 组件卸载
onUnmounted(() => {
  if (pdfDoc) {
    pdfDoc.destroy()
    pdfDoc = null
  }
})
</script>

<style scoped lang="scss">
.pdf-preview {
  width: 100%;
  height: 100%;
  min-height: 400px;
  background-color: #f5f5f5;
  border-radius: 8px;
  overflow: hidden;

  .loading,
  .error {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
    min-height: 200px;
    gap: 8px;
    color: #909399;
  }

  .loading-icon {
    font-size: 32px;
    animation: spin 1s linear infinite;
  }

  .error {
    color: #f56c6c;
  }

  .pdf-container {
    width: 100%;
    height: 100%;
    overflow: auto;
    display: flex;
    justify-content: center;
    padding: 16px;

    .pdf-canvas {
      max-width: 100%;
      box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    }
  }
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>
