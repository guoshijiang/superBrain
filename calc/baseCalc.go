package calc

import (
	"superBrain/types"
	"superBrain/timer"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)

//球员个人信息
type PlayerInfo struct {
	FieldID string     //球场ID
	TeamID string      //球队ID
	EquipMentId string //设备ID
	PlayerId string    //球员ID
}

//球员个人信息和速度
type PlayerGameInfo struct {
	PlayerData PlayerInfo                 //球员个人信息
	PersonSpeedCalc types.PersonTimeSpeed //球员的速度信息
}

//球队信息和球队速度
type TeamGameInfo struct {
	TeamID string                 //球队ID
	TeamSpeedCalc types.TeamSpeed //球队的数据信息
}

//速度数据打包结构
type BaseSpeedCalcInfo struct {
	MsgChan string
	PlayGame PlayerGameInfo
	TeamGame TeamGameInfo
}

//预准备阶段是为了查询TeamName
func (this* BaseSpeedCalcInfo)PrePrePare() {
	fmt.Println("数据库前期处理")
}

//准备阶段是做计算
func (this* BaseSpeedCalcInfo)PrePare() error {
	var err error
	var (
		Speed SpeedCalcData
	)

	//做计算
	err = Speed.PesonTimeSpeed(this.MsgChan)
	if err != nil {
		log.Error("calc person time speed err")
		return err
	}

	go func() {
		this.Commit(this.MsgChan)
	}()

	return nil
}

//提交阶段为入库
func (this* BaseSpeedCalcInfo) Commit(buf string) {
	var (
		Timer timer.TimerDuty
	)

	go func() {
		Timer.DutyTimerSpeed(buf)
	}()

	go func() {
		//Timer.DutyElectric(buf)
	}()
}



