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

func GetMyFriend(p *Models.ParmGetFriend) ([]Models.MyFriend, error) {

	FriendRes, err := Mysql.GetMyFriend(p.UserId, p.State)
	if err != nil {
		return FriendRes, err
	}
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

func GetMyFriendReq(p *Models.ParmGetFriend) ([]Models.MyFriendReq, error) {

	FriendRes, err := Mysql.GetMyFriendReq(p.UserId, p.State)
	if err != nil {
		return FriendRes, err
	}
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
