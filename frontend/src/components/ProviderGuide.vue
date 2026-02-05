<template>
  <div class="provider-guide" v-if="provider">
    <div class="guide-header">
      <el-icon class="guide-icon"><QuestionFilled /></el-icon>
      <span class="guide-title">{{ $t(provider.guide.title) }}</span>
    </div>

    <ol class="guide-steps">
      <li v-for="(step, index) in provider.guide.steps" :key="index">
        {{ $t(step) }}
      </li>
    </ol>

    <div class="guide-actions" v-if="provider.guide.link">
      <el-button
        type="primary"
        @click="openLink"
      >
        <el-icon><Link /></el-icon>
        {{ $t(`providers.${provider.id}.openConsole`) }}
      </el-button>
    </div>

    <div class="guide-pricing">
      <el-icon><Coin /></el-icon>
      <span>{{ $t(provider.guide.pricing) }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { QuestionFilled, Link, Coin } from '@element-plus/icons-vue'
import type { Provider } from '../data/providers'

const props = defineProps<{
  provider: Provider | null
}>()

function openLink() {
  if (props.provider?.guide.link) {
    window.open(props.provider.guide.link, '_blank')
  }
}
</script>

<style scoped lang="scss">
$primary: #1e3a5f;
$primary-light: #2d5a87;
$primary-lighter: #e8f0f7;
$accent: #0ea5e9;
$success: #16a34a;
$warning: #d97706;

$bg-tip: #f0f9ff;
$border-tip: #bae6fd;

$text-primary: #1a2332;
$text-secondary: #5a6678;

$radius-md: 10px;

.provider-guide {
  background-color: $bg-tip;
  border: 1px solid $border-tip;
  border-radius: $radius-md;
  padding: 18px;
  margin-top: 24px;
  margin-bottom: 24px;
}

.guide-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 14px;

  .guide-icon {
    font-size: 20px;
    color: $accent;
  }

  .guide-title {
    font-size: 15px;
    font-weight: 600;
    color: $text-primary;
  }
}

.guide-steps {
  margin: 0 0 16px 0;
  padding-left: 24px;

  li {
    font-size: 14px;
    color: $text-secondary;
    line-height: 1.8;
    margin-bottom: 6px;

    &:last-child {
      margin-bottom: 0;
    }
  }
}

.guide-actions {
  margin-bottom: 14px;

  :deep(.el-button) {
    --el-button-bg-color: #{$primary};
    --el-button-border-color: #{$primary};
    --el-button-hover-bg-color: #{$primary-light};
    --el-button-hover-border-color: #{$primary-light};
    font-weight: 500;
  }
}

.guide-pricing {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 14px;
  background-color: rgba(22, 163, 74, 0.08);
  border-radius: 6px;
  font-size: 13px;
  color: $success;
  font-weight: 500;

  .el-icon {
    font-size: 16px;
  }
}
</style>
