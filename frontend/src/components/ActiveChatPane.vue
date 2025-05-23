<script lang="ts" setup>
import Message from "./Message.vue";
import useWebSocket from "../utils/useWebSocket";
import { watch, ref, toRef, nextTick } from "vue";
import { useAuthStore } from "../store/auth";
import { type Message as MessageType } from "../types/messages";

const props = defineProps<{
  chatId: string;
  messages: MessageType[];
  recipientUsername: string;
  mobilePaneActive: boolean;
  isMobile: boolean;
}>();

const chatId = toRef(props, "chatId");
const messageInput = ref("");
const auth = useAuthStore();
const chatPaneRef = ref<HTMLDivElement | null>(null);

const emit = defineEmits<{
  (e: "closeMobilePane"): void;
}>();

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

watch([chatId, () => props.messages.length], async () => {
  await nextTick();
  const pane = chatPaneRef.value;
  if (pane) {
    pane.scrollTop = pane.scrollHeight;
  }
});

console.log(props.mobilePaneActive);
console.log(props.isMobile);
</script>

<template>
  <div
    v-if="chatId == ''"
    class="flex flex-col items-center justify-center right-[-100vw] z-10 h-full min-h-0 absolute w-screen lg:static bg-[hsl(35,76%,92.5%)] lg:w-auto transition-transform transform rounded-md lg:shadow-md"
    :class="{
      'translate-x-0': isMobile && mobilePaneActive,
      'translate-x-full': isMobile && !mobilePaneActive,
    }"
  >
    No chat open...
  </div>
  <div
    v-else-if="auth.user"
    class="flex flex-col right-[-100vw] z-10 h-full min-h-0 absolute w-screen lg:static bg-[hsl(35,76%,92.5%)] lg:w-auto transition-transform transform rounded-md lg:shadow-lg"
    :class="{
      '-translate-x-full': isMobile && mobilePaneActive,
      'translate-x-0': isMobile && !mobilePaneActive,
    }"
  >
    <div
      class="flex gap-4 p-4 text-lg border-b-[1px] border-gray-300/0.5 bg-[hsl(35,76%,92.5%)] font-semibold shadow-sm"
    >
      <button
        v-if="isMobile && mobilePaneActive"
        @click="emit('closeMobilePane')"
      >
        <svg
          class="w-4 h-4"
          data-slot="icon"
          fill="none"
          stroke-width="2"
          stroke="currentColor"
          viewBox="0 0 24 24"
          xmlns="http://www.w3.org/2000/svg"
          aria-hidden="true"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M10.5 19.5 3 12m0 0 7.5-7.5M3 12h18"
          ></path>
        </svg>
      </button>
      {{ recipientUsername }}
    </div>
    <div
      ref="chatPaneRef"
      class="flex flex-col flex-1 p-4 overflow-y-auto gap-2"
    >
      <div class="mb-auto"></div>
      <Message
        v-for="message in messages"
        :message="message"
        :senderId="auth.user.id"
      ></Message>
    </div>
    <div class="p-4">
      <form class="flex gap-2">
        <input
          type="text"
          v-model="messageInput"
          class="flex-1 border border-gray-400 rounded-lg px-4 py-2"
        />
        <button
          class="px-3 py-1 rounded-lg text-primary-text bg-primary-green border-solid border border-primary-text"
          @click="handleSubmit"
        >
          SEND
        </button>
      </form>
    </div>
  </div>
</template>
