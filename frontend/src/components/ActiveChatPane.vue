<script lang="ts" setup>
import Message from "./Message.vue";
import useWebSocket from "../utils/useWebSocket";
import { watch, ref, toRef } from "vue";
import { useAuthStore } from "../store/auth";
import { type Message as MessageType } from "../types/messages";

const props = defineProps<{
  chatId: string;
  messages: MessageType[];
  recipientUsername: string;
}>();
const chatId = toRef(props, "chatId");
const messageInput = ref("");
const auth = useAuthStore();

const { isConnected, send } = useWebSocket(
  "ws://localhost:8080/v1/ws/chat/",
  chatId
);

watch(isConnected, () => {
  console.log(`Connection Status : ${isConnected.value}`);
});

function handleSubmit(e: Event) {
  e.preventDefault();
  if (!auth.user || messageInput.value === "") return;
  send(auth.user.id, messageInput.value, new Date().toISOString());
  messageInput.value = "";
}
</script>

<template>
  <div v-if="chatId == ''">No chat open</div>
  <div v-else-if="auth.user" class="flex flex-col flex-1 min-h-0">
    <div class="bg-blue-500 p-4">{{ recipientUsername }}</div>
    <div class="flex flex-col flex-1 p-2 overflow-y-auto gap-2">
      <div class="mb-auto"></div>
      <Message
        v-for="message in messages"
        :message="message"
        :senderId="auth.user.id"
      ></Message>
    </div>
    <div class="p-2">
      <form class="flex gap-2">
        <input
          type="text"
          v-model="messageInput"
          class="flex-1 border border-gray-400 rounded-lg px-4 py-2"
        />
        <button
          class="border-2 px-4 py-2 rounded-md border-black z-10 bg-white"
          @click="handleSubmit"
        >
          SEND
        </button>
      </form>
    </div>
  </div>
</template>
