package unique

import "github.com/sony/sonyflake"

type Config struct {
	MachineId uint16 `json:'machineId'`
}

func Init(config Config) (err error) {
	//id生成器初始化
	sonyMachineID = config.MachineId //保存用户传来的机器id
	settings := sonyflake.Settings{}
	settings.MachineID = getMachineID //通过回调函数的方式获取机器id
	sonyFlake = sonyflake.NewSonyflake(settings)

	return
}
