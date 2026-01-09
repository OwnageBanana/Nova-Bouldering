<script setup>
import { computed } from 'vue';

const props = defineProps(['area', 'crag', 'zone']);

const areaURL = computed(() => {
    return encodeURIComponent(props.title);
});

</script>

<template>
    <div class="areaDiv">

        <div @click="$router.push({ name: 'boulders', params: { propZoneName: zone, propCragName: crag, propAreaName: area } })">
            <h2>{{ area }}</h2>
        </div>

    </div>
</template>

<style>
    .areaDiv {
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
