<script setup>
import { ref } from 'vue'
import { computed } from 'vue';
import { getAttributeImage } from './ProblemAttributes.js';
import GradeColor from '@assets/grade_colors.json';

const {
  grade = '?',
  lineName = 'No Name',
  beta = 'Theres no beta yet!',
  fa = 'None',
  attr1 = 'strong',
  attr2 = 'highball',
  attr3 = 'sharp',
} = defineProps(['grade', 'lineName', 'beta', 'fa', 'attr1', 'attr2', 'attr3'])

const attr1Icon = getAttributeImage(attr1)
const attr2Icon = getAttributeImage(attr2)
const attr3Icon = getAttributeImage(attr3)

const isOpen = ref(false) //for expanding beta & fa

const gradeBackgroundColor = computed(() => {
  const level = GradeColor.difficultyLevels.find(
    l => grade >= l.min && grade <= l.max
  );
  return level ? level.color : '#bdc3c7'; // Fallback color
});

const toggleRows = () => {
  isOpen.value = !isOpen.value
}
</script>

<template>
  <div class="layout">
  <div class="parent" :class="{ 'is-open': isOpen }">
    <div
  id="grade"
  class="lineGrade" :class="{ 'is-open': isOpen }"
  :style="{ backgroundColor: gradeBackgroundColor }"
>V{{ grade }}</div>

    <div class="lineName" :class="{ 'is-open': isOpen }" @click="toggleRows">
      {{ lineName }}
    </div>

    <div class="attr1" :class="{ 'is-open': isOpen }"><img :src="attr1Icon" width="25px" /></div>
    <div class="attr2" :class="{ 'is-open': isOpen }"><img :src="attr2Icon" width="25px" /></div>
    <div class="attr3" :class="{ 'is-open': isOpen }"><img :src="attr3Icon" width="25px" /></div>

    <div class="lineBeta">
      <div style="margin: 15px">{{ beta }}</div>
    </div>

    <div class="lineFA" :class="{ 'is-open': isOpen }">
      <div style="margin: 15px">FA: {{ fa }}</div>
    </div>
  </div>
</div>
</template>

<style>
.parent {
  display: grid;
  grid-template-columns: 40px 1fr 40px 40px 40px;

  /* Use 0fr for the hidden rows */
  grid-template-rows: 40px 0fr 0fr;

  overflow: hidden;

  place-items: stretch;
  grid-column-gap: 0;
  grid-row-gap: 0;

  border: 1px solid black;
  border-radius: 8px 0px 0px 8px;
  margin-bottom: 5px;

  transition: grid-template-rows 0.25s ease;
}

/* expanded state */
.parent.is-open {
  grid-template-rows: 40px 1fr 1fr;
}

/* prevent content bleed */
.parent > div {
  overflow: hidden;
}

.lineGrade {
  /* column 1 */
  grid-area: 1 / 1 / 2 / 2;
  aspect-ratio: 1 / 1;
  display: grid;
  place-items: center;
  border-radius: 8px 0px 0px 8px;

}

.lineGrade.is-open {
    border-radius: 8px 0px 0px 0px;
    border-bottom: 1px solid black;
  }
.lineName {
  /* column 2 (not square) */
  grid-area: 1 / 2 / 2 / 3;
  background-color: rgb(33, 33, 33);
  display: grid;
  place-items: center;
  cursor: pointer;


  &:hover {
    background-color: var(--complement-light);
  }
}

  .lineName.is-open {
    border-bottom: 1px solid black;
  }

/*Attributes, square*/
.attr1 {
  grid-area: 1 / 3 / 2 / 4;
  aspect-ratio: 1 / 1;
  background-color: var(--complement);
  display: grid;
  place-items: center;
}

  .attr1.is-open {
    border-bottom: 1px solid black;
  }

.attr2 {
  grid-area: 1 / 4 / 2 / 5;
  aspect-ratio: 1 / 1;
  background-color: var(--complement);
  display: grid;
  place-items: center;
}

  .attr2.is-open {
    border-bottom: 1px solid black;
  }

.attr3 {
  grid-area: 1 / 5 / 2 / 6;
  aspect-ratio: 1 / 1;
  background-color: var(--complement);
  display: grid;
  place-items: center;
}

  .attr3.is-open {
    border-bottom: 1px solid black;
  }

/* Row-spanning bars */
.lineBeta {
  display: grid;
  min-height: 0;
  grid-area: 2 / 1 / 3 / 6;
  width: 100%;
  background-color: rgb(35, 35, 35);
}

.lineFA {
  grid-area: 3 / 1 / 4 / 6;
  width: 100%;
  min-height: 0;
  display: grid;
  background-color: rgb(35, 35, 35);
  border-radius: 0px 0px 0px 8px;

}

</style>
