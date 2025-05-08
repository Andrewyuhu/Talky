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
      'py-3 px-5 rounded-full  text-white max-w-full break-all relative',
      isUser
        ? 'self-end bg-blue-600 outgoing-message '
        : 'self-start bg-green-600 incoming-message ',
    ]"
  >
    {{ props.message.message }}
  </div>
</template>

<style></style>
