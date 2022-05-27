package handler

import (
	"log"

	"github.com/clubo-app/protobuf/events"
)

func (s *server) PartyCreated(p *events.PartyCreated) {
	log.Printf("%v", p)
}
