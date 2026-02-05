import { createRouter, createWebHistory } from 'vue-router'

// 路由配置
const routes = [
  {
    path: '/',
    name: 'Projects',
    component: () => import('../views/ProjectListView.vue'),
    meta: { title: '招聘项目' }
  },
  {
    path: '/project/:id',
    name: 'ProjectDetail',
    component: () => import('../views/HomeView.vue'),
    meta: { title: '项目详情' }
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import('../views/SettingsView.vue'),
    meta: { title: '设置' }
  }
]

// 创建路由器
const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  document.title = `${to.meta.title || 'TalentLens'} - TalentLens`
  next()
})

export default router
