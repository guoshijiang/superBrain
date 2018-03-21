package engine

import (
	"superBrain/kafka"
	"github.com/ethereum/go-ethereum/log"
)

type Input struct {
	source string
	content []byte
}
type InputSubscriber struct {
	Engine DealEngine
}

func (iSub *InputSubscriber) Receive(publish kafka.Publish, i interface{})  {
	log.Info(string(publish.(*kafka.DefaultPublish).GetValue()))
	input := &Input{"kafka", publish.(*kafka.DefaultPublish).GetValue()}
	iSub.Engine.Deal(input)
}

func CreateInput() {
	sde := &SpeedDealEngine{}
	dp := kafka.NewDefaultPublish()
	dp.Attach("1", &InputSubscriber{sde})
	kafka.StartConsumer("maisiTest", dp)
}
