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
  <div
    class="flex flex-col p-4 gap-2 flex-1 bg-[hsl(35,76%,92.5%)] rounded-md lg:shadow-lg"
  >
    <form class="flex gap-2 h-[2rem]">
      <input
        v-model="recipientUsernameInput"
        type="text"
        class="flex-1 border border-gray-400 rounded-lg px-4 py-2 text-sm"
      />
      <button
        @click="handleCreateChat"
        class="px-3 py-1 rounded-lg text-primary-text bg-primary-green border-solid border border-primary-text text-sm"
      >
        Create
      </button>
    </form>
    <slot></slot>
  </div>
</template>
