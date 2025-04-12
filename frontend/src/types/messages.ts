export type Message = {
  senderId: string;
  message: string;
  sentAt: string;
  chatId: string;
};

export type Chat = {
  id: number;
  receiverID: string;
  senderID: string;
  recieverUserName: string;
  last_message_content: string;
  last_message_at: string;
};
