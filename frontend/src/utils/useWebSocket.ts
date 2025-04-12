import { onUnmounted, ref, watch, type Ref } from "vue";
import { type Message } from "../types/messages";
import { useChatStore } from "../store/chat";
import { useMessageStore } from "../store/message";

export default function useWebSocket(url: string, chatIdRef: Ref<string>) {
  const socket = ref<null | WebSocket>(null);

  const isConnected = ref(false);
  const error = ref<Event | null>(null);

  const chatStore = useChatStore();
  const messageStore = useMessageStore();

  let currentUrl = "";

  const connect = (url: string) => {
    if (socket.value && currentUrl === url) return;
    socket.value?.close();
    isConnected.value = false;
    const ws = new WebSocket(url);
    socket.value = ws;
    currentUrl = url;

    ws.onopen = () => {
      isConnected.value = true;
      console.log(socket.value);
      console.log("Connected to new websocket");
    };

    ws.onmessage = (event) => {
      let msg = JSON.parse(event.data) as Message;
      let chatId = msg.chatId;
      chatStore.updateChatPreview(msg.message, chatId);
      messageStore.addMessage(msg, chatId);
    };

    ws.onerror = (err) => {
      console.log(err);
      error.value = err;
    };

    ws.onclose = () => {
      isConnected.value = false;
      if (socket.value === ws) {
        socket.value = null;
        currentUrl = "";
      }
    };
  };

  const disconnect = () => {
    socket.value?.close();
  };

  const send = (senderId: string, content: string, sentAt: string) => {
    if (socket.value && isConnected.value) {
      let newMessage: Message = {
        senderId,
        message: content,
        sentAt,
        chatId: chatIdRef.value,
      };
      console.log(JSON.stringify(newMessage));
      socket.value.send(JSON.stringify(newMessage));
    } else {
      console.log("error sending message");
    }
  };

  watch(
    chatIdRef,
    (newChatId, oldChatId) => {
      if (newChatId != oldChatId && newChatId != "") {
        console.log(`${url}${newChatId}`);
        connect(`${url}${newChatId}`);
      }
    },
    { immediate: true }
  );

  onUnmounted(() => {
    disconnect();
  });

  return { socket, isConnected, error, disconnect, send };
}
