<script setup>
import { computed } from 'vue';

const props = defineProps(['title', 'origin']);
</script>

<template>
    <div class="cragDiv">
        <div @click="$router.push({ name: 'areas', params: { zoneName: origin, cragName: title } })">
            <h2>{{ title }}</h2>
        </div>
    </div>
</template>

<style>
    .cragDiv {
        background-color: rgba(128, 128, 128, 0.252);
        border: 1px solid black;
        border-radius: 5px;
        margin: 10px;

        & div {
            margin: 15px;
        }
    }

</style>
