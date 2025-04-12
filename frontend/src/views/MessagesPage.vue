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
const currentChatId = ref("0");

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
}

const currentChatMessages = computed(() => {
  return messageStore.messages.get(currentChatId.value) || [];
});

const currentRecipient = computed(() => {
  return chatStore.getRecipientName(currentChatId.value);
});

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
      <ActiveChatPane
        class="col-span-2"
        :chatId="currentChatId"
        :recipientUsername="currentRecipient"
        :messages="currentChatMessages"
      >
      </ActiveChatPane>
    </div>
  </MainLayout>
</template>
