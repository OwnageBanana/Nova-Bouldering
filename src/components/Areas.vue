<script setup>
import { ref } from 'vue'
import Area from '@components/Area.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import rawData from '@assets/master_list.json'
import { processData } from '@modules/DataProcessor.js'

// using string matching because I need radio button like flow on selected sections
let sectionSelection = ref('tech')

const route = useRoute()

//receive zone and crag name from router
const props = defineProps(['zoneName', 'cragName'])

const dataMap = processData(rawData)

const areaList = computed(() => {
  const zone = dataMap[props.zoneName]
  if (zone && zone[props.cragName]) {
    return Object.keys(zone[props.cragName])
  }
  return []
})
</script>

<template>
  <div class="layout">
    <!-- makes breadcrumb.. .yummy...-->
    <h2>
      <router-link :to="{ name: 'crags' }"> {{ zoneName }} </router-link> > {{ cragName }}
    </h2>

    <div v-if="areaList.length > 0">
      <Area
        v-for="area in areaList"
        :key="area"
        :area="area"
        :crag="cragName"
        :zone="zoneName"
      ></Area>
    </div>
    <p v-else>No areas found for this crag.</p>
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
