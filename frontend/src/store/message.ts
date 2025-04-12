import { defineStore } from "pinia";
import { type Message } from "../types/messages";

export const useMessageStore = defineStore("message", {
  state: () => ({
    messages: new Map<string, Message[]>(),
  }),
  actions: {
    setMessages(messages: Message[], chatId: string) {
      console.log(`Settings message ${messages} with chatId ${chatId}`);
      this.messages.set(chatId, messages);
    },
    addMessage(message: Message, chatId: string) {
      console.log("HELLO");
      if (!this.messages.has(chatId)) {
        this.messages.set(chatId, []);
      }
      this.messages.get(chatId)!.push(message);
    },
  },
});
