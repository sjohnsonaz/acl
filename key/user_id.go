package key

import "uuid"

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func (u UserID) String() string {
	return string(u)
}
