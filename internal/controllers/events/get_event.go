package events

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services/events"
	"net/http"
)

// GetEvent
// @Tags         events
// @Summary      Get an event.
// @Description  Get an event.
// @Param        id           	path      string  true  "Event UUID formatted ID"
// @Success      200            {object}  models.Event
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /events/{id} [get]
func GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, _ := ctx.Value("eventId").(uuid.UUID) // getting key set in context.go

	event, err := events.GetEventById(eventId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(event)
	_, _ = w.Write(body)
	return
}

// GetEvents
// @Tags         events
// @Summary      Get all events.
// @Description  Get all events.
// @Success      200            {array}   models.Event
// @Failure      500            "Something went wrong"
// @Router       /events [get]
func GetEvents(w http.ResponseWriter, r *http.Request) {
	eventsList, err := events.GetAllEvents()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(eventsList)
	_, _ = w.Write(body)
	return
}