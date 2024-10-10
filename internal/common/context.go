package common

type contextkey string

const (
	ContextUser    contextkey = "User"
	ContextSession contextkey = "Session"
)
