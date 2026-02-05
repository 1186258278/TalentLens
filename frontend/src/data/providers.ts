// AI 服务商预设数据
export interface ProviderGuide {
  title: string
  steps: string[]
  link: string
  pricing: string
}

export interface Provider {
  id: string
  name: string
  baseURL: string
  models: Array<{ id: string; name: string; recommended?: boolean }>
  defaultModel: string
  guide: ProviderGuide
  recommended?: boolean
}

// 服务商预设配置
export const providers: Provider[] = [
  {
    id: 'deepseek',
    name: 'DeepSeek',
    baseURL: 'https://api.deepseek.com',
    recommended: true,
    models: [
      { id: 'deepseek-chat', name: 'DeepSeek Chat', recommended: true },
      { id: 'deepseek-coder', name: 'DeepSeek Coder' },
      { id: 'deepseek-reasoner', name: 'DeepSeek Reasoner (R1)' }
    ],
    defaultModel: 'deepseek-chat',
    guide: {
      title: 'providers.deepseek.guideTitle',
      steps: [
        'providers.deepseek.steps.0',
        'providers.deepseek.steps.1',
        'providers.deepseek.steps.2',
        'providers.deepseek.steps.3',
        'providers.deepseek.steps.4'
      ],
      link: 'https://platform.deepseek.com/api_keys',
      pricing: 'providers.deepseek.pricing'
    }
  },
  {
    id: 'openai',
    name: 'OpenAI',
    baseURL: 'https://api.openai.com/v1',
    models: [
      { id: 'gpt-4o', name: 'GPT-4o' },
      { id: 'gpt-4o-mini', name: 'GPT-4o Mini', recommended: true },
      { id: 'gpt-4-turbo', name: 'GPT-4 Turbo' },
      { id: 'gpt-3.5-turbo', name: 'GPT-3.5 Turbo' }
    ],
    defaultModel: 'gpt-4o-mini',
    guide: {
      title: 'providers.openai.guideTitle',
      steps: [
        'providers.openai.steps.0',
        'providers.openai.steps.1',
        'providers.openai.steps.2',
        'providers.openai.steps.3',
        'providers.openai.steps.4'
      ],
      link: 'https://platform.openai.com/api-keys',
      pricing: 'providers.openai.pricing'
    }
  },
  {
    id: 'zhipu',
    name: '智谱 AI',
    baseURL: 'https://open.bigmodel.cn/api/paas/v4',
    models: [
      { id: 'glm-4-plus', name: 'GLM-4 Plus' },
      { id: 'glm-4-flash', name: 'GLM-4 Flash', recommended: true },
      { id: 'glm-4', name: 'GLM-4' },
      { id: 'glm-4-air', name: 'GLM-4 Air' }
    ],
    defaultModel: 'glm-4-flash',
    guide: {
      title: 'providers.zhipu.guideTitle',
      steps: [
        'providers.zhipu.steps.0',
        'providers.zhipu.steps.1',
        'providers.zhipu.steps.2',
        'providers.zhipu.steps.3',
        'providers.zhipu.steps.4'
      ],
      link: 'https://open.bigmodel.cn/usercenter/apikeys',
      pricing: 'providers.zhipu.pricing'
    }
  },
  {
    id: 'moonshot',
    name: 'Moonshot',
    baseURL: 'https://api.moonshot.cn/v1',
    models: [
      { id: 'moonshot-v1-8k', name: 'Moonshot V1 8K', recommended: true },
      { id: 'moonshot-v1-32k', name: 'Moonshot V1 32K' },
      { id: 'moonshot-v1-128k', name: 'Moonshot V1 128K' }
    ],
    defaultModel: 'moonshot-v1-8k',
    guide: {
      title: 'providers.moonshot.guideTitle',
      steps: [
        'providers.moonshot.steps.0',
        'providers.moonshot.steps.1',
        'providers.moonshot.steps.2',
        'providers.moonshot.steps.3',
        'providers.moonshot.steps.4'
      ],
      link: 'https://platform.moonshot.cn/console/api-keys',
      pricing: 'providers.moonshot.pricing'
    }
  },
  {
    id: 'siliconflow',
    name: '硅基流动',
    baseURL: 'https://api.siliconflow.cn/v1',
    models: [
      { id: 'Qwen/Qwen2.5-7B-Instruct', name: 'Qwen2.5-7B', recommended: true },
      { id: 'Qwen/Qwen2.5-72B-Instruct', name: 'Qwen2.5-72B' },
      { id: 'deepseek-ai/DeepSeek-V3', name: 'DeepSeek-V3' },
      { id: 'THUDM/glm-4-9b-chat', name: 'GLM-4-9B' }
    ],
    defaultModel: 'Qwen/Qwen2.5-7B-Instruct',
    guide: {
      title: 'providers.siliconflow.guideTitle',
      steps: [
        'providers.siliconflow.steps.0',
        'providers.siliconflow.steps.1',
        'providers.siliconflow.steps.2',
        'providers.siliconflow.steps.3',
        'providers.siliconflow.steps.4'
      ],
      link: 'https://cloud.siliconflow.cn/account/ak',
      pricing: 'providers.siliconflow.pricing'
    }
  },
  {
    id: 'custom',
    name: '自定义',
    baseURL: '',
    models: [],
    defaultModel: '',
    guide: {
      title: 'providers.custom.guideTitle',
      steps: [
        'providers.custom.steps.0',
        'providers.custom.steps.1',
        'providers.custom.steps.2',
        'providers.custom.steps.3'
      ],
      link: '',
      pricing: 'providers.custom.pricing'
    }
  }
]

// 根据ID获取服务商
export function getProviderById(id: string): Provider | undefined {
  return providers.find(p => p.id === id)
}

// 获取推荐的服务商
export function getRecommendedProvider(): Provider {
  return providers.find(p => p.recommended) || providers[0]
}
