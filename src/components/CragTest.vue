<script setup>
import { ref } from 'vue'

import Boulder from '@components/Boulder.vue'
import Problem from '@components/Problem.vue'

// using string matching because I need radio button like flow on selected sections
let sectionSelection = ref('tech')
</script>

<template>
  <div class="layout">
    <!-- <div class="pfp-container">
        <img class="pfp" src="@assets/images/pfp.jpg" height="208px" width="208px">
      </div> -->

    <div class="boulder">
      <Boulder></Boulder>
      <Problem></Problem>
      <Problem></Problem>
      <Problem></Problem>
    </div>

    <div class="boulder">
      <Boulder></Boulder>
      <Problem></Problem>
      <Problem></Problem>
      <Problem></Problem>
    </div>

    <div class="boulder">
      <Boulder></Boulder>
      <Problem></Problem>
      <Problem></Problem>
      <Problem></Problem>
    </div>
  </div>
</template>

<style scoped>
@import url('@assets/modules/sections.module.css');

.boulder {
  width: 80%;
  margin: 25px;
  display: flexbox;
}

.layout {
  display: flex;
  flex-direction: column;
}
.pfp-container {
  height: 13rem;
  width: 13rem;
  margin: 2rem auto 0;
}
@media (min-width: 640px) {
  .pfp-container {
    float: right;
  }
}
.pfp {
  width: 12rem;
  height: 12rem;
  border-radius: 100%;
  border: var(--contrast-bright) solid;
}
html.dark .pfp {
  border: var(--complement) solid;
}

@media (min-width: 640px) {
  .layout {
    /* flex-direction: row-reverse; */
  }
}

.list-container {
  display: flex;
  flex-direction: column;
  margin-bottom: 2rem;
}

@media (min-width: 360px) {
  .list-container {
    min-height: 582px;
  }
}

@media (min-width: 360px) {
  .section {
    max-height: 600px;
    overflow-y: hidden;
  }
  .list-container {
    min-height: 800px;
  }
}
@media (min-width: 460px) {
  .section {
    max-height: 532px;
  }
}
@media (min-width: 560px) {
  .section {
    max-height: 400px;
  }
}
@media (min-width: 640px) {
  .section {
    max-height: 380px;
  }
}
</style>
