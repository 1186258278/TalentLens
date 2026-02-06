// 繁體中文語言包
export default {
  common: {
    save: '儲存設定',
    reset: '重設',
    back: '返回',
    confirm: '確認',
    cancel: '取消',
    delete: '刪除',
    edit: '編輯',
    add: '新增',
    search: '搜尋',
    loading: '載入中...',
    success: '成功',
    error: '錯誤',
    warning: '警告',
    info: '提示',
    year: '年',
    years: '年',
    custom: '自訂'
  },

  app: {
    title: 'TalentLens',
    subtitle: 'AI 智慧履歷審核工具'
  },

  home: {
    resumeList: '履歷列表',
    preview: '預覽',
    parsedContent: '解析內容',
    parsedContentHint: '以下是 PDF/文件解析後 AI 實際看到的文字內容',
    loadingContent: '正在載入解析內容...',
    noContent: '暫無解析內容，請重新拖入履歷',
    contentLength: '共 {count} 個字元',
    startAnalysis: '開始分析',
    reAnalyzeAll: '全部重新分析',
    exportReport: '匯出報告',
    clear: '清除',
    noAnalyzedResumes: '沒有已分析的履歷',
    dragHint: '將履歷拖曳到此處開始',
    selectToPreview: '選擇一份履歷查看預覽',
    dropFiles: '釋放檔案',
    dropHere: '將履歷拖曳到此處',
    supportedFormats: '支援 PDF / Word / 圖片格式'
  },

  settings: {
    title: '設定',
    aiConfig: 'AI 模型設定',
    aiConfigDesc: '配置 AI 服務提供商和認證資訊，用於履歷智慧分析',
    jobRequirements: '職位需求',
    jobRequirementsDesc: '設定招聘職位的具體要求，AI 將據此評估履歷匹配度',
    analysisSettings: '分析設定',
    analysisSettingsDesc: '調整分析行為和效能參數',
    language: '介面語言',
    saved: '設定已儲存',
    resetDone: '設定已重設',
    resetAllSettings: '重設所有設定',
    dataPath: '資料儲存路徑',
    openDataDir: '開啟目錄'
  },

  ai: {
    provider: '服務商',
    selectProvider: '選擇服務商',
    apiKey: 'API 金鑰',
    apiKeyPlaceholder: '請輸入 API 金鑰',
    model: '模型',
    selectModel: '選擇模型',
    testConnection: '測試連線',
    testing: '測試中...',
    connectionSuccess: '連線成功',
    connectionFailed: '連線失敗',
    recommended: '推薦',
    customProvider: '自訂服務商',
    baseUrl: 'API 位址',
    baseUrlPlaceholder: '請輸入 API 位址'
  },

  providers: {
    deepseek: {
      name: 'DeepSeek',
      description: '中國模型，性價比高',
      guideTitle: '如何取得 DeepSeek API Key',
      steps: [
        '前往 DeepSeek 開放平台',
        '註冊/登入帳號',
        '進入「API Keys」頁面',
        '點擊「建立 API Key」',
        '複製產生的金鑰貼上到此處'
      ],
      pricing: '新用戶贈送 500 萬 tokens，之後按量計費',
      openConsole: '開啟 DeepSeek 控制台'
    },
    openai: {
      name: 'OpenAI',
      description: 'ChatGPT 官方介面',
      guideTitle: '如何取得 OpenAI API Key',
      steps: [
        '前往 OpenAI 官網 platform.openai.com',
        '註冊/登入帳號（可能需要國外手機號）',
        '進入「API Keys」頁面',
        '點擊「Create new secret key」',
        '複製產生的金鑰貼上到此處'
      ],
      pricing: '按量計費，GPT-4o-mini 約 $0.15/100萬tokens',
      openConsole: '開啟 OpenAI 控制台'
    },
    zhipu: {
      name: '智譜 AI',
      description: '清華系中國大模型',
      guideTitle: '如何取得智譜 AI API Key',
      steps: [
        '前往智譜 AI 開放平台 open.bigmodel.cn',
        '註冊/登入帳號',
        '進入「API 金鑰」頁面',
        '建立新的 API 金鑰',
        '複製產生的金鑰貼上到此處'
      ],
      pricing: '新用戶贈送 tokens，GLM-4-Flash 免費',
      openConsole: '開啟智譜 AI 控制台'
    },
    moonshot: {
      name: 'Moonshot',
      description: '月之暗面 Kimi 大模型',
      guideTitle: '如何取得 Moonshot API Key',
      steps: [
        '前往 Moonshot 開放平台 platform.moonshot.cn',
        '註冊/登入帳號',
        '進入「API Key 管理」頁面',
        '建立新的 API Key',
        '複製產生的金鑰貼上到此處'
      ],
      pricing: '新用戶贈送 15 元額度',
      openConsole: '開啟 Moonshot 控制台'
    },
    siliconflow: {
      name: '矽基流動',
      description: '中國模型聚合平台，多種模型可選',
      guideTitle: '如何取得矽基流動 API Key',
      steps: [
        '前往矽基流動官網 siliconflow.cn',
        '註冊/登入帳號',
        '進入「API 金鑰」頁面',
        '建立新的 API 金鑰',
        '複製產生的金鑰貼上到此處'
      ],
      pricing: '部分模型免費，其他按量計費',
      openConsole: '開啟矽基流動控制台'
    },
    custom: {
      name: '自訂',
      description: '使用 OpenAI 相容的 API',
      guideTitle: '自訂服務商設定',
      steps: [
        '確保您的服務商支援 OpenAI 相容介面',
        '取得 API 位址（Base URL）',
        '取得 API 金鑰',
        '填寫到下方設定中'
      ],
      pricing: '根據您的服務商定價'
    }
  },

  job: {
    title: '職位名稱',
    titlePlaceholder: '請輸入職位名稱',
    requiredSkills: '必備技能',
    preferredSkills: '加分技能',
    experienceYears: '經驗要求',
    educationLevel: '學歷要求',
    description: '職位描述',
    quickSelect: '快速選擇範本',
    customJob: '自訂職位',
    currentJob: '目前職位',
    categories: {
      tech: '技術類',
      product: '產品類',
      operation: '營運類',
      design: '設計類',
      sales: '銷售類',
      hr: '人力資源',
      finance: '財務類'
    },
    education: {
      any: '不限',
      college: '專科',
      bachelor: '學士',
      master: '碩士',
      phd: '博士'
    }
  },

  analysis: {
    concurrent: '並行數量',
    concurrentHint: '個同時分析',
    autoStart: '自動開始分析',
    autoStartDesc: '新增履歷後自動開始分析',
    pending: '待分析',
    analyzing: '分析中...',
    preparingPrompt: '準備分析...',
    callingAI: 'AI 分析中...',
    parsingResult: '解析結果...',
    batchProgress: '正在分析 {current}/{total} 份履歷',
    done: '已完成',
    failed: '分析失敗',
    reAnalyze: '重新分析',
    totalScore: '總分',
    overallScore: '綜合評分',
    experienceMatch: '經驗匹配度',
    skillMatch: '技能匹配度',
    educationMatch: '學歷匹配度',
    strengths: '優勢亮點',
    weaknesses: '待改進',
    summary: 'AI 分析總結',
    recommendation: '推薦意見',
    pendingHint: '點擊「開始分析」後將對此履歷進行 AI 智慧分析',
    analyzingHint: 'AI 正在分析此履歷，請稍候...',
    errorHint: '分析出錯，請重試或檢查 AI 設定',
    recommendations: {
      strong_recommend: '強烈推薦',
      recommend: '推薦',
      consider: '可考慮',
      not_recommend: '不推薦'
    }
  },

  file: {
    pdf: 'PDF文件',
    word: 'Word文件',
    image: '圖片',
    unknown: '未知類型'
  },

  guide: {
    title: 'AI 設定提示',
    notConfigured: '尚未設定 AI 服務',
    description: '在開始分析履歷之前，您需要先設定 AI 模型的 API 金鑰。設定非常簡單，只需幾步即可完成。',
    tip1: '支援 DeepSeek、OpenAI、智譜等主流服務商',
    tip2: '大多數服務商提供免費額度，註冊即可使用',
    tip3: '設定一次後，後續無需重複設定',
    recommended: '推薦服務商',
    deepseekDesc: '中國模型，性價比高，新用戶有免費額度',
    freeQuota: '新用戶免費',
    goSettings: '前往設定',
    noResumesToAnalyze: '沒有待分析的履歷',
    configSuccess: 'AI 設定成功，可以開始分析了',
    analysisStarted: '開始分析履歷...'
  },

  about: {
    title: '關於',
    version: '目前版本',
    checkUpdate: '檢查更新',
    checking: '正在檢查...',
    latestVersion: '已是最新版本',
    newVersion: '發現新版本 v{version}',
    noUpdate: '目前已是最新版本',
    download: '前往下載',
    updateError: '檢查更新失敗',
    links: '相關連結',
    github: 'GitHub 倉庫',
    website: '產品官網',
    company: '公司官網',
    copyright: '武漢晴辰天下網路科技有限公司'
  },

  project: {
    title: '招聘專案',
    create: '新建專案',
    createFirst: '建立第一個招聘專案',
    createDialogTitle: '新建招聘專案',
    createSuccess: '專案建立成功',
    deleteConfirm: '確定刪除此專案？專案下的所有履歷也將被刪除。',
    deleteSuccess: '專案已刪除',
    projectName: '專案名稱',
    projectNamePlaceholder: '如：Go高級開發 - 2月批次',
    empty: '還沒有招聘專案，建立一個開始吧',
    statusDraft: '準備中',
    statusAnalyzing: '分析中',
    statusCompleted: '已完成',
    ranking: '排名對比',
    exportExcel: '匯出 Excel',
    exportSuccess: '報告已匯出',
    exportFailed: '匯出失敗',
    openExportDir: '開啟匯出目錄',
    stats: '專案統計',
    totalResumes: '總履歷數',
    analyzedCount: '已分析',
    avgScore: '平均分',
    recommendCount: '推薦人數',
    backToList: '返回專案列表'
  }
}
