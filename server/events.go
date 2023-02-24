package main

import (
	"log"
	"time"

	pb "github.com/kainn9/grpc_game/proto"
)

type event struct {
	*pb.PlayerReq
	stalled bool
}

// Starts a new ticker loop that calls processEventsPerTick with the given world
func newTickLoop(w *world) {
	go func() {
		ticker := time.NewTicker(time.Second / 60)
		defer ticker.Stop()

		for range ticker.C {
			processEventsPerTick(w)
		}
	}()
}



// Process events in the given world, removing each event as it is processed
func processEventsPerTick(w *world) {
	w.mutex.Lock()
  defer w.mutex.Unlock()

	// Log information about the number of events in the world, if it exceeds certain thresholds
	if len(w.events) > 25 {
		log.Printf("WORLD: %v\n", w.name)
		log.Printf("LEN! 25 %v\n", len(w.events) > 25)
		log.Printf("LEN! 50 %v\n", len(w.events) > 50)
		log.Printf("LEN! 100 %v\n", len(w.events) > 100)
		log.Printf("LEN! %v\n", len(w.events))
	}

	stalledEvents := make(map[string]bool)
	stalledEventsToSkip := make(map[string]bool)
	
	type dupeCount int
	dupeEvents := make(map[string]dupeCount)

	eventBatchSize := 100

	// Get all stalled events
	for i := 0; i < eventBatchSize; i++ {


		// Exit function if no events
		if len(w.events) == 0 {
			return
		}

		// If current index is out of range, exit the loop, adjust batch length
		if i > len(w.events)-1 {
			eventBatchSize = len(w.events)
			break
		}

		ev := w.events[i]

		if ev.stalled {
			stalledEvents[ev.Id] = true
		} else {
			dupeEvents[ev.Id]++
		}
	}

	// Determine which stalled events to skip
	for i := 0; i < eventBatchSize; i++ {

		ev := w.events[i]

		if !ev.stalled && stalledEvents[ev.Id] {
			stalledEventsToSkip[ev.Id] = true
		}
	}

	// Process all "valid" events
	for i := 0; i < eventBatchSize; i++ {

		ev := w.events[i]
		
		// If there is a player associated with the event, handle the event with the player and world
		if w.players[ev.Id] != nil && !stalledEventsToSkip[ev.Id] && (dupeEvents[ev.Id] < 2) {
			cp := w.players[ev.Id]
			w.Update(cp, ev.Input)
		}

		dupeEvents[ev.Id]--
		if dupeEvents[ev.Id] <= 0 {
			delete(dupeEvents, ev.Id)
		}

	}
	// Remove the event from the world's events queue
	w.events = w.events[eventBatchSize:]
}



func newEvent(req *pb.PlayerReq, stalled bool) *event {
	return &event{
		PlayerReq: req,
		stalled: stalled,
	}
}

func (e *event) enqueue(w *world) {
	w.mutex.Lock()
	w.events = append(w.events, e)
	w.mutex.Unlock()
}

		