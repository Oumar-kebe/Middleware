package events

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllEvents() ([]models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM events")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	events := []models.Event{}
	for rows.Next() {
		var data models.Event
		err = rows.Scan(
			&data.Id, 
			&data.UID, 
			&data.Summary, 
			&data.Description, 
			&data.Location,
			&data.StartDate,
			&data.EndDate,
			&data.LastModified,
			&data.AgendaId,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, data)
	}
	_ = rows.Close()

	return events, err
}

func GetEventById(id uuid.UUID) (*models.Event, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM events WHERE id=?", id.String())
	helpers.CloseDB(db)

	var event models.Event
	err = row.Scan(
		&event.Id, 
		&event.UID, 
		&event.Summary, 
		&event.Description, 
		&event.Location,
		&event.StartDate,
		&event.EndDate,
		&event.LastModified,
		&event.AgendaId,
	)
	if err != nil {
		return nil, err
	}
	return &event, err
}