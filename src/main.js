import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// On page load or when changing themes, best to add inline in `head` to avoid FOUC
if (
  localStorage.getItem('theme') === 'dark' ||
  (localStorage.getItem('theme') === undefined &&
    window.matchMedia('(prefers-color-scheme: dark)').matches)
) {
  document.documentElement.classList.add('dark')
} else {
  document.documentElement.classList.remove('dark')
}

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
