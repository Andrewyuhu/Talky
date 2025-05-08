<script lang="ts" setup>
import NavBar from "../components/NavBar.vue";
</script>

<template>
  <div class="flex flex-col h-screen bg-primary-background">
    <NavBar></NavBar>
    <slot></slot>
  </div>
</template>
