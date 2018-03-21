package kafka

import (
	"github.com/Shopify/sarama"
	"qiniupkg.com/x/log.v7"
	"superBrain/config"
)

var Consumer sarama.Consumer

func init() {
	consumer, err := sarama.NewConsumer(config.Get().Kafka.KafkaBrokers, nil)

	if err != nil {
		panic(err)
	}
	Consumer = consumer
	log.Info("kafka consumer init success!!")
}

func StartConsumer(topic string, publish *DefaultPublish) {
	log.Infof("kafka consumer StartConsumer is topic : %s", topic)
	partitionList, err := Consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := Consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		go func(partitionConsumer sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				publish.SetValue(msg.Value)
			}
		}(pc)
	}
}
