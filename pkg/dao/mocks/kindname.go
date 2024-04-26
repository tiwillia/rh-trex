package mocks

import (
	"context"

	"gorm.io/gorm"

	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/dao"
	"github.com/openshift-online/rh-trex/pkg/errors"
)

var _ dao.KindNameDao = &kindnameDaoMock{}

type kindnameDaoMock struct {
	kindnames api.KindNameList
}

func NewKindNameDao() *kindnameDaoMock {
	return &kindnameDaoMock{}
}

func (d *kindnameDaoMock) Get(ctx context.Context, id string) (*api.KindName, error) {
	for _, dino := range d.kindnames {
		if dino.ID == id {
			return dino, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (d *kindnameDaoMock) Create(ctx context.Context, kindname *api.KindName) (*api.KindName, error) {
	d.kindnames = append(d.kindnames, kindname)
	return kindname, nil
}

func (d *kindnameDaoMock) Replace(ctx context.Context, kindname *api.KindName) (*api.KindName, error) {
	return nil, errors.NotImplemented("KindName").AsError()
}

func (d *kindnameDaoMock) Delete(ctx context.Context, id string) error {
	return errors.NotImplemented("KindName").AsError()
}

func (d *kindnameDaoMock) FindByIDs(ctx context.Context, ids []string) (api.KindNameList, error) {
	return nil, errors.NotImplemented("KindName").AsError()
}

func (d *kindnameDaoMock) All(ctx context.Context) (api.KindNameList, error) {
	return d.kindnames, nil
}
