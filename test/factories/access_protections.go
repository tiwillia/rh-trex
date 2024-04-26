package factories

import (
	"context"
	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/pkg/api"
)

func (f *Factories) Newaccess_protection(id string) (*api.access_protection, error) {
	access_protectionService := environments.Environment().Services.access_protections()

	access_protection := &api.access_protection{
		Meta:       api.Meta{ID: id},
	}

	sub, err := access_protectionService.Create(context.Background(), access_protection)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (f *Factories) Newaccess_protectionList(name string, count int) ([]*api.access_protection, error) {
	access_protections := []*api.access_protection{}
	for i := 1; i <= count; i++ {
		c, _ := f.Newaccess_protection(f.NewID())
		access_protections = append(access_protections, c)
	}
	return access_protections, nil
}
