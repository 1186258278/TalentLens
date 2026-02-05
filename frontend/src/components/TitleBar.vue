<template>
  <div class="title-bar">
    <!-- 左侧：交通灯按钮 + 返回/标题 -->
    <div class="title-bar-left">
      <!-- macOS 红绿黄交通灯 -->
      <div class="traffic-lights" @mouseenter="trafficHover = true" @mouseleave="trafficHover = false">
        <button class="light close" @click="closeWindow" :title="$t('common.close') || '关闭'">
          <svg v-if="trafficHover" viewBox="0 0 12 12"><path d="M3.5 3.5l5 5M8.5 3.5l-5 5" stroke="rgba(0,0,0,0.6)" stroke-width="1.2" fill="none" stroke-linecap="round"/></svg>
        </button>
        <button class="light minimize" @click="minimizeWindow" :title="$t('common.minimize') || '最小化'">
          <svg v-if="trafficHover" viewBox="0 0 12 12"><rect x="2" y="5.5" width="8" height="1" rx="0.5" fill="rgba(0,0,0,0.6)"/></svg>
        </button>
        <button class="light maximize" @click="maximizeWindow" :title="$t('common.maximize') || '最大化'">
          <svg v-if="trafficHover" viewBox="0 0 12 12"><path d="M3.5 8.5V3.5H8.5V8.5H3.5Z" stroke="rgba(0,0,0,0.5)" stroke-width="1" fill="none"/></svg>
        </button>
      </div>

      <!-- 返回按钮 -->
      <button v-if="showBack" class="back-btn" @click="handleBack">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <path d="M15 18l-6-6 6-6" />
        </svg>
        <span>{{ backText || '' }}</span>
      </button>
    </div>

    <!-- 中间：标题 (居中) -->
    <div class="title-bar-center" data-wails-drag>
      <span class="title-text">{{ title }}</span>
    </div>

    <!-- 右侧：自定义操作区 -->
    <div class="title-bar-right">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const props = withDefaults(defineProps<{
  title?: string
  showBack?: boolean
  backText?: string
}>(), {
  title: 'TalentLens',
  showBack: false,
  backText: ''
})

const emit = defineEmits<{
  (e: 'back'): void
}>()

const router = useRouter()
const trafficHover = ref(false)

// Wails 绑定
let WailsApp: any = null

onMounted(async () => {
  try {
    WailsApp = await import('../../wailsjs/go/main/App')
  } catch {
    // 非 Wails 环境
  }
})

function handleBack() {
  emit('back')
  router.back()
}

async function minimizeWindow() {
  if (WailsApp) await WailsApp.WindowMinimize()
}

async function maximizeWindow() {
  if (WailsApp) await WailsApp.WindowMaximize()
}

async function closeWindow() {
  if (WailsApp) await WailsApp.WindowClose()
}
</script>

<style scoped lang="scss">
@import '../styles/macos-theme.scss';

.title-bar {
  display: flex;
  align-items: center;
  height: $titlebar-height;
  background: $bg-titlebar;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 1px solid $separator;
  user-select: none;
  -webkit-app-region: drag;
  position: relative;
  z-index: 1000;
  font-family: $font-family;
}

// 左侧区域
.title-bar-left {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-left: 14px;
  -webkit-app-region: no-drag;
  min-width: 160px;
}

// macOS 交通灯按钮
.traffic-lights {
  display: flex;
  gap: 8px;
  padding: 0 4px;

  .light {
    width: 12px;
    height: 12px;
    border-radius: 50%;
    border: none;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: opacity $transition-fast;
    padding: 0;

    svg {
      width: 8px;
      height: 8px;
    }

    &.close {
      background: #FF5F57;
    }

    &.minimize {
      background: #FEBC2E;
    }

    &.maximize {
      background: #28C840;
    }
  }
}

// 返回按钮 (macOS 风格)
.back-btn {
  display: flex;
  align-items: center;
  gap: 2px;
  padding: 4px 8px;
  border: none;
  border-radius: $radius-sm;
  background: transparent;
  color: $system-blue;
  font-size: 13px;
  font-weight: 400;
  cursor: pointer;
  transition: background $transition-fast;

  &:hover {
    background: $system-blue-light;
  }

  svg {
    width: 18px;
    height: 18px;
  }
}

// 中间标题 (居中 + 可拖拽)
.title-bar-center {
  flex: 1;
  text-align: center;
  -webkit-app-region: drag;

  .title-text {
    font-size: 13px;
    font-weight: 600;
    color: $text-primary;
    letter-spacing: -0.01em;
  }
}

// 右侧操作区
.title-bar-right {
  display: flex;
  align-items: center;
  gap: 8px;
  padding-right: 16px;
  -webkit-app-region: no-drag;
  min-width: 160px;
  justify-content: flex-end;
}
</style>
