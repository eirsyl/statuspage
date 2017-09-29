package statuspage

import "github.com/go-redis/redis"

type Service struct {
	id int
	name string
	status string
	description string
	group string
	link string
	tags []string
	enabled bool
}

type Services struct {
	db redis.Client
}

func (s *Services) Initialize(db redis.Client) {
	s.db = db
}

func (s *Services) GetServices() []Service {
	services := []Service{}

	services = append(services, Service{
		id: 1,
		name: "API",
		status: "Operational",
		description: "API serving data for the LEGO system.",
		group: "LEGO",
		link: "https://lego.abakus.no",
		tags: []string{"LEGO", "API", "Django", "Rest-Framework"},
		enabled: true,
	})

	services = append(services, Service{
		id: 2,
		name: "Webapp",
		status: "Partial Outage",
		description: "Webapp provides the GUI in the LEGO system.",
		group: "LEGO",
		link: "https://webapp.abakus.no",
		tags: []string{"LEGO", "React", "Redux"},
		enabled: true,
	})

	return services
}