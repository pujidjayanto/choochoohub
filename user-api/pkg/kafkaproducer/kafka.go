package kafkaproducer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// Producer defines the interface for sending messages to Kafka.
type Producer interface {
	// SendMessage publishes a message to a topic and partition.
	SendMessage(ctx context.Context, topic string, partition int, message []byte) error
}

type producer struct {
	broker string
	logger *logrus.Logger
}

// NewProducer returns a new Kafka producer with the given broker address.
// Example: NewProducer("localhost:9092")
func NewProducer(broker string, logger *logrus.Logger) Producer {
	return &producer{broker: broker}
}

func (p *producer) SendMessage(ctx context.Context, topic string, partition int, message []byte) error {
	conn, err := kafka.DialLeader(ctx, "tcp", p.broker, topic, partition)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			p.logger.WithField("err", err).Info("failed to close kafka connection")
		}
	}()

	_ = conn.SetWriteDeadline(time.Now().Add(3 * time.Second))

	_, err = conn.WriteMessages(kafka.Message{
		Value: message,
	})
	return err
}
