<script setup lang="ts">
import { ref } from "vue";
import { useAuthStore } from "../store/auth";
import { useRouter } from "vue-router";

const auth = useAuthStore();
const route = useRouter();
const navItems = ref([
  { label: "Sign Up", route: "/signup" },
  { label: "Login", route: "/login" },
]);

async function logoutHandler() {
  await auth.logout();
  if (auth.error) {
    console.log("Logout error occured");
  }
  route.push("/");
}
</script>

<template>
  <nav
    class="flex justify-between items-center h-[8vh] px-4 flex-shrink-0 bg-primary-background"
  >
    <router-link to="/">
      <div
        class="flex items-center text-lg text-secondary-green tracking-tight font-bold md:text-xl"
      >
        <img src="../assets/logo.png" class="aspect-auto w-8 md:w-20" />
        Talky
      </div>
    </router-link>
    <ul class="text-sm">
      <div v-if="!auth.isAuthenticated" class="flex gap-6 hover:cursor-pointer">
        <router-link :to="navItems[0].route">
          <div class="px-3 py-1 rounded-lg text-primary-text">
            {{ navItems[0].label }}
          </div>
        </router-link>
        <router-link :to="navItems[1].route">
          <div
            class="px-3 py-1 rounded-lg text-primary-text bg-primary-green border-solid border border-primary-text"
          >
            {{ navItems[1].label }}
          </div>
        </router-link>
      </div>
      <div v-else class="flex gap-6 hover:cursor-pointer">
        <div
          class="px-3 py-1 hover:bg-slate-100 rounded-lg"
          @click="logoutHandler"
        >
          Logout
        </div>
      </div>
    </ul>
  </nav>
</template>
