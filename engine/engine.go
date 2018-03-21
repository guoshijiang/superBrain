package engine

import (
	"fmt"
	"superBrain/calc"
	"superBrain/model"
	"superBrain/types"
	"strings"
	"github.com/ethereum/go-ethereum/log"
	//"strconv"
)

var (
	EM_Id []string
	Data = make(map[string] *types.TeamPlayerIdOrName)
)

type DealEngine interface {
	Deal(input *Input)
}

type SpeedDealEngine struct {
	CalcSpeed calc.BaseSpeedCalcInfo
}

func (this *SpeedDealEngine) Deal(input *Input) {
	var (
		is_have int
		PTInfo model.TeamPlayerInfo
	)
	hdtasStr := fmt.Sprintf("%s", input.content)
	htbdata := strings.Split(hdtasStr, ",")
	is_have = 0
	if len(EM_Id) == 0 {
		EM_Id = append(EM_Id, htbdata[0])
		ret := PTInfo.GetTeamPlayerInfo(htbdata[0])
		if ret < 0 {
			log.Error("Call Function GetTeamPlayerInfo Error")
			return
		}
		Data[htbdata[0]] = &PTInfo.TPInfo
	} else {
		for i := 0; i < len(EM_Id); i++ {
			if htbdata[0] == EM_Id[i] {
				is_have = 1
				break
			}
		}
		if is_have != 1 {
			EM_Id = append(EM_Id, htbdata[0])
			ret := PTInfo.GetTeamPlayerInfo(htbdata[0])
			if ret < 0 {
				log.Error("Call Function GetTeamPlayerInfo Error")
				return
			}
			Data[htbdata[0]] = &PTInfo.TPInfo
		}
	}

	for key, value := range Data {
		fmt.Println("Key:", key, "Value:", value)
	}

	//this.CalcSpeed.MsgChan = hdtasStr + strconv.Itoa(PTInfo.TPInfo.PlayerId) + PTInfo.TPInfo.PlayerName + strconv.Itoa(PTInfo.TPInfo.TeamID) + PTInfo.TPInfo.TeamName
	//fmt.Println(this.CalcSpeed.MsgChan)
	return
}

func Run(){
	CreateInput()
	select {}
}


