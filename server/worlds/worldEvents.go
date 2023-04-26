package worlds

import (
	"log"
	"time"

	evt "github.com/kainn9/grpc_game/server/event"
)

// TODO: RESTRICT EVENTS TO ONLY ALLOWED TYPES!!!!

// Starts a new ticker loop that calls processEventsPerTick with the given world
func newTickLoop(w *World) {

	log.Printf("Starting Tick Loop for %v\n", w.name)

	go func() {
		ticker := time.NewTicker(time.Second / 60)
		defer ticker.Stop()

		for range ticker.C {
			processEventsPerTick(w)
		}
	}()
}

// Process events in the given world, removing each event as it is processed
func processEventsPerTick(w *World) {
	w.eventsMutex.Lock()
	defer w.eventsMutex.Unlock()

	logHighEventCount(w)

	type dupeCount int
	dupeEvents := make(map[string]dupeCount)
	eventBatchSize := 100

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
		dupeKey := ev.Id + string(ev.EventCategory)

		dupeEvents[dupeKey]++
	}

	// Process all "valid" events
	for i := 0; i < eventBatchSize; i++ {

		ev := w.events[i]
		dupeKey := ev.Id + string(ev.EventCategory)

		// If there is a player associated with the event, handle the event with the player and world
		w.WPlayersMutex.RLock()
		cp := w.Players[ev.Id]
		w.WPlayersMutex.RUnlock()

		if cp != nil && (dupeEvents[dupeKey] < 2) && ev.Valid() {
			w.Update(cp, ev.Input)

		}

		dupeEvents[dupeKey]--
		if dupeEvents[dupeKey] <= 0 {
			delete(dupeEvents, dupeKey)
		}

	}
	// Remove the event from the world's events queue
	w.events = w.events[eventBatchSize:]
}

/*
TODO:
* Buffering/re-enqueuing dupes experiment...
* needs to a seperate filter to avoid enqueueing
* too may dupes(creates lag)
* something like only one dupe per
* input or something maybe
*/
// func processEventsPerTick(w *world) {
// 	w.eventsMutex.Lock()
// 	defer w.eventsMutex.Unlock()

// 	logHighEventCount(w)

// 	dupesToEnqueue := make([]*event, 0)
// 	dupesToKeep := make(map[string]*event)

// 	eventBatchSize := 100

// 	for i := 0; i < eventBatchSize; i++ {

// 		// Exit function if no events
// 		if len(w.events) == 0 {
// 			return
// 		}

// 		// If current index is out of range, exit the loop, adjust batch length
// 		if i > len(w.events)-1 {
// 			eventBatchSize = len(w.events)
// 			break
// 		}

// 		ev := w.events[i]
// 		dupeKey := ev.Id + string(ev.eventCategory)

// 		_, isDupe := dupesToKeep[dupeKey]

// 		if isDupe && ev.Input != "nada" {
// 			dupesToEnqueue = append(dupesToEnqueue, ev)
// 		} else {
// 			dupesToKeep[dupeKey] = ev
// 		}

// 	}

// 	// Process all "valid" events
// 	for i := 0; i < eventBatchSize; i++ {

// 		ev := w.events[i]
// 		dupeKey := ev.Id + string(ev.eventCategory)

// 		// If there is a player associated with the event, handle the event with the player and world
// 		w.wPlayersMutex.RLock()
// 		cp := w.players[ev.Id]
// 		w.wPlayersMutex.RUnlock()

// 		if cp != nil && (dupesToKeep[dupeKey] == ev) {
// 			w.Update(cp, ev.Input)
// 		}
// 	}

// 	// Remove the event from the world's events queue
// 	w.events = w.events[eventBatchSize:]

// 	// Enqueue dupes at top
// 	if len(dupesToEnqueue) > 0 {
// 		// prepend duplicates to the events slice
// 		w.events = append(dupesToEnqueue, w.events...)
// 	}
// }

func Enqueue(w *World, e *evt.Event) {
	w.eventsMutex.Lock()
	w.events = append(w.events, e)
	w.eventsMutex.Unlock()
}

func logHighEventCount(w *World) {
	if len(w.events) > 25 {
		log.Printf("WORLD: %v\n", w.name)
		log.Printf("LEN! 25 %v\n", len(w.events) > 25)
		log.Printf("LEN! 50 %v\n", len(w.events) > 50)
		log.Printf("LEN! 100 %v\n", len(w.events) > 100)
		log.Printf("LEN! %v\n", len(w.events))
	}
}
