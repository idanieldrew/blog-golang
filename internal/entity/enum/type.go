package enum

type Type uint

const (
	_ Type = iota
	Super
	Admin
	Author
	User
)