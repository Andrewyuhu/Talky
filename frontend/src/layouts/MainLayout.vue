<script lang="ts" setup>
import NavBar from "../components/NavBar.vue";
</script>

<template>
  <div class="flex flex-col h-screen overflow-y-auto">
    <NavBar></NavBar>
    <slot></slot>
  </div>
</template>
