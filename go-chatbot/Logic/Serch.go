package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
)

//此函数用于根据用户名关键词搜索用户信息
func SerchUser(p *Models.ParmSerchUsername) ([]Models.SerchUserInfo, error) {
	//usernameNew := "%"
	//for i := 0; i < len(p.Username); i++ {
	//	usernameNew = usernameNew + string(p.Username[i]) + "%"
	//}

	//根据用户名关键词查询用户信息
	userRes, err := Mysql.SerchByUsername(p.Username)
	//根据当前用户ID和查询到的用户ID,判断是否已经是好友,并将状态 State 填充到返回结果中
	for i := 0; i < len(userRes); i++ {
		State, _ := Mysql.SerchIsFriend(p.UserId, userRes[i].UserId)
		userRes[i].State = State.State
	}
	return userRes, err
}

//此函数用于判断两个用户(通过UserId和FriendId传入)是否已经是好友
func SerchIsFriend(p *Models.ParmSerchIsFriend) (Models.SerchIsFriend, error) {
	return Mysql.SerchIsFriend(p.UserId, p.FriendId)
}

//此函数用于删除两个用户的好友关系
func DelFriend(p *Models.ParmSerchIsFriend) error {
	return Mysql.DelFriendById(p.UserId, p.FriendId)
}

//此函数用于拒绝添加对方为好友
func RejectFriend(p *Models.ParmSerchIsFriend) error {
	return Mysql.RejectFriendById(p.UserId, p.FriendId)
}

//此函数用于接受对方的好友请求,并将二者的关系更新为好友
func AcceptFriend(p *Models.ParmSerchIsFriend) error {
	req := Models.ParmFriendRequest{
		UserId:   p.UserId,
		FriendId: p.FriendId,
		Msg:      "",
	}
	return Mysql.ChangeFriend(&req, 0)
}
