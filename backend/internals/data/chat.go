package data

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatModel struct {
	db *sql.DB
}

type Chat struct {
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

func sortUUIDsLexo(a, b uuid.UUID) (uuid.UUID, uuid.UUID) {
	if a.String() < b.String() {
		return a, b
	}
	return b, a
}
