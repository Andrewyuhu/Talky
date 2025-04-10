package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatModel struct {
	db *sql.DB
}

type Chat struct {
	Id                   int       `json:"id"`
	SenderId             uuid.UUID `json:"senderID"`
	ReceiverID           uuid.UUID `json:"receiverID"`
	Last_message_content string    `json:"last_message_content"`
	Last_message_at      time.Time `json:"last_message_at"`
}

func NewChatModel(db *sql.DB) *ChatModel {
	return &ChatModel{
		db: db,
	}
}

func (m *ChatModel) InsertNewChat(userID uuid.UUID, recipientId uuid.UUID) error {
	id1, id2 := sortUUIDsLexo(userID, recipientId)

	stmt := `INSERT INTO chats (user1_id, user2_id) VALUES ($1, $2)`

	_, err := m.db.Exec(stmt, id1, id2)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // Unique violation
				return ErrDuplicateChat
			}
		}
		return err
	}

	return nil
}

func (m *ChatModel) GetChats(userID uuid.UUID) ([]Chat, error) {

	stmt := `SELECT * FROM chats WHERE user1_id = $1 OR user2_id = $1`

	rows, err := m.db.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats = make([]Chat, 0)

	for rows.Next() {
		var id int
		var senderId uuid.UUID
		var recieverID uuid.UUID
		var last_message_content string
		var last_message_at time.Time

		err := rows.Scan(&id, &senderId, &recieverID, &last_message_content, &last_message_at)
		if err != nil {
			return nil, err
		}

		chats = append(chats, Chat{
			Id:                   id,
			SenderId:             senderId,
			ReceiverID:           recieverID,
			Last_message_content: last_message_content,
			Last_message_at:      last_message_at,
		})

	}

	return chats, nil
}

func sortUUIDsLexo(a, b uuid.UUID) (uuid.UUID, uuid.UUID) {
	if a.String() < b.String() {
		return a, b
	}
	return b, a
}
