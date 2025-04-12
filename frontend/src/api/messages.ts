import axios from "axios";

export async function getChatMessages(chatID: number) {
  return axios.get(`http://localhost:8080/v1/message/${chatID}`, {
    withCredentials: true,
  });
}
