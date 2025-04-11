<script lang="ts" setup></script>

<template>
  <div class="flex flex-col bg-blue-400 p-4 gap-2">
    <form class="flex gap-2">
      <input
        type="text"
        class="flex-1 border border-gray-400 rounded-lg px-4 py-2"
      />
      <button class="border-2 px-4 py-2 rounded-md border-black z-10 bg-white">
        CREATE
      </button>
    </form>
    <slot></slot>
  </div>
</template>
