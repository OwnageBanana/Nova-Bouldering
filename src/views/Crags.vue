<script setup>
import { ref, computed } from 'vue'
import CragObj from '@components/Crag.vue'
import rawData from '@assets/master_list.json'
import { processData } from '@modules/DataProcessor.js'

// using string matching because I need radio button like flow on selected sections
let sectionSelection = ref('tech')

const targetZone = "Peggy's Cove"

const dataMap = processData(rawData)

const cragsList = computed(() => {
  if (dataMap[targetZone]) {
    return Object.keys(dataMap[targetZone])
  }
  return []
})
</script>

<template>
  <div class="layout">
    <h2>{{ targetZone }}</h2>

    <div v-if="cragsList.length > 0">
      <CragObj v-for="crag in cragsList" :key="crag" :crag="crag" :zone="targetZone"></CragObj>
    </div>

    <!--fallback-->
    <p v-else>No crags found for this zone.</p>
  </div>
</template>

<style scoped>
@import url('@assets/modules/sections.module.css');

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
