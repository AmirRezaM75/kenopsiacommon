package models

type User interface {
	GetId() string
	GetEmail() string
	GetUsername() string
	IsVerified() bool
	GetAvatarId() uint8
}
