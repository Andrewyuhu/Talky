package data

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	SenderId uuid.UUID `json:"senderId"`
	Message  string    `json:"message"`
	SentAt   time.Time `json:"sentAt"`
	ChatId   int       `json:"chatId"`
}

type MessageModel struct {
	db *sql.DB
}

func NewMessageModel(db *sql.DB) *MessageModel {
	return &MessageModel{
		db: db,
	}
}

func (m *MessageModel) Insert(msg Message) error {

	stmt := `
	INSERT INTO messages (sender_id, message, sent_at, chat_id)
	VALUES ($1, $2, $3, $4)
	`
	_, err := m.db.Exec(stmt, msg.SenderId, msg.Message, msg.SentAt, msg.ChatId)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *MessageModel) Get(chatID int) ([]Message, error) {
	stmt := `
	SELECT sender_id, message, sent_at, chat_id FROM messages
	WHERE chat_id = $1
	`

	rows, err := m.db.Query(stmt, chatID)

	if err != nil {
		return nil, err
	}

	chats := make([]Message, 0)

	for rows.Next() {
		var SenderId uuid.UUID
		var Msg string
		var SentAt time.Time
		var ChatId int

		err := rows.Scan(&SenderId, &Msg, &SentAt, &ChatId)

		if err != nil {
			return nil, err
		}

		chats = append(chats, Message{SenderId, Msg, SentAt, ChatId})

	}

	return chats, nil
}
