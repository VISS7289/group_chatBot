package Models

// 搜索信息
type SerchUserInfo struct {
	Username  string `db:"username"`
	UserId    string `db:"user_id"`
	UserIdStr string
	Img       string `db:"imgBase64"`
	Email     string `db:"email"`
	State     int64  `db:"state"`
}

// 查看是否为好友
type SerchIsFriend struct {
	State    int64  `db:"state"`
	Markname string `db:"markName"`
}

// 好友信息
type MyFriend struct {
	Friendid   string `db:"friend_id"`
	FriendName string `db:"username"`
	FriendNick string `db:"markName"`
	FriendImg  string `db:"imgBase64"`
	LastTime   string `db:"last_time"`
}

// 好友请求信息
type MyFriendReq struct {
	Friendid   string `db:"user_id"`
	FriendName string `db:"username"`
	FriendNick string `db:"markName"`
	FriendImg  string `db:"imgBase64"`
	LastTime   string `db:"last_time"`
}
