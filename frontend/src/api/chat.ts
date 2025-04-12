import axios from "axios";

async function getUserChats() {
  return await axios.get("http://localhost:8080/v1/chat", {
    withCredentials: true,
  });
}

async function createUserChat(username: string) {
  return await axios.post(
    "http://localhost:8080/v1/chat",
    {
      username: username,
    },
    {
      withCredentials: true,
    }
  );
}

export { getUserChats, createUserChat };
