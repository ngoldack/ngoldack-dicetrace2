package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"golang.org/x/sync/errgroup"
)

type KafkaClient struct {
	brokers   []string
	dialer    *kafka.Dialer
	producers map[string]*kafka.Writer
	consumers map[string]*kafka.Reader
}

func NewKafkaClient(cfg *Config) (*KafkaClient, error) {
	mechanism, err := scram.Mechanism(scram.SHA256, cfg.KafkaUsername, cfg.KafkaPassword)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to stop")
	}
	client := &KafkaClient{
		brokers: []string{cfg.KafkaHost},
		dialer: &kafka.Dialer{
			SASLMechanism: mechanism,
			TLS:           &tls.Config{},
		},
		producers: make(map[string]*kafka.Writer),
		consumers: make(map[string]*kafka.Reader),
	}

	return client, nil
}

// Producer returns a producer, if it not exists already it will be created
func (k KafkaClient) Producer(topic string) *kafka.Writer {
	if _, ok := k.producers[topic]; !ok {
		k.producers[topic] = kafka.NewWriter(kafka.WriterConfig{
			Brokers: k.brokers,
			Topic:   topic,
			Dialer:  k.dialer,
		})
	}

	return k.producers[topic]
}

// Consumer returns a consumer, if it not exists already it will be created
func (k KafkaClient) Consumer(topic string) *kafka.Reader {
	if _, ok := k.consumers[topic]; !ok {
		k.consumers[topic] = kafka.NewReader(kafka.ReaderConfig{
			Brokers: k.brokers,
			Topic:   topic,
			Dialer:  k.dialer,
		})
	}

	return k.consumers[topic]
}

func (k KafkaClient) Start(ctx context.Context) error {
	testTopic := "test"
	testMessage := "test"

	k.producers[testTopic] = k.Producer(testTopic)
	k.consumers[testTopic] = k.Consumer(testTopic)

	err := k.producers[testTopic].WriteMessages(ctx, kafka.Message{
		Value: []byte(testMessage),
	})
	if err != nil {
		return err
	}

	msg, err := k.consumers[testTopic].ReadMessage(ctx)
	if err != nil {
		return err
	}
	if string(msg.Value) != testMessage {
		return fmt.Errorf("test messages are not equal: got: '%s', expected: '%s'", string(msg.Value), testMessage)
	}

	return nil
}

func (k KafkaClient) Stop(ctx context.Context) error {
	errs, _ := errgroup.WithContext(ctx)

	for _, writer := range k.producers {
		errs.Go(func() error {
			return writer.Close()
		})
	}

	for _, reader := range k.consumers {
		errs.Go(func() error {
			return reader.Close()
		})
	}

	err := errs.Wait()
	if err != nil {
		return err
	}

	log.Info().Msgf("successfully closed %v-producers and %v-consumers", len(k.producers), len(k.consumers))
	return nil
}
