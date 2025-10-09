<template>
  <div class="dot-grid-container">
    <DotGrid
      :dot-size="5"
      :gap="30"
      base-color="#475569"
      active-color="#5b9fff"
      :proximity="150"
      :speed-trigger="100"
      :shock-radius="250"
      :shock-strength="5"
      :max-speed="5000"
      :resistance="750"
      :return-duration="1.5"
      class-name="custom-dot-grid"  
    />
  </div>
</template>

<script setup lang="ts">
  import DotGrid from "./DotGrid.vue";
</script>

<style scoped>
  .dot-grid-container {
    width: 100%;
    height: 500px;
    position: relative;
    overflow: hidden;
  }
</style>