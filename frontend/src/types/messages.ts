export type Message = {
  username: string;
  message: string;
  sentAt: string;
};

export type Chat = {
  id: number;
  receiverID: string;
  senderID: string;
  last_message_content: string;
  last_message_at: string;
};
