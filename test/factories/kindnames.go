package factories

import (
	"context"
	"github.com/openshift-online/rh-trex/cmd/trex/environments"
	"github.com/openshift-online/rh-trex/pkg/api"
)

func (f *Factories) NewKindName(id string) (*api.KindName, error) {
	KindNameService := environments.Environment().Services.KindNames()

	KindName := &api.KindName{
		Meta:       api.Meta{ID: id},
	}

	sub, err := KindNameService.Create(context.Background(), KindName)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (f *Factories) NewKindNameList(name string, count int) ([]*api.KindName, error) {
	KindNames := []*api.KindName{}
	for i := 1; i <= count; i++ {
		c, _ := f.NewKindName(f.NewID())
		KindNames = append(KindNames, c)
	}
	return KindNames, nil
}
