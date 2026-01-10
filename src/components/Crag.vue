<script setup>
import { computed } from 'vue';

const props = defineProps(['zone', 'crag']);
</script>

<template>
    <div class="cragDiv">
        <div class="container1" @click="$router.push({ name: 'areas', params: { zoneName: zone, cragName: crag } })">
          <img src="@assets/images/LOC.jpg" style="border-radius: 12px;" />
            <div class="top-left"><h1>{{ crag }}</h1></div>
        </div>
    </div>
</template>

<style>
    .cragDiv {

        border: 1px solid black;
        margin: 10px;
        cursor: pointer;
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
        border-radius: 12px;

        /* div {
            margin: 15px;
        }*/

        &:hover {
          background-color: var(--complement-light);
        }
    }

.container1 {
  position: relative;
  width: 100%;
  height: 150px;
  overflow: hidden;
}



.container1 img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: 50% 60%;
  mask-image: linear-gradient(to bottom, transparent 0%, black 20%)
}

.top-left {
  position: absolute;
  top: 8px;
  left: 16px;
  z-index: 10;
  color: white;
}

</style>
