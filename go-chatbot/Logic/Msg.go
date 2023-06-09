package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
	"strconv"
)

//此函数用于日期时间比较,通过传入两个时间字符串时间戳
func CompareTime(time1 string, time2 string) (res int, err error) {
	// 将日期时间字符串拆分为数组
	time1Arr := [6]string{time1[0:4], time1[5:7],
		time1[8:10], time1[11:13],
		time1[14:16], time1[17:19]}
	time2Arr := [6]string{time2[0:4], time2[5:7],
		time2[8:10], time2[11:13],
		time2[14:16], time2[17:19]}
	var com1, com2 int
	//逐个元素比较
	for i := 0; i < 6; i++ {
		if com1, err = strconv.Atoi(time1Arr[i]); err != nil {
			return 0, err
		}
		if com2, err = strconv.Atoi(time2Arr[i]); err != nil {
			return 0, err
		}
		if com1 > com2 {
			return 1, nil
		}
		if com1 < com2 {
			return 2, nil
		}
	}
	return 0, nil
}

//此函数用于查询某对用户之间的最新一条消息
func GetNewMsgOne(p *Models.ParmGetNewMsgOne) (Models.MsgRes, error) {
	// 根据发送者和接收者id分别查询最新一条消息
	inf1, err := Mysql.GetNewMsgOne(p.SendId, p.AcceptId)
	if err != nil {
		return inf1, err
	}
	inf2, err := Mysql.GetNewMsgOne(p.AcceptId, p.SendId)
	if err != nil {
		return inf2, err
	}
	// 比较两条消息的时间戳
	res, err := CompareTime(inf1.Time, inf2.Time)
	// inf1 时间较大,返回inf1消息
	if err != nil {
		return inf1, err
	}
	// 否则返回inf2消息
	if res == 1 {
		return inf1, nil
	}
	return inf2, nil
}

//此函数用于获取两用户之间未读消息数量
func GetUnReadMsgNum(p *Models.ParmGetNewMsgOne) (int, error) {
	return Mysql.GetUnKnowMsgNum(p.SendId, p.AcceptId)
}

//此函数用于查询指定用户(p)之前的聊天记录
func GetOldMsgF(p *Models.ParmGetOldChat) ([]Models.MsgOldF, error) {
	info, err := Mysql.GetOldMsgF(p)
	//根据传入的参数查询之前的消息记录
	if err != nil {
		return info, err
	}
	//补充信息
	for i := 0; i < len(info); i++ {
		//根据发件人ID获取对应的头像信息
		info[i].Img, err = Mysql.GetImgById(info[i].SendId)
		if err != nil {
			return info, err
		}
	}
	return info, nil
}
