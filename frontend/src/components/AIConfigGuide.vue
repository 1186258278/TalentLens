<template>
  <el-dialog
    v-model="visible"
    :title="$t('guide.title')"
    width="480px"
    :close-on-click-modal="false"
    class="ai-config-guide"
  >
    <div class="guide-content">
      <!-- 图标和提示 -->
      <div class="guide-icon">
        <el-icon><Warning /></el-icon>
      </div>
      <h3 class="guide-subtitle">{{ $t('guide.notConfigured') }}</h3>
      <p class="guide-description">{{ $t('guide.description') }}</p>

      <!-- 快速配置提示 -->
      <div class="quick-tips">
        <div class="tip-item">
          <el-icon><CircleCheck /></el-icon>
          <span>{{ $t('guide.tip1') }}</span>
        </div>
        <div class="tip-item">
          <el-icon><CircleCheck /></el-icon>
          <span>{{ $t('guide.tip2') }}</span>
        </div>
        <div class="tip-item">
          <el-icon><CircleCheck /></el-icon>
          <span>{{ $t('guide.tip3') }}</span>
        </div>
      </div>

      <!-- 推荐服务商 -->
      <div class="recommended-provider">
        <span class="label">{{ $t('guide.recommended') }}</span>
        <div class="provider-card">
          <div class="provider-info">
            <strong>DeepSeek</strong>
            <span class="provider-desc">{{ $t('guide.deepseekDesc') }}</span>
          </div>
          <el-tag type="success" size="small">{{ $t('guide.freeQuota') }}</el-tag>
        </div>
      </div>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleCancel">{{ $t('common.cancel') }}</el-button>
        <el-button type="primary" @click="handleGoSettings">
          <el-icon><Setting /></el-icon>
          {{ $t('guide.goSettings') }}
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { Warning, CircleCheck, Setting } from '@element-plus/icons-vue'

const props = defineProps<{
  modelValue: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
}>()

const router = useRouter()

const visible = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

function handleCancel() {
  visible.value = false
}

function handleGoSettings() {
  visible.value = false
  router.push('/settings')
}
</script>

<style scoped lang="scss">
$primary: #1e3a5f;
$primary-light: #2d5a87;
$accent: #0ea5e9;
$success: #16a34a;
$warning: #d97706;
$text-primary: #1a2332;
$text-secondary: #5a6678;
$text-muted: #8b95a5;
$bg-section: #f8f9fb;
$border: #e2e7ed;
$radius-md: 10px;

.guide-content {
  text-align: center;
  padding: 10px 0;

  .guide-icon {
    width: 72px;
    height: 72px;
    margin: 0 auto 16px;
    background: linear-gradient(135deg, #fef3c7, #fde68a);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    
    .el-icon {
      font-size: 36px;
      color: $warning;
    }
  }

  .guide-subtitle {
    font-size: 18px;
    font-weight: 600;
    color: $text-primary;
    margin: 0 0 8px 0;
  }

  .guide-description {
    font-size: 14px;
    color: $text-secondary;
    margin: 0 0 24px 0;
    line-height: 1.6;
  }
}

.quick-tips {
  background: $bg-section;
  border-radius: $radius-md;
  padding: 16px;
  margin-bottom: 20px;
  text-align: left;

  .tip-item {
    display: flex;
    align-items: center;
    gap: 10px;
    font-size: 13px;
    color: $text-secondary;
    margin-bottom: 10px;

    &:last-child {
      margin-bottom: 0;
    }

    .el-icon {
      color: $success;
      font-size: 16px;
    }
  }
}

.recommended-provider {
  text-align: left;

  .label {
    font-size: 12px;
    color: $text-muted;
    display: block;
    margin-bottom: 8px;
  }

  .provider-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px 16px;
    background: linear-gradient(135deg, #f0f9ff, #e0f2fe);
    border: 1px solid #bae6fd;
    border-radius: $radius-md;

    .provider-info {
      display: flex;
      flex-direction: column;
      gap: 2px;

      strong {
        font-size: 14px;
        color: $primary;
      }

      .provider-desc {
        font-size: 12px;
        color: $text-muted;
      }
    }
  }
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

:deep(.el-dialog__header) {
  padding: 20px 24px 16px;
  border-bottom: 1px solid $border;
}

:deep(.el-dialog__body) {
  padding: 24px;
}

:deep(.el-dialog__footer) {
  padding: 16px 24px 20px;
  border-top: 1px solid $border;
}
</style>
