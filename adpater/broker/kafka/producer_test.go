package kafka

import (
	"testing"

	ckakfa "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/maxhariel/gateway/adpater/presenter/transaction"
	"github.com/maxhariel/gateway/domain/entity"
	"github.com/maxhariel/gateway/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECT,
		ErrorMessage: "you can't do transcation with value greater than 1000.0",
	}

	// outputJSON, _ := json.Marshal(expectOutput)

	configMap := ckakfa.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransacationKafkaPresente())

	err := producer.Publish(expectOutput, []byte("1"), "test")

	assert.Nil(t, err)

}
