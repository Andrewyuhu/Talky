<script lang="ts" setup>
import { ref, defineEmits } from "vue";
import { createUserChat } from "../api/chat";
const recipientUsernameInput = ref("");

const emit = defineEmits<{
  (e: "chatCreated"): void;
}>();

async function handleCreateChat(e: Event) {
  e.preventDefault();
  if (recipientUsernameInput.value === "") {
    return;
  }
  try {
    await createUserChat(recipientUsernameInput.value);
    emit("chatCreated");
    recipientUsernameInput.value = "";
  } catch (error) {
    console.log(error);
  }
}
</script>

<template>
  <div class="flex flex-col p-4 gap-2 border-r-[1px] border-gray-300 flex-1">
    <form class="flex gap-2">
      <input
        v-model="recipientUsernameInput"
        type="text"
        class="flex-1 border border-gray-400 rounded-lg px-4 py-2"
      />
      <button
        @click="handleCreateChat"
        class="border-2 px-4 py-2 rounded-md border-black bg-white"
      >
        CREATE
      </button>
    </form>
    <slot></slot>
  </div>
</template>
