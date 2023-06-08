package SnowFlake

import (
	"strconv"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// 此函数用于为雪花算法初始化时钟和节点:
func Init(stateTime string, machineID int64) (err error) {
	// 解析状态时间为time.Time
	var st time.Time
	st, err = time.Parse("2006-01-02", stateTime)
	// 解析错误则返回
	if err != nil {
		return err
	}
	// 设置时钟起点
	sf.Epoch = st.UnixNano() / 1000000
	// 初始化节点
	node, err = sf.NewNode(machineID)
	return nil
}

func GenID() string {
	return strconv.FormatInt(node.Generate().Int64(),10)
}
