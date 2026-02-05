// 岗位模板预设数据
export interface JobPreset {
  id: string
  name: string
  nameEn: string
  category: string
  requiredSkills: string[]
  preferredSkills?: string[]
  experienceYears: number
  educationLevel: string
  description?: string
}

export interface JobCategory {
  id: string
  labelKey: string // i18n key
  presets: JobPreset[]
}

// 技术类岗位
const techJobs: JobPreset[] = [
  {
    id: 'frontend-senior',
    name: '高级前端工程师',
    nameEn: 'Senior Frontend Engineer',
    category: 'tech',
    requiredSkills: ['JavaScript', 'TypeScript', 'React', 'Vue', 'CSS', 'HTML'],
    preferredSkills: ['Node.js', 'Webpack', 'Vite', 'GraphQL', '性能优化'],
    experienceYears: 5,
    educationLevel: '本科',
    description: '负责公司核心产品的前端架构设计与开发'
  },
  {
    id: 'frontend-junior',
    name: '前端工程师',
    nameEn: 'Frontend Engineer',
    category: 'tech',
    requiredSkills: ['JavaScript', 'HTML', 'CSS', 'Vue'],
    preferredSkills: ['TypeScript', 'React', 'Element Plus'],
    experienceYears: 2,
    educationLevel: '本科'
  },
  {
    id: 'backend-go',
    name: '高级Go开发工程师',
    nameEn: 'Senior Go Developer',
    category: 'tech',
    requiredSkills: ['Go', 'MySQL', 'Redis', 'Docker', 'Linux'],
    preferredSkills: ['Kubernetes', 'gRPC', 'Kafka', '微服务', '分布式系统'],
    experienceYears: 5,
    educationLevel: '本科',
    description: '负责后端服务架构设计与核心功能开发'
  },
  {
    id: 'backend-java',
    name: '高级Java开发工程师',
    nameEn: 'Senior Java Developer',
    category: 'tech',
    requiredSkills: ['Java', 'Spring Boot', 'MySQL', 'Redis', 'MyBatis'],
    preferredSkills: ['微服务', 'Dubbo', 'RocketMQ', 'Kubernetes'],
    experienceYears: 5,
    educationLevel: '本科'
  },
  {
    id: 'backend-python',
    name: 'Python开发工程师',
    nameEn: 'Python Developer',
    category: 'tech',
    requiredSkills: ['Python', 'Django', 'Flask', 'MySQL', 'Redis'],
    preferredSkills: ['FastAPI', 'Celery', 'Docker', '机器学习'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'fullstack',
    name: '全栈工程师',
    nameEn: 'Full Stack Engineer',
    category: 'tech',
    requiredSkills: ['JavaScript', 'Node.js', 'Vue', 'MySQL', 'Docker'],
    preferredSkills: ['TypeScript', 'React', 'MongoDB', 'AWS'],
    experienceYears: 4,
    educationLevel: '本科'
  },
  {
    id: 'devops',
    name: '运维工程师',
    nameEn: 'DevOps Engineer',
    category: 'tech',
    requiredSkills: ['Linux', 'Docker', 'Kubernetes', 'CI/CD', 'Shell'],
    preferredSkills: ['Terraform', 'Ansible', 'Prometheus', 'ELK'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'ai-engineer',
    name: 'AI算法工程师',
    nameEn: 'AI Engineer',
    category: 'tech',
    requiredSkills: ['Python', '机器学习', '深度学习', 'PyTorch', 'TensorFlow'],
    preferredSkills: ['NLP', 'CV', 'LLM', '模型部署', 'CUDA'],
    experienceYears: 3,
    educationLevel: '硕士'
  },
  {
    id: 'data-engineer',
    name: '数据工程师',
    nameEn: 'Data Engineer',
    category: 'tech',
    requiredSkills: ['SQL', 'Python', 'Spark', 'Hive', 'ETL'],
    preferredSkills: ['Flink', 'Kafka', 'Airflow', '数据仓库'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'mobile-ios',
    name: 'iOS开发工程师',
    nameEn: 'iOS Developer',
    category: 'tech',
    requiredSkills: ['Swift', 'Objective-C', 'iOS SDK', 'Xcode'],
    preferredSkills: ['SwiftUI', 'RxSwift', '性能优化', 'App Store发布'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'mobile-android',
    name: 'Android开发工程师',
    nameEn: 'Android Developer',
    category: 'tech',
    requiredSkills: ['Kotlin', 'Java', 'Android SDK', 'Android Studio'],
    preferredSkills: ['Jetpack Compose', 'Flutter', 'Gradle'],
    experienceYears: 3,
    educationLevel: '本科'
  }
]

// 产品类岗位
const productJobs: JobPreset[] = [
  {
    id: 'product-manager',
    name: '产品经理',
    nameEn: 'Product Manager',
    category: 'product',
    requiredSkills: ['需求分析', 'PRD撰写', '原型设计', '数据分析', '用户调研'],
    preferredSkills: ['SQL', 'Axure', 'Figma', 'Jira'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'product-director',
    name: '产品总监',
    nameEn: 'Product Director',
    category: 'product',
    requiredSkills: ['产品规划', '团队管理', '战略思维', '商业分析', '项目管理'],
    preferredSkills: ['行业洞察', 'OKR', '跨部门协作'],
    experienceYears: 8,
    educationLevel: '本科'
  },
  {
    id: 'product-operation',
    name: '产品运营',
    nameEn: 'Product Operations',
    category: 'product',
    requiredSkills: ['用户运营', '数据分析', '活动策划', '内容运营'],
    preferredSkills: ['增长黑客', 'A/B测试', 'SQL'],
    experienceYears: 2,
    educationLevel: '本科'
  }
]

// 运营类岗位
const operationJobs: JobPreset[] = [
  {
    id: 'operation-manager',
    name: '运营经理',
    nameEn: 'Operations Manager',
    category: 'operation',
    requiredSkills: ['用户运营', '活动策划', '数据分析', '内容策划'],
    preferredSkills: ['社群运营', 'SEO/SEM', '增长黑客'],
    experienceYears: 5,
    educationLevel: '本科'
  },
  {
    id: 'content-operation',
    name: '内容运营',
    nameEn: 'Content Operations',
    category: 'operation',
    requiredSkills: ['内容策划', '文案撰写', '新媒体运营', '数据分析'],
    preferredSkills: ['视频剪辑', 'SEO', '热点追踪'],
    experienceYears: 2,
    educationLevel: '本科'
  },
  {
    id: 'user-operation',
    name: '用户运营',
    nameEn: 'User Operations',
    category: 'operation',
    requiredSkills: ['用户增长', '用户分层', '活动策划', '数据分析'],
    preferredSkills: ['社群运营', 'CRM', '用户画像'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'growth-hacker',
    name: '增长运营',
    nameEn: 'Growth Hacker',
    category: 'operation',
    requiredSkills: ['增长策略', '数据分析', 'A/B测试', '渠道运营'],
    preferredSkills: ['SQL', 'Python', '裂变营销'],
    experienceYears: 3,
    educationLevel: '本科'
  }
]

// 设计类岗位
const designJobs: JobPreset[] = [
  {
    id: 'ui-designer',
    name: 'UI设计师',
    nameEn: 'UI Designer',
    category: 'design',
    requiredSkills: ['Figma', 'Sketch', '视觉设计', '设计规范', 'UI组件'],
    preferredSkills: ['动效设计', '设计系统', 'Principle'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'ux-designer',
    name: 'UX设计师',
    nameEn: 'UX Designer',
    category: 'design',
    requiredSkills: ['用户研究', '交互设计', '原型设计', '可用性测试'],
    preferredSkills: ['数据分析', 'Design Thinking', 'Figma'],
    experienceYears: 3,
    educationLevel: '本科'
  },
  {
    id: 'visual-designer',
    name: '视觉设计师',
    nameEn: 'Visual Designer',
    category: 'design',
    requiredSkills: ['Photoshop', 'Illustrator', '品牌设计', '平面设计'],
    preferredSkills: ['C4D', 'After Effects', '插画'],
    experienceYears: 3,
    educationLevel: '本科'
  }
]

// 销售类岗位
const salesJobs: JobPreset[] = [
  {
    id: 'sales-manager',
    name: '销售经理',
    nameEn: 'Sales Manager',
    category: 'sales',
    requiredSkills: ['销售技巧', '客户管理', '商务谈判', '团队管理'],
    preferredSkills: ['CRM系统', '行业资源', '大客户销售'],
    experienceYears: 5,
    educationLevel: '本科'
  },
  {
    id: 'sales-rep',
    name: '销售代表',
    nameEn: 'Sales Representative',
    category: 'sales',
    requiredSkills: ['销售技巧', '客户开发', '沟通能力', '抗压能力'],
    preferredSkills: ['行业知识', 'CRM'],
    experienceYears: 1,
    educationLevel: '大专'
  },
  {
    id: 'bd-manager',
    name: '商务拓展经理',
    nameEn: 'Business Development Manager',
    category: 'sales',
    requiredSkills: ['商务谈判', '渠道开发', '合作洽谈', '市场分析'],
    preferredSkills: ['行业资源', '战略合作', '商业模式'],
    experienceYears: 5,
    educationLevel: '本科'
  }
]

// 人力资源类岗位
const hrJobs: JobPreset[] = [
  {
    id: 'hr-manager',
    name: '人力资源经理',
    nameEn: 'HR Manager',
    category: 'hr',
    requiredSkills: ['招聘管理', '绩效管理', '员工关系', '薪酬福利'],
    preferredSkills: ['HRBP', '组织发展', '人才盘点'],
    experienceYears: 5,
    educationLevel: '本科'
  },
  {
    id: 'recruiter',
    name: '招聘专员',
    nameEn: 'Recruiter',
    category: 'hr',
    requiredSkills: ['招聘流程', '简历筛选', '面试技巧', '渠道管理'],
    preferredSkills: ['猎头合作', 'ATS系统', '雇主品牌'],
    experienceYears: 2,
    educationLevel: '本科'
  },
  {
    id: 'hrbp',
    name: 'HRBP',
    nameEn: 'HR Business Partner',
    category: 'hr',
    requiredSkills: ['业务理解', '组织诊断', '人才发展', '员工关系'],
    preferredSkills: ['OD', '数据分析', '变革管理'],
    experienceYears: 5,
    educationLevel: '本科'
  }
]

// 财务类岗位
const financeJobs: JobPreset[] = [
  {
    id: 'finance-manager',
    name: '财务经理',
    nameEn: 'Finance Manager',
    category: 'finance',
    requiredSkills: ['财务分析', '预算管理', '成本控制', '财务报表'],
    preferredSkills: ['ERP系统', '税务筹划', '内控管理'],
    experienceYears: 5,
    educationLevel: '本科'
  },
  {
    id: 'accountant',
    name: '会计',
    nameEn: 'Accountant',
    category: 'finance',
    requiredSkills: ['账务处理', '报表编制', '税务申报', '财务软件'],
    preferredSkills: ['成本核算', 'Excel', '审计配合'],
    experienceYears: 2,
    educationLevel: '本科'
  }
]

// 所有岗位分类
export const jobCategories: JobCategory[] = [
  { id: 'tech', labelKey: 'job.categories.tech', presets: techJobs },
  { id: 'product', labelKey: 'job.categories.product', presets: productJobs },
  { id: 'operation', labelKey: 'job.categories.operation', presets: operationJobs },
  { id: 'design', labelKey: 'job.categories.design', presets: designJobs },
  { id: 'sales', labelKey: 'job.categories.sales', presets: salesJobs },
  { id: 'hr', labelKey: 'job.categories.hr', presets: hrJobs },
  { id: 'finance', labelKey: 'job.categories.finance', presets: financeJobs }
]

// 获取所有岗位
export function getAllPresets(): JobPreset[] {
  return jobCategories.flatMap(cat => cat.presets)
}

// 根据ID获取岗位
export function getPresetById(id: string): JobPreset | undefined {
  return getAllPresets().find(p => p.id === id)
}

// 根据分类获取岗位
export function getPresetsByCategory(category: string): JobPreset[] {
  const cat = jobCategories.find(c => c.id === category)
  return cat?.presets || []
}
