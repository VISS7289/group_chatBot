package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
)

//此函数用于发送好友请求
func FriendRequest(p *Models.ParmFriendRequest) error {
	return Mysql.InsertFriend(p, 1)
}

//此函数用于发送新的好友请求
func FriendRequestN(p *Models.ParmFriendRequest) error {
	return Mysql.ChangeFriend(p, 1)
}


//此函数用于获取我的好友列表,通过p传入参数后从数据库读取好友资料。
//对读取到的好友列表进行补全用户名和头像信息。
func GetMyFriend(p *Models.ParmGetFriend) ([]Models.MyFriend, error) {
	// 获取好友列表
	FriendRes, err := Mysql.GetMyFriend(p.UserId, p.State)
	if err != nil {
		return FriendRes, err
	}
	//补全好友信息
	for i := 0; i < len(FriendRes); i++ {
		//根据好友id获取其用户名和头像
		info, err := Mysql.GetNameAndImgById(FriendRes[i].Friendid)
		if err != nil {
			return FriendRes, err
		}
		FriendRes[i].FriendName = info.Username
		FriendRes[i].FriendImg = info.Img
	}
	return FriendRes, nil

}

//此函数用于获取我的好友请求列表
func GetMyFriendReq(p *Models.ParmGetFriend) ([]Models.MyFriendReq, error) {
	//根据条件查询好友请求列表
	FriendRes, err := Mysql.GetMyFriendReq(p.UserId, p.State)
	if err != nil {
		return FriendRes, err
	}
	// 补充好友请求用户信息
	for i := 0; i < len(FriendRes); i++ {
		info, err := Mysql.GetNameAndImgById(FriendRes[i].Friendid)
		if err != nil {
			return FriendRes, err
		}
		FriendRes[i].FriendName = info.Username
		FriendRes[i].FriendImg = info.Img
	}

	return FriendRes, nil

}
