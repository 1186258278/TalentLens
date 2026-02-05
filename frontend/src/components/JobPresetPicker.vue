<template>
  <div class="job-preset-picker">
    <div class="picker-header">
      <span class="picker-title">{{ $t('job.quickSelect') }}</span>
    </div>

    <div class="category-tabs">
      <el-scrollbar>
        <div class="tab-wrapper">
          <div
            v-for="category in jobCategories"
            :key="category.id"
            class="category-tab"
            :class="{ active: activeCategory === category.id }"
            @click="activeCategory = category.id"
          >
            {{ $t(category.labelKey) }}
          </div>
        </div>
      </el-scrollbar>
    </div>

    <div class="preset-list">
      <div
        v-for="preset in currentPresets"
        :key="preset.id"
        class="preset-item"
        :class="{ selected: selectedId === preset.id }"
        @click="selectPreset(preset)"
      >
        <span class="preset-name">{{ preset.name }}</span>
        <span class="preset-info">
          {{ preset.experienceYears }}{{ $t('common.year') }} · {{ preset.educationLevel }}
        </span>
      </div>

      <div
        class="preset-item custom"
        :class="{ selected: selectedId === 'custom' }"
        @click="selectCustom"
      >
        <el-icon><Plus /></el-icon>
        <span>{{ $t('job.customJob') }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { jobCategories, type JobPreset } from '../data/jobPresets'

const props = defineProps<{
  selectedId?: string
}>()

const emit = defineEmits<{
  (e: 'select', preset: JobPreset | null): void
}>()

const activeCategory = ref('tech')

const currentPresets = computed(() => {
  const category = jobCategories.find(c => c.id === activeCategory.value)
  return category?.presets || []
})

function selectPreset(preset: JobPreset) {
  emit('select', preset)
}

function selectCustom() {
  emit('select', null)
}
</script>

<style scoped lang="scss">
$primary: #1e3a5f;
$primary-light: #2d5a87;
$primary-lighter: #e8f0f7;
$accent: #0ea5e9;

$bg-section: #fafbfc;
$border: #e2e7ed;
$divider: #eef1f5;

$text-primary: #1a2332;
$text-secondary: #5a6678;
$text-muted: #8b95a5;

$radius-sm: 6px;
$radius-md: 10px;

.job-preset-picker {
  background-color: $bg-section;
  border: 1px solid $divider;
  border-radius: $radius-md;
  overflow: hidden;
  margin-bottom: 20px;
}

.picker-header {
  padding: 14px 18px;
  border-bottom: 1px solid $divider;

  .picker-title {
    font-size: 14px;
    font-weight: 600;
    color: $text-primary;
  }
}

.category-tabs {
  background-color: white;
  border-bottom: 1px solid $divider;

  .tab-wrapper {
    display: flex;
    gap: 4px;
    padding: 10px 14px;
  }

  .category-tab {
    padding: 6px 14px;
    font-size: 13px;
    color: $text-secondary;
    border-radius: $radius-sm;
    cursor: pointer;
    transition: all 0.2s ease;
    white-space: nowrap;

    &:hover {
      background-color: $primary-lighter;
      color: $primary;
    }

    &.active {
      background-color: $primary;
      color: white;
      font-weight: 500;
    }
  }
}

.preset-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  padding: 14px;
  background-color: white;
}

.preset-item {
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 10px 14px;
  background-color: $bg-section;
  border: 1px solid $divider;
  border-radius: $radius-sm;
  cursor: pointer;
  transition: all 0.2s ease;
  min-width: 140px;

  &:hover {
    border-color: $primary-light;
    background-color: $primary-lighter;
  }

  &.selected {
    border-color: $primary;
    background-color: $primary-lighter;
    box-shadow: 0 0 0 2px rgba(30, 58, 95, 0.1);

    .preset-name {
      color: $primary;
    }
  }

  .preset-name {
    font-size: 13px;
    font-weight: 600;
    color: $text-primary;
  }

  .preset-info {
    font-size: 11px;
    color: $text-muted;
  }

  &.custom {
    flex-direction: row;
    align-items: center;
    gap: 6px;
    color: $text-secondary;
    font-size: 13px;
    border-style: dashed;

    .el-icon {
      font-size: 14px;
    }

    &:hover {
      color: $primary;
    }
  }
}
</style>
