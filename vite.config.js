import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
const config = defineConfig({
  server: { port: 8080 },
  plugins: [
    vue(),
    // vueDevTools(),
  ],
  resolve: {
    alias: [
      {
        find: '@',
        replacement: fileURLToPath(new URL('./src', import.meta.url)),
      },
      {
        find: '@components',
        replacement: fileURLToPath(new URL('./src/components', import.meta.url)),
      },
      {
        find: '@modules',
        replacement: fileURLToPath(new URL('./src/modules', import.meta.url)),
      },
      {
        find: '@views',
        replacement: fileURLToPath(new URL('./src/views', import.meta.url)),
      },
      {
        find: '@assets',
        replacement: fileURLToPath(new URL('./src/assets', import.meta.url)),
      },
      {
        find: '@layout',
        replacement: fileURLToPath(new URL('./src/components/layout', import.meta.url)),
      },
    ],
  },
})

export default config
