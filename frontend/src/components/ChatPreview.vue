<script lang="ts" setup>
import type { Chat } from "../types/messages";
const props = defineProps<{ chat: Chat }>();
console.log(props.chat);
</script>

<template>
  <div class="flex flex-col p-4 border rounded-lg cursor-pointer">
    <div>Chat ID {{ chat.id }} - {{ chat.recieverUserName }}</div>
    <div class="text-gray-600">{{ chat.last_message_content }}</div>
  </div>
</template>
