package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/ponyo877/product-expiry-tracker/internal/shared/message"
)

type MessagingConfig struct {
	Host string
	Port string
}

type MessageBroker interface {
	Publish(ctx context.Context, subject string, msg *message.Message) error
	Subscribe(ctx context.Context, subject string, handler MessageHandler) error
	Close() error
}

func NewMessageBroker(config MessagingConfig) (MessageBroker, error) {
	natsBroker, err := newNATSBroker(config.Host, config.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	return natsBroker, nil
}

type NATSBroker struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

type MessageHandler func(ctx context.Context, msg *message.Message) error

func newNATSBroker(host, port string) (*NATSBroker, error) {
	url := fmt.Sprintf("nats://%s:%s", host, port)
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// Create JetStream context
	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create JetStream context: %w", err)
	}

	return &NATSBroker{
		conn: conn,
		js:   js,
	}, nil
}

func (n *NATSBroker) Publish(ctx context.Context, subject string, msg *message.Message) error {
	// Ensure stream exists
	if err := n.ensureStream(subject); err != nil {
		return fmt.Errorf("failed to ensure stream: %w", err)
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	_, err = n.js.Publish(subject, msgBytes)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (n *NATSBroker) Subscribe(ctx context.Context, subject string, handler MessageHandler) error {
	// Ensure stream exists
	if err := n.ensureStream(subject); err != nil {
		return fmt.Errorf("failed to ensure stream: %w", err)
	}

	// Create durable consumer
	consumerName := fmt.Sprintf("%s_consumer", subject)

	sub, err := n.js.PullSubscribe(subject, consumerName)
	if err != nil {
		return fmt.Errorf("failed to create pull subscription: %w", err)
	}

	// Process messages
	for {
		select {
		case <-ctx.Done():
			sub.Unsubscribe()
			return ctx.Err()
		default:
			msgs, err := sub.Fetch(1, nats.MaxWait(time.Second))
			if err != nil {
				if err == nats.ErrTimeout {
					continue
				}
				log.Printf("Error fetching messages: %v", err)
				continue
			}

			for _, natsMsg := range msgs {
				if err := n.processMessage(ctx, natsMsg, handler); err != nil {
					log.Printf("Error processing message: %v", err)
					natsMsg.Nak()
				} else {
					natsMsg.Ack()
				}
			}
		}
	}
}

func (n *NATSBroker) processMessage(ctx context.Context, natsMsg *nats.Msg, handler MessageHandler) error {
	var msg message.Message
	if err := json.Unmarshal(natsMsg.Data, &msg); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	// Set message ID from NATS metadata if available
	if meta, err := natsMsg.Metadata(); err == nil {
		msg.ID = fmt.Sprintf("%d-%d", meta.Stream, meta.Sequence)
	}

	return handler(ctx, &msg)
}

func (n *NATSBroker) ensureStream(subject string) error {
	streamName := fmt.Sprintf("%s_stream", subject)

	// Check if stream exists
	_, err := n.js.StreamInfo(streamName)
	if err == nil {
		return nil // Stream already exists
	}

	// Create stream if it doesn't exist
	_, err = n.js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{subject},
		Storage:  nats.FileStorage,
		MaxAge:   24 * time.Hour, // Keep messages for 24 hours
	})
	if err != nil {
		return fmt.Errorf("failed to create stream: %w", err)
	}

	return nil
}

func (n *NATSBroker) Close() error {
	if n.conn != nil {
		n.conn.Close()
	}
	return nil
}
