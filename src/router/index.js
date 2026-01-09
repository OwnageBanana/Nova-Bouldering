import { createRouter, createWebHistory } from 'vue-router'

// import Home from "@components/Home.vue"
import About from '@components/About.vue'
import Crags from '@components/Crags.vue'
import Areas from '@components/Areas.vue'
import Boulders from '@components/Boulders.vue'
import CragTest from '@components/CragTest.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // { path: '/', component: Home },
    // { path: '/home', component: Home },
    { path: '/about', component: About },
    { path: '/crags', name: 'crags', component: Crags },
    { path: '/crags/:zoneName/:cragName', name: 'areas', component: Areas, props: true },
    { path: '/crags/:propZoneName/:propCragName/:propAreaName', name: 'boulders', component: Boulders, props: true },
    { path: '/cragtest', component: CragTest }
  ],
})

export default router
