package SnowFlake

/*import (
	"time"

	"github.com/sony/sonyflake"
)

var (
	snowFlake     *sonyflake.Sonyflake
	snowMachineID uint16
)

func GetMachineID() (uint16, error) {
	return snowMachineID, nil
}

func Init2(stateTime string, machineID uint16) (err error) {
	snowMachineID = machineID
	var st time.Time
	st, err = time.Parse("2006-01-02", stateTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: GetMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return nil
}

func GenID2() (id uint16, err error) {
	if sonyFlake == nil {
		err=
	}
}*/
