package kafka

type Publish interface {
	Attach(interface{}, Subscriber)
	Detach(interface{})
	Notify()
}
type Subscriber interface {
	Receive(Publish, interface{})
}

type DefaultPublish struct {
	subs  map[interface{}]Subscriber
	value []byte
}

func NewDefaultPublish() *DefaultPublish {
	dp := new(DefaultPublish)
	dp.subs = make(map[interface{}]Subscriber)
	return dp
}

func (dp *DefaultPublish) Attach(i interface{}, subscriber Subscriber) {
	dp.subs[i] = subscriber
}

func (dp *DefaultPublish) Detach(i interface{}) {
	delete(dp.subs, i)
}

func (dp *DefaultPublish) Notify() {
	for key, value := range dp.subs {
		value.Receive(dp, key)
	}
}

func (dp *DefaultPublish) SetValue(value []byte) {
	dp.value = value
	dp.Notify()
}
func (dp *DefaultPublish) GetValue() []byte {
	return dp.value
}
