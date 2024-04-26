package services

import (
	"context"
	"github.com/openshift-online/rh-trex/pkg/dao"

	"github.com/openshift-online/rh-trex/pkg/api"
	"github.com/openshift-online/rh-trex/pkg/errors"
)

type KindNameService interface {
	Get(ctx context.Context, id string) (*api.KindName, *errors.ServiceError)
	Create(ctx context.Context, kindname *api.KindName) (*api.KindName, *errors.ServiceError)
	Replace(ctx context.Context, kindname *api.KindName) (*api.KindName, *errors.ServiceError)
	Delete(ctx context.Context, id string) *errors.ServiceError
	All(ctx context.Context) (api.KindNameList, *errors.ServiceError)

	FindByIDs(ctx context.Context, ids []string) (api.KindNameList, *errors.ServiceError)
}

func NewKindNameService(kindnameDao dao.KindNameDao, events EventService) KindNameService {
	return &sqlKindNameService{
		kindnameDao: kindnameDao,
		events:      events,
	}
}

var _ KindNameService = &sqlKindNameService{}

type sqlKindNameService struct {
	kindnameDao dao.KindNameDao
	events      EventService
}

func (s *sqlKindNameService) Get(ctx context.Context, id string) (*api.KindName, *errors.ServiceError) {
	kindname, err := s.kindnameDao.Get(ctx, id)
	if err != nil {
		return nil, handleGetError("KindName", "id", id, err)
	}
	return kindname, nil
}

func (s *sqlKindNameService) Create(ctx context.Context, kindname *api.KindName) (*api.KindName, *errors.ServiceError) {
	kindname, err := s.kindnameDao.Create(ctx, kindname)
	if err != nil {
		return nil, handleCreateError("KindName", err)
	}

	_, evErr := s.events.Create(ctx, &api.Event{
		Source:    "KindNames",
		SourceID:  kindname.ID,
		EventType: api.CreateEventType,
	})
	if evErr != nil {
		return nil, handleCreateError("KindName", evErr)
	}

	return kindname, nil
}

func (s *sqlKindNameService) Replace(ctx context.Context, kindname *api.KindName) (*api.KindName, *errors.ServiceError) {
	kindname, err := s.kindnameDao.Replace(ctx, kindname)
	if err != nil {
		return nil, handleUpdateError("KindName", err)
	}

	_, evErr := s.events.Create(ctx, &api.Event{
		Source:    "KindNames",
		SourceID:  kindname.ID,
		EventType: api.UpdateEventType,
	})
	if evErr != nil {
		return nil, handleUpdateError("KindName", evErr)
	}

	return kindname, nil
}

func (s *sqlKindNameService) Delete(ctx context.Context, id string) *errors.ServiceError {
	if err := s.kindnameDao.Delete(ctx, id); err != nil {
		return handleDeleteError("KindName", errors.GeneralError("Unable to delete kindname: %s", err))
	}

	_, evErr := s.events.Create(ctx, &api.Event{
		Source:    "KindNames",
		SourceID:  id,
		EventType: api.DeleteEventType,
	})
	if evErr != nil {
		return handleDeleteError("KindName", evErr)
	}

	return nil
}

func (s *sqlKindNameService) FindByIDs(ctx context.Context, ids []string) (api.KindNameList, *errors.ServiceError) {
	kindnames, err := s.kindnameDao.FindByIDs(ctx, ids)
	if err != nil {
		return nil, errors.GeneralError("Unable to get all kindnames: %s", err)
	}
	return kindnames, nil
}

func (s *sqlKindNameService) All(ctx context.Context) (api.KindNameList, *errors.ServiceError) {
	kindnames, err := s.kindnameDao.All(ctx)
	if err != nil {
		return nil, errors.GeneralError("Unable to get all kindnames: %s", err)
	}
	return kindnames, nil
}
