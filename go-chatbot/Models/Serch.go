package Models

type SerchUserInfo struct {
	Username  string `db:"username"`
	UserId    string `db:"user_id"`
	UserIdStr string
	Img       string `db:"imgBase64"`
	Email     string `db:"email"`
	State     int64  `db:"state"`
}

type SerchIsFriend struct {
	State    int64  `db:"state"`
	Markname string `db:"markName"`
}

type MyFriend struct {
	Friendid   string `db:"friend_id"`
	FriendName string `db:"username"`
	FriendNick string `db:"markName"`
	FriendImg  string `db:"imgBase64"`
	LastTime   string `db:"last_time"`
}

type MyFriendReq struct {
	Friendid   string `db:"user_id"`
	FriendName string `db:"username"`
	FriendNick string `db:"markName"`
	FriendImg  string `db:"imgBase64"`
	LastTime   string `db:"last_time"`
}