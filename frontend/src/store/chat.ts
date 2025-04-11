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
      this.chats = chats;
    },
    addChat(chat: Chat) {
      this.chats.push(chat);
    },
  },
});
