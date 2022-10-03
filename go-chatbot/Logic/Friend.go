package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
)

func FriendRequest(p *Models.ParmFriendRequest) error {
	return Mysql.InsertFriend(p, 1)
}

func FriendRequestN(p *Models.ParmFriendRequest) error {
	return Mysql.ChangeFriend(p, 1)
}
