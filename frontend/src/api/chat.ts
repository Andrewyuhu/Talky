import axios from "axios";

async function getUserChats() {
  return await axios.get("http://localhost:8080/v1/chat", {
    withCredentials: true,
  });
}

export { getUserChats };
