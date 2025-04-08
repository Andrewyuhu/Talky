<script setup lang="ts">
import { ref } from "vue";

const navItems = ref([
  { label: "Home", route: "/" },
  { label: "Messages", route: "/messages" },
  { label: "Sign Up", route: "/signup" },
  { label: "Login", route: "/login" },
]);
</script>

<template>
  <nav
    class="flex justify-between items-center h-[75px] px-8 border-b border-gray-300"
  >
    <div class="flex items-center gap-4 text-xl">
      <img src="../assets/logo.png" class="aspect-auto w-20" />
      Talky
    </div>
    <ul class="flex gap-6 hover:cursor-pointer">
      <router-link v-for="item in navItems" :to="item.route">
        <div class="px-3 py-1 hover:bg-slate-100 rounded-lg">
          {{ item.label }}
        </div>
      </router-link>
    </ul>
  </nav>
</template>
