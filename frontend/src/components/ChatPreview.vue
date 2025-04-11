<script lang="ts" setup>
import type { Chat } from "../types/messages";
const props = defineProps<{ chat: Chat }>();
console.log(props.chat);
</script>

<template>
  <div class="flex flex-col bg-fuchsia-300">
    <div>Chat ID {{ chat.recieverUserName }}</div>
    <div>{{ chat.last_message_content }}</div>
  </div>
</template>
