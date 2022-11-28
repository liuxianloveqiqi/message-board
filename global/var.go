package global

import (
	"database/sql"
	"redrock/web_git/model"
)

var DB *sql.DB

var RegisterID, LoginID, PostID, SenderID, ReceiverID int
var RegisterName, RegisterPassword, RegisterSecretProtection string
var LoginName, LoginPassword, LoginSecretProtection string
var Message, PostUerName string
var CommendWords string
var M model.Message
