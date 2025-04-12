package data

import (
	"database/sql"
	"errors"
	"fmt"
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
	RecieverUserName     string    `json:"recieverUserName"`
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

	stmt := `
	SELECT chats.id, 
  user1_id, 
  user2_id, 
  CASE 
    WHEN user1_id = $1 THEN user2.username
    ELSE user1.username
  END as receiver_username,
  last_message_content, 
  last_message_at
	FROM chats
	JOIN users user1 ON user1.id = chats.user1_id
	JOIN users user2 ON user2.id = chats.user2_id
	WHERE user1_id = $1 OR user2_id = $1;
	`

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
		var recieverUserName string
		var last_message_content string
		var last_message_at time.Time

		err := rows.Scan(&id, &senderId, &recieverID, &recieverUserName, &last_message_content, &last_message_at)
		if err != nil {
			return nil, err
		}

		chats = append(chats, Chat{
			Id:                   id,
			SenderId:             senderId,
			ReceiverID:           recieverID,
			RecieverUserName:     recieverUserName,
			Last_message_content: last_message_content,
			Last_message_at:      last_message_at,
		})

	}

	return chats, nil
}

func (m *ChatModel) Exists(chatID int) (bool, error) {
	stmt := `SELECT EXISTS(SELECT 1 FROM chats WHERE id = $1)`
	var exists bool
	err := m.db.QueryRow(stmt, chatID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (m *ChatModel) UpdatePreview(message string, sentAt time.Time, chatID int) error {
	stmt := `
	UPDATE chats
	SET  last_message_content = $1,
    	last_message_at = $2
	WHERE id = $3;
	`
	_, err := m.db.Exec(stmt, message, sentAt, chatID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (m *ChatModel) UserBelongsToChat(chatID int, userID uuid.UUID) (bool, error) {
	stmt := `SELECT EXISTS(SELECT 1 FROM chats WHERE id = $1 AND (user1_id = $2 OR user2_id = $2))`
	var exists bool
	err := m.db.QueryRow(stmt, chatID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (m *ChatModel) ValidateChat(userID uuid.UUID, chatID int) (bool, error) {

	exists, err := m.Exists(chatID)

	if err != nil {
		return false, err
	}

	if !exists {
		return false, nil
	}

	belongs, err := m.UserBelongsToChat(chatID, userID)

	if err != nil {
		return false, err
	}

	if !belongs {
		return false, nil
	}

	return true, nil
}

func sortUUIDsLexo(a, b uuid.UUID) (uuid.UUID, uuid.UUID) {
	if a.String() < b.String() {
		return a, b
	}
	return b, a
}
