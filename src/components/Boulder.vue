<script setup>
import { ref } from 'vue'
import Problem from './Problem.vue'
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import rawData from '@assets/master_list.json';
import { processData } from './DataProcessor';

const props = defineProps(['zoneName', 'cragName', 'areaName', 'boulderName']);

const route = useRoute();

const dataMap = processData(rawData);

const lineList = computed(() => {
  const zone = dataMap[props.zoneName];

  //this is probably silly
  if (zone && zone[props.cragName] && zone[props.cragName][props.areaName] && zone[props.cragName][props.areaName][props.boulderName]) {
    const boulder = zone[props.cragName][props.areaName][props.boulderName];
    return boulder.lines.map(lineObj => lineObj.Line);
  }

  return [];
});
</script>

<template>
  <div class="boulder">
    <div class="parent">
      <div class="boulderName">{{ boulderName }}</div>
      <div class="boulderPhoto"><img src="@assets/images/Boulder_Image.jpg" /></div>
    </div>
      <div v-if="lineList.length > 0">
        <Problem v-for="line in lineList" :key="area" :areaName="area" :cragName="cragName" :zoneName="zoneName" :boulderName="bloc" :lineName="line"></Problem>
    </div>
  </div>
</template>

<style scoped>
.parent {
  display: flex;
  flex-direction:column;
  width: 100%;
  border: 1px solid black;
}

.boulder {
    width: 80%;
    margin: 25px;
    display: flexbox;
}

.boulderName {
  width: 100%;
  height: 40px;
  display: grid;
  align-items: center;
  justify-content: center;
  background-color: var(--complement-light);
}

.boulderPhoto {
  width: 100%;
  flex-grow: 1;
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
