package kafka

var KafkaBrokers[] string

func InitKafuka(kafkaBrokers[] string) error {
	KafkaBrokers = kafkaBrokers
	return nil
}
