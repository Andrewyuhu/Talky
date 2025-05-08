<script lang="ts" setup>
import type { Chat } from "../types/messages";
import randomColor from "../utils/randomColor";
import { ref, computed } from "vue";

const props = defineProps<{ chat: Chat }>();
const color = ref(randomColor());
const bgColorClass = computed(() => `bg-${color.value}-200`);
</script>

<template>
  <div class="flex p-4 rounded-lg cursor-pointer items-center gap-4">
    <div
      class="w-[2.75rem] aspect-square grid place-content-center rounded-full"
      :class="bgColorClass"
    >
      {{ chat.recieverUserName[0] }}
    </div>
    <div>
      <div class="font-semibold">{{ chat.recieverUserName }}</div>
      <div class="text-sm text-zinc-500 truncate w-[30ch]">
        {{ chat.last_message_content }}
      </div>
    </div>
  </div>
</template>
