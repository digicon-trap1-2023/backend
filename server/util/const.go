package util

import "github.com/google/uuid"

type Key int

const (
	UserKey Key = iota
	RoleKey
)

func SampleUserID() uuid.UUID {
	return uuid.MustParse("00000000-0000-0000-0000-000000000001")
}

func NewID() uuid.UUID {
	return uuid.New()
}

func SampleRole() string {
	return "writer"
}