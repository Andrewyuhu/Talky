<script lang="ts" setup>
import MainLayout from "../layouts/MainLayout.vue";
import ActiveChatPane from "../components/ActiveChatPane.vue";
import ChatsPane from "../components/ChatsPane.vue";
import { getUserChats } from "../api/chat";
import { onMounted, ref } from "vue";
import type { Chat } from "../types/messages";
import { useChatStore } from "../store/chat";
import ChatPreview from "../components/ChatPreview.vue";

const chatStore = useChatStore();
const currentChatId = ref("");

function handleChatSwitch(chatId: number) {
  currentChatId.value = String(chatId);
  console.log("updating chat");
}

onMounted(async () => {
  try {
    const response = await getUserChats();
    chatStore.setChats(response.data.chats as Chat[]);
    console.log(response);
  } catch (error) {
    console.log(error);
  }
});
</script>

<template>
  <MainLayout>
    <div class="grid grid-row-[1fr] grid-cols-3 h-full min-h-0">
      <ChatsPane>
        <ChatPreview
          v-for="chat in chatStore.chats"
          :chat="chat"
          :key="chat.id"
          @click="handleChatSwitch(chat.id)"
        ></ChatPreview>
      </ChatsPane>
      <ActiveChatPane class="col-span-2" :chatId="currentChatId">
      </ActiveChatPane>
    </div>
  </MainLayout>
</template>
