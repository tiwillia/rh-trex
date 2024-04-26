package factories

import (
	"context"
	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/pkg/api"
)

func (f *Factories) NewAsteroid(id string) (*api.Asteroid, error) {
	AsteroidService := environments.Environment().Services.Asteroids()

	Asteroid := &api.Asteroid{
		Meta:       api.Meta{ID: id},
	}

	sub, err := AsteroidService.Create(context.Background(), Asteroid)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (f *Factories) NewAsteroidList(name string, count int) ([]*api.Asteroid, error) {
	Asteroids := []*api.Asteroid{}
	for i := 1; i <= count; i++ {
		c, _ := f.NewAsteroid(f.NewID())
		Asteroids = append(Asteroids, c)
	}
	return Asteroids, nil
}
