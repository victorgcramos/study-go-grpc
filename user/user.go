package user

import (
	"encoding/json"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	//Avatar	 []byte
}

// EncodeUser encodes a user into a JSON byte slice
func EncodeUser(u User) ([]byte, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// DecodeUser decodes a JSON byte slice into a User.
func DecodeUser(payload []byte) (*User, error) {
	var u User

	err := json.Unmarshal(payload, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

type Database interface {
	// New adds a new user
	New(User) error
	// GetById returns a user given some id
	GetById(uuid.UUID) (*User, error)
	// GetAll returns all users
	GetAll() (map[string]User, error)
}
