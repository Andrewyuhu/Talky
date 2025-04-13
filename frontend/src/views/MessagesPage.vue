<script lang="ts" setup>
import MainLayout from "../layouts/MainLayout.vue";
import ActiveChatPane from "../components/ActiveChatPane.vue";
import ChatsPane from "../components/ChatsPane.vue";
import { getUserChats } from "../api/chat";
import { onMounted, ref, computed } from "vue";
import type { Chat } from "../types/messages";
import { useChatStore } from "../store/chat";
import ChatPreview from "../components/ChatPreview.vue";
import { useMessageStore } from "../store/message";
import { type Message } from "../types/messages";
import { getChatMessages } from "../api/messages";

const chatStore = useChatStore();
const messageStore = useMessageStore();
const currentChatId = ref("");
const isMobile = ref(false);
const mobilePaneActive = ref(false);

window.addEventListener("resize", () => {
  console.log(window.innerWidth);
  isMobile.value = window.innerWidth < 1024;
});

async function handleChatSwitch(chatId: number) {
  try {
    const response = await getChatMessages(chatId);
    const messages: Message[] = response.data.messages;
    messageStore.setMessages(messages, String(chatId));
  } catch (err) {
    console.log(err);
    return;
  }
  currentChatId.value = String(chatId);
  mobilePaneActive.value = !mobilePaneActive.value;
  console.log(mobilePaneActive.value);
  console.log(isMobile.value);
}

const currentChatMessages = computed(() => {
  return messageStore.messages.get(currentChatId.value) || [];
});

const currentRecipient = computed(() => {
  return chatStore.getRecipientName(currentChatId.value);
});

async function fetchChats() {
  try {
    const response = await getUserChats();
    chatStore.setChats(response.data.chats as Chat[]);
    console.log(response);
  } catch (error) {
    console.log(error);
  }
}

onMounted(async () => {
  fetchChats();
});
</script>

<template>
  <MainLayout>
    <div
      class="flex lg:grid lg:grid-row-[1fr] lg:grid-cols-3 h-full min-h-0 relative w-screen overflow-hidden"
    >
      <ChatsPane @chatCreated="fetchChats">
        <ChatPreview
          v-for="chat in chatStore.chats"
          :chat="chat"
          :key="chat.id"
          @click="handleChatSwitch(chat.id)"
        ></ChatPreview>
      </ChatsPane>
      <ActiveChatPane
        class="col-span-2"
        :chatId="currentChatId"
        :recipientUsername="currentRecipient"
        :messages="currentChatMessages"
        :mobilePaneActive="mobilePaneActive"
        :isMobile="isMobile"
        @closeMobilePane="mobilePaneActive = false"
      >
      </ActiveChatPane>
    </div>
  </MainLayout>
</template>
