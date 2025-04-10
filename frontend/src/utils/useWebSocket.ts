import { onUnmounted, ref, watch, type Ref } from "vue";
import { type Message } from "../types/messages";

export default function useWebSocket(urlRef: Ref<string>) {
  const socket = ref<null | WebSocket>(null);
  const messages = ref<Message[]>([]);
  const isConnected = ref(false);
  const error = ref<Event | null>(null);
  let currentUrl = "";

  const connect = (url: string) => {
    if (socket.value && currentUrl === url) return;
    const ws = new WebSocket(url);
    socket.value = ws;
    currentUrl = url;
    ws.onopen = () => {
      isConnected.value = true;
    };

    ws.onmessage = (event) => {
      console.log("Raw message:", event.data);
      let msg = JSON.parse(event.data) as Message;
      messages.value.push(msg);
    };

    ws.onerror = (err) => {
      console.log(err);
      error.value = err;
    };

    ws.onclose = () => {
      isConnected.value = false;
      socket.value = null;
    };
  };

  const disconnect = () => {
    socket.value?.close();
  };

  const send = (username: string, content: string, sentAt: string) => {
    if (socket.value && isConnected.value) {
      console.log(`Username : ${username}`);
      let newMessage: Message = {
        username,
        message: content,
        sentAt,
      };

      console.log(JSON.stringify(newMessage));
      socket.value.send(JSON.stringify(newMessage));
    } else {
      console.log("error sending message");
    }
  };

  watch(
    urlRef,
    (newUrl) => {
      connect(newUrl);
    },
    { immediate: true }
  );

  onUnmounted(() => {
    disconnect();
  });

  return { socket, messages, isConnected, error, disconnect, send };
}
