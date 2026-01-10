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
  <div class="parent2" :class="{ 'is-open': isOpen }">
    <div id="grade" class="lineGrade" :class="{ 'is-open': isOpen }" :style="{ backgroundColor: gradeBackgroundColor }">V{{ grade }}</div>

    <div class="lineName" :class="{ 'is-open': isOpen }" @click="toggleRows">
      {{ lineName }}
    </div>

    <div class="attr1" :class="{ 'is-open': isOpen }"><img :src="attr1Icon" width="25px" /></div>
    <div class="attr2" :class="{ 'is-open': isOpen }"><img :src="attr2Icon" width="25px" /></div>
    <div class="attr3" :class="{ 'is-open': isOpen }"><img :src="attr3Icon" width="25px" /></div>

    <div class="lineBeta" :class="{ 'is-open': isOpen }">
      <div style="margin: 15px">{{ beta }}</div>
    </div>

    <div class="lineFA" :class="{ 'is-open': isOpen }">
      <div style="margin: 15px">FA: {{ fa }}</div>
    </div>
  </div>
</div>
</template>

<style>
.parent2 {
  display: grid;
  grid-template-columns: 40px 1fr 40px 40px 40px;

  /* Use 0fr for the hidden rows */
  grid-template-rows: 40px 0fr 0fr;

  overflow: hidden;

  place-items: stretch;
  grid-column-gap: 0;
  grid-row-gap: 0;

  border: 1px solid black;
  border-radius: 8px;
  margin-bottom: 5px;

  transition: grid-template-rows 0.35s ease;
}

/* expanded state */
.parent2.is-open {
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

  border-bottom: 1px solid;
  border-bottom-color: rgba(0, 0, 0, 0);
  transition: border-radius 0.35s ease, border-bottom-color 0.35s ease;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}

.lineGrade.is-open {
    border-radius: 8px 0px 0px 0px;
    border-bottom-color: black;
  }
.lineName {
  /* column 2 (not square) */
  grid-area: 1 / 2 / 2 / 3;
  background-color: var(--complement-light);
  display: grid;
  place-items: center;
  cursor: pointer;
  border-left: 1px solid black;
  border-right: 1px solid black;
  border-bottom: 1px solid;
  border-bottom-color: rgba(0, 0, 0, 0);
  transition: border-bottom-color 0.35s ease;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);

  &:hover {
    background-color: var(--complement-light);
  }
}

  .lineName.is-open {
    border-bottom-color: black;
  }

/*Attributes, square*/
.attr1 {
  grid-area: 1 / 3 / 2 / 4;
  aspect-ratio: 1 / 1;
  background-color: var(--complement-dark);
  display: grid;
  place-items: center;
  border-bottom: 1px solid;
  border-bottom-color: rgba(0, 0, 0, 0);
  transition: border-bottom-color 0.35s ease;
}

  .attr1.is-open {
    border-bottom-color: black;
  }

.attr2 {
  grid-area: 1 / 4 / 2 / 5;
  aspect-ratio: 1 / 1;
  background-color: var(--complement-dark);
  display: grid;
  place-items: center;
  border-bottom: 1px solid;
  border-bottom-color: rgba(0, 0, 0, 0);
  transition: border-bottom-color 0.35s ease;
}

  .attr2.is-open {
    border-bottom-color: black;
  }

.attr3 {
  grid-area: 1 / 5 / 2 / 6;
  aspect-ratio: 1 / 1;
background-color: var(--complement-dark);
  display: grid;
  place-items: center;
  border-bottom: 1px solid;
  border-radius: 0px 8px 8px 0px;
  border-bottom-color: rgba(0, 0, 0, 0);
  transition: border-bottom-color 0.35s, border-radius 0.35s ease;
}

  .attr3.is-open {
    border-bottom: 1px solid;
    border-bottom-color: black;
    border-radius: 0px 8px 0px 0px;
  }

/* Row-spanning bars */
.lineBeta {
  display: grid;
  min-height: 0;
  grid-area: 2 / 1 / 3 / 6;
  width: 100%;
  background-color: var(--complement-lighter);

  border-radius: 8px 0px 0px 0px;
  transition: border-radius 0.35s ease;
}

.lineBeta.is-open {
  border-radius: 0px;
}

.lineFA {
  grid-area: 3 / 1 / 4 / 6;
  width: 100%;
  min-height: 0;
  display: grid;
  background-color: var(--complement-lighter);
  border-radius: 0px 0px 8px 8px;

/*Test*/
  display: flex;
  align-items: left;
  justify-content: left;
  position: relative;
}

</style>
