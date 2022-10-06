package Logic

import (
	"go-chatbot/Dao/Mysql"
	"go-chatbot/Models"
	"strconv"
)

func CompareTime(time1 string, time2 string) (res int, err error) {
	time1Arr := [6]string{time1[0:4], time1[5:7],
		time1[8:10], time1[11:13],
		time1[14:16], time1[17:19]}
	time2Arr := [6]string{time2[0:4], time2[5:7],
		time2[8:10], time2[11:13],
		time2[14:16], time2[17:19]}
	var com1, com2 int
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

func GetNewMsgOne(p *Models.ParmGetNewMsgOne) (Models.MsgRes, error) {
	inf1, err := Mysql.GetNewMsgOne(p.SendId, p.AcceptId)
	if err != nil {
		return inf1, err
	}
	inf2, err := Mysql.GetNewMsgOne(p.AcceptId, p.SendId)
	if err != nil {
		return inf2, err
	}
	res, err := CompareTime(inf1.Time, inf2.Time)
	if err != nil {
		return inf1, err
	}
	if res == 1 {
		return inf1, nil
	}
	return inf2, nil
}

func GetUnReadMsgNum(p *Models.ParmGetNewMsgOne) (int, error) {
	return Mysql.GetUnKnowMsgNum(p.SendId, p.AcceptId)
}

func GetOldMsgF(p *Models.ParmGetOldChat) ([]Models.MsgOldF, error) {
	info, err := Mysql.GetOldMsgF(p)
	if err != nil {
		return info, err
	}
	for i := 0; i < len(info); i++ {
		info[i].Img, err = Mysql.GetImgById(info[i].SendId)
		if err != nil {
			return info, err
		}
	}
	return info, nil
}
