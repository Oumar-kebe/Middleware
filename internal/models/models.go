package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Event struct {
	Id           *uuid.UUID `json:"id"`
	UID          string     `json:"uid"`            // UID de l'événement iCal
	Summary      string     `json:"summary"`        // Titre du cours (ex: "CM Big Data")
	Description  string     `json:"description"`    // Description complète
	Location     string     `json:"location"`       // Salle (ex: "IS_A104")
	StartDate    time.Time  `json:"start_date"`     // Date/heure début (DTSTART)
	EndDate      time.Time  `json:"end_date"`       // Date/heure fin (DTEND)
	LastModified time.Time  `json:"last_modified"`  // Dernière modification
	AgendaId     string     `json:"agenda_id"`      // ID de l'agenda source
}