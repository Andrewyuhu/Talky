<script lang="ts" setup>
import MainLayout from "../layouts/MainLayout.vue";
</script>

<template>
  <MainLayout>
    <main class="flex-grow flex justify-center px-4 pb-4">
      <div
        class="relative isolate flex items-center justify-center bg-[url(./assets/hero-image-2.jpeg)] w-full rounded-lg hero-container bg-cover bg-center"
      >
        <div
          class="hero-text text-white text-center z-10 w-2/3 flex flex-col gap-4"
        >
          <h1
            class="font-bold text-xl text-white text-balance md:text-2xl lg:text-3xl"
          >
            Because everyone deserves to be close.
          </h1>
          <p
            class="text-sm text-neutral-200 md:text-base inline-block md:w-[50ch] mx-auto text-balance"
          >
            From family to friends, colleagues to communities, make connections
            feel effortless, no matter the distance.
          </p>
          <router-link to="/signup">
            <div
              class="px-3 py-1 rounded-lg text-primary-text text-sm bg-primary-green border-solid border border-primary-text w-fit mx-auto"
            >
              Sign Up
            </div>
          </router-link>
        </div>
      </div>
    </main>
  </MainLayout>
</template>

<style>
.hero-container::after {
  position: absolute;
  content: "";
  z-index: 1;
  inset: 0;
  background: black;
  opacity: 0.2;
  border-radius: inherit;
}
</style>
