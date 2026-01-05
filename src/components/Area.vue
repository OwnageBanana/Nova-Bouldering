<script setup>
    defineProps(['title', 'desc'])
</script>

<template>
    <div class="areaDiv">
        <div>
            <h2>{{ title }}</h2>
            <p>{{ desc }}</p>
        </div>
    </div>
</template>

<style>
    .areaDiv {
        background-color: rgba(128, 128, 128, 0.252);
        border: 1px solid black;
        border-radius: 5px;
        margin: 10px;

        & div {
            margin: 15px;
        }
    }

</style>
