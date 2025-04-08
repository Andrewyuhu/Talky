<script lang="ts" setup>
import NavBar from "../components/NavBar.vue";
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <NavBar></NavBar>
    <slot></slot>
  </div>
</template>
