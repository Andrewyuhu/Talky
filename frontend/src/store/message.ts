import { defineStore } from "pinia";
import { type Message } from "../types/messages";

export const useMessageStore = defineStore("message", {
  state: () => ({
    messages: new Map<string, Message[]>(),
  }),
  actions: {
    setMessages(messages: Message[], chatId: string) {
      this.messages.set(chatId, messages);
    },
    addMessage(message: Message, chatId: string) {
      if (!this.messages.has(chatId)) {
        this.messages.set(chatId, []);
      }
      this.messages.get(chatId)!.push(message);
    },
  },
});
