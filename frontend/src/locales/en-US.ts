// English Language Pack
export default {
  common: {
    save: 'Save Settings',
    reset: 'Reset',
    back: 'Back',
    confirm: 'Confirm',
    cancel: 'Cancel',
    delete: 'Delete',
    edit: 'Edit',
    add: 'Add',
    search: 'Search',
    loading: 'Loading...',
    success: 'Success',
    error: 'Error',
    warning: 'Warning',
    info: 'Info',
    year: 'year',
    years: 'years',
    custom: 'Custom'
  },

  app: {
    title: 'TalentLens',
    subtitle: 'AI-Powered Resume Review Tool'
  },

  home: {
    resumeList: 'Resume List',
    preview: 'Preview',
    parsedContent: 'Parsed Content',
    parsedContentHint: 'Text content that AI actually sees after parsing the PDF/document',
    loadingContent: 'Loading parsed content...',
    noContent: 'No parsed content, please re-import the resume',
    contentLength: '{count} characters',
    startAnalysis: 'Start Analysis',
    reAnalyzeAll: 'Re-analyze All',
    exportReport: 'Export Report',
    clear: 'Clear',
    noAnalyzedResumes: 'No analyzed resumes',
    dragHint: 'Drag resumes here to start',
    selectToPreview: 'Select a resume to preview',
    dropFiles: 'Drop files',
    dropHere: 'Drag resumes here',
    supportedFormats: 'Supports PDF / Word / Image formats'
  },

  settings: {
    title: 'Settings',
    aiConfig: 'AI Model Configuration',
    aiConfigDesc: 'Configure AI service provider and authentication for intelligent resume analysis',
    jobRequirements: 'Job Requirements',
    jobRequirementsDesc: 'Set specific job requirements for AI to evaluate resume match',
    analysisSettings: 'Analysis Settings',
    analysisSettingsDesc: 'Adjust analysis behavior and performance parameters',
    language: 'Language',
    saved: 'Settings saved',
    resetDone: 'Settings reset',
    resetAllSettings: 'Reset All Settings',
    dataPath: 'Data Storage Path',
    openDataDir: 'Open Folder'
  },

  ai: {
    provider: 'Provider',
    selectProvider: 'Select Provider',
    apiKey: 'API Key',
    apiKeyPlaceholder: 'Enter your API key',
    model: 'Model',
    selectModel: 'Select Model',
    testConnection: 'Test Connection',
    testing: 'Testing...',
    connectionSuccess: 'Connection successful',
    connectionFailed: 'Connection failed',
    recommended: 'Recommended',
    customProvider: 'Custom Provider',
    baseUrl: 'API URL',
    baseUrlPlaceholder: 'Enter API URL'
  },

  providers: {
    deepseek: {
      name: 'DeepSeek',
      description: 'Cost-effective Chinese AI model',
      guideTitle: 'How to get DeepSeek API Key',
      steps: [
        'Visit DeepSeek platform',
        'Sign up / Sign in',
        'Go to "API Keys" page',
        'Click "Create API Key"',
        'Copy and paste the key here'
      ],
      pricing: 'New users get 5M tokens free, then pay-as-you-go',
      openConsole: 'Open DeepSeek Console'
    },
    openai: {
      name: 'OpenAI',
      description: 'Official ChatGPT API',
      guideTitle: 'How to get OpenAI API Key',
      steps: [
        'Visit platform.openai.com',
        'Sign up / Sign in',
        'Go to "API Keys" page',
        'Click "Create new secret key"',
        'Copy and paste the key here'
      ],
      pricing: 'Pay-as-you-go, GPT-4o-mini ~$0.15/1M tokens',
      openConsole: 'Open OpenAI Console'
    },
    zhipu: {
      name: 'Zhipu AI',
      description: 'Chinese LLM from Tsinghua',
      guideTitle: 'How to get Zhipu AI API Key',
      steps: [
        'Visit open.bigmodel.cn',
        'Sign up / Sign in',
        'Go to "API Keys" page',
        'Create a new API key',
        'Copy and paste the key here'
      ],
      pricing: 'Free tokens for new users, GLM-4-Flash is free',
      openConsole: 'Open Zhipu AI Console'
    },
    moonshot: {
      name: 'Moonshot',
      description: 'Kimi AI model',
      guideTitle: 'How to get Moonshot API Key',
      steps: [
        'Visit platform.moonshot.cn',
        'Sign up / Sign in',
        'Go to "API Key Management" page',
        'Create a new API Key',
        'Copy and paste the key here'
      ],
      pricing: 'New users get ¥15 credit',
      openConsole: 'Open Moonshot Console'
    },
    siliconflow: {
      name: 'SiliconFlow',
      description: 'Chinese AI model aggregation platform',
      guideTitle: 'How to get SiliconFlow API Key',
      steps: [
        'Visit siliconflow.cn',
        'Sign up / Sign in',
        'Go to "API Keys" page',
        'Create a new API key',
        'Copy and paste the key here'
      ],
      pricing: 'Some models free, others pay-as-you-go',
      openConsole: 'Open SiliconFlow Console'
    },
    custom: {
      name: 'Custom',
      description: 'Use OpenAI-compatible API',
      guideTitle: 'Custom Provider Setup',
      steps: [
        'Make sure your provider supports OpenAI-compatible API',
        'Get the API URL (Base URL)',
        'Get your API key',
        'Fill in the configuration below'
      ],
      pricing: 'Based on your provider pricing'
    }
  },

  job: {
    title: 'Job Title',
    titlePlaceholder: 'Enter job title',
    requiredSkills: 'Required Skills',
    preferredSkills: 'Preferred Skills',
    experienceYears: 'Experience Required',
    educationLevel: 'Education Level',
    description: 'Job Description',
    quickSelect: 'Quick Select Template',
    customJob: 'Custom Job',
    currentJob: 'Current Job',
    categories: {
      tech: 'Technology',
      product: 'Product',
      operation: 'Operations',
      design: 'Design',
      sales: 'Sales',
      hr: 'Human Resources',
      finance: 'Finance'
    },
    education: {
      any: 'Any',
      college: 'College',
      bachelor: 'Bachelor',
      master: 'Master',
      phd: 'PhD'
    }
  },

  analysis: {
    concurrent: 'Concurrent',
    concurrentHint: 'simultaneous analyses',
    autoStart: 'Auto-start analysis',
    autoStartDesc: 'Automatically start analysis after adding resumes',
    pending: 'Pending',
    analyzing: 'Analyzing...',
    preparingPrompt: 'Preparing...',
    callingAI: 'AI analyzing...',
    parsingResult: 'Parsing result...',
    batchProgress: 'Analyzing {current}/{total} resumes',
    done: 'Completed',
    failed: 'Analysis failed',
    reAnalyze: 'Re-analyze',
    totalScore: 'Score',
    overallScore: 'Overall Score',
    experienceMatch: 'Experience Match',
    skillMatch: 'Skill Match',
    educationMatch: 'Education Match',
    strengths: 'Strengths',
    weaknesses: 'Areas to Improve',
    summary: 'AI Analysis Summary',
    recommendation: 'Recommendation',
    pendingHint: 'Click "Start Analysis" to analyze this resume with AI',
    analyzingHint: 'AI is analyzing this resume, please wait...',
    errorHint: 'Analysis failed, please retry or check AI configuration',
    recommendations: {
      strong_recommend: 'Strongly Recommended',
      recommend: 'Recommended',
      consider: 'Consider',
      not_recommend: 'Not Recommended'
    }
  },

  file: {
    pdf: 'PDF Document',
    word: 'Word Document',
    image: 'Image',
    unknown: 'Unknown Type'
  },

  guide: {
    title: 'AI Configuration Required',
    notConfigured: 'AI Service Not Configured',
    description: 'Before analyzing resumes, you need to configure the AI model API key. The setup is simple and takes just a few steps.',
    tip1: 'Supports DeepSeek, OpenAI, Zhipu and other providers',
    tip2: 'Most providers offer free credits for new users',
    tip3: 'Configure once, use forever',
    recommended: 'Recommended Provider',
    deepseekDesc: 'Cost-effective, new users get free credits',
    freeQuota: 'Free for new users',
    goSettings: 'Go to Settings',
    noResumesToAnalyze: 'No resumes to analyze',
    configSuccess: 'AI configured successfully, ready to analyze',
    analysisStarted: 'Starting resume analysis...'
  },

  about: {
    title: 'About',
    version: 'Current Version',
    checkUpdate: 'Check for Updates',
    checking: 'Checking...',
    latestVersion: 'Up to date',
    newVersion: 'New version v{version} available',
    noUpdate: 'You are using the latest version',
    download: 'Download',
    updateError: 'Failed to check for updates',
    links: 'Links',
    github: 'GitHub Repository',
    website: 'Product Website',
    company: 'Company Website',
    copyright: 'Wuhan Qingchen Tianxia Network Technology Co., Ltd.'
  },

  project: {
    title: 'Recruitment Projects',
    create: 'New Project',
    createFirst: 'Create your first recruitment project',
    createDialogTitle: 'New Recruitment Project',
    createSuccess: 'Project created',
    deleteConfirm: 'Delete this project? All resumes in the project will also be deleted.',
    deleteSuccess: 'Project deleted',
    projectName: 'Project Name',
    projectNamePlaceholder: 'e.g. Senior Go Dev - Feb Batch',
    empty: 'No recruitment projects yet. Create one to get started.',
    statusDraft: 'Draft',
    statusAnalyzing: 'Analyzing',
    statusCompleted: 'Completed',
    ranking: 'Ranking',
    exportExcel: 'Export Excel',
    exportSuccess: 'Report exported',
    exportFailed: 'Export failed',
    openExportDir: 'Open export folder',
    stats: 'Statistics',
    totalResumes: 'Total Resumes',
    analyzedCount: 'Analyzed',
    avgScore: 'Avg Score',
    recommendCount: 'Recommended',
    backToList: 'Back to Projects'
  }
}
