// 简体中文语言包
export default {
  common: {
    save: '保存设置',
    reset: '重置',
    back: '返回',
    confirm: '确认',
    cancel: '取消',
    delete: '删除',
    edit: '编辑',
    add: '添加',
    search: '搜索',
    loading: '加载中...',
    success: '成功',
    error: '错误',
    warning: '警告',
    info: '提示',
    year: '年',
    years: '年',
    custom: '自定义'
  },

  app: {
    title: 'TalentLens',
    subtitle: 'AI 智能简历审核工具'
  },

  home: {
    resumeList: '简历列表',
    preview: '预览',
    startAnalysis: '开始分析',
    reAnalyzeAll: '全部重新分析',
    exportReport: '导出报告',
    clear: '清除',
    noAnalyzedResumes: '没有已分析的简历',
    dragHint: '将简历拖拽到此处开始',
    selectToPreview: '选择一个简历查看预览',
    dropFiles: '释放文件',
    dropHere: '将简历拖拽到此处',
    supportedFormats: '支持 PDF / Word / 图片格式'
  },

  settings: {
    title: '设置',
    aiConfig: 'AI 模型配置',
    aiConfigDesc: '配置 AI 服务提供商和认证信息，用于简历智能分析',
    jobRequirements: '岗位需求',
    jobRequirementsDesc: '设置招聘岗位的具体要求，AI 将据此评估简历匹配度',
    analysisSettings: '分析设置',
    analysisSettingsDesc: '调整分析行为和性能参数',
    language: '界面语言',
    saved: '设置已保存',
    resetDone: '设置已重置',
    resetAllSettings: '重置所有设置',
    dataPath: '数据存储路径',
    openDataDir: '打开目录'
  },

  ai: {
    provider: '服务商',
    selectProvider: '选择服务商',
    apiKey: 'API 密钥',
    apiKeyPlaceholder: '请输入 API 密钥',
    model: '模型',
    selectModel: '选择模型',
    testConnection: '测试连接',
    testing: '测试中...',
    connectionSuccess: '连接成功',
    connectionFailed: '连接失败',
    recommended: '推荐',
    customProvider: '自定义服务商',
    baseUrl: 'API 地址',
    baseUrlPlaceholder: '请输入 API 地址'
  },

  providers: {
    deepseek: {
      name: 'DeepSeek',
      description: '国产模型，性价比高',
      guideTitle: '如何获取 DeepSeek API Key',
      steps: [
        '访问 DeepSeek 开放平台',
        '注册/登录账号',
        '进入「API Keys」页面',
        '点击「创建 API Key」',
        '复制生成的密钥粘贴到此处'
      ],
      pricing: '新用户赠送 500 万 tokens，之后按量计费',
      openConsole: '打开 DeepSeek 控制台'
    },
    openai: {
      name: 'OpenAI',
      description: 'ChatGPT 官方接口',
      guideTitle: '如何获取 OpenAI API Key',
      steps: [
        '访问 OpenAI 官网 platform.openai.com',
        '注册/登录账号（可能需要国外手机号）',
        '进入「API Keys」页面',
        '点击「Create new secret key」',
        '复制生成的密钥粘贴到此处'
      ],
      pricing: '按量计费，GPT-4o-mini 约 $0.15/100万tokens',
      openConsole: '打开 OpenAI 控制台'
    },
    zhipu: {
      name: '智谱 AI',
      description: '清华系国产大模型',
      guideTitle: '如何获取智谱 AI API Key',
      steps: [
        '访问智谱 AI 开放平台 open.bigmodel.cn',
        '注册/登录账号',
        '进入「API 密钥」页面',
        '创建新的 API 密钥',
        '复制生成的密钥粘贴到此处'
      ],
      pricing: '新用户赠送 tokens，GLM-4-Flash 免费',
      openConsole: '打开智谱 AI 控制台'
    },
    moonshot: {
      name: 'Moonshot',
      description: '月之暗面 Kimi 大模型',
      guideTitle: '如何获取 Moonshot API Key',
      steps: [
        '访问 Moonshot 开放平台 platform.moonshot.cn',
        '注册/登录账号',
        '进入「API Key 管理」页面',
        '创建新的 API Key',
        '复制生成的密钥粘贴到此处'
      ],
      pricing: '新用户赠送 15 元额度',
      openConsole: '打开 Moonshot 控制台'
    },
    siliconflow: {
      name: '硅基流动',
      description: '国产模型聚合平台，多种模型可选',
      guideTitle: '如何获取硅基流动 API Key',
      steps: [
        '访问硅基流动官网 siliconflow.cn',
        '注册/登录账号',
        '进入「API 密钥」页面',
        '创建新的 API 密钥',
        '复制生成的密钥粘贴到此处'
      ],
      pricing: '部分模型免费，其他按量计费',
      openConsole: '打开硅基流动控制台'
    },
    custom: {
      name: '自定义',
      description: '使用 OpenAI 兼容的 API',
      guideTitle: '自定义服务商配置',
      steps: [
        '确保您的服务商支持 OpenAI 兼容接口',
        '获取 API 地址（Base URL）',
        '获取 API 密钥',
        '填写到下方配置中'
      ],
      pricing: '根据您的服务商定价'
    }
  },

  job: {
    title: '岗位名称',
    titlePlaceholder: '请输入岗位名称',
    requiredSkills: '必备技能',
    preferredSkills: '加分技能',
    experienceYears: '经验要求',
    educationLevel: '学历要求',
    description: '岗位描述',
    quickSelect: '快速选择模板',
    customJob: '自定义岗位',
    currentJob: '当前岗位',
    categories: {
      tech: '技术类',
      product: '产品类',
      operation: '运营类',
      design: '设计类',
      sales: '销售类',
      hr: '人力资源',
      finance: '财务类'
    },
    education: {
      any: '不限',
      college: '大专',
      bachelor: '本科',
      master: '硕士',
      phd: '博士'
    }
  },

  analysis: {
    concurrent: '并发数量',
    concurrentHint: '个同时分析',
    autoStart: '自动开始分析',
    autoStartDesc: '添加简历后自动开始分析',
    pending: '待分析',
    analyzing: '分析中...',
    preparingPrompt: '准备分析...',
    callingAI: 'AI 分析中...',
    parsingResult: '解析结果...',
    batchProgress: '正在分析 {current}/{total} 份简历',
    done: '已完成',
    failed: '分析失败',
    reAnalyze: '重新分析',
    totalScore: '总分',
    overallScore: '综合评分',
    experienceMatch: '经验匹配度',
    skillMatch: '技能匹配度',
    educationMatch: '学历匹配度',
    strengths: '优势亮点',
    weaknesses: '待改进',
    summary: 'AI 分析总结',
    recommendation: '推荐意见',
    pendingHint: '点击「开始分析」后将对此简历进行 AI 智能分析',
    analyzingHint: 'AI 正在分析此简历，请稍候...',
    errorHint: '分析出错，请重试或检查 AI 配置',
    recommendations: {
      strong_recommend: '强烈推荐',
      recommend: '推荐',
      consider: '可考虑',
      not_recommend: '不推荐'
    }
  },

  file: {
    pdf: 'PDF文档',
    word: 'Word文档',
    image: '图片',
    unknown: '未知类型'
  },

  guide: {
    title: 'AI 配置提示',
    notConfigured: '尚未配置 AI 服务',
    description: '在开始分析简历之前，您需要先配置 AI 模型的 API 密钥。配置非常简单，只需几步即可完成。',
    tip1: '支持 DeepSeek、OpenAI、智谱等主流服务商',
    tip2: '大多数服务商提供免费额度，注册即可使用',
    tip3: '配置一次后，后续无需重复设置',
    recommended: '推荐服务商',
    deepseekDesc: '国产模型，性价比高，新用户有免费额度',
    freeQuota: '新用户免费',
    goSettings: '前往配置',
    noResumesToAnalyze: '没有待分析的简历',
    configSuccess: 'AI 配置成功，可以开始分析了',
    analysisStarted: '开始分析简历...'
  },

  project: {
    title: '招聘项目',
    create: '新建项目',
    createFirst: '创建第一个招聘项目',
    createDialogTitle: '新建招聘项目',
    createSuccess: '项目创建成功',
    deleteConfirm: '确定删除此项目？项目下的所有简历也将被删除。',
    deleteSuccess: '项目已删除',
    projectName: '项目名称',
    projectNamePlaceholder: '如：Go高级开发 - 2月批次',
    empty: '还没有招聘项目，创建一个开始吧',
    statusDraft: '准备中',
    statusAnalyzing: '分析中',
    statusCompleted: '已完成',
    ranking: '排名对比',
    exportExcel: '导出 Excel',
    exportSuccess: '报告已导出',
    exportFailed: '导出失败',
    openExportDir: '打开导出目录',
    stats: '项目统计',
    totalResumes: '总简历数',
    analyzedCount: '已分析',
    avgScore: '平均分',
    recommendCount: '推荐人数',
    backToList: '返回项目列表'
  }
}
