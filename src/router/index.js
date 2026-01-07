import { createRouter, createWebHistory } from 'vue-router'

// import Home from "@components/Home.vue"
import About from '@components/About.vue'
import Areas from '@components/Areas.vue'
import CragTest from '@components/CragTest.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // { path: '/', component: Home },
    // { path: '/home', component: Home },
    { path: '/about', component: About },
    { path: '/areas', component: Areas },
    { path: '/cragtest', component: CragTest }
  ],
})

export default router
