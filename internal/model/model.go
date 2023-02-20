package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

type EvenType uint8

type EvenStatus uint8

const (
	Created EvenType = iota
	Updated
	Removed

	Deferred EvenStatus = iota
	Processed
)

type UserEvent struct {
	ID     uint64
	Type   EvenType
	Status EvenStatus
	Entity *User
}
