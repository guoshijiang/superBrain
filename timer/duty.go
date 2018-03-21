package timer

import (
	"github.com/robfig/cron"
	"fmt"
)

type TimerDuty struct {

}

//速度调度器
func (this *TimerDuty)DutyTimerSpeed(buf string) {
	c := cron.New()
	spec := "*/3 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println(buf, "数据入库")
	})
	c.Start()
}

func (this* TimerDuty) DutyTimeDistance(buf string) {

}

func (this* TimerDuty) DutyTimeHearRate(buf string) {

}

func (this* TimerDuty) DutyEnterNumPeple(buf string) {

}

func (this* TimerDuty) DutyElectric(buf string) {
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		fmt.Println("护腿板电量值护腿板电量值护腿板电量值护腿板电量值护腿板电量值")
	})
	c.Start()
}