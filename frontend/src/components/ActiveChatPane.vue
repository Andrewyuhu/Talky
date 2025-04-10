<script lang="ts" setup>
import Message from "./Message.vue";
import useWebSocket from "../utils/useWebSocket";
import { watch, ref } from "vue";
import { useAuthStore } from "../store/auth";

const messageInput = ref("");
const wsUrl = ref("ws://localhost:8080/v1/ws/chat/1");
const auth = useAuthStore();
const { isConnected, messages, send } = useWebSocket(wsUrl);

watch(isConnected, () => {
  console.log(`Connection Status : ${isConnected.value}`);
});

function handleSubmit(e: Event) {
  e.preventDefault();
  if (!auth.user || messageInput.value === "") return;
  console.log(auth.user);
  console.log(auth.user.username);
  send(auth.user.username, messageInput.value, new Date().toISOString());
  messageInput.value = "";
}
</script>

<template>
  <div v-if="auth.user" class="flex flex-col flex-1 min-h-0">
    <div class="bg-blue-500 p-4">Chat Header</div>
    <div class="flex flex-col flex-1 p-2 overflow-y-auto gap-2">
      <div class="mb-auto"></div>
      <Message
        v-for="message in messages"
        :message="message"
        :username="auth.user.username"
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
