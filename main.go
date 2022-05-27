package main

import (
	"log"
	"sync"

	"github.com/clubo-app/notification-service/config"
	"github.com/clubo-app/notification-service/handler"
	"github.com/clubo-app/notification-service/mail"
	"github.com/clubo-app/packages/stream"
	"github.com/clubo-app/protobuf/events"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Println("No .env file found")
	}

	opts := []nats.Option{nats.Name("Notification Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	st := stream.New(nc)

	smtp, err := mail.Connect(c)
	if err != nil {
		log.Fatalln(err)
	}

	s := handler.NewServer(smtp)

	var wg sync.WaitGroup
	wg.Add(3)

	go st.SubscribeToEvent("notification.email.friend.requested", events.FriendRequested{}, s.FriendRequested)
	go st.SubscribeToEvent("notification.email.verify", events.Registered{}, s.Registered)

	go st.SubscribeToEvent("notification.push.party.created", events.PartyCreated{}, s.PartyCreated)

	// this will wait until the wg counter is at 0
	wg.Wait()
}
