package src

import (
	"github.com/go-redis/redis"
	"time"
)

type Incident struct {
	Id string
	Time time.Time
	Title string
	Resolved time.Time
}

type IncidentUpdate struct {
	Id string
	Incident string
	Status string
	Message string
}

type Incidents struct {
	db redis.Client
}

func (i *Incidents) Initialize(db redis.Client) {
	i.db = db
}

func (i *Incidents) CreateIncident(incident Incident) error {
	return nil
}

func (i *Incidents) CreateIncidentUpdate(incident string, update IncidentUpdate) error {
	return nil
}