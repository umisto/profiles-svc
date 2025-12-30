package producer

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/umisto/profiles-svc/internal/domain/models"
	"github.com/umisto/profiles-svc/internal/messanger/contracts"
)

func (s Service) WriteProfileUpdated(
	ctx context.Context,
	profile models.Profile,
) error {
	payload, err := json.Marshal(contracts.ProfileUpdatedPayload{
		Profile: profile,
	})
	if err != nil {
		return err
	}

	_, err = s.outbox.CreateOutboxEvent(
		ctx,
		contracts.ProfileUpdatedEvent,
		kafka.Message{
			Topic: contracts.ProfileUpdatedEvent,
			Key:   []byte(profile.AccountID.String()),
			Value: payload,
			Headers: []kafka.Header{
				{Key: "EventID", Value: []byte(uuid.New().String())}, // Outbox will fill this
				{Key: "EventType", Value: []byte(contracts.ProfileUpdatedEvent)},
				{Key: "EventVersion", Value: []byte("1")},
				{Key: "Producer", Value: []byte(contracts.GroupProfilesSvc)},
				{Key: "ContentType", Value: []byte("application/json")},
			},
		},
	)

	return err
}
