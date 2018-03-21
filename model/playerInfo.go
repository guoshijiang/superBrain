package model

import "superBrain/calc"

type PlayerTeamInfo struct {
	PtInfo calc.PlayerInfo
}

func (this* PlayerTeamInfo)GetPlayerTeamId (FieldId string, EquipMentId string){
	this.PtInfo.TeamID = ""
}

