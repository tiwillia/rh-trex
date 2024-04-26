package factories

import (
	"context"
	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/pkg/api"
)

func (f *Factories) NewAccessProtection(id string) (*api.AccessProtection, error) {
	AccessProtectionService := environments.Environment().Services.AccessProtections()

	AccessProtection := &api.AccessProtection{
		Meta:       api.Meta{ID: id},
	}

	sub, err := AccessProtectionService.Create(context.Background(), AccessProtection)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (f *Factories) NewAccessProtectionList(name string, count int) ([]*api.AccessProtection, error) {
	AccessProtections := []*api.AccessProtection{}
	for i := 1; i <= count; i++ {
		c, _ := f.NewAccessProtection(f.NewID())
		AccessProtections = append(AccessProtections, c)
	}
	return AccessProtections, nil
}
