package src


/*
 * Aggregate and group services by service group.
 */

type AggregatedServices map[string][]Service

func AggregateServices (services []Service) AggregatedServices {
	aggregated := AggregatedServices{}

	for _, service := range services {
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
		"Operational": 0,
		"Performance Issues": 1,
		"Partial Outage": 2,
		"Major Outage": 3,
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