<template>
  <el-dropdown trigger="click" @command="handleCommand">
    <span class="language-trigger">
      <el-icon><Place /></el-icon>
      <span class="lang-label">{{ currentLabel }}</span>
      <el-icon class="arrow"><ArrowDown /></el-icon>
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="option in LOCALE_OPTIONS"
          :key="option.value"
          :command="option.value"
          :class="{ 'is-active': option.value === currentLocale }"
        >
          <span class="lang-flag">{{ option.flag }}</span>
          <span class="lang-name">{{ option.label }}</span>
          <el-icon v-if="option.value === currentLocale" class="check-icon"><Check /></el-icon>
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { Place, ArrowDown, Check } from '@element-plus/icons-vue'
import { LOCALE_OPTIONS, setLocale } from '../locales'

const { locale } = useI18n()

const currentLocale = computed(() => locale.value)

const currentLabel = computed(() => {
  const option = LOCALE_OPTIONS.find(opt => opt.value === locale.value)
  return option?.label || '简体中文'
})

function handleCommand(command: string) {
  setLocale(command)
}
</script>

<style scoped lang="scss">
$primary: #1e3a5f;
$primary-light: #2d5a87;
$text-secondary: #5a6678;

.language-trigger {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 13px;
  color: $text-secondary;
  background: rgba(30, 58, 95, 0.04);
  border: 1px solid rgba(30, 58, 95, 0.1);
  transition: all 0.2s ease;

  &:hover {
    background: rgba(30, 58, 95, 0.08);
    border-color: rgba(30, 58, 95, 0.2);
    color: $primary;
  }

  .lang-label {
    font-weight: 500;
  }

  .arrow {
    font-size: 12px;
    margin-left: 2px;
  }
}

:deep(.el-dropdown-menu__item) {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 16px;
  min-width: 140px;

  &.is-active {
    color: $primary;
    background-color: rgba(30, 58, 95, 0.06);
    font-weight: 500;
  }

  .lang-flag {
    font-size: 16px;
  }

  .lang-name {
    flex: 1;
  }

  .check-icon {
    color: $primary;
    font-size: 14px;
  }
}
</style>
