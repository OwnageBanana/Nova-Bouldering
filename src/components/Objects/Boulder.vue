<script setup>
const { boulderName = 'No Name Yet' } = defineProps(['boulderName'])
</script>

<template>
  <div class="parent">
    <div class="boulderName">{{ boulderName }}</div>
    <div class="boulderPhoto">
      <img src="@assets/images/Boulder_Image.jpg" />
    </div>
  </div>
</template>

<style scoped>
.parent {
  display: flex;
  flex-direction: column;
  width: 100%;
  border: 1px solid black;
}

.boulderName {
  width: 100%;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--compliment-lighter);
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
</style>
