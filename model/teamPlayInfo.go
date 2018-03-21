package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"superBrain/types"
	"superBrain/errCode"
)

type TeamPlayerInfo struct {
	TPInfo types.TeamPlayerIdOrName
}

func (this* TeamPlayerInfo)GetTeamPlayerInfo(equipmentId string) (int) {
	if equipmentId == " " {
		log.Error("GetTeamPlayerInfo sql return error")
		return errCode.QUERY_PLAYER_INFO_ERR
	}

	row := DB.QueryRow(fmt.Sprintf(`SELECT player.ID, player.PLAYER_NAME, player.EQUIPMENT_ID, team.ID, team.TEAM_NAME FROM  RD_PLAYER_INFO AS player, RD_TEAM_INFO AS team WHERE player.TEAM_ID = team.ID AND player.EQUIPMENT_ID = '%s'`, equipmentId))
	if row == nil {
		log.Error("GetTeamPlayerInfo sql return error")
		return errCode.QUERY_PLAYER_INFO_ERR
	}

	err := row.Scan(&this.TPInfo.PlayerId, &this.TPInfo.PlayerName, &this.TPInfo.EquipmentId, &this.TPInfo.TeamID, &this.TPInfo.TeamName)
	if err != nil  {
		log.Error("Row Scan return error")
		return errCode.QUERY_PLAYER_INFO_ERR
	}
	return errCode.QUERY_PLAYER_INFO_SUC
}