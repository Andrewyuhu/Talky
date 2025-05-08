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
  isMobile.value = window.innerWidth < 1024;
  fetchChats();
});
</script>

<template>
  <MainLayout>
    <div
      class="flex lg:grid lg:grid-row-[1fr] lg:grid-cols-3 lg:gap-4 flex-grow min-h-0 relative w-full overflow-hidden rounded-lg shadow-md lg:px-4"
    >
      <ChatsPane @chatCreated="fetchChats">
        <template v-if="chatStore.chats && chatStore.chats.length">
          <ChatPreview
            v-for="chat in chatStore.chats"
            :key="chat.id"
            :chat="chat"
            @click="handleChatSwitch(chat.id)"
          />
        </template>
        <div
          v-else
          class="w-full text-zinc-500 flex flex-col items-center text-sm bg-[hsl(35,76%,92.5%)]"
        >
          <img
            class="block h-auto w-1/2 translate-x-[2ch] mt-[200px] mb-2"
            src="../assets/no-messages-2.png"
            alt="A man looking for messages"
          />
          No Messages Yet...
        </div>
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
