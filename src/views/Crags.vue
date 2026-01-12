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

        <div class="boulders">
      <div class="areaName">
        <h2 class="areaTitle">{{ targetZone }}</h2>
        <div class="testLabel">
          <img src="@assets/icons/location-pin.png" style="border-radius: 12px; height: 35px;" />
        </div>
      </div>
          <div class="areaDesc">
          <p>No area description yet!</p>
        </div>
      </div>


    <div v-if="cragsList.length > 0">
      <CragObj v-for="crag in cragsList" :key="crag" :crag="crag" :zone="targetZone"></CragObj>
    </div>

    <!--fallback-->
    <p v-else>No crags found for this zone.</p>
  </div>
</template>

<style scoped>
@import url('@assets/modules/sections.module.css');

.cragDesc {
  border: 1px solid black;
  border-radius: 8px;
  background-color: var(--complement-lighter);
  padding: 12px;
}

.boulders {
  display: flex;
  flex-direction: column;
  align-items: center;

}

.testLabel {
  position: absolute;
  left: 15px;
  font-size: 0.9em;
}

.areaName {
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;      /* Vertical centering */
  justify-content: center;  /* Horizontal centering */
  position: relative;
  background-color: var(--complement-dark);
  border-radius: 8px 8px 0px 0px;
  border: 1px solid black;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
  margin-top: 35px;
}

.areaDesc {
  grid-area: 3 / 1 / 4 / 6;
  width: 100%;
  min-height: 0;
  display: grid;
  background-color: var(--complement-lighter);
  border-radius: 0px 0px 8px 8px;
  padding: 15px;
  margin-bottom: 15px;
  border: 1px solid black;

  /*Test*/
  display: flex;
  align-items: left;
  justify-content: left;
  position: relative;
}

.areaTitle {
  margin: 0;
  font-size: 1.37em;
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
