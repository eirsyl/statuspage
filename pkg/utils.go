package pkg

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	log.Printf("Running with %d CPUs\n", nuCPU)
}

/*
 * Aggregate and group services by service group.
 */

type AggregatedServices map[string][]Service

func AggregateServices(services []Service) AggregatedServices {
	aggregated := AggregatedServices{}

	for _, service := range services {
		if service.Enabled == false {
			continue
		}

		groupName := service.Group

		group := aggregated[groupName]
		if group != nil {
			group = append(group, service)
		} else {
			group = []Service{service}
		}
		aggregated[groupName] = group
	}

	return aggregated
}

/*
 * Extract the most critical service.
 */

func MostCriticalStatus(services []Service) int {
	statusValues := map[string]int{
		"Operational":        0,
		"Performance Issues": 1,
		"Partial Outage":     2,
		"Major Outage":       3,
	}

	mostCritical := 0

	for _, service := range services {
		serviceStatus := statusValues[service.Status]
		if serviceStatus > mostCritical {
			mostCritical = serviceStatus
		}
	}

	return mostCritical
}

/*
 * Aggregate incidents
 */

type AggregatedIncident struct {
	Time      time.Time
	Incidents []Incident
}

type AggregatedIncidents []AggregatedIncident

func AggregateIncidents(incidents []Incident) AggregatedIncidents {
	days := 14
	aggregatedIncidents := AggregatedIncidents{}

	for i := 0; i < days; i++ {
		t := time.Now().Add(-time.Duration(i) * 24 * time.Hour)
		filteredIncidents := []Incident{}

		for _, incident := range incidents {
			if incident.Time.Day() == t.Day() {
				filteredIncidents = append(filteredIncidents, incident)
			}
		}

		aggregatedIncidents = append(aggregatedIncidents, AggregatedIncident{
			Time:      t,
			Incidents: filteredIncidents,
		})
	}

	return aggregatedIncidents
}

/*
 * Migrate DB
 */
func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		&Service{},
		&Incident{},
		&IncidentUpdate{},
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveLastSlash removes the last character in a string if the character is a string.
func RemoveLastSlash(url string) string {
	l := len(url)
	if l > 0 && url[l-1] == '/' {
		return url[:l-2]
	}
	return url
}
