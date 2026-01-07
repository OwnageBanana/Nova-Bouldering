<script setup>
/* temporary code */
// required for accessing router, `this` is not accessible
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const route = useRoute()
const router = useRouter()

defineProps({
  msg: String,
})

const nav = function (path) {
  router.push(path)
}
const themeChange = function () {
  document.documentElement.classList.toggle('dark')
  if (document.documentElement.classList.contains('dark')) {
    localStorage.setItem('theme', 'dark')
  } else {
    localStorage.setItem('theme', 'light')
  }
}
</script>

<template>
  <div class="header">
    <div class="container">
      <div class="heading">
        <h1><a class="invisible" @click="nav('/')">Nova Bouldering</a></h1>
      </div>
      <div class="nav-container">
        <nav>
          <router-link to="/">Home</router-link>
          <router-link to="/about">About</router-link>
          <router-link to="/crags">Crags</router-link>
          <router-link to="/cragtest">Crag Test</router-link>
        </nav>
        <div class="theme-icon">
          <a @click="themeChange">
            <i aria-hidden="true">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 30" class="icon-small sun">
                <path
                  d="M 14.984375 0.98632812 A 1 1 0 0 0 14 2 L 14 5 A 1 1 0 1 0 16 5 L 16 2 A 1 1 0 0 0 14.984375 0.98632812 z M 5.796875 4.7988281 A 1 1 0 0 0 5.1015625 6.515625 L 7.2226562 8.6367188 A 1 1 0 1 0 8.6367188 7.2226562 L 6.515625 5.1015625 A 1 1 0 0 0 5.796875 4.7988281 z M 24.171875 4.7988281 A 1 1 0 0 0 23.484375 5.1015625 L 21.363281 7.2226562 A 1 1 0 1 0 22.777344 8.6367188 L 24.898438 6.515625 A 1 1 0 0 0 24.171875 4.7988281 z M 15 8 A 7 7 0 0 0 8 15 A 7 7 0 0 0 15 22 A 7 7 0 0 0 22 15 A 7 7 0 0 0 15 8 z M 2 14 A 1 1 0 1 0 2 16 L 5 16 A 1 1 0 1 0 5 14 L 2 14 z M 25 14 A 1 1 0 1 0 25 16 L 28 16 A 1 1 0 1 0 28 14 L 25 14 z M 7.9101562 21.060547 A 1 1 0 0 0 7.2226562 21.363281 L 5.1015625 23.484375 A 1 1 0 1 0 6.515625 24.898438 L 8.6367188 22.777344 A 1 1 0 0 0 7.9101562 21.060547 z M 22.060547 21.060547 A 1 1 0 0 0 21.363281 22.777344 L 23.484375 24.898438 A 1 1 0 1 0 24.898438 23.484375 L 22.777344 21.363281 A 1 1 0 0 0 22.060547 21.060547 z M 14.984375 23.986328 A 1 1 0 0 0 14 25 L 14 28 A 1 1 0 1 0 16 28 L 16 25 A 1 1 0 0 0 14.984375 23.986328 z"
                />
              </svg>
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="icon-small moon">
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M17.715 15.15A6.5 6.5 0 0 1 9 6.035C6.106 6.922 4 9.645 4 12.867c0 3.94 3.153 7.136 7.042 7.136 3.101 0 5.734-2.032 6.673-4.853Z"
                  class="fill-sky-400/20"
                ></path>
                <path
                  d="m17.715 15.15.95.316a1 1 0 0 0-1.445-1.185l.495.869ZM9 6.035l.846.534a1 1 0 0 0-1.14-1.49L9 6.035Zm8.221 8.246a5.47 5.47 0 0 1-2.72.718v2a7.47 7.47 0 0 0 3.71-.98l-.99-1.738Zm-2.72.718A5.5 5.5 0 0 1 9 9.5H7a7.5 7.5 0 0 0 7.5 7.5v-2ZM9 9.5c0-1.079.31-2.082.845-2.93L8.153 5.5A7.47 7.47 0 0 0 7 9.5h2Zm-4 3.368C5 10.089 6.815 7.75 9.292 6.99L8.706 5.08C5.397 6.094 3 9.201 3 12.867h2Zm6.042 6.136C7.718 19.003 5 16.268 5 12.867H3c0 4.48 3.588 8.136 8.042 8.136v-2Zm5.725-4.17c-.81 2.433-3.074 4.17-5.725 4.17v2c3.552 0 6.553-2.327 7.622-5.537l-1.897-.632Z"
                  class="fill-sky-500"
                ></path>
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M17 3a1 1 0 0 1 1 1 2 2 0 0 0 2 2 1 1 0 1 1 0 2 2 2 0 0 0-2 2 1 1 0 1 1-2 0 2 2 0 0 0-2-2 1 1 0 1 1 0-2 2 2 0 0 0 2-2 1 1 0 0 1 1-1Z"
                  class="fill-sky-500"
                ></path>
              </svg>
            </i>
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.header {
  position: relative;
  display: flex;
  flex-direction: column;
  padding-bottom: 1rem;
  margin-bottom: 16px;
  /* background-color: var(--main);  */
  border-bottom: 1px solid black;
}
html.dark .header {
  background-color: var(--complement-lighter);
}

@media (min-width: 1024px) {
  .container {
    width: 1024px;
    margin: auto;
  }
}
@media (min-width: 1024px) {
  .container {
    width: 1024px;
    margin: auto;
  }
}

.heading {
  color: var(--main-lighter);
  display: flex;
  flex-wrap: wrap;
  padding-left: 1.5rem;
  margin-right: 1rem;
}
@media (min-width: 640px) {
  .heading {
    padding-left: 2rem;
  }
}

@media (min-width: 1024px) {
  .heading {
    padding-left: 0rem;
  }
}
.heading h1 {
  color: var(--main-darker);
  margin: 0.75rem 1rem 0 0;
}
.heading-subtext {
  /* opacity: 0.75; */
  color: var(--main-darker);
  margin: auto 0 0.25rem 0;
  white-space: nowrap;
}
html.dark .heading h1 {
  color: var(--main-darker);
}
html.dark .heading-subtext {
  color: var(--main-darker);
}
.nav-container {
  display: flex;
}
nav {
  margin-left: 0rem;
  padding: 0 0 0 1rem;
  justify-content: left;
  text-align: left;
  display: flex;
  flex-wrap: wrap;
}

nav a {
  margin: auto 0;
  padding: 0.25rem 1rem;
}

@media (min-width: 640px) {
  nav {
    text-align: center;
  }
  nav a {
    padding: 0 1rem;
  }
}
@media (min-width: 1024px) {
  nav {
    padding-left: 0rem;
  }
}
/* styles to show sun and moon based on themes */
.theme-icon {
  margin: auto 1rem auto 0;
  text-align: right;
  flex-grow: 1;
}

.sun {
  display: none;
}

html.dark .sun {
  display: initial;
}

.moon {
  display: initial;
}

html.dark .moon {
  display: none;
}
i {
  margin: 0;
}
</style>
