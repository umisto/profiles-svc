package outbound

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/netbill/evebox/header"
	"github.com/netbill/profiles-svc/internal/core/modules/profile"
	"github.com/netbill/profiles-svc/internal/messenger/contracts"
	"github.com/segmentio/kafka-go"
)

func (p Outbound) WriteProfileUpdated(
	ctx context.Context,
	accountID uuid.UUID,
	params profile.UpdateParams,
	updatedAt time.Time,
) error {
	payload, err := json.Marshal(contracts.AccountProfileUpdatedPayload{
		AccountID:   accountID,
		Pseudonym:   params.Pseudonym,
		Description: params.Description,
		UpdatedAt:   updatedAt,
	})
	if err != nil {
		return err
	}

	eventID := uuid.New().String()

	_, err = p.outbox.CreateOutboxEvent(
		ctx,
		kafka.Message{
			Topic: contracts.AccountsTopicV1,
			Key:   []byte(accountID.String()),
			Value: payload,
			Headers: []kafka.Header{
				{Key: header.EventID, Value: []byte(eventID)}, // Outbox will fill this
				{Key: header.EventType, Value: []byte(contracts.ProfileUpdatedEvent)},
				{Key: header.EventVersion, Value: []byte("1")},
				{Key: header.Producer, Value: []byte(contracts.ProfilesSvcGroup)},
				{Key: header.ContentType, Value: []byte("application/json")},
			},
		},
	)

	p.log.Debugf("profile updated event queued, account_id: %s, event_id: %s", accountID, eventID)

	return err
}
