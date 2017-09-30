package src

import (
	"github.com/go-redis/redis"
	"encoding/json"
)

const REDIS_MAP = "statuspage_services"

type Service struct {
	ID string
	Name string
	Status string
	Description string
	Group string
	Link string
	Tags []string
	Enabled bool
}

type Services struct {
	db redis.Client
}

func (s *Services) Initialize(db redis.Client) {
	s.db = db
}

func (s *Services) InsertService(id string, service Service) error {
	serialized, err := json.Marshal(service)
	if err != nil { return err }

	if err := s.db.HSet(REDIS_MAP, id, serialized).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Services) RemoveService(id string) error {
	err := s.db.HDel(REDIS_MAP, id).Err()
	return err
}

func (s *Services) GetServices(enabled bool) ([]Service, error) {
	services := []Service{}

	results, err := s.db.HGetAll(REDIS_MAP).Result()
	if err != err {
		return nil, err
	}

	for id, result := range results {
		var service Service
		if err := json.Unmarshal([]byte(result), &service); err != nil {
			return nil, err
		}
		service.ID = id

		if service.Enabled == enabled {
			services = append(services, service)
		}
	}

	return services, nil
}

func (s *Services) GetService(id string) (Service, error){
	result, err := s.db.HGet(REDIS_MAP, id).Result()
	if err != err {
		return Service{}, err
	}

	var service Service
	if err := json.Unmarshal([]byte(result), &service); err != nil {
		return Service{}, err
	}

	return service, nil
}