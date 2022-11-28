package model

type User struct {
	ID               int
	Name             string
	Password         string
	SecretProtection string
}
type Message struct {
	ID             int
	Name           string
	Messgae        string
	CommentderID   int
	CommentderName string
	Comment        string
}
