package model

import (
	"superBrain/types"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)

type PlayerTeamData struct {
	PlayTeam types.PlayerTeamInfo `json:"play_team"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func GetBegainPlay(playId string) (*PlayerTeamData, error) {
	var playTeamInfo []types.PlayerTeamInfo
	hrows, herr := DB.Query(fmt.Sprintf(`SELECT TEAM_NAME, IS_DEL FROM RD_TEAM_INFO WHERE ID ='%s'`, playId))
	if herr != nil {
		log.Error("Get Begin Play Info Sql Return Err %v\n", herr)
		return nil, herr
	}
	for hrows.Next() {
		var hrow = new(types.PlayerTeamInfo)
		if herr = hrows.Scan(&hrow.TeamName, &hrow.IsDel); herr != nil {
			return nil, herr
		}
		playTeamInfo = append(playTeamInfo, *hrow)
	}

	if hrows.Err() != nil {
		log.Error("Get Begin Play Info Sql Rows Err %v\n", herr)
		return nil, hrows.Err()
	}

	if playTeamInfo == nil {
		iRets := &PlayerTeamData{
			Msg:"fail",
			Code:400,
			PlayTeam:playTeamInfo[0],
		}
		return iRets, nil
	}

	iRets := &PlayerTeamData{
		Msg:"success",
		Code:200,
		PlayTeam:playTeamInfo[0],
	}
	return iRets, nil
}

func GetEndPlay(playId string) (*types.EndPlay, error) {
	//收到比赛结束的消息
	iRets := &types.EndPlay {
		Msg:"success request",
		Code:200,
	}
	return iRets, nil
}
