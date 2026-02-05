import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import 'element-plus/dist/index.css'

import App from './App.vue'
import router from './router'
import i18n from './locales'
import './style.css'

// 创建应用
const app = createApp(App)

// 注册Pinia状态管理
app.use(createPinia())

// 注册Vue Router
app.use(router)

// 注册 i18n 国际化
app.use(i18n)

// 注册Element Plus
app.use(ElementPlus)

// 注册所有Element Plus图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 挂载应用
app.mount('#app')
