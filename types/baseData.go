package types

type Hdtas_position_bolt struct {
	SignId string
	FieldId string
	X string
	Y string
	Z string
	Rx string
	Ry string
	Speed string
	Electric string
	HeartReat string
	Ax string
	Ay string
	Az string
	R string
	Gx string
	Gy string
	Gz string
	Zero string
	Time string
}

type Hdtas_speed_bolt struct {
	SignID string
	FieldID string
	Time string
	Diff string
	Dis string
	Speed string
}

type Test struct {
	Id int
	Name string
}

//个人速度结构体存储
type PersonTimeSpeed struct{
	TimeSpeed int      //个人实时速度
	TimeAvgSpeed int   //个人实时平均速度
	AvgSpeed int       //个人平均速度
	HighSpeed int      //个人最快速度
}

type TeamSpeed struct {
	AvgSpeed string   //队平均速度
	HighSpeed string  //队最快速度
}

type EndPlay struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
}

type PlayerTeamInfo struct {
	TeamName string `json:"team_name"`
	IsDel int `json:"is_del"`
}

type TeamPlayerIdOrName struct {
	PlayerId int
	EquipmentId string
	TeamID int
	PlayerName string
	TeamName string
}






