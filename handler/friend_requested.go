package handler

import (
	"log"

	"github.com/clubo-app/protobuf/events"
)

func (s *server) FriendRequested(fr *events.FriendRequested) {
	log.Printf("%v got a friend request from %v", fr.FriendId, fr.UserId)
}
