package unique

import (
	"fmt"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func GetId() (id uint64, err error) { //返回全局唯一的Id
	if sonyFlake == nil { //没有初始化直接返回
		err = fmt.Errorf("snoy flake not inited")
		return
	}

	id, err = sonyFlake.NextID()
	return
}
