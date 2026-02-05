<template>
  <div class="text-preview">
    <div class="content" v-html="formattedContent"></div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  content: string
}>()

// 格式化内容（处理换行）
const formattedContent = computed(() => {
  if (!props.content) return '<p>暂无内容</p>'

  // 将换行符转换为<br>标签
  const withBreaks = props.content
    .replace(/\n/g, '<br>')
    .replace(/\t/g, '&nbsp;&nbsp;&nbsp;&nbsp;')

  return `<div>${withBreaks}</div>`
})
</script>

<style scoped lang="scss">
.text-preview {
  width: 100%;
  height: 100%;
  padding: 16px;
  background-color: #fafafa;
  border-radius: 8px;
  overflow: auto;

  .content {
    font-size: 14px;
    line-height: 1.6;
    color: #303133;
    white-space: pre-wrap;
    word-break: break-word;
  }
}
</style>
