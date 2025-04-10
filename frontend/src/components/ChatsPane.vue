<script lang="ts" setup>
import { getUserChats } from "../api/chat";
import { onMounted } from "vue";
import type { Chat } from "../types/messages";
import ChatPreview from "./ChatPreview.vue";
import { ref } from "vue";

const chats = ref<Chat[]>([]);

onMounted(async () => {
  try {
    const response = await getUserChats();
    chats.value = response.data.chats as Chat[];
  } catch (error) {
    console.log(error);
  }
});
</script>

<template>
  <div class="flex flex-col bg-blue-400 p-4">
    <form class="flex gap-2">
      <input
        type="text"
        class="flex-1 border border-gray-400 rounded-lg px-4 py-2"
      />
      <button class="border-2 px-4 py-2 rounded-md border-black z-10 bg-white">
        CREATE
      </button>
    </form>
    <ChatPreview v-for="chat in chats" :chat="chat"></ChatPreview>
  </div>
</template>
