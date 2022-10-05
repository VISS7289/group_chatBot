package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
)

func SerchUser(p *Models.ParmSerchUsername) ([]Models.SerchUserInfo, error) {
	//usernameNew := "%"
	//for i := 0; i < len(p.Username); i++ {
	//	usernameNew = usernameNew + string(p.Username[i]) + "%"
	//}

	userRes, err := Mysql.SerchByUsername(p.Username)
	for i := 0; i < len(userRes); i++ {
		State, _ := Mysql.SerchIsFriend(p.UserId, userRes[i].UserId)
		userRes[i].State = State.State
	}
	return userRes, err
}

func SerchIsFriend(p *Models.ParmSerchIsFriend) (Models.SerchIsFriend, error) {
	return Mysql.SerchIsFriend(p.UserId, p.FriendId)
}

func DelFriend(p *Models.ParmSerchIsFriend) error {
	return Mysql.DelFriendById(p.UserId, p.FriendId)
}

func RejectFriend(p *Models.ParmSerchIsFriend) error {
	return Mysql.RejectFriendById(p.UserId, p.FriendId)
}

func AcceptFriend(p *Models.ParmSerchIsFriend) error {
	req := Models.ParmFriendRequest{
		UserId:   p.UserId,
		FriendId: p.FriendId,
		Msg:      "",
	}
	return Mysql.ChangeFriend(&req, 0)
}
