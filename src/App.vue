<!-- <script lang="ts">
import Heading from '@layout/Heading.vue'
import Footer from '@layout/Footer.vue'
export default {
  name: 'Nova-Bouldering',
  components: [Heading, Footer],
}
</script> -->
<script setup>
import Heading from '@layout/Heading.vue'
import Footer from '@layout/Footer.vue'
</script>

<template>
  <Heading />
  <div id="app" class="site-background">

  <div class="main-view">
    <router-view></router-view>
  </div>
  </div>
  <Footer />
</template>

<style scoped>
.main-view {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  margin: 0 0.5rem;
  /* text-align: justify; */
}


.site-background {
  position: relative;
  min-height: 100vh;
  width: 100%;
  overflow-x: hidden;
}

.site-background::before {
  content: "";
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;

  background-image: url('@assets/images/boulders_BG.jpg');
  background-size: cover;
  background-repeat: no-repeat;

  background-position: 100% 10%;

  mask-image: linear-gradient(to top, black 20%, transparent 80%);
}

/* 4. Keep your content layer clean */
.main-view {
  position: relative;
  z-index: 1; /* Ensures content stays on top */
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  margin: 0 0.5rem;
}



@media (min-width: 640px) {
  .main-view {
    margin: 0 2rem;
  }
}

@media (min-width: 1024px) {
  .main-view {
    margin: 0 auto;
    width: 1024px;
    /* max-width: 75%; */
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.scroll-top-button {
  display: block;
  cursor: pointer;
  position: fixed;
  right: 2rem;
  bottom: 2.5rem;
  height: 3rem;
  width: 3rem;
  border-radius: 9999px;
  background-color: var(--main-lighter);
  border: 2px solid var(--main-light);
  /* box-shadow: ; */
}
.icon-container {
  margin: auto;
}
.icon-container svg {
  width: 1.5rem;
}
@media (min-width: 640px) {
  .scroll-top-button {
    width: 2.5rem;
    height: 2.5rem;
    bottom: 3rem;
    right: 5%;
  }
  .icon-container svg {
    width: 1.5rem;
  }
}
@media (min-width: 768px) {
  .scroll-top-button {
    width: 3rem;
    height: 3rem;
    bottom: 4rem;
    right: 5%;
  }
  .icon-container svg {
    width: 2rem;
  }
}
@media (min-width: 1024px) {
  .scroll-top-button {
    right: 5%;
  }
}
@media (min-width: 1440px) {
  .scroll-top-button {
    width: 3.5rem;
    height: 3.5rem;
    bottom: 4rem;
    right: 10%;
  }
  .icon-container svg {
    width: 2rem;
  }
}
@media (min-width: 2000px) {
  .scroll-top-button {
    bottom: 4rem;
    right: 20%;
  }
}
</style>
