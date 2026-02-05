<template>
  <div class="analysis-result">
    <!-- 综合评分 -->
    <div class="overall-score">
      <ScoreRing
        :score="analysis.overallScore"
        :size="100"
        :stroke-width="8"
      />
      <div class="score-info">
        <span class="score-label">综合评分</span>
        <span class="score-value">{{ analysis.overallScore }}分</span>
      </div>
    </div>

    <!-- 维度评分 -->
    <div class="dimension-scores">
      <div class="dimension-item">
        <span class="dimension-label">经验匹配</span>
        <el-progress
          :percentage="analysis.experienceMatch"
          :stroke-width="6"
          :color="getProgressColor(analysis.experienceMatch)"
        />
      </div>
      <div class="dimension-item">
        <span class="dimension-label">技能匹配</span>
        <el-progress
          :percentage="analysis.skillMatch"
          :stroke-width="6"
          :color="getProgressColor(analysis.skillMatch)"
        />
      </div>
      <div class="dimension-item">
        <span class="dimension-label">学历匹配</span>
        <el-progress
          :percentage="analysis.educationMatch"
          :stroke-width="6"
          :color="getProgressColor(analysis.educationMatch)"
        />
      </div>
    </div>

    <!-- 一句话总结 -->
    <div class="summary">
      <el-icon><ChatLineRound /></el-icon>
      <span>{{ analysis.summary }}</span>
    </div>

    <!-- 优势 -->
    <div class="strengths" v-if="analysis.strengths && analysis.strengths.length > 0">
      <div class="section-title">
        <el-icon><CircleCheck /></el-icon>
        优势
      </div>
      <ul class="list">
        <li v-for="(item, index) in analysis.strengths" :key="index">
          {{ item }}
        </li>
      </ul>
    </div>

    <!-- 不足 -->
    <div class="weaknesses" v-if="analysis.weaknesses && analysis.weaknesses.length > 0">
      <div class="section-title">
        <el-icon><WarningFilled /></el-icon>
        不足
      </div>
      <ul class="list">
        <li v-for="(item, index) in analysis.weaknesses" :key="index">
          {{ item }}
        </li>
      </ul>
    </div>

    <!-- 推荐意见 -->
    <div class="recommendation" :class="recommendationClass">
      <div class="rec-title">
        <el-icon><Trophy /></el-icon>
        推荐意见
      </div>
      <div class="rec-content">
        {{ recommendationText }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import {
  ChatLineRound,
  CircleCheck,
  WarningFilled,
  Trophy
} from '@element-plus/icons-vue'
import type { Resume, Analysis } from '../composables/useResumeStore'
import ScoreRing from './ScoreRing.vue'

const props = defineProps<{
  resume: Resume
}>()

// 分析结果
const analysis = computed<Analysis>(() => props.resume.analysis || {
  overallScore: 0,
  experienceMatch: 0,
  skillMatch: 0,
  educationMatch: 0,
  summary: '',
  strengths: [],
  weaknesses: [],
  recommendation: ''
})

// 推荐意见文本
const recommendationText = computed(() => {
  const rec = analysis.value.recommendation
  const texts: Record<string, string> = {
    strongly_recommended: '强烈推荐',
    recommended: '推荐',
    consider: '考虑',
    not_recommended: '不推荐'
  }
  return texts[rec] || rec || '暂无建议'
})

// 推荐意见样式
const recommendationClass = computed(() => {
  const rec = analysis.value.recommendation
  if (rec === 'strongly_recommended' || rec === 'recommended') return 'positive'
  if (rec === 'not_recommended') return 'negative'
  return 'neutral'
})

// 获取进度条颜色
function getProgressColor(score: number): string {
  if (score >= 80) return '#67c23a'
  if (score >= 60) return '#e6a23c'
  return '#f56c6c'
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

$bg-card: #ffffff;
$bg-section: #f8f9fb;
$border: #e2e7ed;
$divider: #eef1f5;

$text-primary: #1a2332;
$text-secondary: #5a6678;
$text-muted: #8b95a5;

$success: #16a34a;
$success-light: #dcfce7;
$warning: #d97706;
$warning-light: #fef3c7;
$error: #dc2626;
$error-light: #fee2e2;

$shadow-sm: 0 1px 3px rgba(30, 58, 95, 0.06);
$shadow-card: 0 2px 8px rgba(30, 58, 95, 0.06), 0 0 1px rgba(30, 58, 95, 0.1);

$radius-sm: 6px;
$radius-md: 10px;
$radius-lg: 14px;

// ==========================================
// 分析结果容器
// ==========================================
.analysis-result {
  padding: 24px;
  background-color: $bg-card;
  border-radius: $radius-lg;
  box-shadow: $shadow-card;
  border: 1px solid $divider;
}

// ==========================================
// 综合评分区域
// ==========================================
.overall-score {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 24px;
  padding-bottom: 20px;
  border-bottom: 1px solid $divider;

  .score-info {
    display: flex;
    flex-direction: column;
    gap: 4px;

    .score-label {
      font-size: 13px;
      color: $text-muted;
      font-weight: 500;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }

    .score-value {
      font-size: 28px;
      font-weight: 700;
      color: $text-primary;
      letter-spacing: -0.5px;
    }
  }
}

// ==========================================
// 维度评分
// ==========================================
.dimension-scores {
  margin-bottom: 20px;
  padding: 16px;
  background-color: $bg-section;
  border-radius: $radius-md;

  .dimension-item {
    display: flex;
    align-items: center;
    gap: 14px;
    margin-bottom: 14px;

    &:last-child {
      margin-bottom: 0;
    }

    .dimension-label {
      width: 80px;
      font-size: 13px;
      font-weight: 500;
      color: $text-secondary;
    }

    :deep(.el-progress) {
      flex: 1;

      .el-progress-bar__outer {
        background-color: rgba(30, 58, 95, 0.08);
        border-radius: 4px;
      }

      .el-progress-bar__inner {
        border-radius: 4px;
      }

      .el-progress__text {
        font-weight: 600;
        color: $text-primary;
      }
    }
  }
}

// ==========================================
// 一句话总结
// ==========================================
.summary {
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px;
  background-color: $primary-lighter;
  border-radius: $radius-md;
  margin-bottom: 20px;
  font-size: 14px;
  color: $text-secondary;
  line-height: 1.6;
  border-left: 4px solid $primary;

  .el-icon {
    margin-top: 2px;
    color: $primary;
    font-size: 18px;
    flex-shrink: 0;
  }
}

// ==========================================
// 优势区块
// ==========================================
.strengths {
  margin-bottom: 20px;
  padding: 16px;
  background-color: $success-light;
  border-radius: $radius-md;
  border-left: 4px solid $success;

  .section-title {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    font-weight: 600;
    color: $success;
    margin-bottom: 12px;

    .el-icon {
      font-size: 18px;
    }
  }

  .list {
    margin: 0;
    padding-left: 18px;

    li {
      font-size: 13px;
      color: darken($success, 15%);
      line-height: 1.7;
      margin-bottom: 6px;

      &:last-child {
        margin-bottom: 0;
      }

      &::marker {
        color: $success;
      }
    }
  }
}

// ==========================================
// 不足区块
// ==========================================
.weaknesses {
  margin-bottom: 20px;
  padding: 16px;
  background-color: $warning-light;
  border-radius: $radius-md;
  border-left: 4px solid $warning;

  .section-title {
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    font-weight: 600;
    color: $warning;
    margin-bottom: 12px;

    .el-icon {
      font-size: 18px;
    }
  }

  .list {
    margin: 0;
    padding-left: 18px;

    li {
      font-size: 13px;
      color: darken($warning, 15%);
      line-height: 1.7;
      margin-bottom: 6px;

      &:last-child {
        margin-bottom: 0;
      }

      &::marker {
        color: $warning;
      }
    }
  }
}

// ==========================================
// 推荐意见
// ==========================================
.recommendation {
  padding: 18px;
  border-radius: $radius-md;
  text-align: center;

  .rec-title {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    font-size: 13px;
    font-weight: 500;
    margin-bottom: 10px;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    opacity: 0.85;

    .el-icon {
      font-size: 18px;
    }
  }

  .rec-content {
    font-size: 20px;
    font-weight: 700;
    letter-spacing: 1px;
  }

  &.positive {
    background: linear-gradient(135deg, $success-light, lighten($success-light, 3%));
    border: 1px solid rgba($success, 0.2);

    .rec-title,
    .rec-content {
      color: $success;
    }
  }

  &.negative {
    background: linear-gradient(135deg, $error-light, lighten($error-light, 2%));
    border: 1px solid rgba($error, 0.2);

    .rec-title,
    .rec-content {
      color: $error;
    }
  }

  &.neutral {
    background: linear-gradient(135deg, $warning-light, lighten($warning-light, 3%));
    border: 1px solid rgba($warning, 0.2);

    .rec-title,
    .rec-content {
      color: $warning;
    }
  }
}
</style>
