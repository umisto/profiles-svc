package cmd

import (
	"context"
	"database/sql"
	"sync"

	"github.com/netbill/logium"
	"github.com/netbill/profiles-svc/internal"
	"github.com/netbill/profiles-svc/internal/core/modules/profile"
	"github.com/netbill/profiles-svc/internal/messenger"
	"github.com/netbill/profiles-svc/internal/messenger/inbound"
	"github.com/netbill/profiles-svc/internal/messenger/outbound"
	"github.com/netbill/profiles-svc/internal/repository"
	"github.com/netbill/restkit/mdlv"

	"github.com/netbill/profiles-svc/internal/rest"
	"github.com/netbill/profiles-svc/internal/rest/controller"
)

func StartServices(ctx context.Context, cfg internal.Config, log logium.Logger, wg *sync.WaitGroup) {
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	pg, err := sql.Open("postgres", cfg.Database.SQL.URL)
	if err != nil {
		log.Fatal("failed to connect to database", "error", err)
	}

	repo := repository.New(pg)

	kafkaOutbound := outbound.New(log, pg)

	profileSvc := profile.New(repo, kafkaOutbound)

	ctrl := controller.New(log, profileSvc)
	mdll := mdlv.New(cfg.JWT.User.AccessToken.SecretKey, rest.AccountDataCtxKey, log)
	router := rest.New(log, mdll, ctrl)

	msgx := messenger.New(log, pg, cfg.Kafka.Brokers...)

	run(func() { router.Run(ctx, cfg) })

	log.Infof("starting kafka brokers %s", cfg.Kafka.Brokers)

	run(func() { msgx.RunProducer(ctx) })

	run(func() { msgx.RunConsumer(ctx, inbound.New(log, profileSvc)) })
}
