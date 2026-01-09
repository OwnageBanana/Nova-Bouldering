<script setup>
import { computed } from 'vue';

const props = defineProps(['zone', 'crag']);
</script>

<template>
    <div class="cragDiv">
        <div @click="$router.push({ name: 'areas', params: { zoneName: zone, cragName: crag } })">
            <h2>{{ crag }}</h2>
        </div>
    </div>
</template>

<style>
    .cragDiv {
        background-color: rgba(128, 128, 128, 0.252);
        border: 1px solid black;
        border-radius: 5px;
        margin: 10px;
        cursor: pointer;
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);


        & div {
            margin: 15px;
        }

        &:hover {
          background-color: var(--complement-light);
        }
    }

</style>
