package src

import (
	"github.com/go-pg/pg"
)


type Service struct {
	ID int64
	Name string
	Status string
	Description string
	Group string
	Link string
	Tags []string
	Enabled bool
}

type Services struct {
	db pg.DB
}

func (s *Services) Initialize(db pg.DB) {
	s.db = db
}

func (s *Services) InsertService(service Service) error {
	if service.Group == "" {
		service.Group = "Other"
	}

	err := s.db.Insert(&service)
	return err
}

func (s *Services) GetServices(enabled bool) ([]Service, error) {
	var services []Service

	err := s.db.Model(&services).
		Where("service.enabled = ?", true).
		Select()

	return services, err
}

func (s *Services) GetService(id int64) (Service, error){
	service := Service{
		ID: id,
	}

	err := s.db.Select(&service)

	return service, err
}

func (s *Services) UpdateService(id int64, service Service) error {
	service.ID = id
	if service.Group == "" {
		service.Group = "Other"
	}

	err := s.db.Update(&service)
	return err
}

func (s *Services) DeleteService(id int64) error {
	service := Service{
		ID: id,
	}

	err := s.db.Delete(&service)

	return err
}