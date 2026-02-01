package inbound

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/netbill/ape"
	"github.com/netbill/evebox/box/inbox"
	"github.com/netbill/profiles-svc/internal/messenger/contracts"
)

func (i *Inbound) AccountCreated(
	ctx context.Context,
	event inbox.Event,
) inbox.EventStatus {
	var payload contracts.AccountCreatedPayload
	if err := json.Unmarshal(event.Payload, &payload); err != nil {
		i.log.Errorf("bad payload for %s, key %s, id: %s, error: %v", event.Type, event.Key, event.ID, err)
		return inbox.EventStatusFailed
	}

	if _, err := i.domain.CreateProfile(ctx, payload.AccountID, payload.Username); err != nil {
		var ae *ape.Error
		if errors.As(err, &ae) {
			i.log.Errorf("failed to create profile, key %s, id: %s, error: %v", event.Key, event.ID, err)
			return inbox.EventStatusFailed
		}

		i.log.Errorf("failed to create profile due to internal error, key %s, id: %s, error: %v", event.Key, event.ID, err)
		return inbox.EventStatusPending
	}

	return inbox.EventStatusProcessed
}
