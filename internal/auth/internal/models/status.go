package models

type ErrorCode uint8

const (
	NOT_FOUND ErrorCode = iota
	INTERNAL
	EXISTS
	CREATED
	OK
)
