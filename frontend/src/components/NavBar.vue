<script setup lang="ts">
import { ref } from "vue";
import { useAuthStore } from "../store/auth";
import { useRouter } from "vue-router";

const auth = useAuthStore();
const route = useRouter();
const navItems = ref([
  { label: "Home", route: "/" },
  { label: "Sign Up", route: "/signup" },
  { label: "Login", route: "/login" },
]);

const authNavItems = ref([{ label: "Messages", route: "/messages" }]);

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
    class="flex justify-between items-center h-[75px] px-8 border-b border-gray-300"
  >
    <div class="flex items-center gap-4 text-xl">
      <img src="../assets/logo.png" class="aspect-auto w-20" />
      Talky
    </div>
    <ul>
      <div v-if="!auth.isAuthenticated" class="flex gap-6 hover:cursor-pointer">
        <router-link v-for="item in navItems" :to="item.route">
          <div class="px-3 py-1 hover:bg-slate-100 rounded-lg">
            {{ item.label }}
          </div>
        </router-link>
      </div>
      <div v-else class="flex gap-6 hover:cursor-pointer">
        <router-link v-for="item in authNavItems" :to="item.route">
          <div class="px-3 py-1 hover:bg-slate-100 rounded-lg">
            {{ item.label }}
          </div>
        </router-link>
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
