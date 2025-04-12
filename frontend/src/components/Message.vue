<script lang="ts" setup>
import { ref } from "vue";
import { type Message } from "../types/messages";
const props = defineProps<{
  message: Message;
  senderId: string;
}>();

const isUser = ref(props.senderId == props.message.senderId);
</script>

<template>
  <div
    :class="[
      'p-3 rounded-lg text-white',
      isUser ? 'self-end bg-blue-600' : 'self-start bg-green-600',
    ]"
  >
    {{ props.message.message }}
  </div>
</template>
