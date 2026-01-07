<script setup>
import { ref } from 'vue'

const { grade = "?", lineName = "No Name", beta = "Theres no beta yet!", fa = "None", attr1, attr2, attr3 } = defineProps([
  'grade', 'lineName', 'beta', 'fa', 'attr1', 'attr2', 'attr3'
]);

const isOpen = ref(false) //for expanding beta & fa

const toggleRows = () => {
  isOpen.value = !isOpen.value
}
</script>

<template>
  <div class="parent" :class="{ 'is-open': isOpen }">
    
    <div class="lineGrade" @click="toggleRows">V{{ grade}}</div>

    <div class="lineName" @click="toggleRows">
      {{lineName}}
    </div>

    <div class="attr1" @click="toggleRows">{{ attr1 }}</div>
    <div class="attr2" @click="toggleRows">{{ attr2 }}</div>
    <div class="attr3" @click="toggleRows">{{attr3}}</div>

    <div class="lineBeta">
      <div style="margin: 15px;">{{ beta }}</div>
    </div>

    <div class="lineFA"><div style="margin: 15px;">FA: {{fa}}</div></div>

  </div>
</template>

<style>
.parent {
  display: grid;

  grid-template-columns:
    40px
    1fr
    40px
    40px
    40px;

  /* default: rows 2 & 3 hidden */
  grid-template-rows:
    40px   /* row 1 */
    0px    /* row 2 */
    0px;   /* row 3 */

  place-items: stretch;
  grid-column-gap: 0;
  grid-row-gap: 0;

  transition: grid-template-rows 0.75s ease;
}

/* expanded state */
.parent.is-open {
  grid-template-rows:
    40px
    auto
    auto;
}

/* prevent content bleed */
.parent > div {
  overflow: hidden;
}

.lineGrade { /* column 1 */
  grid-area: 1 / 1 / 2 / 2; aspect-ratio: 1 / 1; 
  background-color: red;
  display: grid;
  place-items: center;
  cursor: pointer;
  } 
.lineName {  /* column 2 (not square) */
  grid-area: 1 / 2 / 2 / 3; 
  background-color: rgb(34, 34, 34);
  display: grid;
  place-items: center;
  cursor: pointer;
  }

/*Attributes, square*/
.attr1 { 
  grid-area: 1 / 3 / 2 / 4; 
  aspect-ratio: 1 / 1;
  background-color: rgb(34, 34, 34);
  display: grid;
  place-items: center;  
  cursor: pointer;
  } 

.attr2 { 
  grid-area: 1 / 4 / 2 / 5; 
  aspect-ratio: 1 / 1;
  background-color: rgb(34, 34, 34);
  display: grid;
  place-items: center;
  cursor: pointer;
  } 
.attr3 { 
  grid-area: 1 / 5 / 2 / 6; 
  aspect-ratio: 1 / 1; 
  background-color: rgb(34, 34, 34);
  display: grid;
  place-items: center;
  cursor: pointer;
  } /* column 5 */

/* Row-spanning bars */
.lineBeta { 
  grid-area: 2 / 1 / 3 / 6; 
  width: 100%; 
background-color: rgb(20, 20, 20);
  }

.lineFA { 
  grid-area: 3 / 1 / 4 / 6; 
  width: 100%; 
  display: grid;
  background-color: rgb(20, 20, 20);
  }
</style>






