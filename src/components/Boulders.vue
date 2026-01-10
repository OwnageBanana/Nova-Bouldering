<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import rawData from '@assets/master_list.json'
import Boulder from '@components/Boulder.vue'
import { processData } from '@modules/DataProcessor.js'

// using string matching because I need radio button like flow on selected sections
let sectionSelection = ref('tech')

const route = useRoute()

//receive zone and crag name from router
const props = defineProps(['propZoneName', 'propCragName', 'propAreaName'])

const dataMap = processData(rawData)

const faceList = computed(() => {
  const area = dataMap[props.propZoneName]?.[props.propCragName]?.[props.propAreaName]
  if (!area) return []

  const flattenedFaces = []

  Object.keys(area).forEach((blocName) => {
    Object.keys(area[blocName]).forEach((faceName) => {
      if (faceName !== 'lines') {
        flattenedFaces.push({
          blocName: blocName,
          faceName: faceName,
          id: `${blocName}-${faceName}`,
        })
      }
    })
  })

  return flattenedFaces
})
</script>

<template>
  <div class="layout">
    <!-- makes breadcrumb.. .yummy...-->
    <h2>
      <router-link :to="{ name: 'crags' }"> {{ propZoneName }} </router-link> >
      <router-link
        :to="{ name: 'areas', params: { zoneName: propZoneName, cragName: propCragName } }"
      >
        {{ propCragName }}
      </router-link>
      > {{ propAreaName }}
    </h2>

    <div class="boulders" v-if="faceList.length > 0">
      <Boulder
        v-for="face in faceList"
        :key="bloc"
        :areaName="propAreaName"
        :cragName="propCragName"
        :zoneName="propZoneName"
        :boulderName="face.blocName"
        :face="face.faceName"
      ></Boulder>
    </div>
    <p v-else>No blocs found for this area.</p>
  </div>
</template>

<style scoped>
@import url('@assets/modules/sections.module.css');

.layout {
  display: flex;
  flex-direction: column;
  background-color: color-mix(in srgb, var(--complement-darker), transparent 80%);
  border: 1px solid var(--complement-darkest);
  padding: 15px;
  border-radius: 12px;
  backdrop-filter: blur(8px) brightness(1.4) saturate(120%);
}

.boulders {
  display: flex;
  flex-direction: column;
  align-items: center;
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
