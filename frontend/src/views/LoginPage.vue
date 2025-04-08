<script setup lang="ts">
import MainLayout from "../layouts/MainLayout.vue";
import LoginForm from "../components/LoginForm.vue";
</script>

<template>
  <MainLayout>
    <div class="flex flex-grow justify-center gap-32 px-16 py-[100px]">
      <LoginForm></LoginForm>
    </div>
  </MainLayout>
</template>
