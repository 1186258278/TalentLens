<template>
  <div class="image-preview">
    <img
      v-if="imageUrl"
      :src="imageUrl"
      :alt="filePath"
      class="preview-image"
      @load="onImageLoad"
      @error="onImageError"
    />
    <div v-else class="error">
      <el-icon><Picture /></el-icon>
      <span>无法加载图片</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { Picture } from '@element-plus/icons-vue'

const props = defineProps<{
  filePath: string
}>()

// 图片URL
const imageUrl = computed(() => {
  if (!props.filePath) return null
  return `file://${props.filePath}`
})

// 图片加载成功
function onImageLoad() {
  console.log('图片加载成功')
}

// 图片加载失败
function onImageError() {
  console.error('图片加载失败:', props.filePath)
}
</script>

<style scoped lang="scss">
.image-preview {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: #f5f5f5;
  border-radius: 8px;
  overflow: hidden;

  .preview-image {
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
  }

  .error {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8px;
    color: #909399;

    .el-icon {
      font-size: 48px;
    }
  }
}
</style>
