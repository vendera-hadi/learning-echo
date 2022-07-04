package common

type (
	UserRole string
)

const (
	Admin     UserRole = "ADMIN"
	Moderator UserRole = "MODERATOR"
	Writer    UserRole = "WRITER"
)
