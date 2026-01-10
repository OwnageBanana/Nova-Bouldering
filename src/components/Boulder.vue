<script setup>
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import Problem from '@components/Problem.vue'
import rawData from '@assets/master_list.json'
import { processData } from '@modules/DataProcessor.js'

const props = defineProps(['zoneName', 'cragName', 'areaName', 'boulderName', 'face'])

const route = useRoute()

const dataMap = processData(rawData)

const lineList = computed(() => {
  const faceData =
    dataMap[props.zoneName]?.[props.cragName]?.[props.areaName]?.[props.boulderName]?.[props.face]

  return faceData?.lines
    ? faceData.lines.map((lineObj) => ({
        name: lineObj.Line,
        grade: lineObj.Grade,
        beta: lineObj.Beta,
      }))
    : []
})
</script>

<template>
  <div class="boulder">
    <div class="parent">
      <div class="boulderName">
        <div>
          <h3>{{ boulderName }} {{ face }}</h3>
        </div>
      </div>
      <div class="drawing-container">
        <div class="boulderPhoto">
          <img src="@assets/images/Boulder_Image.jpg" style="border-radius: 12px" />
        </div>
        <svg class="overlay-svg" viewBox="0 0 1000 1000" preserveAspectRatio="none">
          <!-- Lines and points will be rendered here via Vue -->
          <polyline points="100,100 400,250" stroke="red" fill="none" stroke-width="4" />
        </svg>
      </div>
    </div>
    <div v-if="lineList.length > 0">
      <Problem
        v-for="line in lineList"
        :key="line"
        :grade="line.grade"
        :lineName="line.name"
        :beta="line.beta"
      ></Problem>
    </div>
  </div>
</template>

<style scoped>
.parent {
  display: flex;
  flex-direction: column;
  width: 100%;
  place-items: center;
}

.drawing-container {
  position: relative;
  display: inline-block;
  width: 100%;
  max-width: 800px;
}

.overlay-svg {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: all;
  z-index: 10;
}

.boulder {
  width: 80%;
  margin: 25px;
  display: flexbox;
  background-color: rgb(18, 18, 18);
  border: 1px solid black;
  border-radius: 8px;
  padding: 8px;
}

.boulderName {
  width: 100%;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  background-color: var(--complement-light);
  border-radius: 8px 8px 0px 0px;
  border: 1px solid black;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.boulderPhoto {
  width: 100%;
  flex-grow: 1;
  align-items: center;
  padding: 8px;
}

.boulderPhoto img {
  display: block;
  width: 100%;
  height: auto;
  object-fit: cover;
}

.boulderPhoto:empty {
  display: none;
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
