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

func (m *MessageModel) Get(chatId int) ([]Message, error) {
	return nil, nil
}
