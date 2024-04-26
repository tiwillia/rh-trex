package dao

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/db"
)

type KindNameDao interface {
	Get(ctx context.Context, id string) (*api.KindName, error)
	Create(ctx context.Context, kindname *api.KindName) (*api.KindName, error)
	Replace(ctx context.Context, kindname *api.KindName) (*api.KindName, error)
	Delete(ctx context.Context, id string) error
	FindByIDs(ctx context.Context, ids []string) (api.KindNameList, error)
	All(ctx context.Context) (api.KindNameList, error)
}

var _ KindNameDao = &sqlKindNameDao{}

type sqlKindNameDao struct {
	sessionFactory *db.SessionFactory
}

func NewKindNameDao(sessionFactory *db.SessionFactory) KindNameDao {
	return &sqlKindNameDao{sessionFactory: sessionFactory}
}

func (d *sqlKindNameDao) Get(ctx context.Context, id string) (*api.KindName, error) {
	g2 := (*d.sessionFactory).New(ctx)
	var kindname api.KindName
	if err := g2.Take(&kindname, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &kindname, nil
}

func (d *sqlKindNameDao) Create(ctx context.Context, kindname *api.KindName) (*api.KindName, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Create(kindname).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return kindname, nil
}

func (d *sqlKindNameDao) Replace(ctx context.Context, kindname *api.KindName) (*api.KindName, error) {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Save(kindname).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return nil, err
	}
	return kindname, nil
}

func (d *sqlKindNameDao) Delete(ctx context.Context, id string) error {
	g2 := (*d.sessionFactory).New(ctx)
	if err := g2.Omit(clause.Associations).Delete(&api.KindName{Meta: api.Meta{ID: id}}).Error; err != nil {
		db.MarkForRollback(ctx, err)
		return err
	}
	return nil
}

func (d *sqlKindNameDao) FindByIDs(ctx context.Context, ids []string) (api.KindNameList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	kindnames := api.KindNameList{}
	if err := g2.Where("id in (?)", ids).Find(&kindnames).Error; err != nil {
		return nil, err
	}
	return kindnames, nil
}

func (d *sqlKindNameDao) All(ctx context.Context) (api.KindNameList, error) {
	g2 := (*d.sessionFactory).New(ctx)
	kindnames := api.KindNameList{}
	if err := g2.Find(&kindnames).Error; err != nil {
		return nil, err
	}
	return kindnames, nil
}
