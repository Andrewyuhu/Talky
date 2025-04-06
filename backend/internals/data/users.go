package data

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	db *sql.DB
}

type PublicUser struct {
	Id       uuid.UUID
	Username string
	Email    string
}

func New(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) Insert(username, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	stmt := `
	INSERT INTO users (username, email, hashed_password)
	VALUES ($1, $2, $3);
	`

	_, err = m.db.Exec(stmt, username, email, hashedPassword)

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // Unique violation
				switch pgErr.Constraint {
				case "unique_username":
					return ErrDuplicateUsername
				case "unique_email":
					return ErrDuplicateEmail
				}
			}
			return err
		}
	}

	return nil
}

func (m *UserModel) Authenticate(username, password string) (uuid.UUID, error) {

	var id uuid.UUID
	var hashedPassword string

	stmt := `
	SELECT id, hashed_password FROM users
	WHERE username = $1;
	`

	err := m.db.QueryRow(stmt, username).Scan(&id, &hashedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, ErrInvalidCredentials
		}
		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return uuid.Nil, ErrInvalidCredentials
	}

	return id, nil
}

func (m *UserModel) Get(id uuid.UUID) (*PublicUser, error) {
	var user PublicUser

	stmt := `
	SELECT id, username, email FROM users
	WHERE id = $1;
	`

	err := m.db.QueryRow(stmt, id).Scan(&user.Id, &user.Username, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil

}

func (m *UserModel) Exists(id uuid.UUID) (bool, error) {
	stmt := `SELECT 1 FROM users WHERE id = $1;`

	var exists int
	err := m.db.QueryRow(stmt, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ErrRecordNotFound
		}
		return false, err
	}
	return true, nil
}
