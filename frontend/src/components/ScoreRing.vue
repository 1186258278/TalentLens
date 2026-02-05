<template>
  <svg :width="size" :height="size" class="score-ring">
    <!-- 背景圆 -->
    <circle
      :cx="radius"
      :cy="radius"
      :r="radius - strokeWidth / 2"
      fill="none"
      stroke="#ebeef5"
      :stroke-width="strokeWidth"
    />
    <!-- 进度圆 -->
    <circle
      :cx="radius"
      :cy="radius"
      :r="radius - strokeWidth / 2"
      fill="none"
      :stroke="color"
      :stroke-width="strokeWidth"
      :stroke-dasharray="circumference"
      :stroke-dashoffset="dashOffset"
      stroke-linecap="round"
      transform="rotate(-90, radius, radius)"
    />
    <!-- 分数文本 -->
    <text
      :x="radius"
      :y="radius"
      class="score-text"
      text-anchor="middle"
      dominant-baseline="middle"
    >
      {{ score }}
    </text>
  </svg>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  score: number
  size?: number
  strokeWidth?: number
}>()

// 默认值
const size = computed(() => props.size || 80)
const strokeWidth = computed(() => props.strokeWidth || 6)

// 计算属性
const radius = computed(() => size.value / 2)
const circumference = computed(() => 2 * Math.PI * (radius.value - strokeWidth.value / 2))
const dashOffset = computed(() => {
  const progress = props.score / 100
  return circumference.value * (1 - progress)
})

// 颜色 - 使用优化后的配色
const color = computed(() => {
  if (props.score >= 80) return '#16a34a'  // 成功绿
  if (props.score >= 60) return '#d97706'  // 警告橙
  return '#dc2626'  // 错误红
})
</script>

<style scoped>
.score-ring {
  display: block;
  filter: drop-shadow(0 2px 4px rgba(30, 58, 95, 0.1));
}

.score-text {
  font-size: 22px;
  font-weight: 700;
  fill: #1a2332;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}
</style>
