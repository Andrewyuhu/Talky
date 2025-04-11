import { defineStore } from "pinia";
import { type Chat } from "../types/messages";

type ChatState = {
  chats: Chat[];
};

export const useChatStore = defineStore("chats", {
  state: (): ChatState => ({
    chats: [],
  }),
  actions: {
    setChats(chats: Chat[]) {
      this.chats = [...chats];
    },
    updateChatPreview(message: string, chatId: string) {
      const c = this.chats.find((chat) => String(chat.id) === chatId);
      if (!c) {
        return;
      }
      c.last_message_content = message;
    },
  },
});
