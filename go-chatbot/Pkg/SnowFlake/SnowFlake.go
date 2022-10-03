package SnowFlake

import (
	"strconv"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init(stateTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", stateTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return nil
}

func GenID() string {
	return strconv.FormatInt(node.Generate().Int64(),10)
}
