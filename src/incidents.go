package src

import (
	"time"
	"github.com/go-pg/pg"
)

type Incident struct {
	Id int64
	Time time.Time
	Title string
	Updates []*IncidentUpdate
}

type IncidentUpdate struct {
	Id int64
	Time time.Time
	IncidentId int64
	Status string
	Message string
}

type Incidents struct {
	db pg.DB
}

func (i *Incidents) Initialize(db pg.DB) {
	i.db = db
}

func (i *Incidents) InsertIncident(incident Incident) error {
	if incident.Time.IsZero() {
		now := time.Now()
		incident.Time = now
	}
	err := i.db.Insert(&incident)
	return err
}

func (i *Incidents) InsertIncidentUpdate(incident int64, update IncidentUpdate) error {
	update.IncidentId = incident

	if update.Time.IsZero() {
		now := time.Now()
		update.Time = now
	}

	err := i.db.Insert(&update)
	return err
}

func (i *Incidents) GetLatestIncidents() ([]Incident, error) {
	to := time.Now()
	from := to.Add(-14 * 24 * time.Hour).Truncate(24 * time.Hour)

	var incidents []Incident

	err := i.db.Model(&incidents).
		Column("incident.*", "Updates").
		Where("time > ?", from).
		Where("time < ?", to).
		Select()

	return incidents, err
}