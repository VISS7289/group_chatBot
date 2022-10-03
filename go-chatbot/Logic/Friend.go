package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
)

func FriendRequest(p *Models.ParmFriendRequest) error {
	return Mysql.InsertFriend(p, 1)
}
