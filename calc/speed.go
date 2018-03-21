package calc

import (
	"superBrain/types"
	"strconv"
	"github.com/ethereum/go-ethereum/log"
	"strings"
	//"fmt"
)

var PersonSpeed []int   //个人速度
var ProBufArr []string

type CalcSpeed interface {
	PesonTimeSpeed(buf string)
	PersonTimeAvgSpeed(buf string)
	PersonAvgSpeed(buf string)
	PersonHighSpeed(buf string)
	TeamAvgSpeed(buf string)
	TeamHighSpeed(buf string)
}

type SpeedCalcData struct {
	CalcPersonSpeed types.PersonTimeSpeed
	CalcTeamSpeed types.TeamSpeed
}

func SpeedData() {
	//naturals := make(chan int)
	//squares := make(chan int)
	//go counter(naturals)
	//go squarer(squares, naturals)
	//printer(squares)
}

//个人实时速度
func (this *SpeedCalcData) PesonTimeSpeed(buf string) error{
	ProBufArr = strings.Split(buf, ",")
	//fmt.Println(ProBufArr)
	return nil
}

//个人实时平均速速
func (this *SpeedCalcData) PersonTimeAvgSpeed(buf string) error{
	var speedInt int
	var err error
	var speedNum int
    if speedInt, err = strconv.Atoi(buf); err !=nil {
    	log.Error("function PersonTimeAvgSpeed get the speed of person error %v\n", err)
    	return err
	}
	PersonSpeed = append(PersonSpeed, speedInt)
	sliceLen := len(PersonSpeed)
	for i := 0; i < sliceLen; i++ {
		speedNum += PersonSpeed[i]
	}
	this.CalcPersonSpeed.TimeAvgSpeed = speedNum / sliceLen
	return nil
}

//个人平均速度
func (this *SpeedCalcData) PersonAvgSpeed(buf string) error{
	var speedInt int
	var err error
	var speedNum int
	if speedInt, err = strconv.Atoi(buf); err !=nil {
		log.Error("function PersonAvgSpeed get the speed of person error %v\n", err)
		return err
	}
	PersonSpeed = append(PersonSpeed, speedInt)
	sliceLen := len(PersonSpeed)
	for i := 0; i < sliceLen; i++ {
		speedNum += PersonSpeed[i]
	}
	this.CalcPersonSpeed.AvgSpeed = speedNum / sliceLen
	return nil
}

//个人最快速度
func (this *SpeedCalcData) PersonHighSpeed(buf string) error{
	return nil
}

//队内平均速度
func (CalcSpeed SpeedCalcData) TeamAvgSpeed(buf string) error{
	return nil
}

//队内最快速度
func (CalcSpeed SpeedCalcData) TeamHighSpeed(buf string) error {
	var(
		//TeamMaxSpeed int
	)
	return nil
}

