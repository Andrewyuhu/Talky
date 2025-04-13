<script lang="ts" setup>
import MainLayout from "../layouts/MainLayout.vue";
</script>

<template>
  <MainLayout>
    <div
      class="flex flex-col items-center lg:items-stretch lg:flex-row flex-grow justify-center gap-32 px-16 py-[100px]"
    >
      <div class="flex flex-col gap-10 items-start">
        <h1 class="font-bold text-[32px] mb-4">
          Amazing, never before seen, fantastic messaging service
        </h1>
        <p>
          Send messages to your friend(s) blazingly fast. Best of all, your
          messages are <br />
          <span class="font-bold bg-clip-text text-transparent rainbow-bg">
            NOT ENCRYPTED</span
          >
          either so you never have to worry if your data is private or not!
          <br />
        </p>
        <div class="relative group">
          <button
            class="relative border-2 px-4 py-2 rounded-md border-black z-10 bg-white"
          >
            <router-link to="/signup">SIGN UP NOW</router-link>
          </button>
          <div
            class="absolute top-0 right-0 text-3xl z-0 group-hover:translate-x-10 ease-in-out duration-200"
          >
            👍
          </div>
        </div>
      </div>
      <div>
        <img src="../assets/epic.png" alt="sad" />
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.rainbow-bg {
  background-image: linear-gradient(
    to right,
    red,
    orange,
    yellow,
    green,
    blue,
    indigo,
    violet
  );
}
</style>
