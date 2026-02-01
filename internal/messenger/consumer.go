package messenger

import (
	"context"
	"sync"
	"time"

	"github.com/netbill/evebox/box/inbox"
	"github.com/netbill/evebox/consumer"
	"github.com/netbill/profiles-svc/internal/messenger/contracts"
	"github.com/segmentio/kafka-go"
)

type handlers interface {
	AccountCreated(
		ctx context.Context,
		event inbox.Event,
	) inbox.EventStatus
	AccountDeleted(
		ctx context.Context,
		event inbox.Event,
	) inbox.EventStatus
	AccountUsernameUpdated(
		ctx context.Context,
		event inbox.Event,
	) inbox.EventStatus
}

func (m *Messenger) RunConsumer(ctx context.Context, handlers handlers) {
	wg := &sync.WaitGroup{}
	run := func(f func()) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f()
		}()
	}

	accountConsumer := consumer.New(consumer.NewConsumerParams{
		Log:  m.log,
		DB:   m.db,
		Name: "profiles-svc-account-consumer",
		Addr: m.addr,
		OnUnknown: func(ctx context.Context, m kafka.Message, eventType string) error {
			return nil
		},
	})

	accountConsumer.Handle(contracts.AccountCreatedEvent, handlers.AccountCreated)
	accountConsumer.Handle(contracts.AccountDeletedEvent, handlers.AccountDeleted)
	accountConsumer.Handle(contracts.AccountUsernameUpdatedEvent, handlers.AccountUsernameUpdated)

	inboxer1 := consumer.NewInboxer(
		consumer.NewInboxerParams{
			Log:        m.log,
			Pool:       m.db,
			Name:       "profiles-svc-inbox-worker-1",
			BatchSize:  10,
			RetryDelay: 1 * time.Minute,
			MinSleep:   100 * time.Millisecond,
			MaxSleep:   1 * time.Second,
			Unknown: func(ctx context.Context, ev inbox.Event) inbox.EventStatus {
				return inbox.EventStatusFailed
			},
		},
	)

	inboxer1.Handle(contracts.AccountCreatedEvent, handlers.AccountCreated)
	inboxer1.Handle(contracts.AccountDeletedEvent, handlers.AccountDeleted)
	inboxer1.Handle(contracts.AccountUsernameUpdatedEvent, handlers.AccountUsernameUpdated)

	inboxer2 := consumer.NewInboxer(
		consumer.NewInboxerParams{
			Log:        m.log,
			Pool:       m.db,
			Name:       "profiles-svc-inbox-worker-2",
			BatchSize:  10,
			RetryDelay: 1 * time.Minute,
			MinSleep:   100 * time.Millisecond,
			MaxSleep:   1 * time.Second,
			Unknown: func(ctx context.Context, ev inbox.Event) inbox.EventStatus {
				return inbox.EventStatusFailed
			},
		},
	)

	inboxer2.Handle(contracts.AccountCreatedEvent, handlers.AccountCreated)
	inboxer2.Handle(contracts.AccountDeletedEvent, handlers.AccountDeleted)
	inboxer2.Handle(contracts.AccountUsernameUpdatedEvent, handlers.AccountUsernameUpdated)

	run(func() {
		accountConsumer.Run(ctx, contracts.ProfilesSvcGroup, contracts.AccountsTopicV1, m.addr...)
	})

	run(func() {
		inboxer1.Run(ctx)
	})

	run(func() {
		inboxer2.Run(ctx)
	})

	wg.Wait()
}
