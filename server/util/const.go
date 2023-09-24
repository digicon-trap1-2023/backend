package util

import "github.com/google/uuid"

type Key int

const (
	UserKey Key = iota
)

func SampleUserID() uuid.UUID {
	return uuid.MustParse("00000000-0000-0000-0000-000000000001")
}
