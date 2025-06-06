CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    hashed_password BYTEA NOT NULL,
    created TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT unique_username UNIQUE (username),
    CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    user1_id UUID NOT NULL,
    user2_id UUID NOT NULL,
    last_message_content TEXT DEFAULT 'No messages yet',
    last_message_at TIMESTAMPTZ DEFAULT NOW(),
    
    CONSTRAINT fk_user1 FOREIGN KEY (user1_id) REFERENCES users(id),
    CONSTRAINT fk_user2 FOREIGN KEY (user2_id) REFERENCES users(id),
    CONSTRAINT unique_pair UNIQUE (user1_id, user2_id)
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id UUID NOT NULL,
    message TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL,
    chat_id INTEGER NOT NULL,

    CONSTRAINT fk_sender_id FOREIGN KEY (sender_id) REFERENCES users(id),
    CONSTRAINT fk_chat_id FOREIGN KEY (chat_id) REFERENCES chats(id)
);